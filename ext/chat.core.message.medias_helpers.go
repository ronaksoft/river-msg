// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.core.message.medias.proto

package msg

import (
	fmt "fmt"
	pbytes "github.com/gobwas/pool/pbytes"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const C_DocumentAttributeAudio int64 = 309707708

type poolDocumentAttributeAudio struct {
	pool sync.Pool
}

func (p poolDocumentAttributeAudio) Get() *DocumentAttributeAudio {
	x, ok := p.pool.Get().(*DocumentAttributeAudio)
	if !ok {
		return &DocumentAttributeAudio{}
	}
	x.Waveform = nil
	return x
}

func (p poolDocumentAttributeAudio) Put(x *DocumentAttributeAudio) {
	p.pool.Put(x)
}

var PoolDocumentAttributeAudio = poolDocumentAttributeAudio{}

func ResultDocumentAttributeAudio(out *MessageEnvelope, res *DocumentAttributeAudio) {
	out.Constructor = C_DocumentAttributeAudio
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_DocumentAttributeVideo int64 = 1993289477

type poolDocumentAttributeVideo struct {
	pool sync.Pool
}

func (p poolDocumentAttributeVideo) Get() *DocumentAttributeVideo {
	x, ok := p.pool.Get().(*DocumentAttributeVideo)
	if !ok {
		return &DocumentAttributeVideo{}
	}
	return x
}

func (p poolDocumentAttributeVideo) Put(x *DocumentAttributeVideo) {
	p.pool.Put(x)
}

var PoolDocumentAttributeVideo = poolDocumentAttributeVideo{}

func ResultDocumentAttributeVideo(out *MessageEnvelope, res *DocumentAttributeVideo) {
	out.Constructor = C_DocumentAttributeVideo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_DocumentAttributePhoto int64 = 515862833

type poolDocumentAttributePhoto struct {
	pool sync.Pool
}

func (p poolDocumentAttributePhoto) Get() *DocumentAttributePhoto {
	x, ok := p.pool.Get().(*DocumentAttributePhoto)
	if !ok {
		return &DocumentAttributePhoto{}
	}
	return x
}

func (p poolDocumentAttributePhoto) Put(x *DocumentAttributePhoto) {
	p.pool.Put(x)
}

var PoolDocumentAttributePhoto = poolDocumentAttributePhoto{}

func ResultDocumentAttributePhoto(out *MessageEnvelope, res *DocumentAttributePhoto) {
	out.Constructor = C_DocumentAttributePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_DocumentAttributeFile int64 = 2227452062

type poolDocumentAttributeFile struct {
	pool sync.Pool
}

func (p poolDocumentAttributeFile) Get() *DocumentAttributeFile {
	x, ok := p.pool.Get().(*DocumentAttributeFile)
	if !ok {
		return &DocumentAttributeFile{}
	}
	return x
}

func (p poolDocumentAttributeFile) Put(x *DocumentAttributeFile) {
	p.pool.Put(x)
}

var PoolDocumentAttributeFile = poolDocumentAttributeFile{}

func ResultDocumentAttributeFile(out *MessageEnvelope, res *DocumentAttributeFile) {
	out.Constructor = C_DocumentAttributeFile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_DocumentAttribute int64 = 4146719643

type poolDocumentAttribute struct {
	pool sync.Pool
}

func (p poolDocumentAttribute) Get() *DocumentAttribute {
	x, ok := p.pool.Get().(*DocumentAttribute)
	if !ok {
		return &DocumentAttribute{}
	}
	x.Data = nil
	return x
}

func (p poolDocumentAttribute) Put(x *DocumentAttribute) {
	p.pool.Put(x)
}

var PoolDocumentAttribute = poolDocumentAttribute{}

func ResultDocumentAttribute(out *MessageEnvelope, res *DocumentAttribute) {
	out.Constructor = C_DocumentAttribute
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_Document int64 = 555739168

type poolDocument struct {
	pool sync.Pool
}

func (p poolDocument) Get() *Document {
	x, ok := p.pool.Get().(*Document)
	if !ok {
		return &Document{}
	}
	x.Attributes = x.Attributes[:0]
	x.Thumbnail = nil
	x.MD5Checksum = ""
	return x
}

func (p poolDocument) Put(x *Document) {
	p.pool.Put(x)
}

var PoolDocument = poolDocument{}

func ResultDocument(out *MessageEnvelope, res *Document) {
	out.Constructor = C_Document
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaUploadedPhoto int64 = 849930963

type poolInputMediaUploadedPhoto struct {
	pool sync.Pool
}

func (p poolInputMediaUploadedPhoto) Get() *InputMediaUploadedPhoto {
	x, ok := p.pool.Get().(*InputMediaUploadedPhoto)
	if !ok {
		return &InputMediaUploadedPhoto{}
	}
	x.Stickers = x.Stickers[:0]
	x.Attributes = x.Attributes[:0]
	return x
}

func (p poolInputMediaUploadedPhoto) Put(x *InputMediaUploadedPhoto) {
	p.pool.Put(x)
}

var PoolInputMediaUploadedPhoto = poolInputMediaUploadedPhoto{}

func ResultInputMediaUploadedPhoto(out *MessageEnvelope, res *InputMediaUploadedPhoto) {
	out.Constructor = C_InputMediaUploadedPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaPhoto int64 = 2201579839

type poolInputMediaPhoto struct {
	pool sync.Pool
}

func (p poolInputMediaPhoto) Get() *InputMediaPhoto {
	x, ok := p.pool.Get().(*InputMediaPhoto)
	if !ok {
		return &InputMediaPhoto{}
	}
	return x
}

func (p poolInputMediaPhoto) Put(x *InputMediaPhoto) {
	p.pool.Put(x)
}

var PoolInputMediaPhoto = poolInputMediaPhoto{}

func ResultInputMediaPhoto(out *MessageEnvelope, res *InputMediaPhoto) {
	out.Constructor = C_InputMediaPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaContact int64 = 1534117184

type poolInputMediaContact struct {
	pool sync.Pool
}

func (p poolInputMediaContact) Get() *InputMediaContact {
	x, ok := p.pool.Get().(*InputMediaContact)
	if !ok {
		return &InputMediaContact{}
	}
	x.VCard = ""
	return x
}

func (p poolInputMediaContact) Put(x *InputMediaContact) {
	p.pool.Put(x)
}

var PoolInputMediaContact = poolInputMediaContact{}

func ResultInputMediaContact(out *MessageEnvelope, res *InputMediaContact) {
	out.Constructor = C_InputMediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaUploadedDocument int64 = 870692909

type poolInputMediaUploadedDocument struct {
	pool sync.Pool
}

func (p poolInputMediaUploadedDocument) Get() *InputMediaUploadedDocument {
	x, ok := p.pool.Get().(*InputMediaUploadedDocument)
	if !ok {
		return &InputMediaUploadedDocument{}
	}
	x.Thumbnail = nil
	x.Stickers = x.Stickers[:0]
	x.Attributes = x.Attributes[:0]
	return x
}

func (p poolInputMediaUploadedDocument) Put(x *InputMediaUploadedDocument) {
	p.pool.Put(x)
}

var PoolInputMediaUploadedDocument = poolInputMediaUploadedDocument{}

func ResultInputMediaUploadedDocument(out *MessageEnvelope, res *InputMediaUploadedDocument) {
	out.Constructor = C_InputMediaUploadedDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaDocument int64 = 2258657627

type poolInputMediaDocument struct {
	pool sync.Pool
}

func (p poolInputMediaDocument) Get() *InputMediaDocument {
	x, ok := p.pool.Get().(*InputMediaDocument)
	if !ok {
		return &InputMediaDocument{}
	}
	return x
}

func (p poolInputMediaDocument) Put(x *InputMediaDocument) {
	p.pool.Put(x)
}

var PoolInputMediaDocument = poolInputMediaDocument{}

func ResultInputMediaDocument(out *MessageEnvelope, res *InputMediaDocument) {
	out.Constructor = C_InputMediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaGeoLocation int64 = 185664060

type poolInputMediaGeoLocation struct {
	pool sync.Pool
}

func (p poolInputMediaGeoLocation) Get() *InputMediaGeoLocation {
	x, ok := p.pool.Get().(*InputMediaGeoLocation)
	if !ok {
		return &InputMediaGeoLocation{}
	}
	return x
}

func (p poolInputMediaGeoLocation) Put(x *InputMediaGeoLocation) {
	p.pool.Put(x)
}

var PoolInputMediaGeoLocation = poolInputMediaGeoLocation{}

func ResultInputMediaGeoLocation(out *MessageEnvelope, res *InputMediaGeoLocation) {
	out.Constructor = C_InputMediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaPhoto int64 = 71894788

type poolMediaPhoto struct {
	pool sync.Pool
}

func (p poolMediaPhoto) Get() *MediaPhoto {
	x, ok := p.pool.Get().(*MediaPhoto)
	if !ok {
		return &MediaPhoto{}
	}
	return x
}

func (p poolMediaPhoto) Put(x *MediaPhoto) {
	p.pool.Put(x)
}

var PoolMediaPhoto = poolMediaPhoto{}

func ResultMediaPhoto(out *MessageEnvelope, res *MediaPhoto) {
	out.Constructor = C_MediaPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaDocument int64 = 2281620705

type poolMediaDocument struct {
	pool sync.Pool
}

func (p poolMediaDocument) Get() *MediaDocument {
	x, ok := p.pool.Get().(*MediaDocument)
	if !ok {
		return &MediaDocument{}
	}
	return x
}

func (p poolMediaDocument) Put(x *MediaDocument) {
	p.pool.Put(x)
}

var PoolMediaDocument = poolMediaDocument{}

func ResultMediaDocument(out *MessageEnvelope, res *MediaDocument) {
	out.Constructor = C_MediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaContact int64 = 3735320833

type poolMediaContact struct {
	pool sync.Pool
}

func (p poolMediaContact) Get() *MediaContact {
	x, ok := p.pool.Get().(*MediaContact)
	if !ok {
		return &MediaContact{}
	}
	x.VCard = ""
	return x
}

func (p poolMediaContact) Put(x *MediaContact) {
	p.pool.Put(x)
}

var PoolMediaContact = poolMediaContact{}

func ResultMediaContact(out *MessageEnvelope, res *MediaContact) {
	out.Constructor = C_MediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaGeoLocation int64 = 2625326500

type poolMediaGeoLocation struct {
	pool sync.Pool
}

func (p poolMediaGeoLocation) Get() *MediaGeoLocation {
	x, ok := p.pool.Get().(*MediaGeoLocation)
	if !ok {
		return &MediaGeoLocation{}
	}
	return x
}

func (p poolMediaGeoLocation) Put(x *MediaGeoLocation) {
	p.pool.Put(x)
}

var PoolMediaGeoLocation = poolMediaGeoLocation{}

func ResultMediaGeoLocation(out *MessageEnvelope, res *MediaGeoLocation) {
	out.Constructor = C_MediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaWebPage int64 = 148034084

type poolMediaWebPage struct {
	pool sync.Pool
}

func (p poolMediaWebPage) Get() *MediaWebPage {
	x, ok := p.pool.Get().(*MediaWebPage)
	if !ok {
		return &MediaWebPage{}
	}
	return x
}

func (p poolMediaWebPage) Put(x *MediaWebPage) {
	p.pool.Put(x)
}

var PoolMediaWebPage = poolMediaWebPage{}

func ResultMediaWebPage(out *MessageEnvelope, res *MediaWebPage) {
	out.Constructor = C_MediaWebPage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[309707708] = "DocumentAttributeAudio"
	ConstructorNames[1993289477] = "DocumentAttributeVideo"
	ConstructorNames[515862833] = "DocumentAttributePhoto"
	ConstructorNames[2227452062] = "DocumentAttributeFile"
	ConstructorNames[4146719643] = "DocumentAttribute"
	ConstructorNames[555739168] = "Document"
	ConstructorNames[849930963] = "InputMediaUploadedPhoto"
	ConstructorNames[2201579839] = "InputMediaPhoto"
	ConstructorNames[1534117184] = "InputMediaContact"
	ConstructorNames[870692909] = "InputMediaUploadedDocument"
	ConstructorNames[2258657627] = "InputMediaDocument"
	ConstructorNames[185664060] = "InputMediaGeoLocation"
	ConstructorNames[71894788] = "MediaPhoto"
	ConstructorNames[2281620705] = "MediaDocument"
	ConstructorNames[3735320833] = "MediaContact"
	ConstructorNames[2625326500] = "MediaGeoLocation"
	ConstructorNames[148034084] = "MediaWebPage"
}
