// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tstore.proto

package threads_pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// AddrBookRecord represents a record for a log in the address book.
type AddrBookRecord struct {
	// Thread ID.
	ThreadID *ProtoThreadID `protobuf:"bytes,1,opt,name=threadID,proto3,customtype=ProtoThreadID" json:"threadID,omitempty"`
	// The peer ID.
	PeerID *ProtoPeerID `protobuf:"bytes,2,opt,name=peerID,proto3,customtype=ProtoPeerID" json:"peerID,omitempty"`
	// The multiaddresses. This is a sorted list where element 0 expires the soonest.
	Addrs []*AddrBookRecord_AddrEntry `protobuf:"bytes,3,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (m *AddrBookRecord) Reset()         { *m = AddrBookRecord{} }
func (m *AddrBookRecord) String() string { return proto.CompactTextString(m) }
func (*AddrBookRecord) ProtoMessage()    {}
func (*AddrBookRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_74b294050729ac9e, []int{0}
}
func (m *AddrBookRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddrBookRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddrBookRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddrBookRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrBookRecord.Merge(m, src)
}
func (m *AddrBookRecord) XXX_Size() int {
	return m.Size()
}
func (m *AddrBookRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrBookRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AddrBookRecord proto.InternalMessageInfo

func (m *AddrBookRecord) GetAddrs() []*AddrBookRecord_AddrEntry {
	if m != nil {
		return m.Addrs
	}
	return nil
}

// AddrEntry represents a single multiaddress.
type AddrBookRecord_AddrEntry struct {
	Addr *ProtoAddr `protobuf:"bytes,1,opt,name=addr,proto3,customtype=ProtoAddr" json:"addr,omitempty"`
	// The point in time when this address expires.
	Expiry int64 `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// The original TTL of this address.
	Ttl int64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (m *AddrBookRecord_AddrEntry) Reset()         { *m = AddrBookRecord_AddrEntry{} }
func (m *AddrBookRecord_AddrEntry) String() string { return proto.CompactTextString(m) }
func (*AddrBookRecord_AddrEntry) ProtoMessage()    {}
func (*AddrBookRecord_AddrEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_74b294050729ac9e, []int{0, 0}
}
func (m *AddrBookRecord_AddrEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddrBookRecord_AddrEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddrBookRecord_AddrEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddrBookRecord_AddrEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrBookRecord_AddrEntry.Merge(m, src)
}
func (m *AddrBookRecord_AddrEntry) XXX_Size() int {
	return m.Size()
}
func (m *AddrBookRecord_AddrEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrBookRecord_AddrEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AddrBookRecord_AddrEntry proto.InternalMessageInfo

func (m *AddrBookRecord_AddrEntry) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *AddrBookRecord_AddrEntry) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// HeadBookRecord represents the list of heads currently in a log
type HeadBookRecord struct {
	// List of current heads of a log.
	Heads []*HeadBookRecord_HeadEntry `protobuf:"bytes,1,rep,name=heads,proto3" json:"heads,omitempty"`
}

func (m *HeadBookRecord) Reset()         { *m = HeadBookRecord{} }
func (m *HeadBookRecord) String() string { return proto.CompactTextString(m) }
func (*HeadBookRecord) ProtoMessage()    {}
func (*HeadBookRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_74b294050729ac9e, []int{1}
}
func (m *HeadBookRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadBookRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadBookRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadBookRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadBookRecord.Merge(m, src)
}
func (m *HeadBookRecord) XXX_Size() int {
	return m.Size()
}
func (m *HeadBookRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadBookRecord.DiscardUnknown(m)
}

var xxx_messageInfo_HeadBookRecord proto.InternalMessageInfo

func (m *HeadBookRecord) GetHeads() []*HeadBookRecord_HeadEntry {
	if m != nil {
		return m.Heads
	}
	return nil
}

// HeadEntry represents a single cid.
type HeadBookRecord_HeadEntry struct {
	Cid *ProtoCid `protobuf:"bytes,1,opt,name=cid,proto3,customtype=ProtoCid" json:"cid,omitempty"`
}

func (m *HeadBookRecord_HeadEntry) Reset()         { *m = HeadBookRecord_HeadEntry{} }
func (m *HeadBookRecord_HeadEntry) String() string { return proto.CompactTextString(m) }
func (*HeadBookRecord_HeadEntry) ProtoMessage()    {}
func (*HeadBookRecord_HeadEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_74b294050729ac9e, []int{1, 0}
}
func (m *HeadBookRecord_HeadEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadBookRecord_HeadEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadBookRecord_HeadEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadBookRecord_HeadEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadBookRecord_HeadEntry.Merge(m, src)
}
func (m *HeadBookRecord_HeadEntry) XXX_Size() int {
	return m.Size()
}
func (m *HeadBookRecord_HeadEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadBookRecord_HeadEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HeadBookRecord_HeadEntry proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddrBookRecord)(nil), "threads.pb.AddrBookRecord")
	proto.RegisterType((*AddrBookRecord_AddrEntry)(nil), "threads.pb.AddrBookRecord.AddrEntry")
	proto.RegisterType((*HeadBookRecord)(nil), "threads.pb.HeadBookRecord")
	proto.RegisterType((*HeadBookRecord_HeadEntry)(nil), "threads.pb.HeadBookRecord.HeadEntry")
}

