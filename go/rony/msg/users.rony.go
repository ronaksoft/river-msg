package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_UsersGet int64 = 1039301579

type poolUsersGet struct {
	pool sync.Pool
}

func (p *poolUsersGet) Get() *UsersGet {
	x, ok := p.pool.Get().(*UsersGet)
	if !ok {
		return &UsersGet{}
	}
	return x
}

func (p *poolUsersGet) Put(x *UsersGet) {
	x.Users = x.Users[:0]
	p.pool.Put(x)
}

var PoolUsersGet = poolUsersGet{}

const C_UsersGetFull int64 = 3343342086

type poolUsersGetFull struct {
	pool sync.Pool
}

func (p *poolUsersGetFull) Get() *UsersGetFull {
	x, ok := p.pool.Get().(*UsersGetFull)
	if !ok {
		return &UsersGetFull{}
	}
	return x
}

func (p *poolUsersGetFull) Put(x *UsersGetFull) {
	x.Users = x.Users[:0]
	p.pool.Put(x)
}

var PoolUsersGetFull = poolUsersGetFull{}

const C_UsersMany int64 = 801733941

type poolUsersMany struct {
	pool sync.Pool
}

func (p *poolUsersMany) Get() *UsersMany {
	x, ok := p.pool.Get().(*UsersMany)
	if !ok {
		return &UsersMany{}
	}
	return x
}

func (p *poolUsersMany) Put(x *UsersMany) {
	x.Users = x.Users[:0]
	x.Empty = false
	p.pool.Put(x)
}

var PoolUsersMany = poolUsersMany{}

func init() {
	registry.RegisterConstructor(1039301579, "UsersGet")
	registry.RegisterConstructor(3343342086, "UsersGetFull")
	registry.RegisterConstructor(801733941, "UsersMany")
}

func (x *UsersGet) DeepCopy(z *UsersGet) {
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolInputUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
}

func (x *UsersGetFull) DeepCopy(z *UsersGetFull) {
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolInputUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
}

func (x *UsersMany) DeepCopy(z *UsersMany) {
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	z.Empty = x.Empty
}

func (x *UsersGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersGet, x)
}

func (x *UsersGetFull) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersGetFull, x)
}

func (x *UsersMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UsersMany, x)
}

func (x *UsersGet) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UsersGetFull) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UsersMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UsersGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UsersGetFull) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UsersMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}