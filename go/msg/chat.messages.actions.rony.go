// Code generated by Rony's protoc plugin; DO NOT EDIT.

package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_MessageActionGroupAddUser int64 = 1949386261

type poolMessageActionGroupAddUser struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupAddUser) Get() *MessageActionGroupAddUser {
	x, ok := p.pool.Get().(*MessageActionGroupAddUser)
	if !ok {
		x = &MessageActionGroupAddUser{}
	}
	return x
}

func (p *poolMessageActionGroupAddUser) Put(x *MessageActionGroupAddUser) {
	if x == nil {
		return
	}
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupAddUser = poolMessageActionGroupAddUser{}

func (x *MessageActionGroupAddUser) DeepCopy(z *MessageActionGroupAddUser) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *MessageActionGroupAddUser) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionGroupAddUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionGroupAddUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionGroupAddUser, x)
}

const C_MessageActionGroupDeleteUser int64 = 1213452128

type poolMessageActionGroupDeleteUser struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupDeleteUser) Get() *MessageActionGroupDeleteUser {
	x, ok := p.pool.Get().(*MessageActionGroupDeleteUser)
	if !ok {
		x = &MessageActionGroupDeleteUser{}
	}
	return x
}

func (p *poolMessageActionGroupDeleteUser) Put(x *MessageActionGroupDeleteUser) {
	if x == nil {
		return
	}
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupDeleteUser = poolMessageActionGroupDeleteUser{}

func (x *MessageActionGroupDeleteUser) DeepCopy(z *MessageActionGroupDeleteUser) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *MessageActionGroupDeleteUser) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionGroupDeleteUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionGroupDeleteUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionGroupDeleteUser, x)
}

const C_MessageActionGroupCreated int64 = 2241024808

type poolMessageActionGroupCreated struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupCreated) Get() *MessageActionGroupCreated {
	x, ok := p.pool.Get().(*MessageActionGroupCreated)
	if !ok {
		x = &MessageActionGroupCreated{}
	}
	return x
}

func (p *poolMessageActionGroupCreated) Put(x *MessageActionGroupCreated) {
	if x == nil {
		return
	}
	x.GroupTitle = ""
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolMessageActionGroupCreated = poolMessageActionGroupCreated{}

func (x *MessageActionGroupCreated) DeepCopy(z *MessageActionGroupCreated) {
	z.GroupTitle = x.GroupTitle
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *MessageActionGroupCreated) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionGroupCreated) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionGroupCreated) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionGroupCreated, x)
}

const C_MessageActionGroupTitleChanged int64 = 2418464749

type poolMessageActionGroupTitleChanged struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupTitleChanged) Get() *MessageActionGroupTitleChanged {
	x, ok := p.pool.Get().(*MessageActionGroupTitleChanged)
	if !ok {
		x = &MessageActionGroupTitleChanged{}
	}
	return x
}

func (p *poolMessageActionGroupTitleChanged) Put(x *MessageActionGroupTitleChanged) {
	if x == nil {
		return
	}
	x.GroupTitle = ""
	p.pool.Put(x)
}

var PoolMessageActionGroupTitleChanged = poolMessageActionGroupTitleChanged{}

func (x *MessageActionGroupTitleChanged) DeepCopy(z *MessageActionGroupTitleChanged) {
	z.GroupTitle = x.GroupTitle
}

func (x *MessageActionGroupTitleChanged) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionGroupTitleChanged) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionGroupTitleChanged) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionGroupTitleChanged, x)
}

const C_MessageActionGroupPhotoChanged int64 = 188265964

type poolMessageActionGroupPhotoChanged struct {
	pool sync.Pool
}

func (p *poolMessageActionGroupPhotoChanged) Get() *MessageActionGroupPhotoChanged {
	x, ok := p.pool.Get().(*MessageActionGroupPhotoChanged)
	if !ok {
		x = &MessageActionGroupPhotoChanged{}
	}
	return x
}

func (p *poolMessageActionGroupPhotoChanged) Put(x *MessageActionGroupPhotoChanged) {
	if x == nil {
		return
	}
	PoolGroupPhoto.Put(x.Photo)
	x.Photo = nil
	p.pool.Put(x)
}

var PoolMessageActionGroupPhotoChanged = poolMessageActionGroupPhotoChanged{}

func (x *MessageActionGroupPhotoChanged) DeepCopy(z *MessageActionGroupPhotoChanged) {
	if x.Photo != nil {
		if z.Photo == nil {
			z.Photo = PoolGroupPhoto.Get()
		}
		x.Photo.DeepCopy(z.Photo)
	} else {
		z.Photo = nil
	}
}

func (x *MessageActionGroupPhotoChanged) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionGroupPhotoChanged) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionGroupPhotoChanged) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionGroupPhotoChanged, x)
}

const C_MessageActionClearHistory int64 = 1270465696

type poolMessageActionClearHistory struct {
	pool sync.Pool
}

func (p *poolMessageActionClearHistory) Get() *MessageActionClearHistory {
	x, ok := p.pool.Get().(*MessageActionClearHistory)
	if !ok {
		x = &MessageActionClearHistory{}
	}
	return x
}

func (p *poolMessageActionClearHistory) Put(x *MessageActionClearHistory) {
	if x == nil {
		return
	}
	x.MaxID = 0
	x.Delete = false
	p.pool.Put(x)
}

var PoolMessageActionClearHistory = poolMessageActionClearHistory{}

func (x *MessageActionClearHistory) DeepCopy(z *MessageActionClearHistory) {
	z.MaxID = x.MaxID
	z.Delete = x.Delete
}

