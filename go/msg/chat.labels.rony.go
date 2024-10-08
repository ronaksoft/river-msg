// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: chat.labels.proto

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

const C_LabelsCreate int64 = 2138857068

type poolLabelsCreate struct {
	pool sync.Pool
}

func (p *poolLabelsCreate) Get() *LabelsCreate {
	x, ok := p.pool.Get().(*LabelsCreate)
	if !ok {
		x = &LabelsCreate{}
	}

	return x
}

func (p *poolLabelsCreate) Put(x *LabelsCreate) {
	if x == nil {
		return
	}

	x.RandomID = 0
	x.Name = ""
	x.Colour = ""

	p.pool.Put(x)
}

var PoolLabelsCreate = poolLabelsCreate{}

func (x *LabelsCreate) DeepCopy(z *LabelsCreate) {
	z.RandomID = x.RandomID
	z.Name = x.Name
	z.Colour = x.Colour
}

func (x *LabelsCreate) Clone() *LabelsCreate {
	z := &LabelsCreate{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsCreate) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsCreate) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsCreate) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsCreate) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsCreate) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsCreate, x)
}

const C_LabelsEdit int64 = 2790466877

type poolLabelsEdit struct {
	pool sync.Pool
}

func (p *poolLabelsEdit) Get() *LabelsEdit {
	x, ok := p.pool.Get().(*LabelsEdit)
	if !ok {
		x = &LabelsEdit{}
	}

	return x
}

func (p *poolLabelsEdit) Put(x *LabelsEdit) {
	if x == nil {
		return
	}

	x.LabelID = 0
	x.Name = ""
	x.Colour = ""

	p.pool.Put(x)
}

var PoolLabelsEdit = poolLabelsEdit{}

func (x *LabelsEdit) DeepCopy(z *LabelsEdit) {
	z.LabelID = x.LabelID
	z.Name = x.Name
	z.Colour = x.Colour
}

func (x *LabelsEdit) Clone() *LabelsEdit {
	z := &LabelsEdit{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsEdit) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsEdit) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsEdit) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsEdit) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsEdit) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsEdit, x)
}

const C_LabelsDelete int64 = 3401105936

type poolLabelsDelete struct {
	pool sync.Pool
}

func (p *poolLabelsDelete) Get() *LabelsDelete {
	x, ok := p.pool.Get().(*LabelsDelete)
	if !ok {
		x = &LabelsDelete{}
	}

	return x
}

func (p *poolLabelsDelete) Put(x *LabelsDelete) {
	if x == nil {
		return
	}

	x.LabelIDs = x.LabelIDs[:0]

	p.pool.Put(x)
}

var PoolLabelsDelete = poolLabelsDelete{}

func (x *LabelsDelete) DeepCopy(z *LabelsDelete) {
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
}

func (x *LabelsDelete) Clone() *LabelsDelete {
	z := &LabelsDelete{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsDelete) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsDelete) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsDelete) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsDelete) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsDelete) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsDelete, x)
}

const C_LabelsGet int64 = 2575409921

type poolLabelsGet struct {
	pool sync.Pool
}

func (p *poolLabelsGet) Get() *LabelsGet {
	x, ok := p.pool.Get().(*LabelsGet)
	if !ok {
		x = &LabelsGet{}
	}

	return x
}

func (p *poolLabelsGet) Put(x *LabelsGet) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolLabelsGet = poolLabelsGet{}

func (x *LabelsGet) DeepCopy(z *LabelsGet) {
}

func (x *LabelsGet) Clone() *LabelsGet {
	z := &LabelsGet{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsGet) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsGet) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsGet) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsGet, x)
}

const C_LabelsAddToMessage int64 = 180144503

type poolLabelsAddToMessage struct {
	pool sync.Pool
}

