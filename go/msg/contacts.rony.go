// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: contacts.proto

package msg

import (
	bytes "bytes"
	edge "github.com/ronaksoft/rony/edge"
	pools "github.com/ronaksoft/rony/pools"
	registry "github.com/ronaksoft/rony/registry"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

var _ = pools.Imported

const C_ContactsImport int64 = 3473528730

type poolContactsImport struct {
	pool sync.Pool
}

func (p *poolContactsImport) Get() *ContactsImport {
	x, ok := p.pool.Get().(*ContactsImport)
	if !ok {
		x = &ContactsImport{}
	}

	return x
}

func (p *poolContactsImport) Put(x *ContactsImport) {
	if x == nil {
		return
	}

	for _, z := range x.Contacts {
		PoolPhoneContact.Put(z)
	}
	x.Contacts = x.Contacts[:0]
	x.Replace = false

	p.pool.Put(x)
}

var PoolContactsImport = poolContactsImport{}

func (x *ContactsImport) DeepCopy(z *ContactsImport) {
	for idx := range x.Contacts {
		if x.Contacts[idx] == nil {
			continue
		}
		xx := PoolPhoneContact.Get()
		x.Contacts[idx].DeepCopy(xx)
		z.Contacts = append(z.Contacts, xx)
	}
	z.Replace = x.Replace
}

func (x *ContactsImport) Clone() *ContactsImport {
	z := &ContactsImport{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsImport) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsImport) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsImport) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsImport) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsImport) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsImport, x)
}

const C_ContactsAdd int64 = 1410714478

type poolContactsAdd struct {
	pool sync.Pool
}

func (p *poolContactsAdd) Get() *ContactsAdd {
	x, ok := p.pool.Get().(*ContactsAdd)
	if !ok {
		x = &ContactsAdd{}
	}

	x.User = PoolInputUser.Get()

	return x
}

func (p *poolContactsAdd) Put(x *ContactsAdd) {
	if x == nil {
		return
	}

	PoolInputUser.Put(x.User)
	x.FirstName = ""
	x.LastName = ""
	x.Phone = ""

	p.pool.Put(x)
}

var PoolContactsAdd = poolContactsAdd{}

func (x *ContactsAdd) DeepCopy(z *ContactsAdd) {
	if x.User != nil {
		if z.User == nil {
			z.User = PoolInputUser.Get()
		}
		x.User.DeepCopy(z.User)
	} else {
		PoolInputUser.Put(z.User)
		z.User = nil
	}
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Phone = x.Phone
}

func (x *ContactsAdd) Clone() *ContactsAdd {
	z := &ContactsAdd{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsAdd) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsAdd) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsAdd) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsAdd) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsAdd) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsAdd, x)
}

const C_ContactsGet int64 = 1412732665

type poolContactsGet struct {
	pool sync.Pool
}

func (p *poolContactsGet) Get() *ContactsGet {
	x, ok := p.pool.Get().(*ContactsGet)
	if !ok {
		x = &ContactsGet{}
	}

	return x
}

func (p *poolContactsGet) Put(x *ContactsGet) {
	if x == nil {
		return
	}

	x.Crc32Hash = 0

	p.pool.Put(x)
}

var PoolContactsGet = poolContactsGet{}

func (x *ContactsGet) DeepCopy(z *ContactsGet) {
	z.Crc32Hash = x.Crc32Hash
}

func (x *ContactsGet) Clone() *ContactsGet {
	z := &ContactsGet{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGet) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsGet) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsGet) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGet, x)
}

const C_ContactsDelete int64 = 1750426880

type poolContactsDelete struct {
	pool sync.Pool
}

func (p *poolContactsDelete) Get() *ContactsDelete {
	x, ok := p.pool.Get().(*ContactsDelete)
	if !ok {
		x = &ContactsDelete{}
	}

	return x
}

func (p *poolContactsDelete) Put(x *ContactsDelete) {
	if x == nil {
		return
	}

	x.UserIDs = x.UserIDs[:0]

	p.pool.Put(x)
}

var PoolContactsDelete = poolContactsDelete{}

func (x *ContactsDelete) DeepCopy(z *ContactsDelete) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *ContactsDelete) Clone() *ContactsDelete {
	z := &ContactsDelete{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsDelete) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsDelete) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsDelete) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsDelete) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsDelete) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsDelete, x)
}

