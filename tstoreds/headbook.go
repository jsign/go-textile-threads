package tstoreds

import (
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-textile-core/thread"
	tstore "github.com/textileio/go-textile-core/threadstore"
)

type dsHeadBook struct {
	ds ds.Datastore
}

var _ tstore.HeadBook = (*dsHeadBook)(nil)

func NewHeadBook(ds ds.Datastore) tstore.HeadBook {
	return &dsHeadBook{
		ds: ds,
	}
}

func (hb *dsHeadBook) AddHead(_ thread.ID, _ peer.ID, _ cid.Cid) {
	panic("not implemented")
}

func (hb *dsHeadBook) AddHeads(_ thread.ID, _ peer.ID, _ []cid.Cid) {
	panic("not implemented")
}

func (hb *dsHeadBook) SetHead(_ thread.ID, _ peer.ID, _ cid.Cid) {
	panic("not implemented")
}

func (hb *dsHeadBook) SetHeads(_ thread.ID, _ peer.ID, _ []cid.Cid) {
	panic("not implemented")
}

func (hb *dsHeadBook) Heads(_ thread.ID, _ peer.ID) []cid.Cid {
	panic("not implemented")
}

func (hb *dsHeadBook) ClearHeads(_ thread.ID, _ peer.ID) {
	panic("not implemented")
}
