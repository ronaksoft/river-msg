package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_ClientUpdatePendingMessageDelivery int64 = 3828722061

type poolClientUpdatePendingMessageDelivery struct {
	pool sync.Pool
}

func (p *poolClientUpdatePendingMessageDelivery) Get() *ClientUpdatePendingMessageDelivery {
	x, ok := p.pool.Get().(*ClientUpdatePendingMessageDelivery)
	if !ok {
		return &ClientUpdatePendingMessageDelivery{}
	}
	return x
}

func (p *poolClientUpdatePendingMessageDelivery) Put(x *ClientUpdatePendingMessageDelivery) {
	if x.Messages != nil {
		PoolUserMessage.Put(x.Messages)
		x.Messages = nil
	}
	if x.PendingMessage != nil {
		PoolClientPendingMessage.Put(x.PendingMessage)
		x.PendingMessage = nil
	}
	x.Success = false
	p.pool.Put(x)
}

var PoolClientUpdatePendingMessageDelivery = poolClientUpdatePendingMessageDelivery{}

const C_ClientUpdateMessagesDeleted int64 = 3060926862

type poolClientUpdateMessagesDeleted struct {
	pool sync.Pool
}

func (p *poolClientUpdateMessagesDeleted) Get() *ClientUpdateMessagesDeleted {
	x, ok := p.pool.Get().(*ClientUpdateMessagesDeleted)
	if !ok {
		return &ClientUpdateMessagesDeleted{}
	}
	return x
}

func (p *poolClientUpdateMessagesDeleted) Put(x *ClientUpdateMessagesDeleted) {
	x.PeerID = 0
	x.PeerType = 0
	x.MessageIDs = x.MessageIDs[:0]
	p.pool.Put(x)
}

var PoolClientUpdateMessagesDeleted = poolClientUpdateMessagesDeleted{}

const C_ClientUpdateSynced int64 = 4244270269

type poolClientUpdateSynced struct {
	pool sync.Pool
}

func (p *poolClientUpdateSynced) Get() *ClientUpdateSynced {
	x, ok := p.pool.Get().(*ClientUpdateSynced)
	if !ok {
		return &ClientUpdateSynced{}
	}
	return x
}

func (p *poolClientUpdateSynced) Put(x *ClientUpdateSynced) {
	x.Dialogs = false
	x.Contacts = false
	x.Gifs = false
	p.pool.Put(x)
}

var PoolClientUpdateSynced = poolClientUpdateSynced{}

func init() {
	registry.RegisterConstructor(3828722061, "ClientUpdatePendingMessageDelivery")
	registry.RegisterConstructor(3060926862, "ClientUpdateMessagesDeleted")
	registry.RegisterConstructor(4244270269, "ClientUpdateSynced")
}

func (x *ClientUpdatePendingMessageDelivery) DeepCopy(z *ClientUpdatePendingMessageDelivery) {
	if x.Messages != nil {
		z.Messages = PoolUserMessage.Get()
		x.Messages.DeepCopy(z.Messages)
	}
	if x.PendingMessage != nil {
		z.PendingMessage = PoolClientPendingMessage.Get()
		x.PendingMessage.DeepCopy(z.PendingMessage)
	}
	z.Success = x.Success
}

func (x *ClientUpdateMessagesDeleted) DeepCopy(z *ClientUpdateMessagesDeleted) {
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *ClientUpdateSynced) DeepCopy(z *ClientUpdateSynced) {
	z.Dialogs = x.Dialogs
	z.Contacts = x.Contacts
	z.Gifs = x.Gifs
}
