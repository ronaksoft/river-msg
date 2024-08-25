// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: chat.password.algorithms.proto

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

const C_PasswordAlgorithmVer6A int64 = 341860043

type poolPasswordAlgorithmVer6A struct {
	pool sync.Pool
}

func (p *poolPasswordAlgorithmVer6A) Get() *PasswordAlgorithmVer6A {
	x, ok := p.pool.Get().(*PasswordAlgorithmVer6A)
	if !ok {
		x = &PasswordAlgorithmVer6A{}
	}

	return x
}

func (p *poolPasswordAlgorithmVer6A) Put(x *PasswordAlgorithmVer6A) {
	if x == nil {
		return
	}

	x.Salt1 = x.Salt1[:0]
	x.Salt2 = x.Salt2[:0]
	x.G = 0
	x.P = x.P[:0]

	p.pool.Put(x)
}

var PoolPasswordAlgorithmVer6A = poolPasswordAlgorithmVer6A{}

func (x *PasswordAlgorithmVer6A) DeepCopy(z *PasswordAlgorithmVer6A) {
	z.Salt1 = append(z.Salt1[:0], x.Salt1...)
	z.Salt2 = append(z.Salt2[:0], x.Salt2...)
	z.G = x.G
	z.P = append(z.P[:0], x.P...)
}

func (x *PasswordAlgorithmVer6A) Clone() *PasswordAlgorithmVer6A {
	z := &PasswordAlgorithmVer6A{}
	x.DeepCopy(z)
	return z
}

func (x *PasswordAlgorithmVer6A) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PasswordAlgorithmVer6A) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PasswordAlgorithmVer6A) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *PasswordAlgorithmVer6A) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *PasswordAlgorithmVer6A) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PasswordAlgorithmVer6A, x)
}

func init() {
	registry.RegisterConstructor(341860043, "PasswordAlgorithmVer6A")
}

var _ = bytes.MinRead
