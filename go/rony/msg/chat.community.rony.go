package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_CommunitySendMessage int64 = 3506778488

type poolCommunitySendMessage struct {
	pool sync.Pool
}

func (p *poolCommunitySendMessage) Get() *CommunitySendMessage {
	x, ok := p.pool.Get().(*CommunitySendMessage)
	if !ok {
		return &CommunitySendMessage{}
	}
	return x
}

func (p *poolCommunitySendMessage) Put(x *CommunitySendMessage) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	x.SenderID = 0
	x.SenderMsgID = 0
	p.pool.Put(x)
}

var PoolCommunitySendMessage = poolCommunitySendMessage{}

const C_CommunitySendMedia int64 = 2436824148

type poolCommunitySendMedia struct {
	pool sync.Pool
}

func (p *poolCommunitySendMedia) Get() *CommunitySendMedia {
	x, ok := p.pool.Get().(*CommunitySendMedia)
	if !ok {
		return &CommunitySendMedia{}
	}
	return x
}

func (p *poolCommunitySendMedia) Put(x *CommunitySendMedia) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MediaType = 0
	x.MediaData = x.MediaData[:0]
	x.ReplyTo = 0
	x.ClearDraft = false
	x.SenderID = 0
	x.SenderMsgID = 0
	p.pool.Put(x)
}

var PoolCommunitySendMedia = poolCommunitySendMedia{}

const C_CommunitySetTyping int64 = 3413516595

type poolCommunitySetTyping struct {
	pool sync.Pool
}

func (p *poolCommunitySetTyping) Get() *CommunitySetTyping {
	x, ok := p.pool.Get().(*CommunitySetTyping)
	if !ok {
		return &CommunitySetTyping{}
	}
	return x
}

func (p *poolCommunitySetTyping) Put(x *CommunitySetTyping) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Action = 0
	x.SenderID = 0
	p.pool.Put(x)
}

var PoolCommunitySetTyping = poolCommunitySetTyping{}

const C_CommunityGetUpdates int64 = 2021391963

type poolCommunityGetUpdates struct {
	pool sync.Pool
}

func (p *poolCommunityGetUpdates) Get() *CommunityGetUpdates {
	x, ok := p.pool.Get().(*CommunityGetUpdates)
	if !ok {
		return &CommunityGetUpdates{}
	}
	return x
}

func (p *poolCommunityGetUpdates) Put(x *CommunityGetUpdates) {
	x.WaitAfterInMS = 0
	x.WaitMaxInMS = 0
	x.SizeLimit = 0
	x.OffsetID = 0
	p.pool.Put(x)
}

var PoolCommunityGetUpdates = poolCommunityGetUpdates{}

const C_CommunityGetMembers int64 = 2022915988

type poolCommunityGetMembers struct {
	pool sync.Pool
}

func (p *poolCommunityGetMembers) Get() *CommunityGetMembers {
	x, ok := p.pool.Get().(*CommunityGetMembers)
	if !ok {
		return &CommunityGetMembers{}
	}
	return x
}

func (p *poolCommunityGetMembers) Put(x *CommunityGetMembers) {
	x.Offset = 0
	x.Limit = 0
	p.pool.Put(x)
}

var PoolCommunityGetMembers = poolCommunityGetMembers{}

const C_CommunityRecall int64 = 890349574

type poolCommunityRecall struct {
	pool sync.Pool
}

func (p *poolCommunityRecall) Get() *CommunityRecall {
	x, ok := p.pool.Get().(*CommunityRecall)
	if !ok {
		return &CommunityRecall{}
	}
	return x
}

func (p *poolCommunityRecall) Put(x *CommunityRecall) {
	x.TeamID = 0
	x.AccessKey = x.AccessKey[:0]
	p.pool.Put(x)
}

var PoolCommunityRecall = poolCommunityRecall{}

const C_CommunityAuthorizeUser int64 = 1452766231

type poolCommunityAuthorizeUser struct {
	pool sync.Pool
}

func (p *poolCommunityAuthorizeUser) Get() *CommunityAuthorizeUser {
	x, ok := p.pool.Get().(*CommunityAuthorizeUser)
	if !ok {
		return &CommunityAuthorizeUser{}
	}
	return x
}

func (p *poolCommunityAuthorizeUser) Put(x *CommunityAuthorizeUser) {
	x.Phone = ""
	x.FirstName = ""
	x.LastName = ""
	x.Provider = ""
	p.pool.Put(x)
}

var PoolCommunityAuthorizeUser = poolCommunityAuthorizeUser{}

const C_CommunityUser int64 = 3812837958

type poolCommunityUser struct {
	pool sync.Pool
}

func (p *poolCommunityUser) Get() *CommunityUser {
	x, ok := p.pool.Get().(*CommunityUser)
	if !ok {
		return &CommunityUser{}
	}
	return x
}

func (p *poolCommunityUser) Put(x *CommunityUser) {
	x.UserID = 0
	x.FirstName = ""
	x.LastName = ""
	x.Phone = ""
	p.pool.Put(x)
}

var PoolCommunityUser = poolCommunityUser{}

const C_CommunityUpdateEnvelope int64 = 1076119993

