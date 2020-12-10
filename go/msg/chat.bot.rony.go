package msg

import (
	fmt "fmt"
	rony "github.com/ronaksoft/rony"
	edge "github.com/ronaksoft/rony/edge"
	edgec "github.com/ronaksoft/rony/edgec"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_BotStart int64 = 3448758942

type poolBotStart struct {
	pool sync.Pool
}

func (p *poolBotStart) Get() *BotStart {
	x, ok := p.pool.Get().(*BotStart)
	if !ok {
		return &BotStart{}
	}
	return x
}

func (p *poolBotStart) Put(x *BotStart) {
	if x.Bot != nil {
		PoolInputPeer.Put(x.Bot)
		x.Bot = nil
	}
	x.RandomID = 0
	x.StartParam = ""
	p.pool.Put(x)
}

var PoolBotStart = poolBotStart{}

const C_BotRecall int64 = 567464755

type poolBotRecall struct {
	pool sync.Pool
}

func (p *poolBotRecall) Get() *BotRecall {
	x, ok := p.pool.Get().(*BotRecall)
	if !ok {
		return &BotRecall{}
	}
	return x
}

func (p *poolBotRecall) Put(x *BotRecall) {
	x.Version = 0
	p.pool.Put(x)
}

var PoolBotRecall = poolBotRecall{}

const C_BotSetInfo int64 = 1487199122

type poolBotSetInfo struct {
	pool sync.Pool
}

func (p *poolBotSetInfo) Get() *BotSetInfo {
	x, ok := p.pool.Get().(*BotSetInfo)
	if !ok {
		return &BotSetInfo{}
	}
	return x
}

func (p *poolBotSetInfo) Put(x *BotSetInfo) {
	x.BotID = 0
	x.RandomID = 0
	x.Owner = 0
	x.BotCommands = x.BotCommands[:0]
	x.Description = ""
	x.InlinePlaceholder = ""
	x.InlineGeo = false
	p.pool.Put(x)
}

var PoolBotSetInfo = poolBotSetInfo{}

const C_BotGet int64 = 3431546487

type poolBotGet struct {
	pool sync.Pool
}

func (p *poolBotGet) Get() *BotGet {
	x, ok := p.pool.Get().(*BotGet)
	if !ok {
		return &BotGet{}
	}
	return x
}

func (p *poolBotGet) Put(x *BotGet) {
	x.UserID = 0
	x.Limit = 0
	p.pool.Put(x)
}

var PoolBotGet = poolBotGet{}

const C_BotSendMessage int64 = 3116228191

type poolBotSendMessage struct {
	pool sync.Pool
}

func (p *poolBotSendMessage) Get() *BotSendMessage {
	x, ok := p.pool.Get().(*BotSendMessage)
	if !ok {
		return &BotSendMessage{}
	}
	return x
}

func (p *poolBotSendMessage) Put(x *BotSendMessage) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Body = ""
	x.ReplyTo = 0
	x.ClearDraft = false
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	p.pool.Put(x)
}

var PoolBotSendMessage = poolBotSendMessage{}

const C_BotEditMessage int64 = 149350155

type poolBotEditMessage struct {
	pool sync.Pool
}

func (p *poolBotEditMessage) Get() *BotEditMessage {
	x, ok := p.pool.Get().(*BotEditMessage)
	if !ok {
		return &BotEditMessage{}
	}
	return x
}

func (p *poolBotEditMessage) Put(x *BotEditMessage) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Body = ""
	x.MessageID = 0
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	p.pool.Put(x)
}

var PoolBotEditMessage = poolBotEditMessage{}

const C_BotSendMedia int64 = 2100881420

type poolBotSendMedia struct {
	pool sync.Pool
}

func (p *poolBotSendMedia) Get() *BotSendMedia {
	x, ok := p.pool.Get().(*BotSendMedia)
	if !ok {
		return &BotSendMedia{}
	}
	return x
}

func (p *poolBotSendMedia) Put(x *BotSendMedia) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MediaType = 0
	x.MediaData = x.MediaData[:0]
	x.ReplyTo = 0
	p.pool.Put(x)
}

var PoolBotSendMedia = poolBotSendMedia{}

const C_BotSaveFilePart int64 = 598724192

type poolBotSaveFilePart struct {
	pool sync.Pool
}

func (p *poolBotSaveFilePart) Get() *BotSaveFilePart {
	x, ok := p.pool.Get().(*BotSaveFilePart)
	if !ok {
		return &BotSaveFilePart{}
	}
	return x
}