const C_ContactsDeleteAll int64 = 14552524

type poolContactsDeleteAll struct {
	pool sync.Pool
}

func (p *poolContactsDeleteAll) Get() *ContactsDeleteAll {
	x, ok := p.pool.Get().(*ContactsDeleteAll)
	if !ok {
		x = &ContactsDeleteAll{}
	}

	return x
}

func (p *poolContactsDeleteAll) Put(x *ContactsDeleteAll) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolContactsDeleteAll = poolContactsDeleteAll{}

func (x *ContactsDeleteAll) DeepCopy(z *ContactsDeleteAll) {
}

func (x *ContactsDeleteAll) Clone() *ContactsDeleteAll {
	z := &ContactsDeleteAll{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsDeleteAll) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsDeleteAll) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsDeleteAll) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsDeleteAll) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsDeleteAll) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsDeleteAll, x)
}

const C_ContactsBlock int64 = 2900371089

type poolContactsBlock struct {
	pool sync.Pool
}

func (p *poolContactsBlock) Get() *ContactsBlock {
	x, ok := p.pool.Get().(*ContactsBlock)
	if !ok {
		x = &ContactsBlock{}
	}

	x.User = PoolInputUser.Get()

	return x
}

func (p *poolContactsBlock) Put(x *ContactsBlock) {
	if x == nil {
		return
	}

	PoolInputUser.Put(x.User)

	p.pool.Put(x)
}

var PoolContactsBlock = poolContactsBlock{}

func (x *ContactsBlock) DeepCopy(z *ContactsBlock) {
	if x.User != nil {
		if z.User == nil {
			z.User = PoolInputUser.Get()
		}
		x.User.DeepCopy(z.User)
	} else {
		PoolInputUser.Put(z.User)
		z.User = nil
	}
}

func (x *ContactsBlock) Clone() *ContactsBlock {
	z := &ContactsBlock{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsBlock) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsBlock) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsBlock) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsBlock) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsBlock) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsBlock, x)
}

const C_ContactsUnblock int64 = 662011773

type poolContactsUnblock struct {
	pool sync.Pool
}

func (p *poolContactsUnblock) Get() *ContactsUnblock {
	x, ok := p.pool.Get().(*ContactsUnblock)
	if !ok {
		x = &ContactsUnblock{}
	}

	x.User = PoolInputUser.Get()

	return x
}

func (p *poolContactsUnblock) Put(x *ContactsUnblock) {
	if x == nil {
		return
	}

	PoolInputUser.Put(x.User)

	p.pool.Put(x)
}

var PoolContactsUnblock = poolContactsUnblock{}

func (x *ContactsUnblock) DeepCopy(z *ContactsUnblock) {
	if x.User != nil {
		if z.User == nil {
			z.User = PoolInputUser.Get()
		}
		x.User.DeepCopy(z.User)
	} else {
		PoolInputUser.Put(z.User)
		z.User = nil
	}
}

func (x *ContactsUnblock) Clone() *ContactsUnblock {
	z := &ContactsUnblock{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsUnblock) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsUnblock) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsUnblock) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsUnblock) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsUnblock) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsUnblock, x)
}

const C_ContactsGetBlocked int64 = 1073733371

type poolContactsGetBlocked struct {
	pool sync.Pool
}

func (p *poolContactsGetBlocked) Get() *ContactsGetBlocked {
	x, ok := p.pool.Get().(*ContactsGetBlocked)
	if !ok {
		x = &ContactsGetBlocked{}
	}

	return x
}

func (p *poolContactsGetBlocked) Put(x *ContactsGetBlocked) {
	if x == nil {
		return
	}

	x.Offset = 0
	x.Limit = 0

	p.pool.Put(x)
}

var PoolContactsGetBlocked = poolContactsGetBlocked{}

func (x *ContactsGetBlocked) DeepCopy(z *ContactsGetBlocked) {
	z.Offset = x.Offset
	z.Limit = x.Limit
}

func (x *ContactsGetBlocked) Clone() *ContactsGetBlocked {
	z := &ContactsGetBlocked{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsGetBlocked) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGetBlocked) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsGetBlocked) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsGetBlocked) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsGetBlocked) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGetBlocked, x)
}

const C_ContactsSearch int64 = 3870802464

type poolContactsSearch struct {
	pool sync.Pool
}

