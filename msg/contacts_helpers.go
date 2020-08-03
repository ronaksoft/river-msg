// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contacts.proto

package msg

import (
	fmt "fmt"
	pbytes "github.com/gobwas/pool/pbytes"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const C_ContactsImport int64 = 3473528730

type poolContactsImport struct {
	pool sync.Pool
}

func (p *poolContactsImport) Get() *ContactsImport {
	x, ok := p.pool.Get().(*ContactsImport)
	if !ok {
		return &ContactsImport{}
	}
	return x
}

func (p *poolContactsImport) Put(x *ContactsImport) {
	x.Contacts = x.Contacts[:0]
	p.pool.Put(x)
}

var PoolContactsImport = poolContactsImport{}

func ResultContactsImport(out *MessageEnvelope, res *ContactsImport) {
	out.Constructor = C_ContactsImport
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsAdd int64 = 1410714478

type poolContactsAdd struct {
	pool sync.Pool
}

func (p *poolContactsAdd) Get() *ContactsAdd {
	x, ok := p.pool.Get().(*ContactsAdd)
	if !ok {
		return &ContactsAdd{}
	}
	return x
}

func (p *poolContactsAdd) Put(x *ContactsAdd) {
	x.LastName = ""
	x.Phone = ""
	p.pool.Put(x)
}

var PoolContactsAdd = poolContactsAdd{}

func ResultContactsAdd(out *MessageEnvelope, res *ContactsAdd) {
	out.Constructor = C_ContactsAdd
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsGet int64 = 1412732665

type poolContactsGet struct {
	pool sync.Pool
}

func (p *poolContactsGet) Get() *ContactsGet {
	x, ok := p.pool.Get().(*ContactsGet)
	if !ok {
		return &ContactsGet{}
	}
	return x
}

func (p *poolContactsGet) Put(x *ContactsGet) {
	p.pool.Put(x)
}

var PoolContactsGet = poolContactsGet{}

func ResultContactsGet(out *MessageEnvelope, res *ContactsGet) {
	out.Constructor = C_ContactsGet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsDelete int64 = 1750426880

type poolContactsDelete struct {
	pool sync.Pool
}

func (p *poolContactsDelete) Get() *ContactsDelete {
	x, ok := p.pool.Get().(*ContactsDelete)
	if !ok {
		return &ContactsDelete{}
	}
	return x
}

func (p *poolContactsDelete) Put(x *ContactsDelete) {
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolContactsDelete = poolContactsDelete{}

func ResultContactsDelete(out *MessageEnvelope, res *ContactsDelete) {
	out.Constructor = C_ContactsDelete
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsDeleteAll int64 = 14552524

type poolContactsDeleteAll struct {
	pool sync.Pool
}

func (p *poolContactsDeleteAll) Get() *ContactsDeleteAll {
	x, ok := p.pool.Get().(*ContactsDeleteAll)
	if !ok {
		return &ContactsDeleteAll{}
	}
	return x
}

func (p *poolContactsDeleteAll) Put(x *ContactsDeleteAll) {
	p.pool.Put(x)
}

var PoolContactsDeleteAll = poolContactsDeleteAll{}

func ResultContactsDeleteAll(out *MessageEnvelope, res *ContactsDeleteAll) {
	out.Constructor = C_ContactsDeleteAll
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsBlock int64 = 2900371089

type poolContactsBlock struct {
	pool sync.Pool
}

func (p *poolContactsBlock) Get() *ContactsBlock {
	x, ok := p.pool.Get().(*ContactsBlock)
	if !ok {
		return &ContactsBlock{}
	}
	return x
}

func (p *poolContactsBlock) Put(x *ContactsBlock) {
	p.pool.Put(x)
}

var PoolContactsBlock = poolContactsBlock{}

func ResultContactsBlock(out *MessageEnvelope, res *ContactsBlock) {
	out.Constructor = C_ContactsBlock
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsUnblock int64 = 662011773

type poolContactsUnblock struct {
	pool sync.Pool
}

func (p *poolContactsUnblock) Get() *ContactsUnblock {
	x, ok := p.pool.Get().(*ContactsUnblock)
	if !ok {
		return &ContactsUnblock{}
	}
	return x
}

func (p *poolContactsUnblock) Put(x *ContactsUnblock) {
	p.pool.Put(x)
}

var PoolContactsUnblock = poolContactsUnblock{}

func ResultContactsUnblock(out *MessageEnvelope, res *ContactsUnblock) {
	out.Constructor = C_ContactsUnblock
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsGetBlocked int64 = 1073733371

type poolContactsGetBlocked struct {
	pool sync.Pool
}

func (p *poolContactsGetBlocked) Get() *ContactsGetBlocked {
	x, ok := p.pool.Get().(*ContactsGetBlocked)
	if !ok {
		return &ContactsGetBlocked{}
	}
	return x
}

func (p *poolContactsGetBlocked) Put(x *ContactsGetBlocked) {
	p.pool.Put(x)
}

var PoolContactsGetBlocked = poolContactsGetBlocked{}

func ResultContactsGetBlocked(out *MessageEnvelope, res *ContactsGetBlocked) {
	out.Constructor = C_ContactsGetBlocked
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsSearch int64 = 3870802464

type poolContactsSearch struct {
	pool sync.Pool
}

func (p *poolContactsSearch) Get() *ContactsSearch {
	x, ok := p.pool.Get().(*ContactsSearch)
	if !ok {
		return &ContactsSearch{}
	}
	return x
}

func (p *poolContactsSearch) Put(x *ContactsSearch) {
	p.pool.Put(x)
}

var PoolContactsSearch = poolContactsSearch{}

func ResultContactsSearch(out *MessageEnvelope, res *ContactsSearch) {
	out.Constructor = C_ContactsSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsGetTopPeers int64 = 1378126220

type poolContactsGetTopPeers struct {
	pool sync.Pool
}

func (p *poolContactsGetTopPeers) Get() *ContactsGetTopPeers {
	x, ok := p.pool.Get().(*ContactsGetTopPeers)
	if !ok {
		return &ContactsGetTopPeers{}
	}
	return x
}

func (p *poolContactsGetTopPeers) Put(x *ContactsGetTopPeers) {
	p.pool.Put(x)
}

var PoolContactsGetTopPeers = poolContactsGetTopPeers{}

func ResultContactsGetTopPeers(out *MessageEnvelope, res *ContactsGetTopPeers) {
	out.Constructor = C_ContactsGetTopPeers
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsResetTopPeer int64 = 1114887378

type poolContactsResetTopPeer struct {
	pool sync.Pool
}

func (p *poolContactsResetTopPeer) Get() *ContactsResetTopPeer {
	x, ok := p.pool.Get().(*ContactsResetTopPeer)
	if !ok {
		return &ContactsResetTopPeer{}
	}
	return x
}

func (p *poolContactsResetTopPeer) Put(x *ContactsResetTopPeer) {
	p.pool.Put(x)
}

var PoolContactsResetTopPeer = poolContactsResetTopPeer{}

func ResultContactsResetTopPeer(out *MessageEnvelope, res *ContactsResetTopPeer) {
	out.Constructor = C_ContactsResetTopPeer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsTopPeers int64 = 2243919622

type poolContactsTopPeers struct {
	pool sync.Pool
}

func (p *poolContactsTopPeers) Get() *ContactsTopPeers {
	x, ok := p.pool.Get().(*ContactsTopPeers)
	if !ok {
		return &ContactsTopPeers{}
	}
	return x
}

func (p *poolContactsTopPeers) Put(x *ContactsTopPeers) {
	x.Peers = x.Peers[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	p.pool.Put(x)
}

var PoolContactsTopPeers = poolContactsTopPeers{}

func ResultContactsTopPeers(out *MessageEnvelope, res *ContactsTopPeers) {
	out.Constructor = C_ContactsTopPeers
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TopPeer int64 = 1763100161

type poolTopPeer struct {
	pool sync.Pool
}

func (p *poolTopPeer) Get() *TopPeer {
	x, ok := p.pool.Get().(*TopPeer)
	if !ok {
		return &TopPeer{}
	}
	return x
}

func (p *poolTopPeer) Put(x *TopPeer) {
	x.TeamID = 0
	x.LastUpdate = 0
	p.pool.Put(x)
}

var PoolTopPeer = poolTopPeer{}

func ResultTopPeer(out *MessageEnvelope, res *TopPeer) {
	out.Constructor = C_TopPeer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_BlockedContactsMany int64 = 2067026404

type poolBlockedContactsMany struct {
	pool sync.Pool
}

func (p *poolBlockedContactsMany) Get() *BlockedContactsMany {
	x, ok := p.pool.Get().(*BlockedContactsMany)
	if !ok {
		return &BlockedContactsMany{}
	}
	return x
}

func (p *poolBlockedContactsMany) Put(x *BlockedContactsMany) {
	x.Contacts = x.Contacts[:0]
	x.Users = x.Users[:0]
	x.Total = 0
	p.pool.Put(x)
}

var PoolBlockedContactsMany = poolBlockedContactsMany{}

func ResultBlockedContactsMany(out *MessageEnvelope, res *BlockedContactsMany) {
	out.Constructor = C_BlockedContactsMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_BlockedContact int64 = 53788553

type poolBlockedContact struct {
	pool sync.Pool
}

func (p *poolBlockedContact) Get() *BlockedContact {
	x, ok := p.pool.Get().(*BlockedContact)
	if !ok {
		return &BlockedContact{}
	}
	return x
}

func (p *poolBlockedContact) Put(x *BlockedContact) {
	p.pool.Put(x)
}

var PoolBlockedContact = poolBlockedContact{}

func ResultBlockedContact(out *MessageEnvelope, res *BlockedContact) {
	out.Constructor = C_BlockedContact
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsImported int64 = 2157298354

type poolContactsImported struct {
	pool sync.Pool
}

func (p *poolContactsImported) Get() *ContactsImported {
	x, ok := p.pool.Get().(*ContactsImported)
	if !ok {
		return &ContactsImported{}
	}
	return x
}

func (p *poolContactsImported) Put(x *ContactsImported) {
	x.ContactUsers = x.ContactUsers[:0]
	x.Users = x.Users[:0]
	x.Empty = false
	p.pool.Put(x)
}

var PoolContactsImported = poolContactsImported{}

func ResultContactsImported(out *MessageEnvelope, res *ContactsImported) {
	out.Constructor = C_ContactsImported
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactsMany int64 = 3883395672

type poolContactsMany struct {
	pool sync.Pool
}

func (p *poolContactsMany) Get() *ContactsMany {
	x, ok := p.pool.Get().(*ContactsMany)
	if !ok {
		return &ContactsMany{}
	}
	return x
}

func (p *poolContactsMany) Put(x *ContactsMany) {
	x.Contacts = x.Contacts[:0]
	x.ContactUsers = x.ContactUsers[:0]
	x.Users = x.Users[:0]
	x.Empty = false
	x.Hash = 0
	p.pool.Put(x)
}

var PoolContactsMany = poolContactsMany{}

func ResultContactsMany(out *MessageEnvelope, res *ContactsMany) {
	out.Constructor = C_ContactsMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

func init() {
	ConstructorNames[3473528730] = "ContactsImport"
	ConstructorNames[1410714478] = "ContactsAdd"
	ConstructorNames[1412732665] = "ContactsGet"
	ConstructorNames[1750426880] = "ContactsDelete"
	ConstructorNames[14552524] = "ContactsDeleteAll"
	ConstructorNames[2900371089] = "ContactsBlock"
	ConstructorNames[662011773] = "ContactsUnblock"
	ConstructorNames[1073733371] = "ContactsGetBlocked"
	ConstructorNames[3870802464] = "ContactsSearch"
	ConstructorNames[1378126220] = "ContactsGetTopPeers"
	ConstructorNames[1114887378] = "ContactsResetTopPeer"
	ConstructorNames[2243919622] = "ContactsTopPeers"
	ConstructorNames[1763100161] = "TopPeer"
	ConstructorNames[2067026404] = "BlockedContactsMany"
	ConstructorNames[53788553] = "BlockedContact"
	ConstructorNames[2157298354] = "ContactsImported"
	ConstructorNames[3883395672] = "ContactsMany"
}