func (p *poolBotSaveFilePart) Put(x *BotSaveFilePart) {
	x.FileID = 0
	x.PartID = 0
	x.TotalParts = 0
	x.Bytes = x.Bytes[:0]
	p.pool.Put(x)
}

var PoolBotSaveFilePart = poolBotSaveFilePart{}

const C_BotUpdateProfile int64 = 1624560842

type poolBotUpdateProfile struct {
	pool sync.Pool
}

func (p *poolBotUpdateProfile) Get() *BotUpdateProfile {
	x, ok := p.pool.Get().(*BotUpdateProfile)
	if !ok {
		return &BotUpdateProfile{}
	}
	return x
}

func (p *poolBotUpdateProfile) Put(x *BotUpdateProfile) {
	x.BotID = 0
	x.Name = ""
	x.Bio = ""
	p.pool.Put(x)
}

var PoolBotUpdateProfile = poolBotUpdateProfile{}

const C_BotUpdatePhoto int64 = 4201069063

type poolBotUpdatePhoto struct {
	pool sync.Pool
}

func (p *poolBotUpdatePhoto) Get() *BotUpdatePhoto {
	x, ok := p.pool.Get().(*BotUpdatePhoto)
	if !ok {
		return &BotUpdatePhoto{}
	}
	return x
}

func (p *poolBotUpdatePhoto) Put(x *BotUpdatePhoto) {
	if x.File != nil {
		PoolInputFileLocation.Put(x.File)
		x.File = nil
	}
	x.BotID = 0
	p.pool.Put(x)
}

var PoolBotUpdatePhoto = poolBotUpdatePhoto{}

const C_BotRevokeToken int64 = 1601295593

type poolBotRevokeToken struct {
	pool sync.Pool
}

func (p *poolBotRevokeToken) Get() *BotRevokeToken {
	x, ok := p.pool.Get().(*BotRevokeToken)
	if !ok {
		return &BotRevokeToken{}
	}
	return x
}

func (p *poolBotRevokeToken) Put(x *BotRevokeToken) {
	x.BotID = 0
	x.GetNew = false
	p.pool.Put(x)
}

var PoolBotRevokeToken = poolBotRevokeToken{}

const C_BotDeleteMessage int64 = 423422518

type poolBotDeleteMessage struct {
	pool sync.Pool
}

func (p *poolBotDeleteMessage) Get() *BotDeleteMessage {
	x, ok := p.pool.Get().(*BotDeleteMessage)
	if !ok {
		return &BotDeleteMessage{}
	}
	return x
}

func (p *poolBotDeleteMessage) Put(x *BotDeleteMessage) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MessageIDs = x.MessageIDs[:0]
	p.pool.Put(x)
}

var PoolBotDeleteMessage = poolBotDeleteMessage{}

const C_BotSetCallbackAnswer int64 = 759509107

type poolBotSetCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotSetCallbackAnswer) Get() *BotSetCallbackAnswer {
	x, ok := p.pool.Get().(*BotSetCallbackAnswer)
	if !ok {
		return &BotSetCallbackAnswer{}
	}
	return x
}

func (p *poolBotSetCallbackAnswer) Put(x *BotSetCallbackAnswer) {
	x.QueryID = 0
	x.Url = ""
	x.Message = ""
	x.CacheTime = 0
	p.pool.Put(x)
}

var PoolBotSetCallbackAnswer = poolBotSetCallbackAnswer{}

const C_BotGetCallbackAnswer int64 = 1226608321

type poolBotGetCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotGetCallbackAnswer) Get() *BotGetCallbackAnswer {
	x, ok := p.pool.Get().(*BotGetCallbackAnswer)
	if !ok {
		return &BotGetCallbackAnswer{}
	}
	return x
}

func (p *poolBotGetCallbackAnswer) Put(x *BotGetCallbackAnswer) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MessageID = 0
	x.Data = x.Data[:0]
	p.pool.Put(x)
}

var PoolBotGetCallbackAnswer = poolBotGetCallbackAnswer{}

const C_BotGetInlineResults int64 = 375113086

type poolBotGetInlineResults struct {
	pool sync.Pool
}

func (p *poolBotGetInlineResults) Get() *BotGetInlineResults {
	x, ok := p.pool.Get().(*BotGetInlineResults)
	if !ok {
		return &BotGetInlineResults{}
	}
	return x
}

func (p *poolBotGetInlineResults) Put(x *BotGetInlineResults) {
	if x.Bot != nil {
		PoolInputUser.Put(x.Bot)
		x.Bot = nil
	}
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Query = ""
	x.Offset = ""
	if x.Location != nil {
		PoolInputGeoLocation.Put(x.Location)
		x.Location = nil
	}
	p.pool.Put(x)
}