func (p *poolContactsSearch) Get() *ContactsSearch {
	x, ok := p.pool.Get().(*ContactsSearch)
	if !ok {
		x = &ContactsSearch{}
	}

	return x
}

func (p *poolContactsSearch) Put(x *ContactsSearch) {
	if x == nil {
		return
	}

	x.Q = ""

	p.pool.Put(x)
}

var PoolContactsSearch = poolContactsSearch{}

func (x *ContactsSearch) DeepCopy(z *ContactsSearch) {
	z.Q = x.Q
}

func (x *ContactsSearch) Clone() *ContactsSearch {
	z := &ContactsSearch{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsSearch) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsSearch) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsSearch) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsSearch) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsSearch) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsSearch, x)
}

const C_ContactsGetTopPeers int64 = 1378126220

type poolContactsGetTopPeers struct {
	pool sync.Pool
}

func (p *poolContactsGetTopPeers) Get() *ContactsGetTopPeers {
	x, ok := p.pool.Get().(*ContactsGetTopPeers)
	if !ok {
		x = &ContactsGetTopPeers{}
	}

	return x
}

func (p *poolContactsGetTopPeers) Put(x *ContactsGetTopPeers) {
	if x == nil {
		return
	}

	x.Offset = 0
	x.Limit = 0
	x.Category = 0

	p.pool.Put(x)
}

var PoolContactsGetTopPeers = poolContactsGetTopPeers{}

func (x *ContactsGetTopPeers) DeepCopy(z *ContactsGetTopPeers) {
	z.Offset = x.Offset
	z.Limit = x.Limit
	z.Category = x.Category
}

func (x *ContactsGetTopPeers) Clone() *ContactsGetTopPeers {
	z := &ContactsGetTopPeers{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsGetTopPeers) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGetTopPeers) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsGetTopPeers) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsGetTopPeers) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsGetTopPeers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGetTopPeers, x)
}

const C_ContactsResetTopPeer int64 = 1114887378

type poolContactsResetTopPeer struct {
	pool sync.Pool
}

func (p *poolContactsResetTopPeer) Get() *ContactsResetTopPeer {
	x, ok := p.pool.Get().(*ContactsResetTopPeer)
	if !ok {
		x = &ContactsResetTopPeer{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolContactsResetTopPeer) Put(x *ContactsResetTopPeer) {
	if x == nil {
		return
	}

	x.Category = 0
	PoolInputPeer.Put(x.Peer)

	p.pool.Put(x)
}

var PoolContactsResetTopPeer = poolContactsResetTopPeer{}

func (x *ContactsResetTopPeer) DeepCopy(z *ContactsResetTopPeer) {
	z.Category = x.Category
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
}

func (x *ContactsResetTopPeer) Clone() *ContactsResetTopPeer {
	z := &ContactsResetTopPeer{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsResetTopPeer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsResetTopPeer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsResetTopPeer) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsResetTopPeer) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsResetTopPeer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsResetTopPeer, x)
}

const C_ContactsTopPeers int64 = 2243919622

type poolContactsTopPeers struct {
	pool sync.Pool
}

func (p *poolContactsTopPeers) Get() *ContactsTopPeers {
	x, ok := p.pool.Get().(*ContactsTopPeers)
	if !ok {
		x = &ContactsTopPeers{}
	}

	return x
}

func (p *poolContactsTopPeers) Put(x *ContactsTopPeers) {
	if x == nil {
		return
	}

	x.Category = 0
	x.Count = 0
	for _, z := range x.Peers {
		PoolTopPeer.Put(z)
	}
	x.Peers = x.Peers[:0]
	for _, z := range x.Users {
		PoolUser.Put(z)
	}
	x.Users = x.Users[:0]
	for _, z := range x.Groups {
		PoolGroup.Put(z)
	}
	x.Groups = x.Groups[:0]

	p.pool.Put(x)
}

var PoolContactsTopPeers = poolContactsTopPeers{}

func (x *ContactsTopPeers) DeepCopy(z *ContactsTopPeers) {
	z.Category = x.Category
	z.Count = x.Count
	for idx := range x.Peers {
		if x.Peers[idx] == nil {
			continue
		}
		xx := PoolTopPeer.Get()
		x.Peers[idx].DeepCopy(xx)
		z.Peers = append(z.Peers, xx)
	}
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
	for idx := range x.Groups {
		if x.Groups[idx] == nil {
			continue
		}
		xx := PoolGroup.Get()
		x.Groups[idx].DeepCopy(xx)
		z.Groups = append(z.Groups, xx)
	}
}

