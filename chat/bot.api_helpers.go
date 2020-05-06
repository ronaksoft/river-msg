// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bot.api.proto

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

const C_BotStart int64 = 2068702388

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
	p.pool.Put(x)
}

var PoolBotStart = poolBotStart{}

func ResultBotStart(out *MessageEnvelope, res *BotStart) {
	out.Constructor = C_BotStart
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotRecall int64 = 4208974051

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
	p.pool.Put(x)
}

var PoolBotRecall = poolBotRecall{}

func ResultBotRecall(out *MessageEnvelope, res *BotRecall) {
	out.Constructor = C_BotRecall
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSetInfo int64 = 3735815245

type poolBotSetInfo struct {
	pool sync.Pool
}

func (p *poolBotSetInfo) Get() *BotSetInfo {
	x, ok := p.pool.Get().(*BotSetInfo)
	if !ok {
		return &BotSetInfo{}
	}
	x.BotCommands = x.BotCommands[:0]
	return x
}

func (p *poolBotSetInfo) Put(x *BotSetInfo) {
	p.pool.Put(x)
}

var PoolBotSetInfo = poolBotSetInfo{}

func ResultBotSetInfo(out *MessageEnvelope, res *BotSetInfo) {
	out.Constructor = C_BotSetInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotGet int64 = 911895569

type poolBotGet struct {
	pool sync.Pool
}

func (p *poolBotGet) Get() *BotGet {
	x, ok := p.pool.Get().(*BotGet)
	if !ok {
		return &BotGet{}
	}
	x.Limit = 0
	return x
}

func (p *poolBotGet) Put(x *BotGet) {
	p.pool.Put(x)
}

var PoolBotGet = poolBotGet{}

func ResultBotGet(out *MessageEnvelope, res *BotGet) {
	out.Constructor = C_BotGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSendMessage int64 = 2371725696

type poolBotSendMessage struct {
	pool sync.Pool
}

func (p *poolBotSendMessage) Get() *BotSendMessage {
	x, ok := p.pool.Get().(*BotSendMessage)
	if !ok {
		return &BotSendMessage{}
	}
	x.ReplyTo = 0
	x.ClearDraft = false
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	return x
}

func (p *poolBotSendMessage) Put(x *BotSendMessage) {
	p.pool.Put(x)
}

var PoolBotSendMessage = poolBotSendMessage{}

func ResultBotSendMessage(out *MessageEnvelope, res *BotSendMessage) {
	out.Constructor = C_BotSendMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotEditMessage int64 = 1007063252

type poolBotEditMessage struct {
	pool sync.Pool
}

func (p *poolBotEditMessage) Get() *BotEditMessage {
	x, ok := p.pool.Get().(*BotEditMessage)
	if !ok {
		return &BotEditMessage{}
	}
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	return x
}

func (p *poolBotEditMessage) Put(x *BotEditMessage) {
	p.pool.Put(x)
}

var PoolBotEditMessage = poolBotEditMessage{}

func ResultBotEditMessage(out *MessageEnvelope, res *BotEditMessage) {
	out.Constructor = C_BotEditMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSendMedia int64 = 1844738193

type poolBotSendMedia struct {
	pool sync.Pool
}

func (p *poolBotSendMedia) Get() *BotSendMedia {
	x, ok := p.pool.Get().(*BotSendMedia)
	if !ok {
		return &BotSendMedia{}
	}
	x.ReplyTo = 0
	return x
}

func (p *poolBotSendMedia) Put(x *BotSendMedia) {
	p.pool.Put(x)
}

var PoolBotSendMedia = poolBotSendMedia{}

func ResultBotSendMedia(out *MessageEnvelope, res *BotSendMedia) {
	out.Constructor = C_BotSendMedia
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSaveFilePart int64 = 905437522

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
	p.pool.Put(x)
}

var PoolBotSaveFilePart = poolBotSaveFilePart{}

func ResultBotSaveFilePart(out *MessageEnvelope, res *BotSaveFilePart) {
	out.Constructor = C_BotSaveFilePart
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotUpdateProfile int64 = 2820005221

type poolBotUpdateProfile struct {
	pool sync.Pool
}

func (p *poolBotUpdateProfile) Get() *BotUpdateProfile {
	x, ok := p.pool.Get().(*BotUpdateProfile)
	if !ok {
		return &BotUpdateProfile{}
	}
	x.Bio = ""
	return x
}

func (p *poolBotUpdateProfile) Put(x *BotUpdateProfile) {
	p.pool.Put(x)
}

var PoolBotUpdateProfile = poolBotUpdateProfile{}

func ResultBotUpdateProfile(out *MessageEnvelope, res *BotUpdateProfile) {
	out.Constructor = C_BotUpdateProfile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotUpdatePhoto int64 = 3464973784

type poolBotUpdatePhoto struct {
	pool sync.Pool
}

func (p *poolBotUpdatePhoto) Get() *BotUpdatePhoto {
	x, ok := p.pool.Get().(*BotUpdatePhoto)
	if !ok {
		return &BotUpdatePhoto{}
	}
	x.File = nil
	return x
}

func (p *poolBotUpdatePhoto) Put(x *BotUpdatePhoto) {
	p.pool.Put(x)
}

var PoolBotUpdatePhoto = poolBotUpdatePhoto{}

func ResultBotUpdatePhoto(out *MessageEnvelope, res *BotUpdatePhoto) {
	out.Constructor = C_BotUpdatePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotRevokeToken int64 = 1804706614

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
	p.pool.Put(x)
}

var PoolBotRevokeToken = poolBotRevokeToken{}

func ResultBotRevokeToken(out *MessageEnvelope, res *BotRevokeToken) {
	out.Constructor = C_BotRevokeToken
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotDeleteMessage int64 = 3523077017

type poolBotDeleteMessage struct {
	pool sync.Pool
}

func (p *poolBotDeleteMessage) Get() *BotDeleteMessage {
	x, ok := p.pool.Get().(*BotDeleteMessage)
	if !ok {
		return &BotDeleteMessage{}
	}
	x.MessageIDs = x.MessageIDs[:0]
	return x
}

func (p *poolBotDeleteMessage) Put(x *BotDeleteMessage) {
	p.pool.Put(x)
}

var PoolBotDeleteMessage = poolBotDeleteMessage{}

func ResultBotDeleteMessage(out *MessageEnvelope, res *BotDeleteMessage) {
	out.Constructor = C_BotDeleteMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSetCallbackAnswer int64 = 1891806754

type poolBotSetCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotSetCallbackAnswer) Get() *BotSetCallbackAnswer {
	x, ok := p.pool.Get().(*BotSetCallbackAnswer)
	if !ok {
		return &BotSetCallbackAnswer{}
	}
	x.Url = ""
	x.Message = ""
	return x
}

func (p *poolBotSetCallbackAnswer) Put(x *BotSetCallbackAnswer) {
	p.pool.Put(x)
}

var PoolBotSetCallbackAnswer = poolBotSetCallbackAnswer{}

func ResultBotSetCallbackAnswer(out *MessageEnvelope, res *BotSetCallbackAnswer) {
	out.Constructor = C_BotSetCallbackAnswer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotGetCallbackAnswer int64 = 345706640

type poolBotGetCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotGetCallbackAnswer) Get() *BotGetCallbackAnswer {
	x, ok := p.pool.Get().(*BotGetCallbackAnswer)
	if !ok {
		return &BotGetCallbackAnswer{}
	}
	x.MessageID = 0
	return x
}

func (p *poolBotGetCallbackAnswer) Put(x *BotGetCallbackAnswer) {
	p.pool.Put(x)
}

var PoolBotGetCallbackAnswer = poolBotGetCallbackAnswer{}

func ResultBotGetCallbackAnswer(out *MessageEnvelope, res *BotGetCallbackAnswer) {
	out.Constructor = C_BotGetCallbackAnswer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotGetInlineResults int64 = 4192114308

type poolBotGetInlineResults struct {
	pool sync.Pool
}

func (p *poolBotGetInlineResults) Get() *BotGetInlineResults {
	x, ok := p.pool.Get().(*BotGetInlineResults)
	if !ok {
		return &BotGetInlineResults{}
	}
	x.Location = nil
	return x
}

func (p *poolBotGetInlineResults) Put(x *BotGetInlineResults) {
	p.pool.Put(x)
}

var PoolBotGetInlineResults = poolBotGetInlineResults{}

func ResultBotGetInlineResults(out *MessageEnvelope, res *BotGetInlineResults) {
	out.Constructor = C_BotGetInlineResults
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSetInlineResults int64 = 3418940573

type poolBotSetInlineResults struct {
	pool sync.Pool
}

func (p *poolBotSetInlineResults) Get() *BotSetInlineResults {
	x, ok := p.pool.Get().(*BotSetInlineResults)
	if !ok {
		return &BotSetInlineResults{}
	}
	x.NextOffset = ""
	x.Results = x.Results[:0]
	x.SwitchPM = nil
	return x
}

func (p *poolBotSetInlineResults) Put(x *BotSetInlineResults) {
	p.pool.Put(x)
}

var PoolBotSetInlineResults = poolBotSetInlineResults{}

func ResultBotSetInlineResults(out *MessageEnvelope, res *BotSetInlineResults) {
	out.Constructor = C_BotSetInlineResults
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotSendInlineResults int64 = 923160988

type poolBotSendInlineResults struct {
	pool sync.Pool
}

func (p *poolBotSendInlineResults) Get() *BotSendInlineResults {
	x, ok := p.pool.Get().(*BotSendInlineResults)
	if !ok {
		return &BotSendInlineResults{}
	}
	x.Silent = false
	x.HideVia = false
	return x
}

func (p *poolBotSendInlineResults) Put(x *BotSendInlineResults) {
	p.pool.Put(x)
}

var PoolBotSendInlineResults = poolBotSendInlineResults{}

func ResultBotSendInlineResults(out *MessageEnvelope, res *BotSendInlineResults) {
	out.Constructor = C_BotSendInlineResults
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotResults int64 = 527920130

type poolBotResults struct {
	pool sync.Pool
}

func (p *poolBotResults) Get() *BotResults {
	x, ok := p.pool.Get().(*BotResults)
	if !ok {
		return &BotResults{}
	}
	x.NextOffset = ""
	x.SwitchPM = nil
	x.Results = x.Results[:0]
	return x
}

func (p *poolBotResults) Put(x *BotResults) {
	p.pool.Put(x)
}

var PoolBotResults = poolBotResults{}

func ResultBotResults(out *MessageEnvelope, res *BotResults) {
	out.Constructor = C_BotResults
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotInlineSwitchPM int64 = 3014743726

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
	p.pool.Put(x)
}

var PoolBotInlineSwitchPM = poolBotInlineSwitchPM{}

func ResultBotInlineSwitchPM(out *MessageEnvelope, res *BotInlineSwitchPM) {
	out.Constructor = C_BotInlineSwitchPM
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotInlineResult int64 = 942846933

type poolBotInlineResult struct {
	pool sync.Pool
}

func (p *poolBotInlineResult) Get() *BotInlineResult {
	x, ok := p.pool.Get().(*BotInlineResult)
	if !ok {
		return &BotInlineResult{}
	}
	x.Title = ""
	x.Description = ""
	x.Url = ""
	x.Thumb = nil
	return x
}

func (p *poolBotInlineResult) Put(x *BotInlineResult) {
	p.pool.Put(x)
}

var PoolBotInlineResult = poolBotInlineResult{}

func ResultBotInlineResult(out *MessageEnvelope, res *BotInlineResult) {
	out.Constructor = C_BotInlineResult
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputBotInlineResult int64 = 2158273502

type poolInputBotInlineResult struct {
	pool sync.Pool
}

func (p *poolInputBotInlineResult) Get() *InputBotInlineResult {
	x, ok := p.pool.Get().(*InputBotInlineResult)
	if !ok {
		return &InputBotInlineResult{}
	}
	x.Title = ""
	x.Description = ""
	x.Url = ""
	x.Thumb = nil
	return x
}

func (p *poolInputBotInlineResult) Put(x *InputBotInlineResult) {
	p.pool.Put(x)
}

var PoolInputBotInlineResult = poolInputBotInlineResult{}

func ResultInputBotInlineResult(out *MessageEnvelope, res *InputBotInlineResult) {
	out.Constructor = C_InputBotInlineResult
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotInlineMessage int64 = 3297841032

type poolBotInlineMessage struct {
	pool sync.Pool
}

func (p *poolBotInlineMessage) Get() *BotInlineMessage {
	x, ok := p.pool.Get().(*BotInlineMessage)
	if !ok {
		return &BotInlineMessage{}
	}
	x.Contact = nil
	x.Geo = nil
	x.Doc = nil
	x.Poll = nil
	x.Invoice = nil
	x.WebDoc = nil
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyTo = 0
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	return x
}

func (p *poolBotInlineMessage) Put(x *BotInlineMessage) {
	p.pool.Put(x)
}

var PoolBotInlineMessage = poolBotInlineMessage{}

func ResultBotInlineMessage(out *MessageEnvelope, res *BotInlineMessage) {
	out.Constructor = C_BotInlineMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_InputBotInlineMessage int64 = 1408974864

type poolInputBotInlineMessage struct {
	pool sync.Pool
}

func (p *poolInputBotInlineMessage) Get() *InputBotInlineMessage {
	x, ok := p.pool.Get().(*InputBotInlineMessage)
	if !ok {
		return &InputBotInlineMessage{}
	}
	x.Contact = nil
	x.Geo = nil
	x.Doc = nil
	x.Poll = nil
	x.Invoice = nil
	x.UploadedDocument = nil
	x.WebDoc = nil
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyTo = 0
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	return x
}

func (p *poolInputBotInlineMessage) Put(x *InputBotInlineMessage) {
	p.pool.Put(x)
}

var PoolInputBotInlineMessage = poolInputBotInlineMessage{}

func ResultInputBotInlineMessage(out *MessageEnvelope, res *InputBotInlineMessage) {
	out.Constructor = C_InputBotInlineMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotToken int64 = 3137540096

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
	p.pool.Put(x)
}

var PoolBotToken = poolBotToken{}

func ResultBotToken(out *MessageEnvelope, res *BotToken) {
	out.Constructor = C_BotToken
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotRecalled int64 = 4007077962

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
	p.pool.Put(x)
}

var PoolBotRecalled = poolBotRecalled{}

func ResultBotRecalled(out *MessageEnvelope, res *BotRecalled) {
	out.Constructor = C_BotRecalled
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotCallbackAnswer int64 = 3344545062

type poolBotCallbackAnswer struct {
	pool sync.Pool
}

func (p *poolBotCallbackAnswer) Get() *BotCallbackAnswer {
	x, ok := p.pool.Get().(*BotCallbackAnswer)
	if !ok {
		return &BotCallbackAnswer{}
	}
	x.Url = ""
	x.Message = ""
	return x
}

func (p *poolBotCallbackAnswer) Put(x *BotCallbackAnswer) {
	p.pool.Put(x)
}

var PoolBotCallbackAnswer = poolBotCallbackAnswer{}

func ResultBotCallbackAnswer(out *MessageEnvelope, res *BotCallbackAnswer) {
	out.Constructor = C_BotCallbackAnswer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotsMany int64 = 2942918011

type poolBotsMany struct {
	pool sync.Pool
}

func (p *poolBotsMany) Get() *BotsMany {
	x, ok := p.pool.Get().(*BotsMany)
	if !ok {
		return &BotsMany{}
	}
	x.Bots = x.Bots[:0]
	return x
}

func (p *poolBotsMany) Put(x *BotsMany) {
	p.pool.Put(x)
}

var PoolBotsMany = poolBotsMany{}

func ResultBotsMany(out *MessageEnvelope, res *BotsMany) {
	out.Constructor = C_BotsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotGetCommands int64 = 473628905

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
	p.pool.Put(x)
}

var PoolBotGetCommands = poolBotGetCommands{}

func ResultBotGetCommands(out *MessageEnvelope, res *BotGetCommands) {
	out.Constructor = C_BotGetCommands
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_BotCommandsMany int64 = 6153347

type poolBotCommandsMany struct {
	pool sync.Pool
}

func (p *poolBotCommandsMany) Get() *BotCommandsMany {
	x, ok := p.pool.Get().(*BotCommandsMany)
	if !ok {
		return &BotCommandsMany{}
	}
	x.Commands = x.Commands[:0]
	return x
}

func (p *poolBotCommandsMany) Put(x *BotCommandsMany) {
	p.pool.Put(x)
}

var PoolBotCommandsMany = poolBotCommandsMany{}

func ResultBotCommandsMany(out *MessageEnvelope, res *BotCommandsMany) {
	out.Constructor = C_BotCommandsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[2068702388] = "BotStart"
	ConstructorNames[4208974051] = "BotRecall"
	ConstructorNames[3735815245] = "BotSetInfo"
	ConstructorNames[911895569] = "BotGet"
	ConstructorNames[2371725696] = "BotSendMessage"
	ConstructorNames[1007063252] = "BotEditMessage"
	ConstructorNames[1844738193] = "BotSendMedia"
	ConstructorNames[905437522] = "BotSaveFilePart"
	ConstructorNames[2820005221] = "BotUpdateProfile"
	ConstructorNames[3464973784] = "BotUpdatePhoto"
	ConstructorNames[1804706614] = "BotRevokeToken"
	ConstructorNames[3523077017] = "BotDeleteMessage"
	ConstructorNames[1891806754] = "BotSetCallbackAnswer"
	ConstructorNames[345706640] = "BotGetCallbackAnswer"
	ConstructorNames[4192114308] = "BotGetInlineResults"
	ConstructorNames[3418940573] = "BotSetInlineResults"
	ConstructorNames[923160988] = "BotSendInlineResults"
	ConstructorNames[527920130] = "BotResults"
	ConstructorNames[3014743726] = "BotInlineSwitchPM"
	ConstructorNames[942846933] = "BotInlineResult"
	ConstructorNames[2158273502] = "InputBotInlineResult"
	ConstructorNames[3297841032] = "BotInlineMessage"
	ConstructorNames[1408974864] = "InputBotInlineMessage"
	ConstructorNames[3137540096] = "BotToken"
	ConstructorNames[4007077962] = "BotRecalled"
	ConstructorNames[3344545062] = "BotCallbackAnswer"
	ConstructorNames[2942918011] = "BotsMany"
	ConstructorNames[473628905] = "BotGetCommands"
	ConstructorNames[6153347] = "BotCommandsMany"
}
