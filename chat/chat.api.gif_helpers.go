// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.gif.proto

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

const C_GifGetSaved int64 = 35292745

type poolGifGetSaved struct {
	pool sync.Pool
}

func (p *poolGifGetSaved) Get() *GifGetSaved {
	x, ok := p.pool.Get().(*GifGetSaved)
	if !ok {
		return &GifGetSaved{}
	}
	return x
}

func (p *poolGifGetSaved) Put(x *GifGetSaved) {
	p.pool.Put(x)
}

var PoolGifGetSaved = poolGifGetSaved{}

func ResultGifGetSaved(out *MessageEnvelope, res *GifGetSaved) {
	out.Constructor = C_GifGetSaved
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_GifSave int64 = 4049142282

type poolGifSave struct {
	pool sync.Pool
}

func (p *poolGifSave) Get() *GifSave {
	x, ok := p.pool.Get().(*GifSave)
	if !ok {
		return &GifSave{}
	}
	return x
}

func (p *poolGifSave) Put(x *GifSave) {
	p.pool.Put(x)
}

var PoolGifSave = poolGifSave{}

func ResultGifSave(out *MessageEnvelope, res *GifSave) {
	out.Constructor = C_GifSave
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_GifSearch int64 = 2040973085

type poolGifSearch struct {
	pool sync.Pool
}

func (p *poolGifSearch) Get() *GifSearch {
	x, ok := p.pool.Get().(*GifSearch)
	if !ok {
		return &GifSearch{}
	}
	return x
}

func (p *poolGifSearch) Put(x *GifSearch) {
	p.pool.Put(x)
}

var PoolGifSearch = poolGifSearch{}

func ResultGifSearch(out *MessageEnvelope, res *GifSearch) {
	out.Constructor = C_GifSearch
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_FoundGifs int64 = 423157907

type poolFoundGifs struct {
	pool sync.Pool
}

func (p *poolFoundGifs) Get() *FoundGifs {
	x, ok := p.pool.Get().(*FoundGifs)
	if !ok {
		return &FoundGifs{}
	}
	x.Gifs = x.Gifs[:0]
	return x
}

func (p *poolFoundGifs) Put(x *FoundGifs) {
	p.pool.Put(x)
}

var PoolFoundGifs = poolFoundGifs{}

func ResultFoundGifs(out *MessageEnvelope, res *FoundGifs) {
	out.Constructor = C_FoundGifs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_FoundGif int64 = 1539084995

type poolFoundGif struct {
	pool sync.Pool
}

func (p *poolFoundGif) Get() *FoundGif {
	x, ok := p.pool.Get().(*FoundGif)
	if !ok {
		return &FoundGif{}
	}
	return x
}

func (p *poolFoundGif) Put(x *FoundGif) {
	p.pool.Put(x)
}

var PoolFoundGif = poolFoundGif{}

func ResultFoundGif(out *MessageEnvelope, res *FoundGif) {
	out.Constructor = C_FoundGif
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_SavedGifs int64 = 2982813633

type poolSavedGifs struct {
	pool sync.Pool
}

func (p *poolSavedGifs) Get() *SavedGifs {
	x, ok := p.pool.Get().(*SavedGifs)
	if !ok {
		return &SavedGifs{}
	}
	x.Docs = x.Docs[:0]
	return x
}

func (p *poolSavedGifs) Put(x *SavedGifs) {
	p.pool.Put(x)
}

var PoolSavedGifs = poolSavedGifs{}

func ResultSavedGifs(out *MessageEnvelope, res *SavedGifs) {
	out.Constructor = C_SavedGifs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[35292745] = "GifGetSaved"
	ConstructorNames[4049142282] = "GifSave"
	ConstructorNames[2040973085] = "GifSearch"
	ConstructorNames[423157907] = "FoundGifs"
	ConstructorNames[1539084995] = "FoundGif"
	ConstructorNames[2982813633] = "SavedGifs"
}
