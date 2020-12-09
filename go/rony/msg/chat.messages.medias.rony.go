package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_DocumentAttributeAudio int64 = 309707708

type poolDocumentAttributeAudio struct {
	pool sync.Pool
}

func (p *poolDocumentAttributeAudio) Get() *DocumentAttributeAudio {
	x, ok := p.pool.Get().(*DocumentAttributeAudio)
	if !ok {
		return &DocumentAttributeAudio{}
	}
	return x
}

func (p *poolDocumentAttributeAudio) Put(x *DocumentAttributeAudio) {
	x.Voice = false
	x.Duration = 0
	x.Title = ""
	x.Performer = ""
	x.Album = ""
	x.Waveform = x.Waveform[:0]
	p.pool.Put(x)
}

var PoolDocumentAttributeAudio = poolDocumentAttributeAudio{}

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
	x.Width = 0
	x.Height = 0
	x.Duration = 0
	x.Round = false
	p.pool.Put(x)
}

var PoolDocumentAttributeVideo = poolDocumentAttributeVideo{}

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
	x.Width = 0
	x.Height = 0
	p.pool.Put(x)
}

var PoolDocumentAttributePhoto = poolDocumentAttributePhoto{}

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
	x.Filename = ""
	p.pool.Put(x)
}

var PoolDocumentAttributeFile = poolDocumentAttributeFile{}

const C_DocumentAttributeAnimated int64 = 1040723836

type poolDocumentAttributeAnimated struct {
	pool sync.Pool
}

func (p *poolDocumentAttributeAnimated) Get() *DocumentAttributeAnimated {
	x, ok := p.pool.Get().(*DocumentAttributeAnimated)
	if !ok {
		return &DocumentAttributeAnimated{}
	}
	return x
}

func (p *poolDocumentAttributeAnimated) Put(x *DocumentAttributeAnimated) {
	x.Sound = false
	p.pool.Put(x)
}

var PoolDocumentAttributeAnimated = poolDocumentAttributeAnimated{}

const C_DocumentAttribute int64 = 4146719643

type poolDocumentAttribute struct {
	pool sync.Pool
}

func (p *poolDocumentAttribute) Get() *DocumentAttribute {
	x, ok := p.pool.Get().(*DocumentAttribute)
	if !ok {
		return &DocumentAttribute{}
	}
	return x
}

func (p *poolDocumentAttribute) Put(x *DocumentAttribute) {
	x.Type = 0
	x.Data = x.Data[:0]
	p.pool.Put(x)
}

var PoolDocumentAttribute = poolDocumentAttribute{}

const C_Document int64 = 555739168

type poolDocument struct {
	pool sync.Pool
}

func (p *poolDocument) Get() *Document {
	x, ok := p.pool.Get().(*Document)
	if !ok {
		return &Document{}
	}
	return x
}

func (p *poolDocument) Put(x *Document) {
	x.ID = 0
	x.AccessHash = 0
	x.Date = 0
	x.MimeType = ""
	x.FileSize = 0
	x.Version = 0
	x.ClusterID = 0
	x.Attributes = x.Attributes[:0]
	if x.Thumbnail != nil {
		PoolFileLocation.Put(x.Thumbnail)
		x.Thumbnail = nil
	}
	x.MD5Checksum = ""
	x.TinyThumbnail = x.TinyThumbnail[:0]
	p.pool.Put(x)
}

var PoolDocument = poolDocument{}

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
	x.Url = ""
	x.DocSize = 0
	x.MimeType = ""
	x.Attributes = x.Attributes[:0]
	p.pool.Put(x)
}

var PoolInputMediaWebDocument = poolInputMediaWebDocument{}

const C_MediaWebDocument int64 = 1161129349

type poolMediaWebDocument struct {
	pool sync.Pool
}

func (p *poolMediaWebDocument) Get() *MediaWebDocument {
	x, ok := p.pool.Get().(*MediaWebDocument)
	if !ok {
		return &MediaWebDocument{}
	}
	return x
}