var PoolBotGetInlineResults = poolBotGetInlineResults{}

const C_BotSetInlineResults int64 = 609072999

type poolBotSetInlineResults struct {
	pool sync.Pool
}

func (p *poolBotSetInlineResults) Get() *BotSetInlineResults {
	x, ok := p.pool.Get().(*BotSetInlineResults)
	if !ok {
		return &BotSetInlineResults{}
	}
	return x
}

func (p *poolBotSetInlineResults) Put(x *BotSetInlineResults) {
	x.Gallery = false
	x.Private = false
	x.CacheTime = 0
	x.NextOffset = ""
	x.Results = x.Results[:0]
	if x.SwitchPM != nil {
		PoolBotInlineSwitchPM.Put(x.SwitchPM)
		x.SwitchPM = nil
	}
	x.QueryID = 0
	p.pool.Put(x)
}

var PoolBotSetInlineResults = poolBotSetInlineResults{}

const C_BotSendInlineResults int64 = 1786892237

type poolBotSendInlineResults struct {
	pool sync.Pool
}

func (p *poolBotSendInlineResults) Get() *BotSendInlineResults {
	x, ok := p.pool.Get().(*BotSendInlineResults)
	if !ok {
		return &BotSendInlineResults{}
	}
	return x
}

func (p *poolBotSendInlineResults) Put(x *BotSendInlineResults) {
	x.RandomID = 0
	x.QueryID = 0
	x.ResultID = ""
	x.ClearDraft = false
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.ReplyTo = 0
	x.Silent = false
	x.HideVia = false
	p.pool.Put(x)
}

var PoolBotSendInlineResults = poolBotSendInlineResults{}

const C_BotUploadWallPaper int64 = 2519914079

type poolBotUploadWallPaper struct {
	pool sync.Pool
}

func (p *poolBotUploadWallPaper) Get() *BotUploadWallPaper {
	x, ok := p.pool.Get().(*BotUploadWallPaper)
	if !ok {
		return &BotUploadWallPaper{}
	}
	return x
}

func (p *poolBotUploadWallPaper) Put(x *BotUploadWallPaper) {
	if x.File != nil {
		PoolInputFileLocation.Put(x.File)
		x.File = nil
	}
	x.Dark = false
	x.Pattern = false
	if x.Settings != nil {
		PoolWallPaperSettings.Put(x.Settings)
		x.Settings = nil
	}
	p.pool.Put(x)
}

var PoolBotUploadWallPaper = poolBotUploadWallPaper{}

const C_BotUploadGif int64 = 3754133337

type poolBotUploadGif struct {
	pool sync.Pool
}

func (p *poolBotUploadGif) Get() *BotUploadGif {
	x, ok := p.pool.Get().(*BotUploadGif)
	if !ok {
		return &BotUploadGif{}
	}
	return x
}

func (p *poolBotUploadGif) Put(x *BotUploadGif) {
	x.Token = ""
	if x.File != nil {
		PoolInputFile.Put(x.File)
		x.File = nil
	}
	if x.Thumb != nil {
		PoolInputFile.Put(x.Thumb)
		x.Thumb = nil
	}
	x.Width = 0
	x.Height = 0
	x.MimeType = ""
	p.pool.Put(x)
}

var PoolBotUploadGif = poolBotUploadGif{}

const C_BotResults int64 = 2575283165

type poolBotResults struct {
	pool sync.Pool
}

func (p *poolBotResults) Get() *BotResults {
	x, ok := p.pool.Get().(*BotResults)
	if !ok {
		return &BotResults{}
	}
	return x
}

func (p *poolBotResults) Put(x *BotResults) {
	x.Gallery = false
	x.QueryID = 0
	x.NextOffset = ""
	if x.SwitchPM != nil {
		PoolBotInlineSwitchPM.Put(x.SwitchPM)
		x.SwitchPM = nil
	}
	x.Results = x.Results[:0]
	p.pool.Put(x)
}

var PoolBotResults = poolBotResults{}

const C_BotInlineSwitchPM int64 = 4111477214

type poolBotInlineSwitchPM struct {
	pool sync.Pool
}

func (p *poolBotInlineSwitchPM) Get() *BotInlineSwitchPM {
	x, ok := p.pool.Get().(*BotInlineSwitchPM)
	if !ok {
		return &BotInlineSwitchPM{}
	}
	return x
}

func (p *poolBotInlineSwitchPM) Put(x *BotInlineSwitchPM) {
	x.Text = ""
	x.StartParam = ""
	p.pool.Put(x)
}

