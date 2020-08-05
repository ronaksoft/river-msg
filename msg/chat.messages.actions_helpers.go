// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.messages.actions.proto

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

const C_MessageActionGroupAddUser int64 = 1949386261

type poolMessageActionGroupAddUser struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupAddUser) Get() *MessageActionGroupAddUser {
	x, ok := p.pool.Get().(*MessageActionGroupAddUser)
	if !ok {
		return &MessageActionGroupAddUser{}
	}
	return x
}

func (p *poolMessageActionGroupAddUser) Put(x *MessageActionGroupAddUser) {
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupAddUser = poolMessageActionGroupAddUser{}

func ResultMessageActionGroupAddUser(out *MessageEnvelope, res *MessageActionGroupAddUser) {
	out.Constructor = C_MessageActionGroupAddUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionGroupDeleteUser int64 = 1213452128

type poolMessageActionGroupDeleteUser struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupDeleteUser) Get() *MessageActionGroupDeleteUser {
	x, ok := p.pool.Get().(*MessageActionGroupDeleteUser)
	if !ok {
		return &MessageActionGroupDeleteUser{}
	}
	return x
}

func (p *poolMessageActionGroupDeleteUser) Put(x *MessageActionGroupDeleteUser) {
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupDeleteUser = poolMessageActionGroupDeleteUser{}

func ResultMessageActionGroupDeleteUser(out *MessageEnvelope, res *MessageActionGroupDeleteUser) {
	out.Constructor = C_MessageActionGroupDeleteUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionGroupCreated int64 = 2241024808

type poolMessageActionGroupCreated struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupCreated) Get() *MessageActionGroupCreated {
	x, ok := p.pool.Get().(*MessageActionGroupCreated)
	if !ok {
		return &MessageActionGroupCreated{}
	}
	return x
}

func (p *poolMessageActionGroupCreated) Put(x *MessageActionGroupCreated) {
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupCreated = poolMessageActionGroupCreated{}

func ResultMessageActionGroupCreated(out *MessageEnvelope, res *MessageActionGroupCreated) {
	out.Constructor = C_MessageActionGroupCreated
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionGroupTitleChanged int64 = 2418464749

type poolMessageActionGroupTitleChanged struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupTitleChanged) Get() *MessageActionGroupTitleChanged {
	x, ok := p.pool.Get().(*MessageActionGroupTitleChanged)
	if !ok {
		return &MessageActionGroupTitleChanged{}
	}
	return x
}

func (p *poolMessageActionGroupTitleChanged) Put(x *MessageActionGroupTitleChanged) {
	p.pool.Put(x)
}

var PoolMessageActionGroupTitleChanged = poolMessageActionGroupTitleChanged{}

func ResultMessageActionGroupTitleChanged(out *MessageEnvelope, res *MessageActionGroupTitleChanged) {
	out.Constructor = C_MessageActionGroupTitleChanged
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionGroupPhotoChanged int64 = 188265964

type poolMessageActionGroupPhotoChanged struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupPhotoChanged) Get() *MessageActionGroupPhotoChanged {
	x, ok := p.pool.Get().(*MessageActionGroupPhotoChanged)
	if !ok {
		return &MessageActionGroupPhotoChanged{}
	}
	return x
}

func (p *poolMessageActionGroupPhotoChanged) Put(x *MessageActionGroupPhotoChanged) {
	if x.Photo != nil {
		*x.Photo = GroupPhoto{}
	}

	p.pool.Put(x)
}

var PoolMessageActionGroupPhotoChanged = poolMessageActionGroupPhotoChanged{}

func ResultMessageActionGroupPhotoChanged(out *MessageEnvelope, res *MessageActionGroupPhotoChanged) {
	out.Constructor = C_MessageActionGroupPhotoChanged
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionClearHistory int64 = 1270465696

type poolMessageActionClearHistory struct {
	pool sync.Pool
}

func (p *poolMessageActionClearHistory) Get() *MessageActionClearHistory {
	x, ok := p.pool.Get().(*MessageActionClearHistory)
	if !ok {
		return &MessageActionClearHistory{}
	}
	return x
}

func (p *poolMessageActionClearHistory) Put(x *MessageActionClearHistory) {
	p.pool.Put(x)
}

var PoolMessageActionClearHistory = poolMessageActionClearHistory{}

func ResultMessageActionClearHistory(out *MessageEnvelope, res *MessageActionClearHistory) {
	out.Constructor = C_MessageActionClearHistory
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionContactRegistered int64 = 2399156016

type poolMessageActionContactRegistered struct {
	pool sync.Pool
}

func (p *poolMessageActionContactRegistered) Get() *MessageActionContactRegistered {
	x, ok := p.pool.Get().(*MessageActionContactRegistered)
	if !ok {
		return &MessageActionContactRegistered{}
	}
	return x
}

func (p *poolMessageActionContactRegistered) Put(x *MessageActionContactRegistered) {
	p.pool.Put(x)
}

var PoolMessageActionContactRegistered = poolMessageActionContactRegistered{}

func ResultMessageActionContactRegistered(out *MessageEnvelope, res *MessageActionContactRegistered) {
	out.Constructor = C_MessageActionContactRegistered
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionScreenShotTaken int64 = 2637201461

type poolMessageActionScreenShotTaken struct {
	pool sync.Pool
}

func (p *poolMessageActionScreenShotTaken) Get() *MessageActionScreenShotTaken {
	x, ok := p.pool.Get().(*MessageActionScreenShotTaken)
	if !ok {
		return &MessageActionScreenShotTaken{}
	}
	return x
}

func (p *poolMessageActionScreenShotTaken) Put(x *MessageActionScreenShotTaken) {
	x.MinID = 0
	x.MaxID = 0
	p.pool.Put(x)
}

var PoolMessageActionScreenShotTaken = poolMessageActionScreenShotTaken{}

func ResultMessageActionScreenShotTaken(out *MessageEnvelope, res *MessageActionScreenShotTaken) {
	out.Constructor = C_MessageActionScreenShotTaken
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageActionThreadClosed int64 = 1366382890

type poolMessageActionThreadClosed struct {
	pool sync.Pool
}

func (p *poolMessageActionThreadClosed) Get() *MessageActionThreadClosed {
	x, ok := p.pool.Get().(*MessageActionThreadClosed)
	if !ok {
		return &MessageActionThreadClosed{}
	}
	return x
}

func (p *poolMessageActionThreadClosed) Put(x *MessageActionThreadClosed) {
	x.ThreadID = 0
	p.pool.Put(x)
}

var PoolMessageActionThreadClosed = poolMessageActionThreadClosed{}

func ResultMessageActionThreadClosed(out *MessageEnvelope, res *MessageActionThreadClosed) {
	out.Constructor = C_MessageActionThreadClosed
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
	ConstructorNames[1949386261] = "MessageActionGroupAddUser"
	ConstructorNames[1213452128] = "MessageActionGroupDeleteUser"
	ConstructorNames[2241024808] = "MessageActionGroupCreated"
	ConstructorNames[2418464749] = "MessageActionGroupTitleChanged"
	ConstructorNames[188265964] = "MessageActionGroupPhotoChanged"
	ConstructorNames[1270465696] = "MessageActionClearHistory"
	ConstructorNames[2399156016] = "MessageActionContactRegistered"
	ConstructorNames[2637201461] = "MessageActionScreenShotTaken"
	ConstructorNames[1366382890] = "MessageActionThreadClosed"
}