func (x *MessageActionClearHistory) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionClearHistory) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionClearHistory) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionClearHistory, x)
}

const C_MessageActionContactRegistered int64 = 2399156016

type poolMessageActionContactRegistered struct {
	pool sync.Pool
}

func (p *poolMessageActionContactRegistered) Get() *MessageActionContactRegistered {
	x, ok := p.pool.Get().(*MessageActionContactRegistered)
	if !ok {
		x = &MessageActionContactRegistered{}
	}
	return x
}

func (p *poolMessageActionContactRegistered) Put(x *MessageActionContactRegistered) {
	if x == nil {
		return
	}
	p.pool.Put(x)
}

var PoolMessageActionContactRegistered = poolMessageActionContactRegistered{}

func (x *MessageActionContactRegistered) DeepCopy(z *MessageActionContactRegistered) {
}

func (x *MessageActionContactRegistered) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionContactRegistered) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionContactRegistered) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionContactRegistered, x)
}

const C_MessageActionScreenShotTaken int64 = 2637201461

type poolMessageActionScreenShotTaken struct {
	pool sync.Pool
}

func (p *poolMessageActionScreenShotTaken) Get() *MessageActionScreenShotTaken {
	x, ok := p.pool.Get().(*MessageActionScreenShotTaken)
	if !ok {
		x = &MessageActionScreenShotTaken{}
	}
	return x
}

func (p *poolMessageActionScreenShotTaken) Put(x *MessageActionScreenShotTaken) {
	if x == nil {
		return
	}
	x.MinID = 0
	x.MaxID = 0
	p.pool.Put(x)
}

var PoolMessageActionScreenShotTaken = poolMessageActionScreenShotTaken{}

func (x *MessageActionScreenShotTaken) DeepCopy(z *MessageActionScreenShotTaken) {
	z.MinID = x.MinID
	z.MaxID = x.MaxID
}

func (x *MessageActionScreenShotTaken) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionScreenShotTaken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionScreenShotTaken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionScreenShotTaken, x)
}

const C_MessageActionThreadClosed int64 = 1366382890

type poolMessageActionThreadClosed struct {
	pool sync.Pool
}

func (p *poolMessageActionThreadClosed) Get() *MessageActionThreadClosed {
	x, ok := p.pool.Get().(*MessageActionThreadClosed)
	if !ok {
		x = &MessageActionThreadClosed{}
	}
	return x
}

func (p *poolMessageActionThreadClosed) Put(x *MessageActionThreadClosed) {
	if x == nil {
		return
	}
	x.ThreadID = 0
	p.pool.Put(x)
}

var PoolMessageActionThreadClosed = poolMessageActionThreadClosed{}

func (x *MessageActionThreadClosed) DeepCopy(z *MessageActionThreadClosed) {
	z.ThreadID = x.ThreadID
}

func (x *MessageActionThreadClosed) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionThreadClosed) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionThreadClosed) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionThreadClosed, x)
}

const C_MessageActionCallStarted int64 = 3970382785

type poolMessageActionCallStarted struct {
	pool sync.Pool
}

func (p *poolMessageActionCallStarted) Get() *MessageActionCallStarted {
	x, ok := p.pool.Get().(*MessageActionCallStarted)
	if !ok {
		x = &MessageActionCallStarted{}
	}
	return x
}

func (p *poolMessageActionCallStarted) Put(x *MessageActionCallStarted) {
	if x == nil {
		return
	}
	x.CallID = 0
	p.pool.Put(x)
}

var PoolMessageActionCallStarted = poolMessageActionCallStarted{}

func (x *MessageActionCallStarted) DeepCopy(z *MessageActionCallStarted) {
	z.CallID = x.CallID
}

func (x *MessageActionCallStarted) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionCallStarted) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionCallStarted) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionCallStarted, x)
}

const C_MessageActionCallEnded int64 = 62258227

type poolMessageActionCallEnded struct {
	pool sync.Pool
}

func (p *poolMessageActionCallEnded) Get() *MessageActionCallEnded {
	x, ok := p.pool.Get().(*MessageActionCallEnded)
	if !ok {
		x = &MessageActionCallEnded{}
	}
	return x
}

func (p *poolMessageActionCallEnded) Put(x *MessageActionCallEnded) {
	if x == nil {
		return
	}
	x.CallID = 0
	x.Reason = 0
	x.Duration = 0
	p.pool.Put(x)
}

var PoolMessageActionCallEnded = poolMessageActionCallEnded{}

func (x *MessageActionCallEnded) DeepCopy(z *MessageActionCallEnded) {
	z.CallID = x.CallID
	z.Reason = x.Reason
	z.Duration = x.Duration
}

func (x *MessageActionCallEnded) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageActionCallEnded) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageActionCallEnded) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageActionCallEnded, x)
}

func init() {
	registry.RegisterConstructor(1949386261, "MessageActionGroupAddUser")
	registry.RegisterConstructor(1213452128, "MessageActionGroupDeleteUser")
	registry.RegisterConstructor(2241024808, "MessageActionGroupCreated")
	registry.RegisterConstructor(2418464749, "MessageActionGroupTitleChanged")
	registry.RegisterConstructor(188265964, "MessageActionGroupPhotoChanged")
	registry.RegisterConstructor(1270465696, "MessageActionClearHistory")
	registry.RegisterConstructor(2399156016, "MessageActionContactRegistered")
	registry.RegisterConstructor(2637201461, "MessageActionScreenShotTaken")
	registry.RegisterConstructor(1366382890, "MessageActionThreadClosed")
	registry.RegisterConstructor(3970382785, "MessageActionCallStarted")
	registry.RegisterConstructor(62258227, "MessageActionCallEnded")
}
