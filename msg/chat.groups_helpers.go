// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.groups.proto

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

const C_GroupsCreate int64 = 1271969037

type poolGroupsCreate struct {
	pool sync.Pool
}

func (p *poolGroupsCreate) Get() *GroupsCreate {
	x, ok := p.pool.Get().(*GroupsCreate)
	if !ok {
		return &GroupsCreate{}
	}
	x.Users = x.Users[:0]
	return x
}

func (p *poolGroupsCreate) Put(x *GroupsCreate) {
	p.pool.Put(x)
}

var PoolGroupsCreate = poolGroupsCreate{}

func ResultGroupsCreate(out *MessageEnvelope, res *GroupsCreate) {
	out.Constructor = C_GroupsCreate
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsAddUser int64 = 394654713

type poolGroupsAddUser struct {
	pool sync.Pool
}

func (p *poolGroupsAddUser) Get() *GroupsAddUser {
	x, ok := p.pool.Get().(*GroupsAddUser)
	if !ok {
		return &GroupsAddUser{}
	}
	return x
}

func (p *poolGroupsAddUser) Put(x *GroupsAddUser) {
	p.pool.Put(x)
}

var PoolGroupsAddUser = poolGroupsAddUser{}

func ResultGroupsAddUser(out *MessageEnvelope, res *GroupsAddUser) {
	out.Constructor = C_GroupsAddUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsEditTitle int64 = 2582813461

type poolGroupsEditTitle struct {
	pool sync.Pool
}

func (p *poolGroupsEditTitle) Get() *GroupsEditTitle {
	x, ok := p.pool.Get().(*GroupsEditTitle)
	if !ok {
		return &GroupsEditTitle{}
	}
	return x
}

func (p *poolGroupsEditTitle) Put(x *GroupsEditTitle) {
	p.pool.Put(x)
}

var PoolGroupsEditTitle = poolGroupsEditTitle{}

func ResultGroupsEditTitle(out *MessageEnvelope, res *GroupsEditTitle) {
	out.Constructor = C_GroupsEditTitle
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsDeleteUser int64 = 3172322223

type poolGroupsDeleteUser struct {
	pool sync.Pool
}

func (p *poolGroupsDeleteUser) Get() *GroupsDeleteUser {
	x, ok := p.pool.Get().(*GroupsDeleteUser)
	if !ok {
		return &GroupsDeleteUser{}
	}
	return x
}

func (p *poolGroupsDeleteUser) Put(x *GroupsDeleteUser) {
	p.pool.Put(x)
}

var PoolGroupsDeleteUser = poolGroupsDeleteUser{}

func ResultGroupsDeleteUser(out *MessageEnvelope, res *GroupsDeleteUser) {
	out.Constructor = C_GroupsDeleteUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsGetFull int64 = 2986704909

type poolGroupsGetFull struct {
	pool sync.Pool
}

func (p *poolGroupsGetFull) Get() *GroupsGetFull {
	x, ok := p.pool.Get().(*GroupsGetFull)
	if !ok {
		return &GroupsGetFull{}
	}
	return x
}

func (p *poolGroupsGetFull) Put(x *GroupsGetFull) {
	p.pool.Put(x)
}

var PoolGroupsGetFull = poolGroupsGetFull{}

func ResultGroupsGetFull(out *MessageEnvelope, res *GroupsGetFull) {
	out.Constructor = C_GroupsGetFull
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsToggleAdmins int64 = 1581076909

type poolGroupsToggleAdmins struct {
	pool sync.Pool
}

func (p *poolGroupsToggleAdmins) Get() *GroupsToggleAdmins {
	x, ok := p.pool.Get().(*GroupsToggleAdmins)
	if !ok {
		return &GroupsToggleAdmins{}
	}
	return x
}

func (p *poolGroupsToggleAdmins) Put(x *GroupsToggleAdmins) {
	p.pool.Put(x)
}

var PoolGroupsToggleAdmins = poolGroupsToggleAdmins{}

func ResultGroupsToggleAdmins(out *MessageEnvelope, res *GroupsToggleAdmins) {
	out.Constructor = C_GroupsToggleAdmins
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsUpdateAdmin int64 = 1345991011

type poolGroupsUpdateAdmin struct {
	pool sync.Pool
}

func (p *poolGroupsUpdateAdmin) Get() *GroupsUpdateAdmin {
	x, ok := p.pool.Get().(*GroupsUpdateAdmin)
	if !ok {
		return &GroupsUpdateAdmin{}
	}
	return x
}

func (p *poolGroupsUpdateAdmin) Put(x *GroupsUpdateAdmin) {
	p.pool.Put(x)
}

var PoolGroupsUpdateAdmin = poolGroupsUpdateAdmin{}

func ResultGroupsUpdateAdmin(out *MessageEnvelope, res *GroupsUpdateAdmin) {
	out.Constructor = C_GroupsUpdateAdmin
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsUploadPhoto int64 = 2624284907

type poolGroupsUploadPhoto struct {
	pool sync.Pool
}

func (p *poolGroupsUploadPhoto) Get() *GroupsUploadPhoto {
	x, ok := p.pool.Get().(*GroupsUploadPhoto)
	if !ok {
		return &GroupsUploadPhoto{}
	}
	x.ReturnObject = false
	return x
}

func (p *poolGroupsUploadPhoto) Put(x *GroupsUploadPhoto) {
	p.pool.Put(x)
}

var PoolGroupsUploadPhoto = poolGroupsUploadPhoto{}

func ResultGroupsUploadPhoto(out *MessageEnvelope, res *GroupsUploadPhoto) {
	out.Constructor = C_GroupsUploadPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsRemovePhoto int64 = 176771682

type poolGroupsRemovePhoto struct {
	pool sync.Pool
}

func (p *poolGroupsRemovePhoto) Get() *GroupsRemovePhoto {
	x, ok := p.pool.Get().(*GroupsRemovePhoto)
	if !ok {
		return &GroupsRemovePhoto{}
	}
	x.PhotoID = 0
	return x
}

func (p *poolGroupsRemovePhoto) Put(x *GroupsRemovePhoto) {
	p.pool.Put(x)
}

var PoolGroupsRemovePhoto = poolGroupsRemovePhoto{}

func ResultGroupsRemovePhoto(out *MessageEnvelope, res *GroupsRemovePhoto) {
	out.Constructor = C_GroupsRemovePhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupsUpdatePhoto int64 = 3431184397

type poolGroupsUpdatePhoto struct {
	pool sync.Pool
}

func (p *poolGroupsUpdatePhoto) Get() *GroupsUpdatePhoto {
	x, ok := p.pool.Get().(*GroupsUpdatePhoto)
	if !ok {
		return &GroupsUpdatePhoto{}
	}
	return x
}

func (p *poolGroupsUpdatePhoto) Put(x *GroupsUpdatePhoto) {
	p.pool.Put(x)
}

var PoolGroupsUpdatePhoto = poolGroupsUpdatePhoto{}

func ResultGroupsUpdatePhoto(out *MessageEnvelope, res *GroupsUpdatePhoto) {
	out.Constructor = C_GroupsUpdatePhoto
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
	ConstructorNames[1271969037] = "GroupsCreate"
	ConstructorNames[394654713] = "GroupsAddUser"
	ConstructorNames[2582813461] = "GroupsEditTitle"
	ConstructorNames[3172322223] = "GroupsDeleteUser"
	ConstructorNames[2986704909] = "GroupsGetFull"
	ConstructorNames[1581076909] = "GroupsToggleAdmins"
	ConstructorNames[1345991011] = "GroupsUpdateAdmin"
	ConstructorNames[2624284907] = "GroupsUploadPhoto"
	ConstructorNames[176771682] = "GroupsRemovePhoto"
	ConstructorNames[3431184397] = "GroupsUpdatePhoto"
}
