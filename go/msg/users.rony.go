// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v3.15.8
// Rony ver. v0.12.22
// Source: users.proto

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

const C_UsersGet int64 = 1039301579

type poolUsersGet struct {
	pool sync.Pool
}

func (p *poolUsersGet) Get() *UsersGet {
	x, ok := p.pool.Get().(*UsersGet)
	if !ok {
		x = &UsersGet{}
	}

	return x
}

func (p *poolUsersGet) Put(x *UsersGet) {
	if x == nil {
		return
	}

	for _, z := range x.Users {
		PoolInputUser.Put(z)
	}
	x.Users = x.Users[:0]

	p.pool.Put(x)
}

var PoolUsersGet = poolUsersGet{}

func (x *UsersGet) DeepCopy(z *UsersGet) {
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolInputUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
}

func (x *UsersGet) Clone() *UsersGet {
	z := &UsersGet{}
	x.DeepCopy(z)
	return z
}

func (x *UsersGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UsersGet) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UsersGet) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *UsersGet) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *UsersGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersGet, x)
}

const C_UsersGetFull int64 = 3343342086

type poolUsersGetFull struct {
	pool sync.Pool
}

func (p *poolUsersGetFull) Get() *UsersGetFull {
	x, ok := p.pool.Get().(*UsersGetFull)
	if !ok {
		x = &UsersGetFull{}
	}

	return x
}

func (p *poolUsersGetFull) Put(x *UsersGetFull) {
	if x == nil {
		return
	}

	for _, z := range x.Users {
		PoolInputUser.Put(z)
	}
	x.Users = x.Users[:0]

	p.pool.Put(x)
}

var PoolUsersGetFull = poolUsersGetFull{}

func (x *UsersGetFull) DeepCopy(z *UsersGetFull) {
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolInputUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
}

func (x *UsersGetFull) Clone() *UsersGetFull {
	z := &UsersGetFull{}
	x.DeepCopy(z)
	return z
}

func (x *UsersGetFull) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UsersGetFull) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UsersGetFull) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *UsersGetFull) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *UsersGetFull) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersGetFull, x)
}

const C_UsersMany int64 = 801733941

type poolUsersMany struct {
	pool sync.Pool
}

func (p *poolUsersMany) Get() *UsersMany {
	x, ok := p.pool.Get().(*UsersMany)
	if !ok {
		x = &UsersMany{}
	}

	return x
}

func (p *poolUsersMany) Put(x *UsersMany) {
	if x == nil {
		return
	}

	for _, z := range x.Users {
		PoolUser.Put(z)
	}
	x.Users = x.Users[:0]
	x.Empty = false

	p.pool.Put(x)
}

var PoolUsersMany = poolUsersMany{}

func (x *UsersMany) DeepCopy(z *UsersMany) {
	for idx := range x.Users {
		if x.Users[idx] == nil {
			continue
		}
		xx := PoolUser.Get()
		x.Users[idx].DeepCopy(xx)
		z.Users = append(z.Users, xx)
	}
	z.Empty = x.Empty
}

func (x *UsersMany) Clone() *UsersMany {
	z := &UsersMany{}
	x.DeepCopy(z)
	return z
}

func (x *UsersMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UsersMany) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *UsersMany) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *UsersMany) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *UsersMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersMany, x)
}

func init() {
	registry.RegisterConstructor(1039301579, "UsersGet")
	registry.RegisterConstructor(3343342086, "UsersGetFull")
	registry.RegisterConstructor(801733941, "UsersMany")
}

var _ = bytes.MinRead
