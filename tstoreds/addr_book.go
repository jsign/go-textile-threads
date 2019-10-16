package tstoreds

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-textile-core/thread"
	"github.com/textileio/go-textile-core/threadstore"
	pb "github.com/textileio/go-textile-threads/pb"
	"github.com/whyrusleeping/base32"
)

type ttlWriteMode int

const (
	ttlOverride ttlWriteMode = iota
	ttlExtend
)

var (
	log = logging.Logger("threadstore")

	// Thread addresses are stored db key pattern:
	// /thread/addrs/<b32 thread id no padding>/<b32 log id no padding>
	threadBookBase = ds.NewKey("/thread/addrs")
)

type DsAddrBook struct {
	ctx         context.Context
	opts        Options
	ds          ds.Batching
	cache       cache
	gc          *dsAddrBookGc
	subsManager *pstoremem.AddrSubManager

	// controls children goroutine lifetime.
	childrenDone sync.WaitGroup
	cancelFn     func()
}

var _ threadstore.AddrBook = (*DsAddrBook)(nil)

// addrsRecord decorates the AddrBookRecord with locks and metadata.
type addrsRecord struct {
	sync.RWMutex
	*pb.AddrBookRecord
	dirty bool
}

// cacheKey is a comparable struct used as a key in the book cache
type cacheKey struct {
	threadID thread.ID
	peerID   peer.ID
}

// NewAddrBook initializes a new datastore-backed address book. It serves as a drop-in replacement for pstoremem
// (memory-backed peerstore), and works with any datastore implementing the ds.Batching interface.
//
// Threads and logs addresses are serialized into protobuf, storing one datastore entry per (thread, log), along with metadata
// to control address expiration. To alleviate disk access and serde overhead, we internally use a read/write-through
// ARC cache, the size of which is adjustable via Options.CacheSize.
//
// The user has a choice of two GC algorithms:
//  - full-purge GC (default): performs a full visit of the store with periodicity Options.GCPurgeInterval. Useful when
//    the range of possible TTL values is small and the values themselves are also extreme, e.g. 10 minutes or
//    permanent, popular values used in other libp2p modules. In this cited case, optimizing with lookahead windows
//    makes little sense.
func NewAddrBook(ctx context.Context, ds ds.Batching, opts Options) (*DsAddrBook, error) {
	ctx, cancelFn := context.WithCancel(ctx)
	ab := &DsAddrBook{
		ctx:         ctx,
		ds:          ds,
		opts:        opts,
		subsManager: pstoremem.NewAddrSubManager(),
		cancelFn:    cancelFn,
	}

	var err error
	if opts.CacheSize > 0 {
		if ab.cache, err = lru.NewARC(int(opts.CacheSize)); err != nil {
			return nil, err
		}
	} else {
		ab.cache = new(noopCache)
	}

	if ab.gc, err = newAddressBookGc(ctx, ab); err != nil {
		return nil, err
	}

	return ab, nil
}

// AddAddr will add a new address if it's not already in the AddrBook.
func (ab *DsAddrBook) AddAddr(t thread.ID, p peer.ID, addr ma.Multiaddr, ttl time.Duration) error {
	return ab.AddAddrs(t, p, []ma.Multiaddr{addr}, ttl)
}

// AddAddrs will add many multiple addresses if they aren't already in the AddrBook.
func (ab *DsAddrBook) AddAddrs(t thread.ID, p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}
	addrs = cleanAddrs(addrs)
	if err := ab.setAddrs(t, p, addrs, ttl, ttlExtend); err != nil {
		return err
	}
	return nil
}

func (ab *DsAddrBook) SetAddr(t thread.ID, p peer.ID, addr ma.Multiaddr, ttl time.Duration) error {
	return ab.SetAddrs(t, p, []ma.Multiaddr{addr}, ttl)
}

func (ab *DsAddrBook) SetAddrs(t thread.ID, p peer.ID, addrs []ma.Multiaddr, ttl time.Duration) error {
	addrs = cleanAddrs(addrs)
	if ttl <= 0 {
		err := ab.deleteAddrs(t, p, addrs)
		return err
	}
	if err := ab.setAddrs(t, p, addrs, ttl, ttlOverride); err != nil {
		return err
	}
	return nil
}

func (ab *DsAddrBook) UpdateAddrs(t thread.ID, p peer.ID, oldTTL time.Duration, newTTL time.Duration) error {
	pr, err := ab.loadRecord(t, p, true, false)
	if err != nil {
		return fmt.Errorf("failed to update ttls for log %s: %w", p.Pretty(), err)

	}

	pr.Lock()
	defer pr.Unlock()

	newExp := time.Now().Add(newTTL).Unix()
	for _, entry := range pr.Addrs {
		if entry.Ttl != int64(oldTTL) {
			continue
		}
		entry.Ttl, entry.Expiry = int64(newTTL), newExp
		pr.dirty = true
	}

	if pr.clean() {
		if err := pr.flush(ab.ds); err != nil {
			return err
		}
	}
	return nil
}