func (p *poolMediaWebDocument) Put(x *MediaWebDocument) {
	x.Url = ""
	x.DocSize = 0
	x.MimeType = ""
	x.Attributes = x.Attributes[:0]
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolMediaWebDocument = poolMediaWebDocument{}

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

const C_InputMediaContact int64 = 1534117184

type poolInputMediaContact struct {
	pool sync.Pool
}

func (p *poolInputMediaContact) Get() *InputMediaContact {
	x, ok := p.pool.Get().(*InputMediaContact)
	if !ok {
		return &InputMediaContact{}
	}
	return x
}

func (p *poolInputMediaContact) Put(x *InputMediaContact) {
	x.Phone = ""
	x.FirstName = ""
	x.LastName = ""
	x.VCard = ""
	p.pool.Put(x)
}

var PoolInputMediaContact = poolInputMediaContact{}

const C_MediaContact int64 = 3735320833

type poolMediaContact struct {
	pool sync.Pool
}

func (p *poolMediaContact) Get() *MediaContact {
	x, ok := p.pool.Get().(*MediaContact)
	if !ok {
		return &MediaContact{}
	}
	return x
}

func (p *poolMediaContact) Put(x *MediaContact) {
	x.Phone = ""
	x.FirstName = ""
	x.LastName = ""
	x.VCard = ""
	p.pool.Put(x)
}

var PoolMediaContact = poolMediaContact{}

const C_InputMediaUploadedDocument int64 = 870692909

type poolInputMediaUploadedDocument struct {
	pool sync.Pool
}

func (p *poolInputMediaUploadedDocument) Get() *InputMediaUploadedDocument {
	x, ok := p.pool.Get().(*InputMediaUploadedDocument)
	if !ok {
		return &InputMediaUploadedDocument{}
	}
	return x
}

func (p *poolInputMediaUploadedDocument) Put(x *InputMediaUploadedDocument) {
	if x.File != nil {
		PoolInputFile.Put(x.File)
		x.File = nil
	}
	if x.Thumbnail != nil {
		PoolInputFile.Put(x.Thumbnail)
		x.Thumbnail = nil
	}
	x.MimeType = ""
	x.Caption = ""
	x.Stickers = x.Stickers[:0]
	x.Attributes = x.Attributes[:0]
	x.Entities = x.Entities[:0]
	x.TinyThumbnail = x.TinyThumbnail[:0]
	p.pool.Put(x)
}

var PoolInputMediaUploadedDocument = poolInputMediaUploadedDocument{}

const C_InputMediaDocument int64 = 2258657627

type poolInputMediaDocument struct {
	pool sync.Pool
}

func (p *poolInputMediaDocument) Get() *InputMediaDocument {
	x, ok := p.pool.Get().(*InputMediaDocument)
	if !ok {
		return &InputMediaDocument{}
	}
	return x
}

func (p *poolInputMediaDocument) Put(x *InputMediaDocument) {
	x.Caption = ""
	if x.Document != nil {
		PoolInputDocument.Put(x.Document)
		x.Document = nil
	}
	x.Entities = x.Entities[:0]
	if x.Thumbnail != nil {
		PoolInputFile.Put(x.Thumbnail)
		x.Thumbnail = nil
	}
	x.Attributes = x.Attributes[:0]
	x.TinyThumbnail = x.TinyThumbnail[:0]
	p.pool.Put(x)
}

var PoolInputMediaDocument = poolInputMediaDocument{}

const C_InputMediaMessageDocument int64 = 3638653559

type poolInputMediaMessageDocument struct {
	pool sync.Pool
}

func (p *poolInputMediaMessageDocument) Get() *InputMediaMessageDocument {
	x, ok := p.pool.Get().(*InputMediaMessageDocument)
	if !ok {
		return &InputMediaMessageDocument{}
	}
	return x
}

func (p *poolInputMediaMessageDocument) Put(x *InputMediaMessageDocument) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MessageID = 0
	x.Caption = ""
	x.Entities = x.Entities[:0]
	p.pool.Put(x)
}

var PoolInputMediaMessageDocument = poolInputMediaMessageDocument{}

const C_MediaDocument int64 = 2281620705

type poolMediaDocument struct {
	pool sync.Pool
}

func (p *poolMediaDocument) Get() *MediaDocument {
	x, ok := p.pool.Get().(*MediaDocument)
	if !ok {
		return &MediaDocument{}
	}
	return x
}

func (p *poolMediaDocument) Put(x *MediaDocument) {
	x.Caption = ""
	x.TTLinSeconds = 0
	if x.Doc != nil {
		PoolDocument.Put(x.Doc)
		x.Doc = nil
	}
	x.Entities = x.Entities[:0]
	p.pool.Put(x)
}

var PoolMediaDocument = poolMediaDocument{}

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
	x.Lat = 0
	x.Long = 0
	x.Caption = ""
	x.Entities = x.Entities[:0]
	p.pool.Put(x)
}

var PoolInputMediaGeoLocation = poolInputMediaGeoLocation{}

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
	x.Lat = 0
	x.Long = 0
	x.Caption = ""
	x.Entities = x.Entities[:0]
	p.pool.Put(x)
}

var PoolMediaGeoLocation = poolMediaGeoLocation{}

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
	if x.Poll != nil {
		PoolMediaPoll.Put(x.Poll)
		x.Poll = nil
	}
	p.pool.Put(x)
}

var PoolInputMediaPoll = poolInputMediaPoll{}

const C_MediaPoll int64 = 2688924895

