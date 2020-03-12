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
	ConstructorNames[2820005221] = "BotUpdateProfile"
	ConstructorNames[1891806754] = "BotSetCallbackAnswer"
	ConstructorNames[345706640] = "BotGetCallbackAnswer"
	ConstructorNames[4007077962] = "BotRecalled"
	ConstructorNames[3344545062] = "BotCallbackAnswer"
	ConstructorNames[2942918011] = "BotsMany"
	ConstructorNames[473628905] = "BotGetCommands"
	ConstructorNames[6153347] = "BotCommandsMany"
}
