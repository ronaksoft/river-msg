package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_ServiceSendMessage int64 = 824547051

type poolServiceSendMessage struct {
	pool sync.Pool
}

func (p *poolServiceSendMessage) Get() *ServiceSendMessage {
	x, ok := p.pool.Get().(*ServiceSendMessage)
	if !ok {
		return &ServiceSendMessage{}
	}
	return x
}

func (p *poolServiceSendMessage) Put(x *ServiceSendMessage) {
	x.OnBehalf = 0
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Body = ""
	x.ReplyTo = 0
	x.ClearDraft = false
	x.Entities = x.Entities[:0]
	p.pool.Put(x)
}

var PoolServiceSendMessage = poolServiceSendMessage{}

func init() {
	registry.RegisterConstructor(824547051, "ServiceSendMessage")
}

func (x *ServiceSendMessage) DeepCopy(z *ServiceSendMessage) {
	z.OnBehalf = x.OnBehalf
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Body = x.Body
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
}

func (x *ServiceSendMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ServiceSendMessage, x)
}