type poolMediaPoll struct {
	pool sync.Pool
}

func (p *poolMediaPoll) Get() *MediaPoll {
	x, ok := p.pool.Get().(*MediaPoll)
	if !ok {
		return &MediaPoll{}
	}
	return x
}

func (p *poolMediaPoll) Put(x *MediaPoll) {
	x.ID = 0
	x.Closed = false
	x.PublicVoters = false
	x.MultiChoice = false
	x.Quiz = false
	x.Question = ""
	x.Answers = x.Answers[:0]
	p.pool.Put(x)
}

var PoolMediaPoll = poolMediaPoll{}

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
	x.Text = ""
	x.Option = x.Option[:0]
	p.pool.Put(x)
}

var PoolPollAnswer = poolPollAnswer{}

const C_PollResults int64 = 3283416711

type poolPollResults struct {
	pool sync.Pool
}

func (p *poolPollResults) Get() *PollResults {
	x, ok := p.pool.Get().(*PollResults)
	if !ok {
		return &PollResults{}
	}
	return x
}

func (p *poolPollResults) Put(x *PollResults) {
	x.Results = x.Results[:0]
	x.TotalVoters = 0
	p.pool.Put(x)
}

var PoolPollResults = poolPollResults{}

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
	x.Chosen = false
	x.Correct = false
	x.Option = x.Option[:0]
	x.Voters = 0
	p.pool.Put(x)
}

var PoolPollAnswerVoters = poolPollAnswerVoters{}

func init() {
	registry.RegisterConstructor(309707708, "DocumentAttributeAudio")
	registry.RegisterConstructor(1993289477, "DocumentAttributeVideo")
	registry.RegisterConstructor(515862833, "DocumentAttributePhoto")
	registry.RegisterConstructor(2227452062, "DocumentAttributeFile")
	registry.RegisterConstructor(1040723836, "DocumentAttributeAnimated")
	registry.RegisterConstructor(4146719643, "DocumentAttribute")
	registry.RegisterConstructor(555739168, "Document")
	registry.RegisterConstructor(3529450013, "InputMediaWebDocument")
	registry.RegisterConstructor(1161129349, "MediaWebDocument")
	registry.RegisterConstructor(148034084, "MediaWebPage")
	registry.RegisterConstructor(1534117184, "InputMediaContact")
	registry.RegisterConstructor(3735320833, "MediaContact")
	registry.RegisterConstructor(870692909, "InputMediaUploadedDocument")
	registry.RegisterConstructor(2258657627, "InputMediaDocument")
	registry.RegisterConstructor(3638653559, "InputMediaMessageDocument")
	registry.RegisterConstructor(2281620705, "MediaDocument")
	registry.RegisterConstructor(185664060, "InputMediaGeoLocation")
	registry.RegisterConstructor(2625326500, "MediaGeoLocation")
	registry.RegisterConstructor(3633337678, "InputMediaPoll")
	registry.RegisterConstructor(2688924895, "MediaPoll")
	registry.RegisterConstructor(2124799390, "PollAnswer")
	registry.RegisterConstructor(3283416711, "PollResults")
	registry.RegisterConstructor(2095107985, "PollAnswerVoters")
}

func (x *DocumentAttributeAudio) DeepCopy(z *DocumentAttributeAudio) {
	z.Voice = x.Voice
	z.Duration = x.Duration
	z.Title = x.Title
	z.Performer = x.Performer
	z.Album = x.Album
	z.Waveform = append(z.Waveform[:0], x.Waveform...)
}

func (x *DocumentAttributeVideo) DeepCopy(z *DocumentAttributeVideo) {
	z.Width = x.Width
	z.Height = x.Height
	z.Duration = x.Duration
	z.Round = x.Round
}

func (x *DocumentAttributePhoto) DeepCopy(z *DocumentAttributePhoto) {
	z.Width = x.Width
	z.Height = x.Height
}

func (x *DocumentAttributeFile) DeepCopy(z *DocumentAttributeFile) {
	z.Filename = x.Filename
}

func (x *DocumentAttributeAnimated) DeepCopy(z *DocumentAttributeAnimated) {
	z.Sound = x.Sound
}

func (x *DocumentAttribute) DeepCopy(z *DocumentAttribute) {
	z.Type = x.Type
	z.Data = append(z.Data[:0], x.Data...)
}

func (x *Document) DeepCopy(z *Document) {
	z.ID = x.ID
	z.AccessHash = x.AccessHash
	z.Date = x.Date
	z.MimeType = x.MimeType
	z.FileSize = x.FileSize
	z.Version = x.Version
	z.ClusterID = x.ClusterID
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
	if x.Thumbnail != nil {
		z.Thumbnail = PoolFileLocation.Get()
		x.Thumbnail.DeepCopy(z.Thumbnail)
	}
	z.MD5Checksum = x.MD5Checksum
	z.TinyThumbnail = append(z.TinyThumbnail[:0], x.TinyThumbnail...)
}