func (x *ContactsTopPeers) Clone() *ContactsTopPeers {
	z := &ContactsTopPeers{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsTopPeers) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsTopPeers) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsTopPeers) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsTopPeers) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsTopPeers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsTopPeers, x)
}

const C_TopPeer int64 = 1763100161

type poolTopPeer struct {
	pool sync.Pool
}

func (p *poolTopPeer) Get() *TopPeer {
	x, ok := p.pool.Get().(*TopPeer)
	if !ok {
		x = &TopPeer{}
	}

	x.Peer = PoolPeer.Get()

	return x
}

func (p *poolTopPeer) Put(x *TopPeer) {
	if x == nil {
		return
	}

	x.TeamID = 0
	PoolPeer.Put(x.Peer)
	x.Rate = 0
	x.LastUpdate = 0

	p.pool.Put(x)
}

var PoolTopPeer = poolTopPeer{}

func (x *TopPeer) DeepCopy(z *TopPeer) {
	z.TeamID = x.TeamID
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.Rate = x.Rate
	z.LastUpdate = x.LastUpdate
}

func (x *TopPeer) Clone() *TopPeer {
	z := &TopPeer{}
	x.DeepCopy(z)
	return z
}

func (x *TopPeer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *TopPeer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *TopPeer) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *TopPeer) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *TopPeer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_TopPeer, x)
}

const C_BlockedContactsMany int64 = 2067026404

type poolBlockedContactsMany struct {
	pool sync.Pool
}

func (p *poolBlockedContactsMany) Get() *BlockedContactsMany {
	x, ok := p.pool.Get().(*BlockedContactsMany)
	if !ok {
		x = &BlockedContactsMany{}
	}

	return x
}

func (p *poolBlockedContactsMany) Put(x *BlockedContactsMany) {
	if x == nil {
		return
	}

	for _, z := range x.Contacts {
		PoolBlockedContact.Put(z)
	}
	x.Contacts = x.Contacts[:0]
	for _, z := range x.Users {
		PoolUser.Put(z)
	}
	x.Users = x.Users[:0]
	x.Total = 0

	p.pool.Put(x)
}

var PoolBlockedContactsMany = poolBlockedContactsMany{}

func (x *BlockedContactsMany) DeepCopy(z *BlockedContactsMany) {
	for idx := range x.Contacts {
		if x.Contacts[idx] == nil {
			continue
		}
		xx := PoolBlockedContact.Get()
		x.Contacts[idx].DeepCopy(xx)
		z.Contacts = append(z.Contacts, xx)
	}
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
	z.Total = x.Total
}

func (x *BlockedContactsMany) Clone() *BlockedContactsMany {
	z := &BlockedContactsMany{}
	x.DeepCopy(z)
	return z
}

func (x *BlockedContactsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BlockedContactsMany) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *BlockedContactsMany) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *BlockedContactsMany) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *BlockedContactsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BlockedContactsMany, x)
}

const C_BlockedContact int64 = 53788553

type poolBlockedContact struct {
	pool sync.Pool
}

func (p *poolBlockedContact) Get() *BlockedContact {
	x, ok := p.pool.Get().(*BlockedContact)
	if !ok {
		x = &BlockedContact{}
	}

	return x
}

func (p *poolBlockedContact) Put(x *BlockedContact) {
	if x == nil {
		return
	}

	x.UserID = 0
	x.Date = 0

	p.pool.Put(x)
}

var PoolBlockedContact = poolBlockedContact{}

func (x *BlockedContact) DeepCopy(z *BlockedContact) {
	z.UserID = x.UserID
	z.Date = x.Date
}

func (x *BlockedContact) Clone() *BlockedContact {
	z := &BlockedContact{}
	x.DeepCopy(z)
	return z
}

func (x *BlockedContact) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BlockedContact) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *BlockedContact) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *BlockedContact) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *BlockedContact) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BlockedContact, x)
}

const C_ContactsImported int64 = 2157298354

type poolContactsImported struct {
	pool sync.Pool
}

