package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

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
	x.Replace = false
	p.pool.Put(x)
}

var PoolContactsImport = poolContactsImport{}

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
	if x.User != nil {
		PoolInputUser.Put(x.User)
		x.User = nil
	}
	x.FirstName = ""
	x.LastName = ""
	x.Phone = ""
	p.pool.Put(x)
}

var PoolContactsAdd = poolContactsAdd{}

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
	x.Crc32Hash = 0
	p.pool.Put(x)
}

var PoolContactsGet = poolContactsGet{}

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
	if x.User != nil {
		PoolInputUser.Put(x.User)
		x.User = nil
	}
	p.pool.Put(x)
}

var PoolContactsBlock = poolContactsBlock{}

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
	if x.User != nil {
		PoolInputUser.Put(x.User)
		x.User = nil
	}
	p.pool.Put(x)
}

var PoolContactsUnblock = poolContactsUnblock{}

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
	x.Offset = 0
	x.Limit = 0
	p.pool.Put(x)
}

var PoolContactsGetBlocked = poolContactsGetBlocked{}

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
	x.Q = ""
	p.pool.Put(x)
}

var PoolContactsSearch = poolContactsSearch{}

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
	x.Offset = 0
	x.Limit = 0
	x.Category = 0
	p.pool.Put(x)
}

var PoolContactsGetTopPeers = poolContactsGetTopPeers{}

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
	x.Category = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolContactsResetTopPeer = poolContactsResetTopPeer{}

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
	x.Category = 0
	x.Count = 0
	x.Peers = x.Peers[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	p.pool.Put(x)
}

var PoolContactsTopPeers = poolContactsTopPeers{}

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
	if x.Peer != nil {
		PoolPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Rate = 0
	x.LastUpdate = 0
	p.pool.Put(x)
}

var PoolTopPeer = poolTopPeer{}

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
	x.UserID = 0
	x.Date = 0
	p.pool.Put(x)
}

var PoolBlockedContact = poolBlockedContact{}

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
	x.Modified = false
	x.Users = x.Users[:0]
	x.Empty = false
	x.Hash = 0
	p.pool.Put(x)
}

var PoolContactsMany = poolContactsMany{}

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

func (x *ContactsImport) DeepCopy(z *ContactsImport) {
	for idx := range x.Contacts {
		if x.Contacts[idx] != nil {
			xx := PoolPhoneContact.Get()
			x.Contacts[idx].DeepCopy(xx)
			z.Contacts = append(z.Contacts, xx)
		}
	}
	z.Replace = x.Replace
}

func (x *ContactsAdd) DeepCopy(z *ContactsAdd) {
	if x.User != nil {
		z.User = PoolInputUser.Get()
		x.User.DeepCopy(z.User)
	}
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Phone = x.Phone
}

func (x *ContactsGet) DeepCopy(z *ContactsGet) {
	z.Crc32Hash = x.Crc32Hash
}

func (x *ContactsDelete) DeepCopy(z *ContactsDelete) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *ContactsDeleteAll) DeepCopy(z *ContactsDeleteAll) {
}

func (x *ContactsBlock) DeepCopy(z *ContactsBlock) {
	if x.User != nil {
		z.User = PoolInputUser.Get()
		x.User.DeepCopy(z.User)
	}
}

func (x *ContactsUnblock) DeepCopy(z *ContactsUnblock) {
	if x.User != nil {
		z.User = PoolInputUser.Get()
		x.User.DeepCopy(z.User)
	}
}

func (x *ContactsGetBlocked) DeepCopy(z *ContactsGetBlocked) {
	z.Offset = x.Offset
	z.Limit = x.Limit
}

func (x *ContactsSearch) DeepCopy(z *ContactsSearch) {
	z.Q = x.Q
}

func (x *ContactsGetTopPeers) DeepCopy(z *ContactsGetTopPeers) {
	z.Offset = x.Offset
	z.Limit = x.Limit
	z.Category = x.Category
}