var PoolBotInlineSwitchPM = poolBotInlineSwitchPM{}

const C_BotInlineResult int64 = 778738919

type poolBotInlineResult struct {
	pool sync.Pool
}

func (p *poolBotInlineResult) Get() *BotInlineResult {
	x, ok := p.pool.Get().(*BotInlineResult)
	if !ok {
		return &BotInlineResult{}
	}
	return x
}

func (p *poolBotInlineResult) Put(x *BotInlineResult) {
	x.ID = ""
	x.Type = 0
	x.Title = ""
	x.Description = ""
	x.Url = ""
	if x.Thumb != nil {
		PoolMediaWebDocument.Put(x.Thumb)
		x.Thumb = nil
	}
	if x.Message != nil {
		PoolBotInlineMessage.Put(x.Message)
		x.Message = nil
	}
	p.pool.Put(x)
}

var PoolBotInlineResult = poolBotInlineResult{}

const C_InputBotInlineResult int64 = 3710074255

type poolInputBotInlineResult struct {
	pool sync.Pool
}

func (p *poolInputBotInlineResult) Get() *InputBotInlineResult {
	x, ok := p.pool.Get().(*InputBotInlineResult)
	if !ok {
		return &InputBotInlineResult{}
	}
	return x
}

func (p *poolInputBotInlineResult) Put(x *InputBotInlineResult) {
	x.ID = ""
	x.Type = 0
	x.Title = ""
	x.Description = ""
	x.Url = ""
	if x.Thumb != nil {
		PoolInputMediaWebDocument.Put(x.Thumb)
		x.Thumb = nil
	}
	if x.Message != nil {
		PoolInputBotInlineMessage.Put(x.Message)
		x.Message = nil
	}
	p.pool.Put(x)
}

var PoolInputBotInlineResult = poolInputBotInlineResult{}

const C_BotInlineMessage int64 = 206571047

type poolBotInlineMessage struct {
	pool sync.Pool
}

func (p *poolBotInlineMessage) Get() *BotInlineMessage {
	x, ok := p.pool.Get().(*BotInlineMessage)
	if !ok {
		return &BotInlineMessage{}
	}
	return x
}

func (p *poolBotInlineMessage) Put(x *BotInlineMessage) {
	x.MediaData = x.MediaData[:0]
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyTo = 0
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	p.pool.Put(x)
}

var PoolBotInlineMessage = poolBotInlineMessage{}

const C_InputBotInlineMessage int64 = 1338681068

type poolInputBotInlineMessage struct {
	pool sync.Pool
}

func (p *poolInputBotInlineMessage) Get() *InputBotInlineMessage {
	x, ok := p.pool.Get().(*InputBotInlineMessage)
	if !ok {
		return &InputBotInlineMessage{}
	}
	return x
}

func (p *poolInputBotInlineMessage) Put(x *InputBotInlineMessage) {
	x.InputMediaData = x.InputMediaData[:0]
	x.NoWebPage = false
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyTo = 0
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	p.pool.Put(x)
}

var PoolInputBotInlineMessage = poolInputBotInlineMessage{}

const C_BotToken int64 = 230757930

type poolBotToken struct {
	pool sync.Pool
}

func (p *poolBotToken) Get() *BotToken {
	x, ok := p.pool.Get().(*BotToken)
	if !ok {
		return &BotToken{}
	}
	return x
}

func (p *poolBotToken) Put(x *BotToken) {
	x.Token = x.Token[:0]
	p.pool.Put(x)
}

var PoolBotToken = poolBotToken{}

const C_BotRecalled int64 = 4164808656

type poolBotRecalled struct {
	pool sync.Pool
}

func (p *poolBotRecalled) Get() *BotRecalled {
	x, ok := p.pool.Get().(*BotRecalled)
	if !ok {
		return &BotRecalled{}
	}
	return x
}

func (p *poolBotRecalled) Put(x *BotRecalled) {
	x.ID = 0
	x.Username = ""
	p.pool.Put(x)
}

var PoolBotRecalled = poolBotRecalled{}

const C_BotCallbackAnswer int64 = 2180565590

type poolBotCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotCallbackAnswer) Get() *BotCallbackAnswer {
	x, ok := p.pool.Get().(*BotCallbackAnswer)
	if !ok {
		return &BotCallbackAnswer{}
	}
	return x
}

func (p *poolBotCallbackAnswer) Put(x *BotCallbackAnswer) {
	x.Url = ""
	x.Message = ""
	x.CacheTime = 0
	p.pool.Put(x)
}

var PoolBotCallbackAnswer = poolBotCallbackAnswer{}

const C_BotsMany int64 = 430660433