type poolCommunityUpdateEnvelope struct {
	pool sync.Pool
}

func (p *poolCommunityUpdateEnvelope) Get() *CommunityUpdateEnvelope {
	x, ok := p.pool.Get().(*CommunityUpdateEnvelope)
	if !ok {
		return &CommunityUpdateEnvelope{}
	}
	return x
}

func (p *poolCommunityUpdateEnvelope) Put(x *CommunityUpdateEnvelope) {
	x.OffsetID = 0
	x.PartitionID = 0
	x.Constructor = 0
	x.Update = x.Update[:0]
	p.pool.Put(x)
}

var PoolCommunityUpdateEnvelope = poolCommunityUpdateEnvelope{}

const C_CommunityUpdateContainer int64 = 918339432

type poolCommunityUpdateContainer struct {
	pool sync.Pool
}

func (p *poolCommunityUpdateContainer) Get() *CommunityUpdateContainer {
	x, ok := p.pool.Get().(*CommunityUpdateContainer)
	if !ok {
		return &CommunityUpdateContainer{}
	}
	return x
}

func (p *poolCommunityUpdateContainer) Put(x *CommunityUpdateContainer) {
	x.Updates = x.Updates[:0]
	x.Empty = false
	p.pool.Put(x)
}

var PoolCommunityUpdateContainer = poolCommunityUpdateContainer{}

func init() {
	registry.RegisterConstructor(3506778488, "CommunitySendMessage")
	registry.RegisterConstructor(2436824148, "CommunitySendMedia")
	registry.RegisterConstructor(3413516595, "CommunitySetTyping")
	registry.RegisterConstructor(2021391963, "CommunityGetUpdates")
	registry.RegisterConstructor(2022915988, "CommunityGetMembers")
	registry.RegisterConstructor(890349574, "CommunityRecall")
	registry.RegisterConstructor(1452766231, "CommunityAuthorizeUser")
	registry.RegisterConstructor(3812837958, "CommunityUser")
	registry.RegisterConstructor(1076119993, "CommunityUpdateEnvelope")
	registry.RegisterConstructor(918339432, "CommunityUpdateContainer")
}

func (x *CommunitySendMessage) DeepCopy(z *CommunitySendMessage) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Body = x.Body
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
	z.SenderID = x.SenderID
	z.SenderMsgID = x.SenderMsgID
}

func (x *CommunitySendMedia) DeepCopy(z *CommunitySendMedia) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MediaType = x.MediaType
	z.MediaData = append(z.MediaData[:0], x.MediaData...)
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	z.SenderID = x.SenderID
	z.SenderMsgID = x.SenderMsgID
}

func (x *CommunitySetTyping) DeepCopy(z *CommunitySetTyping) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Action = x.Action
	z.SenderID = x.SenderID
}

func (x *CommunityGetUpdates) DeepCopy(z *CommunityGetUpdates) {
	z.WaitAfterInMS = x.WaitAfterInMS
	z.WaitMaxInMS = x.WaitMaxInMS
	z.SizeLimit = x.SizeLimit
	z.OffsetID = x.OffsetID
}

func (x *CommunityGetMembers) DeepCopy(z *CommunityGetMembers) {
	z.Offset = x.Offset
	z.Limit = x.Limit
}

func (x *CommunityRecall) DeepCopy(z *CommunityRecall) {
	z.TeamID = x.TeamID
	z.AccessKey = append(z.AccessKey[:0], x.AccessKey...)
}

func (x *CommunityAuthorizeUser) DeepCopy(z *CommunityAuthorizeUser) {
	z.Phone = x.Phone
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Provider = x.Provider
}

func (x *CommunityUser) DeepCopy(z *CommunityUser) {
	z.UserID = x.UserID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Phone = x.Phone
}

func (x *CommunityUpdateEnvelope) DeepCopy(z *CommunityUpdateEnvelope) {
	z.OffsetID = x.OffsetID
	z.PartitionID = x.PartitionID
	z.Constructor = x.Constructor
	z.Update = append(z.Update[:0], x.Update...)
}

func (x *CommunityUpdateContainer) DeepCopy(z *CommunityUpdateContainer) {
	for idx := range x.Updates {
		if x.Updates[idx] != nil {
			xx := PoolCommunityUpdateEnvelope.Get()
			x.Updates[idx].DeepCopy(xx)
			z.Updates = append(z.Updates, xx)
		}
	}
	z.Empty = x.Empty
}

func (x *CommunitySendMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySendMessage, x)
}

func (x *CommunitySendMedia) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySendMedia, x)
}

func (x *CommunitySetTyping) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySetTyping, x)
}

func (x *CommunityGetUpdates) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityGetUpdates, x)
}

func (x *CommunityGetMembers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityGetMembers, x)
}

func (x *CommunityRecall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityRecall, x)
}

func (x *CommunityAuthorizeUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityAuthorizeUser, x)
}

func (x *CommunityUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUser, x)
}

func (x *CommunityUpdateEnvelope) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUpdateEnvelope, x)
}

func (x *CommunityUpdateContainer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUpdateContainer, x)
}