func (x *ContactsResetTopPeer) DeepCopy(z *ContactsResetTopPeer) {
	z.Category = x.Category
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *ContactsTopPeers) DeepCopy(z *ContactsTopPeers) {
	z.Category = x.Category
	z.Count = x.Count
	for idx := range x.Peers {
		if x.Peers[idx] != nil {
			xx := PoolTopPeer.Get()
			x.Peers[idx].DeepCopy(xx)
			z.Peers = append(z.Peers, xx)
		}
	}
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	for idx := range x.Groups {
		if x.Groups[idx] != nil {
			xx := PoolGroup.Get()
			x.Groups[idx].DeepCopy(xx)
			z.Groups = append(z.Groups, xx)
		}
	}
}

func (x *TopPeer) DeepCopy(z *TopPeer) {
	z.TeamID = x.TeamID
	if x.Peer != nil {
		z.Peer = PoolPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Rate = x.Rate
	z.LastUpdate = x.LastUpdate
}

func (x *BlockedContactsMany) DeepCopy(z *BlockedContactsMany) {
	for idx := range x.Contacts {
		if x.Contacts[idx] != nil {
			xx := PoolBlockedContact.Get()
			x.Contacts[idx].DeepCopy(xx)
			z.Contacts = append(z.Contacts, xx)
		}
	}
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	z.Total = x.Total
}

func (x *BlockedContact) DeepCopy(z *BlockedContact) {
	z.UserID = x.UserID
	z.Date = x.Date
}

func (x *ContactsImported) DeepCopy(z *ContactsImported) {
	for idx := range x.ContactUsers {
		if x.ContactUsers[idx] != nil {
			xx := PoolContactUser.Get()
			x.ContactUsers[idx].DeepCopy(xx)
			z.ContactUsers = append(z.ContactUsers, xx)
		}
	}
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	z.Empty = x.Empty
}

func (x *ContactsMany) DeepCopy(z *ContactsMany) {
	for idx := range x.Contacts {
		if x.Contacts[idx] != nil {
			xx := PoolPhoneContact.Get()
			x.Contacts[idx].DeepCopy(xx)
			z.Contacts = append(z.Contacts, xx)
		}
	}
	for idx := range x.ContactUsers {
		if x.ContactUsers[idx] != nil {
			xx := PoolContactUser.Get()
			x.ContactUsers[idx].DeepCopy(xx)
			z.ContactUsers = append(z.ContactUsers, xx)
		}
	}
	z.Modified = x.Modified
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	z.Empty = x.Empty
	z.Hash = x.Hash
}

func (x *ContactsImport) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsImport, x)
}

func (x *ContactsAdd) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsAdd, x)
}

func (x *ContactsGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGet, x)
}

func (x *ContactsDelete) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsDelete, x)
}

func (x *ContactsDeleteAll) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsDeleteAll, x)
}

func (x *ContactsBlock) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsBlock, x)
}

func (x *ContactsUnblock) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsUnblock, x)
}

func (x *ContactsGetBlocked) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGetBlocked, x)
}

func (x *ContactsSearch) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsSearch, x)
}

func (x *ContactsGetTopPeers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsGetTopPeers, x)
}

func (x *ContactsResetTopPeer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsResetTopPeer, x)
}

func (x *ContactsTopPeers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsTopPeers, x)
}

func (x *TopPeer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_TopPeer, x)
}

func (x *BlockedContactsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BlockedContactsMany, x)
}

func (x *BlockedContact) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BlockedContact, x)
}

func (x *ContactsImported) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsImported, x)
}

func (x *ContactsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactsMany, x)
}

func (x *ContactsImport) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsAdd) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsGet) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsDelete) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsDeleteAll) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsBlock) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsUnblock) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsGetBlocked) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsSearch) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsGetTopPeers) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsResetTopPeer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsTopPeers) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *TopPeer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BlockedContactsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BlockedContact) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsImported) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactsImport) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsAdd) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsDelete) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsDeleteAll) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsBlock) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsUnblock) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGetBlocked) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsSearch) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsGetTopPeers) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsResetTopPeer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsTopPeers) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *TopPeer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BlockedContactsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BlockedContact) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsImported) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}