func (p *poolContactsImported) Get() *ContactsImported {
	x, ok := p.pool.Get().(*ContactsImported)
	if !ok {
		x = &ContactsImported{}
	}

	return x
}

func (p *poolContactsImported) Put(x *ContactsImported) {
	if x == nil {
		return
	}

	for _, z := range x.ContactUsers {
		PoolContactUser.Put(z)
	}
	x.ContactUsers = x.ContactUsers[:0]
	for _, z := range x.Users {
		PoolUser.Put(z)
	}
	x.Users = x.Users[:0]
	x.Empty = false

	p.pool.Put(x)
}

var PoolContactsImported = poolContactsImported{}

func (x *ContactsImported) DeepCopy(z *ContactsImported) {
	for idx := range x.ContactUsers {
		if x.ContactUsers[idx] == nil {
			continue
		}
		xx := PoolContactUser.Get()
		x.ContactUsers[idx].DeepCopy(xx)
		z.ContactUsers = append(z.ContactUsers, xx)
	}
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
	z.Empty = x.Empty
}

func (x *ContactsImported) Clone() *ContactsImported {
	z := &ContactsImported{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsImported) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsImported) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsImported) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsImported) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsImported) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsImported, x)
}

const C_ContactsMany int64 = 3883395672

type poolContactsMany struct {
	pool sync.Pool
}

func (p *poolContactsMany) Get() *ContactsMany {
	x, ok := p.pool.Get().(*ContactsMany)
	if !ok {
		x = &ContactsMany{}
	}

	return x
}

func (p *poolContactsMany) Put(x *ContactsMany) {
	if x == nil {
		return
	}

	for _, z := range x.Contacts {
		PoolPhoneContact.Put(z)
	}
	x.Contacts = x.Contacts[:0]
	for _, z := range x.ContactUsers {
		PoolContactUser.Put(z)
	}
	x.ContactUsers = x.ContactUsers[:0]
	x.Modified = false
	for _, z := range x.Users {
		PoolUser.Put(z)
	}
	x.Users = x.Users[:0]
	x.Empty = false
	x.Hash = 0

	p.pool.Put(x)
}

var PoolContactsMany = poolContactsMany{}

func (x *ContactsMany) DeepCopy(z *ContactsMany) {
	for idx := range x.Contacts {
		if x.Contacts[idx] == nil {
			continue
		}
		xx := PoolPhoneContact.Get()
		x.Contacts[idx].DeepCopy(xx)
		z.Contacts = append(z.Contacts, xx)
	}
	for idx := range x.ContactUsers {
		if x.ContactUsers[idx] == nil {
			continue
		}
		xx := PoolContactUser.Get()
		x.ContactUsers[idx].DeepCopy(xx)
		z.ContactUsers = append(z.ContactUsers, xx)
	}
	z.Modified = x.Modified
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
	z.Empty = x.Empty
	z.Hash = x.Hash
}

func (x *ContactsMany) Clone() *ContactsMany {
	z := &ContactsMany{}
	x.DeepCopy(z)
	return z
}

func (x *ContactsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsMany) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ContactsMany) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ContactsMany) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ContactsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsMany, x)
}

func init() {
	registry.RegisterConstructor(3473528730, "ContactsImport")
	registry.RegisterConstructor(1410714478, "ContactsAdd")
	registry.RegisterConstructor(1412732665, "ContactsGet")
	registry.RegisterConstructor(1750426880, "ContactsDelete")
	registry.RegisterConstructor(14552524, "ContactsDeleteAll")
	registry.RegisterConstructor(2900371089, "ContactsBlock")
	registry.RegisterConstructor(662011773, "ContactsUnblock")
	registry.RegisterConstructor(1073733371, "ContactsGetBlocked")
	registry.RegisterConstructor(3870802464, "ContactsSearch")
	registry.RegisterConstructor(1378126220, "ContactsGetTopPeers")
	registry.RegisterConstructor(1114887378, "ContactsResetTopPeer")
	registry.RegisterConstructor(2243919622, "ContactsTopPeers")
	registry.RegisterConstructor(1763100161, "TopPeer")
	registry.RegisterConstructor(2067026404, "BlockedContactsMany")
	registry.RegisterConstructor(53788553, "BlockedContact")
	registry.RegisterConstructor(2157298354, "ContactsImported")
	registry.RegisterConstructor(3883395672, "ContactsMany")
}

var _ = bytes.MinRead
