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

func (p *poolDocumentAttributeAudio) Get() *DocumentAttributeAudio {
	x, ok := p.pool.Get().(*DocumentAttributeAudio)
	if !ok {
		return &DocumentAttributeAudio{}
	}
	x.Waveform = nil
	return x
}

func (p *poolDocumentAttributeAudio) Put(x *DocumentAttributeAudio) {
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

func (p *poolDocumentAttributeVideo) Get() *DocumentAttributeVideo {
	x, ok := p.pool.Get().(*DocumentAttributeVideo)
	if !ok {
		return &DocumentAttributeVideo{}
	}
	return x
}

func (p *poolDocumentAttributeVideo) Put(x *DocumentAttributeVideo) {
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

func (p *poolDocumentAttributePhoto) Get() *DocumentAttributePhoto {
	x, ok := p.pool.Get().(*DocumentAttributePhoto)
	if !ok {
		return &DocumentAttributePhoto{}
	}
	return x
}

func (p *poolDocumentAttributePhoto) Put(x *DocumentAttributePhoto) {
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

func (p *poolDocumentAttributeFile) Get() *DocumentAttributeFile {
	x, ok := p.pool.Get().(*DocumentAttributeFile)
	if !ok {
		return &DocumentAttributeFile{}
	}
	return x
}

func (p *poolDocumentAttributeFile) Put(x *DocumentAttributeFile) {
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

func (p *poolDocumentAttribute) Get() *DocumentAttribute {
	x, ok := p.pool.Get().(*DocumentAttribute)
	if !ok {
		return &DocumentAttribute{}
	}
	x.Data = nil
	return x
}

func (p *poolDocumentAttribute) Put(x *DocumentAttribute) {
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

func (p *poolDocument) Get() *Document {
	x, ok := p.pool.Get().(*Document)
	if !ok {
		return &Document{}
	}
	x.Attributes = x.Attributes[:0]
	x.Thumbnail = nil
	x.MD5Checksum = ""
	return x
}

func (p *poolDocument) Put(x *Document) {
	p.pool.Put(x)
}

var PoolDocument = poolDocument{}

func ResultDocument(out *MessageEnvelope, res *Document) {
	out.Constructor = C_Document
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaInvoice int64 = 2272736316

type poolInputMediaInvoice struct {
	pool sync.Pool
}

func (p *poolInputMediaInvoice) Get() *InputMediaInvoice {
	x, ok := p.pool.Get().(*InputMediaInvoice)
	if !ok {
		return &InputMediaInvoice{}
	}
	return x
}

func (p *poolInputMediaInvoice) Put(x *InputMediaInvoice) {
	p.pool.Put(x)
}

var PoolInputMediaInvoice = poolInputMediaInvoice{}

func ResultInputMediaInvoice(out *MessageEnvelope, res *InputMediaInvoice) {
	out.Constructor = C_InputMediaInvoice
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaInvoice int64 = 44271741

type poolMediaInvoice struct {
	pool sync.Pool
}

func (p *poolMediaInvoice) Get() *MediaInvoice {
	x, ok := p.pool.Get().(*MediaInvoice)
	if !ok {
		return &MediaInvoice{}
	}
	return x
}

func (p *poolMediaInvoice) Put(x *MediaInvoice) {
	p.pool.Put(x)
}

var PoolMediaInvoice = poolMediaInvoice{}

func ResultMediaInvoice(out *MessageEnvelope, res *MediaInvoice) {
	out.Constructor = C_MediaInvoice
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaWebDocument int64 = 3529450013

type poolInputMediaWebDocument struct {
	pool sync.Pool
}

func (p *poolInputMediaWebDocument) Get() *InputMediaWebDocument {
	x, ok := p.pool.Get().(*InputMediaWebDocument)
	if !ok {
		return &InputMediaWebDocument{}
	}
	return x
}

func (p *poolInputMediaWebDocument) Put(x *InputMediaWebDocument) {
	p.pool.Put(x)
}

var PoolInputMediaWebDocument = poolInputMediaWebDocument{}

func ResultInputMediaWebDocument(out *MessageEnvelope, res *InputMediaWebDocument) {
	out.Constructor = C_InputMediaWebDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaWebDocument int64 = 1161129349

type poolMediaWebDocument struct {
	pool sync.Pool
}

func (p *poolMediaWebDocument) Get() *MediaWebDocument {
	x, ok := p.pool.Get().(*MediaWebDocument)
	if !ok {
		return &MediaWebDocument{}
	}
	x.SiteTitle = ""
	x.Title = ""
	x.EmbedUrl = ""
	x.EmbedType = ""
	x.EmbedHeight = 0
	x.EmbedWidth = 0
	x.Duration = 0
	x.Author = ""
	x.Doc = nil
	return x
}

func (p *poolMediaWebDocument) Put(x *MediaWebDocument) {
	p.pool.Put(x)
}

var PoolMediaWebDocument = poolMediaWebDocument{}

func ResultMediaWebDocument(out *MessageEnvelope, res *MediaWebDocument) {
	out.Constructor = C_MediaWebDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaWebPage int64 = 148034084

type poolMediaWebPage struct {
	pool sync.Pool
}

func (p *poolMediaWebPage) Get() *MediaWebPage {
	x, ok := p.pool.Get().(*MediaWebPage)
	if !ok {
		return &MediaWebPage{}
	}
	return x
}

func (p *poolMediaWebPage) Put(x *MediaWebPage) {
	p.pool.Put(x)
}

var PoolMediaWebPage = poolMediaWebPage{}

func ResultMediaWebPage(out *MessageEnvelope, res *MediaWebPage) {
	out.Constructor = C_MediaWebPage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaContact int64 = 1534117184

type poolInputMediaContact struct {
	pool sync.Pool
}

func (p *poolInputMediaContact) Get() *InputMediaContact {
	x, ok := p.pool.Get().(*InputMediaContact)
	if !ok {
		return &InputMediaContact{}
	}
	x.VCard = ""
	return x
}

func (p *poolInputMediaContact) Put(x *InputMediaContact) {
	p.pool.Put(x)
}

var PoolInputMediaContact = poolInputMediaContact{}

func ResultInputMediaContact(out *MessageEnvelope, res *InputMediaContact) {
	out.Constructor = C_InputMediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaContact int64 = 3735320833

type poolMediaContact struct {
	pool sync.Pool
}

func (p *poolMediaContact) Get() *MediaContact {
	x, ok := p.pool.Get().(*MediaContact)
	if !ok {
		return &MediaContact{}
	}
	x.VCard = ""
	return x
}

func (p *poolMediaContact) Put(x *MediaContact) {
	p.pool.Put(x)
}

var PoolMediaContact = poolMediaContact{}

func ResultMediaContact(out *MessageEnvelope, res *MediaContact) {
	out.Constructor = C_MediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaUploadedDocument int64 = 870692909

type poolInputMediaUploadedDocument struct {
	pool sync.Pool
}

func (p *poolInputMediaUploadedDocument) Get() *InputMediaUploadedDocument {
	x, ok := p.pool.Get().(*InputMediaUploadedDocument)
	if !ok {
		return &InputMediaUploadedDocument{}
	}
	x.Thumbnail = nil
	x.Stickers = x.Stickers[:0]
	x.Attributes = x.Attributes[:0]
	x.Entities = x.Entities[:0]
	return x
}

func (p *poolInputMediaUploadedDocument) Put(x *InputMediaUploadedDocument) {
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

func (p *poolInputMediaDocument) Get() *InputMediaDocument {
	x, ok := p.pool.Get().(*InputMediaDocument)
	if !ok {
		return &InputMediaDocument{}
	}
	x.Entities = x.Entities[:0]
	x.Thumbnail = nil
	x.Attributes = x.Attributes[:0]
	return x
}

func (p *poolInputMediaDocument) Put(x *InputMediaDocument) {
	p.pool.Put(x)
}

var PoolInputMediaDocument = poolInputMediaDocument{}

func ResultInputMediaDocument(out *MessageEnvelope, res *InputMediaDocument) {
	out.Constructor = C_InputMediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaDocument int64 = 2281620705

type poolMediaDocument struct {
	pool sync.Pool
}

func (p *poolMediaDocument) Get() *MediaDocument {
	x, ok := p.pool.Get().(*MediaDocument)
	if !ok {
		return &MediaDocument{}
	}
	x.Entities = x.Entities[:0]
	return x
}

func (p *poolMediaDocument) Put(x *MediaDocument) {
	p.pool.Put(x)
}

var PoolMediaDocument = poolMediaDocument{}

func ResultMediaDocument(out *MessageEnvelope, res *MediaDocument) {
	out.Constructor = C_MediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaGeoLocation int64 = 185664060

type poolInputMediaGeoLocation struct {
	pool sync.Pool
}

func (p *poolInputMediaGeoLocation) Get() *InputMediaGeoLocation {
	x, ok := p.pool.Get().(*InputMediaGeoLocation)
	if !ok {
		return &InputMediaGeoLocation{}
	}
	return x
}

func (p *poolInputMediaGeoLocation) Put(x *InputMediaGeoLocation) {
	p.pool.Put(x)
}

var PoolInputMediaGeoLocation = poolInputMediaGeoLocation{}

func ResultInputMediaGeoLocation(out *MessageEnvelope, res *InputMediaGeoLocation) {
	out.Constructor = C_InputMediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaGeoLocation int64 = 2625326500

type poolMediaGeoLocation struct {
	pool sync.Pool
}

func (p *poolMediaGeoLocation) Get() *MediaGeoLocation {
	x, ok := p.pool.Get().(*MediaGeoLocation)
	if !ok {
		return &MediaGeoLocation{}
	}
	return x
}

func (p *poolMediaGeoLocation) Put(x *MediaGeoLocation) {
	p.pool.Put(x)
}

var PoolMediaGeoLocation = poolMediaGeoLocation{}

func ResultMediaGeoLocation(out *MessageEnvelope, res *MediaGeoLocation) {
	out.Constructor = C_MediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputMediaPoll int64 = 3633337678

type poolInputMediaPoll struct {
	pool sync.Pool
}

func (p *poolInputMediaPoll) Get() *InputMediaPoll {
	x, ok := p.pool.Get().(*InputMediaPoll)
	if !ok {
		return &InputMediaPoll{}
	}
	return x
}

func (p *poolInputMediaPoll) Put(x *InputMediaPoll) {
	p.pool.Put(x)
}

var PoolInputMediaPoll = poolInputMediaPoll{}

func ResultInputMediaPoll(out *MessageEnvelope, res *InputMediaPoll) {
	out.Constructor = C_InputMediaPoll
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaPoll int64 = 2688924895

type poolMediaPoll struct {
	pool sync.Pool
}

func (p *poolMediaPoll) Get() *MediaPoll {
	x, ok := p.pool.Get().(*MediaPoll)
	if !ok {
		return &MediaPoll{}
	}
	x.Closed = false
	x.PublicVoters = false
	x.MultiChoice = false
	x.Quiz = false
	x.Answers = x.Answers[:0]
	return x
}

func (p *poolMediaPoll) Put(x *MediaPoll) {
	p.pool.Put(x)
}

var PoolMediaPoll = poolMediaPoll{}

func ResultMediaPoll(out *MessageEnvelope, res *MediaPoll) {
	out.Constructor = C_MediaPoll
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_PollAnswer int64 = 2124799390

type poolPollAnswer struct {
	pool sync.Pool
}

func (p *poolPollAnswer) Get() *PollAnswer {
	x, ok := p.pool.Get().(*PollAnswer)
	if !ok {
		return &PollAnswer{}
	}
	return x
}

func (p *poolPollAnswer) Put(x *PollAnswer) {
	p.pool.Put(x)
}

var PoolPollAnswer = poolPollAnswer{}

func ResultPollAnswer(out *MessageEnvelope, res *PollAnswer) {
	out.Constructor = C_PollAnswer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_PollResults int64 = 3283416711

type poolPollResults struct {
	pool sync.Pool
}

func (p *poolPollResults) Get() *PollResults {
	x, ok := p.pool.Get().(*PollResults)
	if !ok {
		return &PollResults{}
	}
	x.Results = x.Results[:0]
	return x
}

func (p *poolPollResults) Put(x *PollResults) {
	p.pool.Put(x)
}

var PoolPollResults = poolPollResults{}

func ResultPollResults(out *MessageEnvelope, res *PollResults) {
	out.Constructor = C_PollResults
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_PollAnswerVoters int64 = 2095107985

type poolPollAnswerVoters struct {
	pool sync.Pool
}

func (p *poolPollAnswerVoters) Get() *PollAnswerVoters {
	x, ok := p.pool.Get().(*PollAnswerVoters)
	if !ok {
		return &PollAnswerVoters{}
	}
	return x
}

func (p *poolPollAnswerVoters) Put(x *PollAnswerVoters) {
	p.pool.Put(x)
}

var PoolPollAnswerVoters = poolPollAnswerVoters{}

func ResultPollAnswerVoters(out *MessageEnvelope, res *PollAnswerVoters) {
	out.Constructor = C_PollAnswerVoters
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
	ConstructorNames[2272736316] = "InputMediaInvoice"
	ConstructorNames[44271741] = "MediaInvoice"
	ConstructorNames[3529450013] = "InputMediaWebDocument"
	ConstructorNames[1161129349] = "MediaWebDocument"
	ConstructorNames[148034084] = "MediaWebPage"
	ConstructorNames[1534117184] = "InputMediaContact"
	ConstructorNames[3735320833] = "MediaContact"
	ConstructorNames[870692909] = "InputMediaUploadedDocument"
	ConstructorNames[2258657627] = "InputMediaDocument"
	ConstructorNames[2281620705] = "MediaDocument"
	ConstructorNames[185664060] = "InputMediaGeoLocation"
	ConstructorNames[2625326500] = "MediaGeoLocation"
	ConstructorNames[3633337678] = "InputMediaPoll"
	ConstructorNames[2688924895] = "MediaPoll"
	ConstructorNames[2124799390] = "PollAnswer"
	ConstructorNames[3283416711] = "PollResults"
	ConstructorNames[2095107985] = "PollAnswerVoters"
}
