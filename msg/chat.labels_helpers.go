// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.labels.proto

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

const C_LabelsCreate int64 = 2138857068

type poolLabelsCreate struct {
	pool sync.Pool
}

func (p *poolLabelsCreate) Get() *LabelsCreate {
	x, ok := p.pool.Get().(*LabelsCreate)
	if !ok {
		return &LabelsCreate{}
	}
	return x
}

func (p *poolLabelsCreate) Put(x *LabelsCreate) {
	p.pool.Put(x)
}

var PoolLabelsCreate = poolLabelsCreate{}

func ResultLabelsCreate(out *MessageEnvelope, res *LabelsCreate) {
	out.Constructor = C_LabelsCreate
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsEdit int64 = 2790466877

type poolLabelsEdit struct {
	pool sync.Pool
}

func (p *poolLabelsEdit) Get() *LabelsEdit {
	x, ok := p.pool.Get().(*LabelsEdit)
	if !ok {
		return &LabelsEdit{}
	}
	return x
}

func (p *poolLabelsEdit) Put(x *LabelsEdit) {
	p.pool.Put(x)
}

var PoolLabelsEdit = poolLabelsEdit{}

func ResultLabelsEdit(out *MessageEnvelope, res *LabelsEdit) {
	out.Constructor = C_LabelsEdit
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsDelete int64 = 3401105936

type poolLabelsDelete struct {
	pool sync.Pool
}

func (p *poolLabelsDelete) Get() *LabelsDelete {
	x, ok := p.pool.Get().(*LabelsDelete)
	if !ok {
		return &LabelsDelete{}
	}
	x.LabelIDs = x.LabelIDs[:0]
	return x
}

func (p *poolLabelsDelete) Put(x *LabelsDelete) {
	p.pool.Put(x)
}

var PoolLabelsDelete = poolLabelsDelete{}

func ResultLabelsDelete(out *MessageEnvelope, res *LabelsDelete) {
	out.Constructor = C_LabelsDelete
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsGet int64 = 2575409921

type poolLabelsGet struct {
	pool sync.Pool
}

func (p *poolLabelsGet) Get() *LabelsGet {
	x, ok := p.pool.Get().(*LabelsGet)
	if !ok {
		return &LabelsGet{}
	}
	return x
}

func (p *poolLabelsGet) Put(x *LabelsGet) {
	p.pool.Put(x)
}

var PoolLabelsGet = poolLabelsGet{}

func ResultLabelsGet(out *MessageEnvelope, res *LabelsGet) {
	out.Constructor = C_LabelsGet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsAddToMessage int64 = 180144503

type poolLabelsAddToMessage struct {
	pool sync.Pool
}

func (p *poolLabelsAddToMessage) Get() *LabelsAddToMessage {
	x, ok := p.pool.Get().(*LabelsAddToMessage)
	if !ok {
		return &LabelsAddToMessage{}
	}
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]
	x.Team = nil
	return x
}

func (p *poolLabelsAddToMessage) Put(x *LabelsAddToMessage) {
	p.pool.Put(x)
}

var PoolLabelsAddToMessage = poolLabelsAddToMessage{}

func ResultLabelsAddToMessage(out *MessageEnvelope, res *LabelsAddToMessage) {
	out.Constructor = C_LabelsAddToMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsRemoveFromMessage int64 = 4195197703

type poolLabelsRemoveFromMessage struct {
	pool sync.Pool
}

func (p *poolLabelsRemoveFromMessage) Get() *LabelsRemoveFromMessage {
	x, ok := p.pool.Get().(*LabelsRemoveFromMessage)
	if !ok {
		return &LabelsRemoveFromMessage{}
	}
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]
	x.Team = nil
	return x
}

func (p *poolLabelsRemoveFromMessage) Put(x *LabelsRemoveFromMessage) {
	p.pool.Put(x)
}

var PoolLabelsRemoveFromMessage = poolLabelsRemoveFromMessage{}

func ResultLabelsRemoveFromMessage(out *MessageEnvelope, res *LabelsRemoveFromMessage) {
	out.Constructor = C_LabelsRemoveFromMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsListItems int64 = 2351763198

type poolLabelsListItems struct {
	pool sync.Pool
}

func (p *poolLabelsListItems) Get() *LabelsListItems {
	x, ok := p.pool.Get().(*LabelsListItems)
	if !ok {
		return &LabelsListItems{}
	}
	x.MinID = 0
	x.MaxID = 0
	x.Limit = 0
	x.Team = nil
	return x
}

func (p *poolLabelsListItems) Put(x *LabelsListItems) {
	p.pool.Put(x)
}

var PoolLabelsListItems = poolLabelsListItems{}

func ResultLabelsListItems(out *MessageEnvelope, res *LabelsListItems) {
	out.Constructor = C_LabelsListItems
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelItems int64 = 4271841358

type poolLabelItems struct {
	pool sync.Pool
}

func (p *poolLabelItems) Get() *LabelItems {
	x, ok := p.pool.Get().(*LabelItems)
	if !ok {
		return &LabelItems{}
	}
	x.Messages = x.Messages[:0]
	x.Dialogs = x.Dialogs[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	return x
}

func (p *poolLabelItems) Put(x *LabelItems) {
	p.pool.Put(x)
}

var PoolLabelItems = poolLabelItems{}

func ResultLabelItems(out *MessageEnvelope, res *LabelItems) {
	out.Constructor = C_LabelItems
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
	ConstructorNames[2138857068] = "LabelsCreate"
	ConstructorNames[2790466877] = "LabelsEdit"
	ConstructorNames[3401105936] = "LabelsDelete"
	ConstructorNames[2575409921] = "LabelsGet"
	ConstructorNames[180144503] = "LabelsAddToMessage"
	ConstructorNames[4195197703] = "LabelsRemoveFromMessage"
	ConstructorNames[2351763198] = "LabelsListItems"
	ConstructorNames[4271841358] = "LabelItems"
}