type poolBotsMany struct {
	pool sync.Pool
}

func (p *poolBotsMany) Get() *BotsMany {
	x, ok := p.pool.Get().(*BotsMany)
	if !ok {
		return &BotsMany{}
	}
	return x
}

func (p *poolBotsMany) Put(x *BotsMany) {
	x.Bots = x.Bots[:0]
	p.pool.Put(x)
}

var PoolBotsMany = poolBotsMany{}

const C_BotGetCommands int64 = 685471542

type poolBotGetCommands struct {
	pool sync.Pool
}

func (p *poolBotGetCommands) Get() *BotGetCommands {
	x, ok := p.pool.Get().(*BotGetCommands)
	if !ok {
		return &BotGetCommands{}
	}
	return x
}

func (p *poolBotGetCommands) Put(x *BotGetCommands) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolBotGetCommands = poolBotGetCommands{}

const C_BotCommandsMany int64 = 369478577

type poolBotCommandsMany struct {
	pool sync.Pool
}

func (p *poolBotCommandsMany) Get() *BotCommandsMany {
	x, ok := p.pool.Get().(*BotCommandsMany)
	if !ok {
		return &BotCommandsMany{}
	}
	return x
}

func (p *poolBotCommandsMany) Put(x *BotCommandsMany) {
	x.Commands = x.Commands[:0]
	x.Empty = false
	p.pool.Put(x)
}

var PoolBotCommandsMany = poolBotCommandsMany{}

func init() {
	registry.RegisterConstructor(3448758942, "msg.BotStart")
	registry.RegisterConstructor(567464755, "msg.BotRecall")
	registry.RegisterConstructor(1487199122, "msg.BotSetInfo")
	registry.RegisterConstructor(3431546487, "msg.BotGet")
	registry.RegisterConstructor(3116228191, "msg.BotSendMessage")
	registry.RegisterConstructor(149350155, "msg.BotEditMessage")
	registry.RegisterConstructor(2100881420, "msg.BotSendMedia")
	registry.RegisterConstructor(598724192, "msg.BotSaveFilePart")
	registry.RegisterConstructor(1624560842, "msg.BotUpdateProfile")
	registry.RegisterConstructor(4201069063, "msg.BotUpdatePhoto")
	registry.RegisterConstructor(1601295593, "msg.BotRevokeToken")
	registry.RegisterConstructor(423422518, "msg.BotDeleteMessage")
	registry.RegisterConstructor(759509107, "msg.BotSetCallbackAnswer")
	registry.RegisterConstructor(1226608321, "msg.BotGetCallbackAnswer")
	registry.RegisterConstructor(375113086, "msg.BotGetInlineResults")
	registry.RegisterConstructor(609072999, "msg.BotSetInlineResults")
	registry.RegisterConstructor(1786892237, "msg.BotSendInlineResults")
	registry.RegisterConstructor(2519914079, "msg.BotUploadWallPaper")
	registry.RegisterConstructor(3754133337, "msg.BotUploadGif")
	registry.RegisterConstructor(2575283165, "msg.BotResults")
	registry.RegisterConstructor(4111477214, "msg.BotInlineSwitchPM")
	registry.RegisterConstructor(778738919, "msg.BotInlineResult")
	registry.RegisterConstructor(3710074255, "msg.InputBotInlineResult")
	registry.RegisterConstructor(206571047, "msg.BotInlineMessage")
	registry.RegisterConstructor(1338681068, "msg.InputBotInlineMessage")
	registry.RegisterConstructor(230757930, "msg.BotToken")
	registry.RegisterConstructor(4164808656, "msg.BotRecalled")
	registry.RegisterConstructor(2180565590, "msg.BotCallbackAnswer")
	registry.RegisterConstructor(430660433, "msg.BotsMany")
	registry.RegisterConstructor(685471542, "msg.BotGetCommands")
	registry.RegisterConstructor(369478577, "msg.BotCommandsMany")
}

func (x *BotStart) DeepCopy(z *BotStart) {
	if x.Bot != nil {
		z.Bot = PoolInputPeer.Get()
		x.Bot.DeepCopy(z.Bot)
	}
	z.RandomID = x.RandomID
	z.StartParam = x.StartParam
}

func (x *BotRecall) DeepCopy(z *BotRecall) {
	z.Version = x.Version
}

func (x *BotSetInfo) DeepCopy(z *BotSetInfo) {
	z.BotID = x.BotID
	z.RandomID = x.RandomID
	z.Owner = x.Owner
	for idx := range x.BotCommands {
		if x.BotCommands[idx] != nil {
			xx := PoolBotCommands.Get()
			x.BotCommands[idx].DeepCopy(xx)
			z.BotCommands = append(z.BotCommands, xx)
		}
	}
	z.Description = x.Description
	z.InlinePlaceholder = x.InlinePlaceholder
	z.InlineGeo = x.InlineGeo
}