func init() { proto.RegisterFile("tstore.proto", fileDescriptor_74b294050729ac9e) }

var fileDescriptor_74b294050729ac9e = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x3d, 0x4f, 0x02, 0x31,
	0x18, 0xc7, 0xa9, 0x27, 0x04, 0x1e, 0x5e, 0xd4, 0x0e, 0xe6, 0xc2, 0xd0, 0x43, 0x62, 0x22, 0x89,
	0xe1, 0x48, 0x74, 0x63, 0x13, 0x31, 0x91, 0x8d, 0x34, 0x0e, 0xae, 0x1c, 0xad, 0x70, 0xf1, 0xa5,
	0x97, 0x52, 0x12, 0xf9, 0x16, 0x7e, 0x24, 0x47, 0x47, 0x46, 0x73, 0xc3, 0x45, 0x8f, 0x2f, 0xe1,
	0x64, 0xcc, 0x3d, 0x34, 0x28, 0x71, 0x7b, 0xfe, 0x2f, 0x4f, 0xfb, 0x6b, 0x0a, 0x15, 0x33, 0x33,
	0x4a, 0x4b, 0x3f, 0xd2, 0xca, 0x28, 0x0a, 0x66, 0xaa, 0xe5, 0x48, 0xcc, 0xfc, 0x28, 0xa8, 0xb7,
	0x27, 0xa1, 0x99, 0xce, 0x03, 0x7f, 0xac, 0x1e, 0x3b, 0x13, 0x35, 0x51, 0x1d, 0xac, 0x04, 0xf3,
	0x3b, 0x54, 0x28, 0x70, 0x5a, 0xaf, 0x36, 0xbf, 0x09, 0xd4, 0x2e, 0x84, 0xd0, 0x3d, 0xa5, 0xee,
	0xb9, 0x1c, 0x2b, 0x2d, 0x68, 0x1b, 0x8a, 0xeb, 0xf3, 0x06, 0x7d, 0x97, 0x34, 0x48, 0xab, 0xd2,
	0x3b, 0x88, 0x13, 0xaf, 0x3a, 0xcc, 0xfa, 0x37, 0x36, 0xe0, 0x9b, 0x0a, 0x3d, 0x81, 0x42, 0x24,
	0xa5, 0x1e, 0xf4, 0xdd, 0x1d, 0x2c, 0xef, 0xc5, 0x89, 0x57, 0xc6, 0xf2, 0x10, 0x6d, 0x6e, 0x63,
	0xda, 0x85, 0xfc, 0x48, 0x08, 0x3d, 0x73, 0x9d, 0x86, 0xd3, 0x2a, 0x9f, 0x1d, 0xfb, 0xbf, 0xd4,
	0xfe, 0x36, 0x02, 0xca, 0xab, 0x27, 0xa3, 0x17, 0x7c, 0xbd, 0x52, 0xbf, 0x85, 0xd2, 0xc6, 0xa3,
	0x47, 0xb0, 0x9b, 0xb9, 0x16, 0xae, 0x1a, 0x27, 0x5e, 0x09, 0xef, 0xcb, 0x1a, 0x1c, 0x23, 0x7a,
	0x08, 0x05, 0xf9, 0x1c, 0x85, 0x7a, 0x81, 0x50, 0x0e, 0xb7, 0x8a, 0xee, 0x83, 0x63, 0xcc, 0x83,
	0xeb, 0xa0, 0x99, 0x8d, 0xcd, 0x05, 0xd4, 0xae, 0xe5, 0x48, 0xfc, 0x79, 0x7f, 0x17, 0xf2, 0xd3,
	0x8c, 0xcb, 0x25, 0xff, 0x39, 0xb7, 0xab, 0x28, 0x2d, 0x27, 0xae, 0xd4, 0x4f, 0xa1, 0xb4, 0xf1,
	0x28, 0x03, 0x67, 0x1c, 0x0a, 0x8b, 0x59, 0x89, 0x13, 0xaf, 0x88, 0x98, 0x97, 0xa1, 0xe0, 0x59,
	0xd0, 0x6b, 0x7c, 0x7d, 0x32, 0xf2, 0x9a, 0x32, 0xf2, 0x96, 0x32, 0xb2, 0x4c, 0x19, 0xf9, 0x48,
	0x19, 0x79, 0x59, 0xb1, 0xdc, 0x72, 0xc5, 0x72, 0xef, 0x2b, 0x96, 0x0b, 0x0a, 0xf8, 0x49, 0xe7,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x24, 0x86, 0x8a, 0xef, 0x01, 0x00, 0x00,
}

