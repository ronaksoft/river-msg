// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: service.proto

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

const C_ServiceSendMessage int64 = 824547051

type poolServiceSendMessage struct {
	pool sync.Pool
}

func (p *poolServiceSendMessage) Get() *ServiceSendMessage {
	x, ok := p.pool.Get().(*ServiceSendMessage)
	if !ok {
		x = &ServiceSendMessage{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolServiceSendMessage) Put(x *ServiceSendMessage) {
	if x == nil {
		return
	}

	x.OnBehalf = 0
	x.RandomID = 0
	PoolInputPeer.Put(x.Peer)
	x.Body = ""
	x.ReplyTo = 0
	x.ClearDraft = false
	for _, z := range x.Entities {
		PoolMessageEntity.Put(z)
	}
	x.Entities = x.Entities[:0]

	p.pool.Put(x)
}

var PoolServiceSendMessage = poolServiceSendMessage{}

func (x *ServiceSendMessage) DeepCopy(z *ServiceSendMessage) {
	z.OnBehalf = x.OnBehalf
	z.RandomID = x.RandomID
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.Body = x.Body
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	for idx := range x.Entities {
		if x.Entities[idx] == nil {
			continue
		}
		xx := PoolMessageEntity.Get()
		x.Entities[idx].DeepCopy(xx)
		z.Entities = append(z.Entities, xx)
	}
}

func (x *ServiceSendMessage) Clone() *ServiceSendMessage {
	z := &ServiceSendMessage{}
	x.DeepCopy(z)
	return z
}

func (x *ServiceSendMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ServiceSendMessage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ServiceSendMessage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ServiceSendMessage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ServiceSendMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ServiceSendMessage, x)
}

func init() {
	registry.RegisterConstructor(824547051, "ServiceSendMessage")
}

var _ = bytes.MinRead
