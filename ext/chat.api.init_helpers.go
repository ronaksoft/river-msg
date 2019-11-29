// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.init.proto

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

const C_InitConnect int64 = 4150793517

type poolInitConnect struct {
	pool sync.Pool
}

func (p poolInitConnect) Get() *InitConnect {
	x, ok := p.pool.Get().(*InitConnect)
	if !ok {
		return &InitConnect{}
	}
	return x
}

func (p poolInitConnect) Put(x *InitConnect) {
	p.pool.Put(x)
}

var PoolInitConnect = poolInitConnect{}

func ResultInitConnect(out *MessageEnvelope, res *InitConnect) {
	out.Constructor = C_InitConnect
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitCompleteAuth int64 = 1583178320

type poolInitCompleteAuth struct {
	pool sync.Pool
}

func (p poolInitCompleteAuth) Get() *InitCompleteAuth {
	x, ok := p.pool.Get().(*InitCompleteAuth)
	if !ok {
		return &InitCompleteAuth{}
	}
	return x
}

func (p poolInitCompleteAuth) Put(x *InitCompleteAuth) {
	p.pool.Put(x)
}

var PoolInitCompleteAuth = poolInitCompleteAuth{}

func ResultInitCompleteAuth(out *MessageEnvelope, res *InitCompleteAuth) {
	out.Constructor = C_InitCompleteAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitConnectTest int64 = 3188015450

type poolInitConnectTest struct {
	pool sync.Pool
}

func (p poolInitConnectTest) Get() *InitConnectTest {
	x, ok := p.pool.Get().(*InitConnectTest)
	if !ok {
		return &InitConnectTest{}
	}
	return x
}

func (p poolInitConnectTest) Put(x *InitConnectTest) {
	p.pool.Put(x)
}

var PoolInitConnectTest = poolInitConnectTest{}

func ResultInitConnectTest(out *MessageEnvelope, res *InitConnectTest) {
	out.Constructor = C_InitConnectTest
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitTestAuth int64 = 2762878006

type poolInitTestAuth struct {
	pool sync.Pool
}

func (p poolInitTestAuth) Get() *InitTestAuth {
	x, ok := p.pool.Get().(*InitTestAuth)
	if !ok {
		return &InitTestAuth{}
	}
	return x
}

func (p poolInitTestAuth) Put(x *InitTestAuth) {
	p.pool.Put(x)
}

var PoolInitTestAuth = poolInitTestAuth{}

func ResultInitTestAuth(out *MessageEnvelope, res *InitTestAuth) {
	out.Constructor = C_InitTestAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitResponse int64 = 4130340247

type poolInitResponse struct {
	pool sync.Pool
}

func (p poolInitResponse) Get() *InitResponse {
	x, ok := p.pool.Get().(*InitResponse)
	if !ok {
		return &InitResponse{}
	}
	return x
}

func (p poolInitResponse) Put(x *InitResponse) {
	p.pool.Put(x)
}

var PoolInitResponse = poolInitResponse{}

func ResultInitResponse(out *MessageEnvelope, res *InitResponse) {
	out.Constructor = C_InitResponse
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitCompleteAuthInternal int64 = 2360982492

type poolInitCompleteAuthInternal struct {
	pool sync.Pool
}

func (p poolInitCompleteAuthInternal) Get() *InitCompleteAuthInternal {
	x, ok := p.pool.Get().(*InitCompleteAuthInternal)
	if !ok {
		return &InitCompleteAuthInternal{}
	}
	return x
}

func (p poolInitCompleteAuthInternal) Put(x *InitCompleteAuthInternal) {
	p.pool.Put(x)
}

var PoolInitCompleteAuthInternal = poolInitCompleteAuthInternal{}

func ResultInitCompleteAuthInternal(out *MessageEnvelope, res *InitCompleteAuthInternal) {
	out.Constructor = C_InitCompleteAuthInternal
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitAuthCompleted int64 = 627708982

type poolInitAuthCompleted struct {
	pool sync.Pool
}

func (p poolInitAuthCompleted) Get() *InitAuthCompleted {
	x, ok := p.pool.Get().(*InitAuthCompleted)
	if !ok {
		return &InitAuthCompleted{}
	}
	return x
}

func (p poolInitAuthCompleted) Put(x *InitAuthCompleted) {
	p.pool.Put(x)
}

var PoolInitAuthCompleted = poolInitAuthCompleted{}

func ResultInitAuthCompleted(out *MessageEnvelope, res *InitAuthCompleted) {
	out.Constructor = C_InitAuthCompleted
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitBindUser int64 = 1933549113

type poolInitBindUser struct {
	pool sync.Pool
}

func (p poolInitBindUser) Get() *InitBindUser {
	x, ok := p.pool.Get().(*InitBindUser)
	if !ok {
		return &InitBindUser{}
	}
	return x
}

func (p poolInitBindUser) Put(x *InitBindUser) {
	p.pool.Put(x)
}

var PoolInitBindUser = poolInitBindUser{}

func ResultInitBindUser(out *MessageEnvelope, res *InitBindUser) {
	out.Constructor = C_InitBindUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InitUserBound int64 = 128391141

type poolInitUserBound struct {
	pool sync.Pool
}

func (p poolInitUserBound) Get() *InitUserBound {
	x, ok := p.pool.Get().(*InitUserBound)
	if !ok {
		return &InitUserBound{}
	}
	return x
}

func (p poolInitUserBound) Put(x *InitUserBound) {
	p.pool.Put(x)
}

var PoolInitUserBound = poolInitUserBound{}

func ResultInitUserBound(out *MessageEnvelope, res *InitUserBound) {
	out.Constructor = C_InitUserBound
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[4150793517] = "InitConnect"
	ConstructorNames[1583178320] = "InitCompleteAuth"
	ConstructorNames[3188015450] = "InitConnectTest"
	ConstructorNames[2762878006] = "InitTestAuth"
	ConstructorNames[4130340247] = "InitResponse"
	ConstructorNames[2360982492] = "InitCompleteAuthInternal"
	ConstructorNames[627708982] = "InitAuthCompleted"
	ConstructorNames[1933549113] = "InitBindUser"
	ConstructorNames[128391141] = "InitUserBound"
}