func (x *InputMediaWebDocument) DeepCopy(z *InputMediaWebDocument) {
	z.Url = x.Url
	z.DocSize = x.DocSize
	z.MimeType = x.MimeType
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
}

func (x *MediaWebDocument) DeepCopy(z *MediaWebDocument) {
	z.Url = x.Url
	z.DocSize = x.DocSize
	z.MimeType = x.MimeType
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
	z.AccessHash = x.AccessHash
}

func (x *MediaWebPage) DeepCopy(z *MediaWebPage) {
}

func (x *InputMediaContact) DeepCopy(z *InputMediaContact) {
	z.Phone = x.Phone
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.VCard = x.VCard
}

func (x *MediaContact) DeepCopy(z *MediaContact) {
	z.Phone = x.Phone
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.VCard = x.VCard
}

func (x *InputMediaUploadedDocument) DeepCopy(z *InputMediaUploadedDocument) {
	if x.File != nil {
		z.File = PoolInputFile.Get()
		x.File.DeepCopy(z.File)
	}
	if x.Thumbnail != nil {
		z.Thumbnail = PoolInputFile.Get()
		x.Thumbnail.DeepCopy(z.Thumbnail)
	}
	z.MimeType = x.MimeType
	z.Caption = x.Caption
	for idx := range x.Stickers {
		if x.Stickers[idx] != nil {
			xx := PoolInputDocument.Get()
			x.Stickers[idx].DeepCopy(xx)
			z.Stickers = append(z.Stickers, xx)
		}
	}
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.TinyThumbnail = append(z.TinyThumbnail[:0], x.TinyThumbnail...)
}

func (x *InputMediaDocument) DeepCopy(z *InputMediaDocument) {
	z.Caption = x.Caption
	if x.Document != nil {
		z.Document = PoolInputDocument.Get()
		x.Document.DeepCopy(z.Document)
	}
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	if x.Thumbnail != nil {
		z.Thumbnail = PoolInputFile.Get()
		x.Thumbnail.DeepCopy(z.Thumbnail)
	}
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
	z.TinyThumbnail = append(z.TinyThumbnail[:0], x.TinyThumbnail...)
}

func (x *InputMediaMessageDocument) DeepCopy(z *InputMediaMessageDocument) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MessageID = x.MessageID
	z.Caption = x.Caption
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
}

func (x *MediaDocument) DeepCopy(z *MediaDocument) {
	z.Caption = x.Caption
	z.TTLinSeconds = x.TTLinSeconds
	if x.Doc != nil {
		z.Doc = PoolDocument.Get()
		x.Doc.DeepCopy(z.Doc)
	}
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
}

func (x *InputMediaGeoLocation) DeepCopy(z *InputMediaGeoLocation) {
	z.Lat = x.Lat
	z.Long = x.Long
	z.Caption = x.Caption
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
}

func (x *MediaGeoLocation) DeepCopy(z *MediaGeoLocation) {
	z.Lat = x.Lat
	z.Long = x.Long
	z.Caption = x.Caption
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
}

func (x *InputMediaPoll) DeepCopy(z *InputMediaPoll) {
	if x.Poll != nil {
		z.Poll = PoolMediaPoll.Get()
		x.Poll.DeepCopy(z.Poll)
	}
}

func (x *MediaPoll) DeepCopy(z *MediaPoll) {
	z.ID = x.ID
	z.Closed = x.Closed
	z.PublicVoters = x.PublicVoters
	z.MultiChoice = x.MultiChoice
	z.Quiz = x.Quiz
	z.Question = x.Question
	for idx := range x.Answers {
		if x.Answers[idx] != nil {
			xx := PoolPollAnswer.Get()
			x.Answers[idx].DeepCopy(xx)
			z.Answers = append(z.Answers, xx)
		}
	}
}

func (x *PollAnswer) DeepCopy(z *PollAnswer) {
	z.Text = x.Text
	z.Option = append(z.Option[:0], x.Option...)
}

func (x *PollResults) DeepCopy(z *PollResults) {
	for idx := range x.Results {
		if x.Results[idx] != nil {
			xx := PoolPollAnswerVoters.Get()
			x.Results[idx].DeepCopy(xx)
			z.Results = append(z.Results, xx)
		}
	}
	z.TotalVoters = x.TotalVoters
}

func (x *PollAnswerVoters) DeepCopy(z *PollAnswerVoters) {
	z.Chosen = x.Chosen
	z.Correct = x.Correct
	z.Option = append(z.Option[:0], x.Option...)
	z.Voters = x.Voters
}