func (x *BotGet) DeepCopy(z *BotGet) {
	z.UserID = x.UserID
	z.Limit = x.Limit
}

func (x *BotSendMessage) DeepCopy(z *BotSendMessage) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Body = x.Body
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
}

func (x *BotEditMessage) DeepCopy(z *BotEditMessage) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Body = x.Body
	z.MessageID = x.MessageID
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
}

func (x *BotSendMedia) DeepCopy(z *BotSendMedia) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MediaType = x.MediaType
	z.MediaData = append(z.MediaData[:0], x.MediaData...)
	z.ReplyTo = x.ReplyTo
}

func (x *BotSaveFilePart) DeepCopy(z *BotSaveFilePart) {
	z.FileID = x.FileID
	z.PartID = x.PartID
	z.TotalParts = x.TotalParts
	z.Bytes = append(z.Bytes[:0], x.Bytes...)
}

func (x *BotUpdateProfile) DeepCopy(z *BotUpdateProfile) {
	z.BotID = x.BotID
	z.Name = x.Name
	z.Bio = x.Bio
}

func (x *BotUpdatePhoto) DeepCopy(z *BotUpdatePhoto) {
	if x.File != nil {
		z.File = PoolInputFileLocation.Get()
		x.File.DeepCopy(z.File)
	}
	z.BotID = x.BotID
}

func (x *BotRevokeToken) DeepCopy(z *BotRevokeToken) {
	z.BotID = x.BotID
	z.GetNew = x.GetNew
}

func (x *BotDeleteMessage) DeepCopy(z *BotDeleteMessage) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *BotSetCallbackAnswer) DeepCopy(z *BotSetCallbackAnswer) {
	z.QueryID = x.QueryID
	z.Url = x.Url
	z.Message = x.Message
	z.CacheTime = x.CacheTime
}

func (x *BotGetCallbackAnswer) DeepCopy(z *BotGetCallbackAnswer) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MessageID = x.MessageID
	z.Data = append(z.Data[:0], x.Data...)
}

func (x *BotGetInlineResults) DeepCopy(z *BotGetInlineResults) {
	if x.Bot != nil {
		z.Bot = PoolInputUser.Get()
		x.Bot.DeepCopy(z.Bot)
	}
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Query = x.Query
	z.Offset = x.Offset
	if x.Location != nil {
		z.Location = PoolInputGeoLocation.Get()
		x.Location.DeepCopy(z.Location)
	}
}

func (x *BotSetInlineResults) DeepCopy(z *BotSetInlineResults) {
	z.Gallery = x.Gallery
	z.Private = x.Private
	z.CacheTime = x.CacheTime
	z.NextOffset = x.NextOffset
	for idx := range x.Results {
		if x.Results[idx] != nil {
			xx := PoolInputBotInlineResult.Get()
			x.Results[idx].DeepCopy(xx)
			z.Results = append(z.Results, xx)
		}
	}
	if x.SwitchPM != nil {
		z.SwitchPM = PoolBotInlineSwitchPM.Get()
		x.SwitchPM.DeepCopy(z.SwitchPM)
	}
	z.QueryID = x.QueryID
}

func (x *BotSendInlineResults) DeepCopy(z *BotSendInlineResults) {
	z.RandomID = x.RandomID
	z.QueryID = x.QueryID
	z.ResultID = x.ResultID
	z.ClearDraft = x.ClearDraft
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.ReplyTo = x.ReplyTo
	z.Silent = x.Silent
	z.HideVia = x.HideVia
}

func (x *BotUploadWallPaper) DeepCopy(z *BotUploadWallPaper) {
	if x.File != nil {
		z.File = PoolInputFileLocation.Get()
		x.File.DeepCopy(z.File)
	}
	z.Dark = x.Dark
	z.Pattern = x.Pattern
	if x.Settings != nil {
		z.Settings = PoolWallPaperSettings.Get()
		x.Settings.DeepCopy(z.Settings)
	}
}

func (x *BotUploadGif) DeepCopy(z *BotUploadGif) {
	z.Token = x.Token
	if x.File != nil {
		z.File = PoolInputFile.Get()
		x.File.DeepCopy(z.File)
	}
	if x.Thumb != nil {
		z.Thumb = PoolInputFile.Get()
		x.Thumb.DeepCopy(z.Thumb)
	}
	z.Width = x.Width
	z.Height = x.Height
	z.MimeType = x.MimeType
}

