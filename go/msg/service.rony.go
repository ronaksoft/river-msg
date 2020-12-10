package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_ServiceSendMessage int64 = 1634092712

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
	registry.RegisterConstructor(1634092712, "msg.ServiceSendMessage")
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

func (x *ServiceSendMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ServiceSendMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}
