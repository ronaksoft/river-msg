// Code generated by Rony's protoc plugin; DO NOT EDIT.

package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_FileSavePart int64 = 3766876582

type poolFileSavePart struct {
	pool sync.Pool
}

func (p *poolFileSavePart) Get() *FileSavePart {
	x, ok := p.pool.Get().(*FileSavePart)
	if !ok {
		x = &FileSavePart{}
	}
	return x
}

func (p *poolFileSavePart) Put(x *FileSavePart) {
	if x == nil {
		return
	}
	x.FileID = 0
	x.PartID = 0
	x.TotalParts = 0
	x.Bytes = x.Bytes[:0]
	p.pool.Put(x)
}

var PoolFileSavePart = poolFileSavePart{}

func (x *FileSavePart) DeepCopy(z *FileSavePart) {
	z.FileID = x.FileID
	z.PartID = x.PartID
	z.TotalParts = x.TotalParts
	z.Bytes = append(z.Bytes[:0], x.Bytes...)
}

func (x *FileSavePart) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *FileSavePart) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *FileSavePart) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_FileSavePart, x)
}

const C_FileGet int64 = 4282510672

type poolFileGet struct {
	pool sync.Pool
}

func (p *poolFileGet) Get() *FileGet {
	x, ok := p.pool.Get().(*FileGet)
	if !ok {
		x = &FileGet{}
	}
	return x
}

func (p *poolFileGet) Put(x *FileGet) {
	if x == nil {
		return
	}
	PoolInputFileLocation.Put(x.Location)
	x.Location = nil
	x.Offset = 0
	x.Limit = 0
	p.pool.Put(x)
}

var PoolFileGet = poolFileGet{}

func (x *FileGet) DeepCopy(z *FileGet) {
	if x.Location != nil {
		if z.Location == nil {
			z.Location = PoolInputFileLocation.Get()
		}
		x.Location.DeepCopy(z.Location)
	} else {
		z.Location = nil
	}
	z.Offset = x.Offset
	z.Limit = x.Limit
}

func (x *FileGet) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *FileGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *FileGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_FileGet, x)
}

const C_FileGetBySha256 int64 = 2768049463

type poolFileGetBySha256 struct {
	pool sync.Pool
}

func (p *poolFileGetBySha256) Get() *FileGetBySha256 {
	x, ok := p.pool.Get().(*FileGetBySha256)
	if !ok {
		x = &FileGetBySha256{}
	}
	return x
}

func (p *poolFileGetBySha256) Put(x *FileGetBySha256) {
	if x == nil {
		return
	}
	x.Sha256 = x.Sha256[:0]
	x.FileSize = 0
	p.pool.Put(x)
}

var PoolFileGetBySha256 = poolFileGetBySha256{}

func (x *FileGetBySha256) DeepCopy(z *FileGetBySha256) {
	z.Sha256 = append(z.Sha256[:0], x.Sha256...)
	z.FileSize = x.FileSize
}

func (x *FileGetBySha256) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *FileGetBySha256) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *FileGetBySha256) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_FileGetBySha256, x)
}

const C_File int64 = 749574446

type poolFile struct {
	pool sync.Pool
}

func (p *poolFile) Get() *File {
	x, ok := p.pool.Get().(*File)
	if !ok {
		x = &File{}
	}
	return x
}

func (p *poolFile) Put(x *File) {
	if x == nil {
		return
	}
	x.Type = 0
	x.ModifiedTime = 0
	x.Bytes = x.Bytes[:0]
	x.MD5Hash = ""
	p.pool.Put(x)
}

var PoolFile = poolFile{}

func (x *File) DeepCopy(z *File) {
	z.Type = x.Type
	z.ModifiedTime = x.ModifiedTime
	z.Bytes = append(z.Bytes[:0], x.Bytes...)
	z.MD5Hash = x.MD5Hash
}

func (x *File) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *File) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *File) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_File, x)
}

func init() {
	registry.RegisterConstructor(3766876582, "FileSavePart")
	registry.RegisterConstructor(4282510672, "FileGet")
	registry.RegisterConstructor(2768049463, "FileGetBySha256")
	registry.RegisterConstructor(749574446, "File")
}