func (x *BotResults) DeepCopy(z *BotResults) {
	z.Gallery = x.Gallery
	z.QueryID = x.QueryID
	z.NextOffset = x.NextOffset
	if x.SwitchPM != nil {
		z.SwitchPM = PoolBotInlineSwitchPM.Get()
		x.SwitchPM.DeepCopy(z.SwitchPM)
	}
	for idx := range x.Results {
		if x.Results[idx] != nil {
			xx := PoolBotInlineResult.Get()
			x.Results[idx].DeepCopy(xx)
			z.Results = append(z.Results, xx)
		}
	}
}

func (x *BotInlineSwitchPM) DeepCopy(z *BotInlineSwitchPM) {
	z.Text = x.Text
	z.StartParam = x.StartParam
}

func (x *BotInlineResult) DeepCopy(z *BotInlineResult) {
	z.ID = x.ID
	z.Type = x.Type
	z.Title = x.Title
	z.Description = x.Description
	z.Url = x.Url
	if x.Thumb != nil {
		z.Thumb = PoolMediaWebDocument.Get()
		x.Thumb.DeepCopy(z.Thumb)
	}
	if x.Message != nil {
		z.Message = PoolBotInlineMessage.Get()
		x.Message.DeepCopy(z.Message)
	}
}

func (x *InputBotInlineResult) DeepCopy(z *InputBotInlineResult) {
	z.ID = x.ID
	z.Type = x.Type
	z.Title = x.Title
	z.Description = x.Description
	z.Url = x.Url
	if x.Thumb != nil {
		z.Thumb = PoolInputMediaWebDocument.Get()
		x.Thumb.DeepCopy(z.Thumb)
	}
	if x.Message != nil {
		z.Message = PoolInputBotInlineMessage.Get()
		x.Message.DeepCopy(z.Message)
	}
}

func (x *BotInlineMessage) DeepCopy(z *BotInlineMessage) {
	z.MediaData = append(z.MediaData[:0], x.MediaData...)
	z.Body = x.Body
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyTo = x.ReplyTo
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
}

func (x *InputBotInlineMessage) DeepCopy(z *InputBotInlineMessage) {
	z.InputMediaData = append(z.InputMediaData[:0], x.InputMediaData...)
	z.NoWebPage = x.NoWebPage
	z.Body = x.Body
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyTo = x.ReplyTo
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
}

func (x *BotToken) DeepCopy(z *BotToken) {
	z.Token = append(z.Token[:0], x.Token...)
}

func (x *BotRecalled) DeepCopy(z *BotRecalled) {
	z.ID = x.ID
	z.Username = x.Username
}

func (x *BotCallbackAnswer) DeepCopy(z *BotCallbackAnswer) {
	z.Url = x.Url
	z.Message = x.Message
	z.CacheTime = x.CacheTime
}

func (x *BotsMany) DeepCopy(z *BotsMany) {
	for idx := range x.Bots {
		if x.Bots[idx] != nil {
			xx := PoolBotInfo.Get()
			x.Bots[idx].DeepCopy(xx)
			z.Bots = append(z.Bots, xx)
		}
	}
}

func (x *BotGetCommands) DeepCopy(z *BotGetCommands) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *BotCommandsMany) DeepCopy(z *BotCommandsMany) {
	for idx := range x.Commands {
		if x.Commands[idx] != nil {
			xx := PoolBotCommands.Get()
			x.Commands[idx].DeepCopy(xx)
			z.Commands = append(z.Commands, xx)
		}
	}
	z.Empty = x.Empty
}

func (x *BotStart) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotStart, x)
}

func (x *BotRecall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotRecall, x)
}

func (x *BotSetInfo) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSetInfo, x)
}

func (x *BotGet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotGet, x)
}

func (x *BotSendMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSendMessage, x)
}

func (x *BotEditMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotEditMessage, x)
}

func (x *BotSendMedia) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSendMedia, x)
}

func (x *BotSaveFilePart) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSaveFilePart, x)
}

func (x *BotUpdateProfile) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotUpdateProfile, x)
}

func (x *BotUpdatePhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotUpdatePhoto, x)
}

func (x *BotRevokeToken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotRevokeToken, x)
}

func (x *BotDeleteMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotDeleteMessage, x)
}

func (x *BotSetCallbackAnswer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSetCallbackAnswer, x)
}

func (x *BotGetCallbackAnswer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotGetCallbackAnswer, x)
}

func (x *BotGetInlineResults) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotGetInlineResults, x)
}