func (p *poolLabelsAddToMessage) Get() *LabelsAddToMessage {
	x, ok := p.pool.Get().(*LabelsAddToMessage)
	if !ok {
		x = &LabelsAddToMessage{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolLabelsAddToMessage) Put(x *LabelsAddToMessage) {
	if x == nil {
		return
	}

	PoolInputPeer.Put(x.Peer)
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]

	p.pool.Put(x)
}

var PoolLabelsAddToMessage = poolLabelsAddToMessage{}

func (x *LabelsAddToMessage) DeepCopy(z *LabelsAddToMessage) {
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *LabelsAddToMessage) Clone() *LabelsAddToMessage {
	z := &LabelsAddToMessage{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsAddToMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsAddToMessage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsAddToMessage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsAddToMessage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsAddToMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsAddToMessage, x)
}

const C_LabelsRemoveFromMessage int64 = 4195197703

type poolLabelsRemoveFromMessage struct {
	pool sync.Pool
}

func (p *poolLabelsRemoveFromMessage) Get() *LabelsRemoveFromMessage {
	x, ok := p.pool.Get().(*LabelsRemoveFromMessage)
	if !ok {
		x = &LabelsRemoveFromMessage{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolLabelsRemoveFromMessage) Put(x *LabelsRemoveFromMessage) {
	if x == nil {
		return
	}

	PoolInputPeer.Put(x.Peer)
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]

	p.pool.Put(x)
}

var PoolLabelsRemoveFromMessage = poolLabelsRemoveFromMessage{}

func (x *LabelsRemoveFromMessage) DeepCopy(z *LabelsRemoveFromMessage) {
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *LabelsRemoveFromMessage) Clone() *LabelsRemoveFromMessage {
	z := &LabelsRemoveFromMessage{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsRemoveFromMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsRemoveFromMessage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsRemoveFromMessage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsRemoveFromMessage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsRemoveFromMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsRemoveFromMessage, x)
}

const C_LabelsListItems int64 = 2351763198

type poolLabelsListItems struct {
	pool sync.Pool
}

func (p *poolLabelsListItems) Get() *LabelsListItems {
	x, ok := p.pool.Get().(*LabelsListItems)
	if !ok {
		x = &LabelsListItems{}
	}

	return x
}

func (p *poolLabelsListItems) Put(x *LabelsListItems) {
	if x == nil {
		return
	}

	x.LabelID = 0
	x.MinID = 0
	x.MaxID = 0
	x.Limit = 0

	p.pool.Put(x)
}

var PoolLabelsListItems = poolLabelsListItems{}

func (x *LabelsListItems) DeepCopy(z *LabelsListItems) {
	z.LabelID = x.LabelID
	z.MinID = x.MinID
	z.MaxID = x.MaxID
	z.Limit = x.Limit
}

func (x *LabelsListItems) Clone() *LabelsListItems {
	z := &LabelsListItems{}
	x.DeepCopy(z)
	return z
}

func (x *LabelsListItems) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsListItems) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelsListItems) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelsListItems) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelsListItems) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsListItems, x)
}

const C_LabelItems int64 = 4271841358

type poolLabelItems struct {
	pool sync.Pool
}

func (p *poolLabelItems) Get() *LabelItems {
	x, ok := p.pool.Get().(*LabelItems)
	if !ok {
		x = &LabelItems{}
	}

	return x
}

func (p *poolLabelItems) Put(x *LabelItems) {
	if x == nil {
		return
	}

	x.LabelID = 0
	for _, z := range x.Messages {
		PoolUserMessage.Put(z)
	}
	x.Messages = x.Messages[:0]
	for _, z := range x.Dialogs {
		PoolDialog.Put(z)
	}
	x.Dialogs = x.Dialogs[:0]
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

var PoolLabelItems = poolLabelItems{}

func (x *LabelItems) DeepCopy(z *LabelItems) {
	z.LabelID = x.LabelID
	for idx := range x.Messages {
		if x.Messages[idx] == nil {
			continue
		}
		xx := PoolUserMessage.Get()
		x.Messages[idx].DeepCopy(xx)
		z.Messages = append(z.Messages, xx)
	}
	for idx := range x.Dialogs {
		if x.Dialogs[idx] == nil {
			continue
		}
		xx := PoolDialog.Get()
		x.Dialogs[idx].DeepCopy(xx)
		z.Dialogs = append(z.Dialogs, xx)
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

func (x *LabelItems) Clone() *LabelItems {
	z := &LabelItems{}
	x.DeepCopy(z)
	return z
}

func (x *LabelItems) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelItems) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LabelItems) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *LabelItems) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *LabelItems) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelItems, x)
}

func init() {
	registry.RegisterConstructor(2138857068, "LabelsCreate")
	registry.RegisterConstructor(2790466877, "LabelsEdit")
	registry.RegisterConstructor(3401105936, "LabelsDelete")
	registry.RegisterConstructor(2575409921, "LabelsGet")
	registry.RegisterConstructor(180144503, "LabelsAddToMessage")
	registry.RegisterConstructor(4195197703, "LabelsRemoveFromMessage")
	registry.RegisterConstructor(2351763198, "LabelsListItems")
	registry.RegisterConstructor(4271841358, "LabelItems")
}

var _ = bytes.MinRead
