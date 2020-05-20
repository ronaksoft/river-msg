// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.wallpaper.proto

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

const C_WallPaperGet int64 = 183906980

type poolWallPaperGet struct {
	pool sync.Pool
}

func (p *poolWallPaperGet) Get() *WallPaperGet {
	x, ok := p.pool.Get().(*WallPaperGet)
	if !ok {
		return &WallPaperGet{}
	}
	return x
}

func (p *poolWallPaperGet) Put(x *WallPaperGet) {
	p.pool.Put(x)
}

var PoolWallPaperGet = poolWallPaperGet{}

func ResultWallPaperGet(out *MessageEnvelope, res *WallPaperGet) {
	out.Constructor = C_WallPaperGet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaperSave int64 = 3559907599

type poolWallPaperSave struct {
	pool sync.Pool
}

func (p *poolWallPaperSave) Get() *WallPaperSave {
	x, ok := p.pool.Get().(*WallPaperSave)
	if !ok {
		return &WallPaperSave{}
	}
	return x
}

func (p *poolWallPaperSave) Put(x *WallPaperSave) {
	p.pool.Put(x)
}

var PoolWallPaperSave = poolWallPaperSave{}

func ResultWallPaperSave(out *MessageEnvelope, res *WallPaperSave) {
	out.Constructor = C_WallPaperSave
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaperDelete int64 = 3006268108

type poolWallPaperDelete struct {
	pool sync.Pool
}

func (p *poolWallPaperDelete) Get() *WallPaperDelete {
	x, ok := p.pool.Get().(*WallPaperDelete)
	if !ok {
		return &WallPaperDelete{}
	}
	return x
}

func (p *poolWallPaperDelete) Put(x *WallPaperDelete) {
	p.pool.Put(x)
}

var PoolWallPaperDelete = poolWallPaperDelete{}

func ResultWallPaperDelete(out *MessageEnvelope, res *WallPaperDelete) {
	out.Constructor = C_WallPaperDelete
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaperUpload int64 = 2661259348

type poolWallPaperUpload struct {
	pool sync.Pool
}

func (p *poolWallPaperUpload) Get() *WallPaperUpload {
	x, ok := p.pool.Get().(*WallPaperUpload)
	if !ok {
		return &WallPaperUpload{}
	}
	x.UploadedFile = nil
	x.File = nil
	return x
}

func (p *poolWallPaperUpload) Put(x *WallPaperUpload) {
	p.pool.Put(x)
}

var PoolWallPaperUpload = poolWallPaperUpload{}

func ResultWallPaperUpload(out *MessageEnvelope, res *WallPaperUpload) {
	out.Constructor = C_WallPaperUpload
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaperReset int64 = 2714244308

type poolWallPaperReset struct {
	pool sync.Pool
}

func (p *poolWallPaperReset) Get() *WallPaperReset {
	x, ok := p.pool.Get().(*WallPaperReset)
	if !ok {
		return &WallPaperReset{}
	}
	return x
}

func (p *poolWallPaperReset) Put(x *WallPaperReset) {
	p.pool.Put(x)
}

var PoolWallPaperReset = poolWallPaperReset{}

func ResultWallPaperReset(out *MessageEnvelope, res *WallPaperReset) {
	out.Constructor = C_WallPaperReset
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputWallPaper int64 = 4000784410

type poolInputWallPaper struct {
	pool sync.Pool
}

func (p *poolInputWallPaper) Get() *InputWallPaper {
	x, ok := p.pool.Get().(*InputWallPaper)
	if !ok {
		return &InputWallPaper{}
	}
	return x
}

func (p *poolInputWallPaper) Put(x *InputWallPaper) {
	p.pool.Put(x)
}

var PoolInputWallPaper = poolInputWallPaper{}

func ResultInputWallPaper(out *MessageEnvelope, res *InputWallPaper) {
	out.Constructor = C_InputWallPaper
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaperSettings int64 = 1098244882

type poolWallPaperSettings struct {
	pool sync.Pool
}

func (p *poolWallPaperSettings) Get() *WallPaperSettings {
	x, ok := p.pool.Get().(*WallPaperSettings)
	if !ok {
		return &WallPaperSettings{}
	}
	x.Blur = false
	x.Motion = false
	x.BackgroundColour = 0
	x.BackgroundSecondColour = 0
	x.Opacity = 0
	x.Rotation = 0
	return x
}

func (p *poolWallPaperSettings) Put(x *WallPaperSettings) {
	p.pool.Put(x)
}

var PoolWallPaperSettings = poolWallPaperSettings{}

func ResultWallPaperSettings(out *MessageEnvelope, res *WallPaperSettings) {
	out.Constructor = C_WallPaperSettings
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPaper int64 = 2527250827

type poolWallPaper struct {
	pool sync.Pool
}

func (p *poolWallPaper) Get() *WallPaper {
	x, ok := p.pool.Get().(*WallPaper)
	if !ok {
		return &WallPaper{}
	}
	x.Creator = false
	x.Default = false
	x.Pattern = false
	x.Dark = false
	x.Document = nil
	x.Settings = nil
	return x
}

func (p *poolWallPaper) Put(x *WallPaper) {
	p.pool.Put(x)
}

var PoolWallPaper = poolWallPaper{}

func ResultWallPaper(out *MessageEnvelope, res *WallPaper) {
	out.Constructor = C_WallPaper
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WallPapersMany int64 = 3121104857

type poolWallPapersMany struct {
	pool sync.Pool
}

func (p *poolWallPapersMany) Get() *WallPapersMany {
	x, ok := p.pool.Get().(*WallPapersMany)
	if !ok {
		return &WallPapersMany{}
	}
	x.WallPapers = x.WallPapers[:0]
	return x
}

func (p *poolWallPapersMany) Put(x *WallPapersMany) {
	p.pool.Put(x)
}

var PoolWallPapersMany = poolWallPapersMany{}

func ResultWallPapersMany(out *MessageEnvelope, res *WallPapersMany) {
	out.Constructor = C_WallPapersMany
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
	ConstructorNames[183906980] = "WallPaperGet"
	ConstructorNames[3559907599] = "WallPaperSave"
	ConstructorNames[3006268108] = "WallPaperDelete"
	ConstructorNames[2661259348] = "WallPaperUpload"
	ConstructorNames[2714244308] = "WallPaperReset"
	ConstructorNames[4000784410] = "InputWallPaper"
	ConstructorNames[1098244882] = "WallPaperSettings"
	ConstructorNames[2527250827] = "WallPaper"
	ConstructorNames[3121104857] = "WallPapersMany"
}
