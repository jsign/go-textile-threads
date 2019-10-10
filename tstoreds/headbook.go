package tstoreds

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
	pb "github.com/textileio/go-textile-threads/pb"
	"github.com/whyrusleeping/base32"
)

type dsHeadBook struct {
	ds ds.TxnDatastore
}

// Heads are stored in db key pattern:
// /thread/heads/<base32 thread id no padding>/<base32 peer id no padding>
var (
	hbBase                 = ds.NewKey("/thread/heads")
	_      tstore.HeadBook = (*dsHeadBook)(nil)
)

// NewHeadBook returns a new HeadBook backed by a datastore.
func NewHeadBook(ds ds.TxnDatastore) tstore.HeadBook {
	return &dsHeadBook{
		ds: ds,
	}
}

// AddHead addes a new head to a log.
func (hb *dsHeadBook) AddHead(t thread.ID, p peer.ID, head cid.Cid) {
	hb.AddHeads(t, p, []cid.Cid{head})
}

// AddHeads adds multiple heads to a log.
func (hb *dsHeadBook) AddHeads(t thread.ID, p peer.ID, heads []cid.Cid) {
	txn, err := hb.ds.NewTransaction(false)
	defer txn.Discard()
	if err != nil {
		panic(fmt.Sprintf("error when creating txn in datastore: %v", err))
	}
	key := genBaseHead(t, p)
	v, err := txn.Get(key)
	if err != nil {
		panic(fmt.Sprintf("error when getting current heads from log %v: %v", key, err))
	}
	hr := pb.HeadBookRecord{}
	if err := proto.Unmarshal(v, &hr); err != nil {
		panic(fmt.Sprintf("error unmarshaling headbookrecord proto: %v", err))
	}

	set := make(map[cid.Cid]struct{})
	for i := range hr.Heads {
		set[hr.Heads[i].Cid.Cid] = struct{}{}
	}
	for i := range heads {
		if !heads[i].Defined() {
			log.Warningf("ignoring head %s is is undefined for %s", heads[i], key)
			continue
		}
		if _, ok := set[heads[i]]; !ok {
			entry := &pb.HeadBookRecord_HeadEntry{Cid: &pb.HeadCid{Cid: heads[i]}}
			hr.Heads = append(hr.Heads, entry)
		}
	}
	data, err := proto.Marshal(&hr)
	if err != nil {
		panic(fmt.Sprintf("error when marshaling headbookrecord proto for %v: %v", key, err))
	}
	if err = txn.Put(key, data); err != nil {
		panic(fmt.Sprintf("error when saving new head record in datastore for %v: %v", key, err))
	}
	txn.Commit()
}

func (hb *dsHeadBook) SetHead(t thread.ID, p peer.ID, c cid.Cid) {
	hb.SetHeads(t, p, []cid.Cid{c})
}

func (hb *dsHeadBook) SetHeads(t thread.ID, p peer.ID, heads []cid.Cid) {
	key := genBaseHead(t, p)
	hr := pb.HeadBookRecord{}
	for i := range heads {
		if !heads[i].Defined() {
			log.Warningf("ignoring head %s is is undefined for %s", heads[i], key)
			continue
		}
		entry := &pb.HeadBookRecord_HeadEntry{Cid: &pb.HeadCid{Cid: heads[i]}}
		hr.Heads = append(hr.Heads, entry)

	}
	data, err := proto.Marshal(&hr)
	if err != nil {
		panic(fmt.Sprintf("error when marshaling headbookrecord proto for %v: %v", key, err))
	}
	if err = hb.ds.Put(key, data); err != nil {
		panic(fmt.Sprintf("error when saving new head record in datastore for %v: %v", key, err))
	}
}

func (hb *dsHeadBook) Heads(t thread.ID, p peer.ID) []cid.Cid {
	key := genBaseHead(t, p)
	v, err := hb.ds.Get(key)
	if err != nil {
		panic(fmt.Sprintf("error when getting current heads from log %v: %v", key, err))
	}
	hr := pb.HeadBookRecord{}
	if err := proto.Unmarshal(v, &hr); err != nil {
		panic(fmt.Sprintf("error unmarshaling headbookrecord proto: %v", err))
	}
	ret := make([]cid.Cid, len(hr.Heads))
	for i := range hr.Heads {
		ret[i] = hr.Heads[i].Cid.Cid
	}
	return ret
}

func (hb *dsHeadBook) ClearHeads(t thread.ID, p peer.ID) {
	key := genBaseHead(t, p)
	if err := hb.ds.Delete(key); err != nil {
		panic(fmt.Sprintf("error when deleting heads from %v", key))
	}
}

func genBaseHead(t thread.ID, p peer.ID) ds.Key {
	key := hbBase.ChildString(base32.RawStdEncoding.EncodeToString(t.Bytes()))
	key = key.ChildString(base32.RawStdEncoding.EncodeToString([]byte(p)))
	return key
}