func (x *BotSetInlineResults) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSetInlineResults, x)
}

func (x *BotSendInlineResults) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotSendInlineResults, x)
}

func (x *BotUploadWallPaper) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotUploadWallPaper, x)
}

func (x *BotUploadGif) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotUploadGif, x)
}

func (x *BotResults) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotResults, x)
}

func (x *BotInlineSwitchPM) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotInlineSwitchPM, x)
}

func (x *BotInlineResult) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotInlineResult, x)
}

func (x *InputBotInlineResult) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputBotInlineResult, x)
}

func (x *BotInlineMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotInlineMessage, x)
}

func (x *InputBotInlineMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputBotInlineMessage, x)
}

func (x *BotToken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotToken, x)
}

func (x *BotRecalled) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotRecalled, x)
}

func (x *BotCallbackAnswer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotCallbackAnswer, x)
}

func (x *BotsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotsMany, x)
}

func (x *BotGetCommands) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotGetCommands, x)
}

func (x *BotCommandsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotCommandsMany, x)
}

func (x *BotStart) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotRecall) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSetInfo) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotGet) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSendMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotEditMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSendMedia) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSaveFilePart) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotUpdateProfile) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotUpdatePhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotRevokeToken) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotDeleteMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSetCallbackAnswer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotGetCallbackAnswer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotGetInlineResults) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSetInlineResults) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotSendInlineResults) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotUploadWallPaper) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotUploadGif) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotResults) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotInlineSwitchPM) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotInlineResult) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputBotInlineResult) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotInlineMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputBotInlineMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotToken) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotRecalled) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotCallbackAnswer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotGetCommands) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotCommandsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotStart) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotRecall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSetInfo) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotGet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSendMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotEditMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSendMedia) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSaveFilePart) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotUpdateProfile) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotUpdatePhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotRevokeToken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotDeleteMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSetCallbackAnswer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotGetCallbackAnswer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotGetInlineResults) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSetInlineResults) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotSendInlineResults) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotUploadWallPaper) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotUploadGif) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotResults) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotInlineSwitchPM) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotInlineResult) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *InputBotInlineResult) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotInlineMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *InputBotInlineMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotToken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotRecalled) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotCallbackAnswer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotGetCommands) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *BotCommandsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

const C_ServeBotRequest int64 = 685561014

type IBotProxy interface {
	ServeBotRequest(ctx *edge.RequestCtx, req *rony.MessageEnvelope, res *rony.MessageEnvelope)
}

type BotProxyWrapper struct {
	h IBotProxy
}

func RegisterBotProxy(h IBotProxy, e *edge.Server) {
	w := BotProxyWrapper{
		h: h,
	}
	w.Register(e)
}

func (sw *BotProxyWrapper) Register(e *edge.Server) {
	e.SetHandlers(C_ServeBotRequest, true, sw.ServeBotRequestWrapper)
}

func (sw *BotProxyWrapper) ServeBotRequestWrapper(ctx *edge.RequestCtx, in *rony.MessageEnvelope) {
	req := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(req)
	res := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(res)
	err := proto.UnmarshalOptions{Merge: true}.Unmarshal(in.Message, req)
	if err != nil {
		ctx.PushError(rony.ErrCodeInvalid, rony.ErrItemRequest)
		return
	}

	sw.h.ServeBotRequest(ctx, req, res)
	if !ctx.Stopped() {
		ctx.PushMessage(rony.C_MessageEnvelope, res)
	}
}

type BotProxyClient struct {
	c edgec.Client
}

func NewBotProxyClient(ec edgec.Client) *BotProxyClient {
	return &BotProxyClient{
		c: ec,
	}
}

func (c *BotProxyClient) ServeBotRequest(req *rony.MessageEnvelope, kvs ...*rony.KeyValue) (*rony.MessageEnvelope, error) {
	out := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(out)
	in := rony.PoolMessageEnvelope.Get()
	defer rony.PoolMessageEnvelope.Put(in)
	out.Fill(c.c.GetRequestID(), C_ServeBotRequest, req, kvs...)
	err := c.c.Send(out, in)
	if err != nil {
		return nil, err
	}
	switch in.GetConstructor() {
	case rony.C_MessageEnvelope:
		x := &rony.MessageEnvelope{}
		_ = proto.Unmarshal(in.Message, x)
		return x, nil
	case rony.C_Error:
		x := &rony.Error{}
		_ = proto.Unmarshal(in.Message, x)
		return nil, fmt.Errorf("%s:%s", x.GetCode(), x.GetItems())
	default:
		return nil, fmt.Errorf("unknown message: %d", in.GetConstructor())
	}
}