func (m *AddrBookRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddrBookRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ThreadID != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.ThreadID.Size()))
		n1, err := m.ThreadID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PeerID != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.PeerID.Size()))
		n2, err := m.PeerID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Addrs) > 0 {
		for _, msg := range m.Addrs {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintTstore(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AddrBookRecord_AddrEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddrBookRecord_AddrEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Addr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.Addr.Size()))
		n3, err := m.Addr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Expiry != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.Expiry))
	}
	if m.Ttl != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.Ttl))
	}
	return i, nil
}

func (m *HeadBookRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadBookRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Heads) > 0 {
		for _, msg := range m.Heads {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTstore(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HeadBookRecord_HeadEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadBookRecord_HeadEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Cid != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTstore(dAtA, i, uint64(m.Cid.Size()))
		n4, err := m.Cid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintTstore(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAddrBookRecord(r randyTstore, easy bool) *AddrBookRecord {
	this := &AddrBookRecord{}
	this.ThreadID = NewPopulatedProtoThreadID(r)
	this.PeerID = NewPopulatedProtoPeerID(r)
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Addrs = make([]*AddrBookRecord_AddrEntry, v1)
		for i := 0; i < v1; i++ {
			this.Addrs[i] = NewPopulatedAddrBookRecord_AddrEntry(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAddrBookRecord_AddrEntry(r randyTstore, easy bool) *AddrBookRecord_AddrEntry {
	this := &AddrBookRecord_AddrEntry{}
	this.Addr = NewPopulatedProtoAddr(r)
	this.Expiry = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Expiry *= -1
	}
	this.Ttl = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Ttl *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHeadBookRecord(r randyTstore, easy bool) *HeadBookRecord {
	this := &HeadBookRecord{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Heads = make([]*HeadBookRecord_HeadEntry, v2)
		for i := 0; i < v2; i++ {
			this.Heads[i] = NewPopulatedHeadBookRecord_HeadEntry(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHeadBookRecord_HeadEntry(r randyTstore, easy bool) *HeadBookRecord_HeadEntry {
	this := &HeadBookRecord_HeadEntry{}
	this.Cid = NewPopulatedProtoCid(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTstore interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTstore(r randyTstore) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTstore(r randyTstore) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTstore(r)
	}
	return string(tmps)
}
func randUnrecognizedTstore(r randyTstore, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTstore(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTstore(dAtA []byte, r randyTstore, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTstore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTstore(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AddrBookRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ThreadID != nil {
		l = m.ThreadID.Size()
		n += 1 + l + sovTstore(uint64(l))
	}
	if m.PeerID != nil {
		l = m.PeerID.Size()
		n += 1 + l + sovTstore(uint64(l))
	}
	if len(m.Addrs) > 0 {
		for _, e := range m.Addrs {
			l = e.Size()
			n += 1 + l + sovTstore(uint64(l))
		}
	}
	return n
}

func (m *AddrBookRecord_AddrEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Addr != nil {
		l = m.Addr.Size()
		n += 1 + l + sovTstore(uint64(l))
	}
	if m.Expiry != 0 {
		n += 1 + sovTstore(uint64(m.Expiry))
	}
	if m.Ttl != 0 {
		n += 1 + sovTstore(uint64(m.Ttl))
	}
	return n
}

func (m *HeadBookRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Heads) > 0 {
		for _, e := range m.Heads {
			l = e.Size()
			n += 1 + l + sovTstore(uint64(l))
		}
	}
	return n
}

func (m *HeadBookRecord_HeadEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cid != nil {
		l = m.Cid.Size()
		n += 1 + l + sovTstore(uint64(l))
	}
	return n
}

func sovTstore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTstore(x uint64) (n int) {
	return sovTstore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddrBookRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddrBookRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddrBookRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreadID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoThreadID
			m.ThreadID = &v
			if err := m.ThreadID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoPeerID
			m.PeerID = &v
			if err := m.PeerID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, &AddrBookRecord_AddrEntry{})
			if err := m.Addrs[len(m.Addrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddrBookRecord_AddrEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddrEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddrEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoAddr
			m.Addr = &v
			if err := m.Addr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeadBookRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeadBookRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadBookRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Heads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Heads = append(m.Heads, &HeadBookRecord_HeadEntry{})
			if err := m.Heads[len(m.Heads)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeadBookRecord_HeadEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTstore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeadEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTstore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTstore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ProtoCid
			m.Cid = &v
			if err := m.Cid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTstore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTstore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTstore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTstore
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTstore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTstore
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTstore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTstore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTstore(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTstore
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTstore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTstore   = fmt.Errorf("proto: integer overflow")
)
