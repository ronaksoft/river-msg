// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.contacts.proto

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
	x.Contacts = x.Contacts[:0]
	return x
}

func (p *poolContactsImport) Put(x *ContactsImport) {
	p.pool.Put(x)
}

var PoolContactsImport = poolContactsImport{}

func ResultContactsImport(out *MessageEnvelope, res *ContactsImport) {
	out.Constructor = C_ContactsImport
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	x.UserIDs = x.UserIDs[:0]
	return x
}

func (p *poolContactsDelete) Put(x *ContactsDelete) {
	p.pool.Put(x)
}

var PoolContactsDelete = poolContactsDelete{}

func ResultContactsDelete(out *MessageEnvelope, res *ContactsDelete) {
	out.Constructor = C_ContactsDelete
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	x.Contacts = x.Contacts[:0]
	x.Users = x.Users[:0]
	return x
}

func (p *poolBlockedContactsMany) Put(x *BlockedContactsMany) {
	p.pool.Put(x)
}

var PoolBlockedContactsMany = poolBlockedContactsMany{}

func ResultBlockedContactsMany(out *MessageEnvelope, res *BlockedContactsMany) {
	out.Constructor = C_BlockedContactsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	x.ContactUsers = x.ContactUsers[:0]
	x.Users = x.Users[:0]
	return x
}

func (p *poolContactsImported) Put(x *ContactsImported) {
	p.pool.Put(x)
}

var PoolContactsImported = poolContactsImported{}

func ResultContactsImported(out *MessageEnvelope, res *ContactsImported) {
	out.Constructor = C_ContactsImported
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	x.Contacts = x.Contacts[:0]
	x.ContactUsers = x.ContactUsers[:0]
	x.Users = x.Users[:0]
	return x
}

func (p *poolContactsMany) Put(x *ContactsMany) {
	p.pool.Put(x)
}

var PoolContactsMany = poolContactsMany{}

func ResultContactsMany(out *MessageEnvelope, res *ContactsMany) {
	out.Constructor = C_ContactsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[3473528730] = "ContactsImport"
	ConstructorNames[1412732665] = "ContactsGet"
	ConstructorNames[1750426880] = "ContactsDelete"
	ConstructorNames[2900371089] = "ContactsBlock"
	ConstructorNames[662011773] = "ContactsUnblock"
	ConstructorNames[1073733371] = "ContactsGetBlocked"
	ConstructorNames[2067026404] = "BlockedContactsMany"
	ConstructorNames[53788553] = "BlockedContact"
	ConstructorNames[2157298354] = "ContactsImported"
	ConstructorNames[3883395672] = "ContactsMany"
}