func (ab *DsAddrBook) Addrs(t thread.ID, p peer.ID) ([]ma.Multiaddr, error) {
	pr, err := ab.loadRecord(t, p, true, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load peerstore entry for log %s while querying addrs: %w", p.Pretty(), err)
	}

	pr.RLock()
	defer pr.RUnlock()
	addrs := make([]ma.Multiaddr, 0, len(pr.Addrs))
	for _, a := range pr.Addrs {
		addrs = append(addrs, a.Addr)
	}
	return addrs, nil
}

func (ab *DsAddrBook) AddrStream(ctx context.Context, t thread.ID, p peer.ID) (<-chan ma.Multiaddr, error) {
	initial, err := ab.Addrs(t, p)
	if err != nil {
		return nil, err
	}
	return ab.subsManager.AddrStream(ctx, p, initial), nil
}

func (ab *DsAddrBook) ClearAddrs(t thread.ID, p peer.ID) error {
	ab.cache.Remove(genCacheKey(t, p))

	key := genDSKey(t, p)
	if err := ab.ds.Delete(key); err != nil {
		return fmt.Errorf("failed to clear addresses for log %s: %w", p.Pretty(), err)
	}
	return nil
}

func (ab *DsAddrBook) LogsWithAddrs(t thread.ID) (peer.IDSlice, error) {
	ids, err := uniqueLogIds(ab.ds, threadBookBase.ChildString(base32.RawStdEncoding.EncodeToString(t.Bytes())), func(result query.Result) string {
		return ds.RawKey(result.Key).Name()
	})
	if err != nil {
		return nil, fmt.Errorf("error while retrieving logs with addresses: %w", err)
	}
	return ids, nil
}

func (ab *DsAddrBook) ThreadsFromAddrs() (thread.IDSlice, error) {
	ids, err := uniqueThreadIds(ab.ds, threadBookBase, func(result query.Result) string {
		return ds.RawKey(result.Key).Parent().Name()
	})
	if err != nil {
		return nil, fmt.Errorf("error while retrieving thread from addresses: %w", err)
	}
	return ids, nil
}

// loadRecord is a read-through fetch. It fetches a record from cache, falling back to the
// datastore upon a miss, and returning a newly initialized record if the peer doesn't exist.
//
// loadRecord calls clean() on an existing record before returning it. If the record changes
// as a result and the update argument is true, the resulting state is saved in the datastore.
//
// If the cache argument is true, the record is inserted in the cache when loaded from the datastore.
func (ab *DsAddrBook) loadRecord(t thread.ID, p peer.ID, cache bool, update bool) (pr *addrsRecord, err error) {
	cacheKey := genCacheKey(t, p)
	if e, ok := ab.cache.Get(cacheKey); ok {
		pr = e.(*addrsRecord)
		pr.Lock()
		defer pr.Unlock()
		if pr.clean() && update {
			err = pr.flush(ab.ds)
		}
		return pr, err
	}

	pr = &addrsRecord{AddrBookRecord: &pb.AddrBookRecord{}}
	key := genDSKey(t, p)
	data, err := ab.ds.Get(key)
	switch err {
	case ds.ErrNotFound:
		err = nil
		pr.ThreadID = &pb.ProtoThreadID{ID: t}
		pr.PeerID = &pb.ProtoPeerID{ID: p}
	case nil:
		if err = pr.Unmarshal(data); err != nil {
			return nil, err
		}
		// this record is new and local for now (not in cache), so we don't need to lock.
		if pr.clean() && update {
			err = pr.flush(ab.ds)
		}
	default:
		return nil, err
	}

	if cache {
		ab.cache.Add(cacheKey, pr)
	}
	return pr, err

}

// clean is called on records to perform housekeeping. The return value indicates if the record was changed
// as a result of this call.
//
// clean does the following:
// * sorts addresses by expiration (soonest expiring first).
// * removes expired addresses.
//
// It short-circuits optimistically when there's nothing to do.
//
// clean is called from several points:
// * when accessing an entry.
// * when performing periodic GC.
// * after an entry has been modified (e.g. addresses have been added or removed, TTLs updated, etc.)
//
// If the return value is true, the caller should perform a flush immediately to sync the record with the store.
func (r *addrsRecord) clean() (chgd bool) {
	now := time.Now().Unix()
	if !r.dirty && len(r.Addrs) > 0 && r.Addrs[0].Expiry > now {
		// record is not dirty, and we have no expired entries to purge.
		return false
	}

	if len(r.Addrs) == 0 {
		// this is a ghost record; let's signal it has to be written.
		// flush() will take care of doing the deletion.
		return true
	}

	if r.dirty && len(r.Addrs) > 1 {
		// the record has been modified, so it may need resorting.
		// we keep addresses sorted by expiration, where 0 is the soonest expiring.
		sort.Slice(r.Addrs, func(i, j int) bool {
			return r.Addrs[i].Expiry < r.Addrs[j].Expiry
		})
	}

	// since addresses are sorted by expiration, we find the first
	// survivor and split the slice on its index.
	pivot := -1
	for i, addr := range r.Addrs {
		if addr.Expiry > now {
			break
		}
		pivot = i
	}

	r.Addrs = r.Addrs[pivot+1:]
	return r.dirty || pivot >= 0
}

// flush writes the record to the datastore by calling ds.Put, unless the record is
// marked for deletion, in which case we call ds.Delete. To be called within a lock.
func (r *addrsRecord) flush(write ds.Write) (err error) {
	key := genDSKey(r.ThreadID.ID, r.PeerID.ID)
	if len(r.Addrs) == 0 {
		if err = write.Delete(key); err == nil {
			r.dirty = false
		}
		return err
	}

	data, err := r.Marshal()
	if err != nil {
		return err
	}
	if err = write.Put(key, data); err != nil {
		return err
	}
	// write succeeded; record is no longer dirty.
	r.dirty = false
	return nil
}

func genDSKey(t thread.ID, p peer.ID) ds.Key {
	return threadBookBase.ChildString(base32.RawStdEncoding.EncodeToString(t.Bytes())).ChildString(base32.RawStdEncoding.EncodeToString([]byte(p)))
}

func genCacheKey(t thread.ID, p peer.ID) cacheKey {
	return cacheKey{threadID: t, peerID: p}
}

func (ab *DsAddrBook) setAddrs(t thread.ID, p peer.ID, addrs []ma.Multiaddr, ttl time.Duration, mode ttlWriteMode) (err error) {
	pr, err := ab.loadRecord(t, p, true, false)
	if err != nil {
		return fmt.Errorf("failed to load peerstore entry for log %v while setting addrs, err: %v", p, err)
	}

	pr.Lock()
	defer pr.Unlock()

	newExp := time.Now().Add(ttl).Unix()
	existed := make([]bool, len(addrs)) // keeps track of which addrs we found.

Outer:
	for i, incoming := range addrs {
		for _, have := range pr.Addrs {
			if incoming.Equal(have.Addr) {
				existed[i] = true
				switch mode {
				case ttlOverride:
					have.Ttl = int64(ttl)
					have.Expiry = newExp
				case ttlExtend:
					if int64(ttl) > have.Ttl {
						have.Ttl = int64(ttl)
					}
					if newExp > have.Expiry {
						have.Expiry = newExp
					}
				default:
					panic("BUG: unimplemented ttl mode")
				}

				// we found the address, and addresses cannot be duplicate,
				// so let's move on to the next.
				continue Outer
			}
		}
	}

	// add addresses we didn't hold.
	var added []*pb.AddrBookRecord_AddrEntry
	for i, e := range existed {
		if e {
			continue
		}
		addr := addrs[i]
		entry := &pb.AddrBookRecord_AddrEntry{
			Addr:   &pb.ProtoAddr{Multiaddr: addr},
			Ttl:    int64(ttl),
			Expiry: newExp,
		}
		added = append(added, entry)
		// note: there's a minor chance that writing the record will fail, in which case we would've broadcast
		// the addresses without persisting them. This is very unlikely and not much of an issue.
		ab.subsManager.BroadcastAddr(p, addr)
	}

	pr.Addrs = append(pr.Addrs, added...)
	pr.dirty = true
	pr.clean()
	return pr.flush(ab.ds)
}

func (ab *DsAddrBook) deleteAddrs(t thread.ID, p peer.ID, addrs []ma.Multiaddr) (err error) {
	pr, err := ab.loadRecord(t, p, false, false)
	if err != nil {
		return fmt.Errorf("failed to load peerstore entry for log %v while deleting addrs, err: %v", p, err)
	}

	if pr.Addrs == nil {
		return nil
	}

	pr.Lock()
	defer pr.Unlock()

	// deletes addresses in place, and avoiding copies until we encounter the first deletion.
	survived := 0
	for i, addr := range pr.Addrs {
		for _, del := range addrs {
			if addr.Addr.Equal(del) {
				continue
			}
			if i != survived {
				pr.Addrs[survived] = pr.Addrs[i]
			}
			survived++
		}
	}
	pr.Addrs = pr.Addrs[:survived]

	pr.dirty = true
	pr.clean()
	return pr.flush(ab.ds)
}

func cleanAddrs(addrs []ma.Multiaddr) []ma.Multiaddr {
	clean := make([]ma.Multiaddr, 0, len(addrs))
	for _, addr := range addrs {
		if addr == nil {
			continue
		}
		clean = append(clean, addr)
	}
	return clean
}

func (ab *DsAddrBook) Close() error {
	ab.cancelFn()
	ab.childrenDone.Wait()
	return nil
}
