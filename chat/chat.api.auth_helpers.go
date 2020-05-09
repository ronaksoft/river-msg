// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.auth.proto

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

const C_AuthRegister int64 = 2228369460

type poolAuthRegister struct {
	pool sync.Pool
}

func (p *poolAuthRegister) Get() *AuthRegister {
	x, ok := p.pool.Get().(*AuthRegister)
	if !ok {
		return &AuthRegister{}
	}
	x.LangCode = ""
	return x
}

func (p *poolAuthRegister) Put(x *AuthRegister) {
	p.pool.Put(x)
}

var PoolAuthRegister = poolAuthRegister{}

func ResultAuthRegister(out *MessageEnvelope, res *AuthRegister) {
	out.Constructor = C_AuthRegister
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthBotRegister int64 = 1579606687

type poolAuthBotRegister struct {
	pool sync.Pool
}

func (p *poolAuthBotRegister) Get() *AuthBotRegister {
	x, ok := p.pool.Get().(*AuthBotRegister)
	if !ok {
		return &AuthBotRegister{}
	}
	return x
}

func (p *poolAuthBotRegister) Put(x *AuthBotRegister) {
	p.pool.Put(x)
}

var PoolAuthBotRegister = poolAuthBotRegister{}

func ResultAuthBotRegister(out *MessageEnvelope, res *AuthBotRegister) {
	out.Constructor = C_AuthBotRegister
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthLogin int64 = 2587620888

type poolAuthLogin struct {
	pool sync.Pool
}

func (p *poolAuthLogin) Get() *AuthLogin {
	x, ok := p.pool.Get().(*AuthLogin)
	if !ok {
		return &AuthLogin{}
	}
	return x
}

func (p *poolAuthLogin) Put(x *AuthLogin) {
	p.pool.Put(x)
}

var PoolAuthLogin = poolAuthLogin{}

func ResultAuthLogin(out *MessageEnvelope, res *AuthLogin) {
	out.Constructor = C_AuthLogin
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthCheckPassword int64 = 3346962908

type poolAuthCheckPassword struct {
	pool sync.Pool
}

func (p *poolAuthCheckPassword) Get() *AuthCheckPassword {
	x, ok := p.pool.Get().(*AuthCheckPassword)
	if !ok {
		return &AuthCheckPassword{}
	}
	return x
}

func (p *poolAuthCheckPassword) Put(x *AuthCheckPassword) {
	p.pool.Put(x)
}

var PoolAuthCheckPassword = poolAuthCheckPassword{}

func ResultAuthCheckPassword(out *MessageEnvelope, res *AuthCheckPassword) {
	out.Constructor = C_AuthCheckPassword
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthRecoverPassword int64 = 2711231991

type poolAuthRecoverPassword struct {
	pool sync.Pool
}

func (p *poolAuthRecoverPassword) Get() *AuthRecoverPassword {
	x, ok := p.pool.Get().(*AuthRecoverPassword)
	if !ok {
		return &AuthRecoverPassword{}
	}
	return x
}

func (p *poolAuthRecoverPassword) Put(x *AuthRecoverPassword) {
	p.pool.Put(x)
}

var PoolAuthRecoverPassword = poolAuthRecoverPassword{}

func ResultAuthRecoverPassword(out *MessageEnvelope, res *AuthRecoverPassword) {
	out.Constructor = C_AuthRecoverPassword
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthLogout int64 = 992431648

type poolAuthLogout struct {
	pool sync.Pool
}

func (p *poolAuthLogout) Get() *AuthLogout {
	x, ok := p.pool.Get().(*AuthLogout)
	if !ok {
		return &AuthLogout{}
	}
	x.AuthIDs = x.AuthIDs[:0]
	return x
}

func (p *poolAuthLogout) Put(x *AuthLogout) {
	p.pool.Put(x)
}

var PoolAuthLogout = poolAuthLogout{}

func ResultAuthLogout(out *MessageEnvelope, res *AuthLogout) {
	out.Constructor = C_AuthLogout
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthLoginByToken int64 = 2851553023

type poolAuthLoginByToken struct {
	pool sync.Pool
}

func (p *poolAuthLoginByToken) Get() *AuthLoginByToken {
	x, ok := p.pool.Get().(*AuthLoginByToken)
	if !ok {
		return &AuthLoginByToken{}
	}
	return x
}

func (p *poolAuthLoginByToken) Put(x *AuthLoginByToken) {
	p.pool.Put(x)
}

var PoolAuthLoginByToken = poolAuthLoginByToken{}

func ResultAuthLoginByToken(out *MessageEnvelope, res *AuthLoginByToken) {
	out.Constructor = C_AuthLoginByToken
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthCheckPhone int64 = 4134648516

type poolAuthCheckPhone struct {
	pool sync.Pool
}

func (p *poolAuthCheckPhone) Get() *AuthCheckPhone {
	x, ok := p.pool.Get().(*AuthCheckPhone)
	if !ok {
		return &AuthCheckPhone{}
	}
	return x
}

func (p *poolAuthCheckPhone) Put(x *AuthCheckPhone) {
	p.pool.Put(x)
}

var PoolAuthCheckPhone = poolAuthCheckPhone{}

func ResultAuthCheckPhone(out *MessageEnvelope, res *AuthCheckPhone) {
	out.Constructor = C_AuthCheckPhone
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthSendCode int64 = 3984043365

type poolAuthSendCode struct {
	pool sync.Pool
}

func (p *poolAuthSendCode) Get() *AuthSendCode {
	x, ok := p.pool.Get().(*AuthSendCode)
	if !ok {
		return &AuthSendCode{}
	}
	x.AppHash = ""
	return x
}

func (p *poolAuthSendCode) Put(x *AuthSendCode) {
	p.pool.Put(x)
}

var PoolAuthSendCode = poolAuthSendCode{}

func ResultAuthSendCode(out *MessageEnvelope, res *AuthSendCode) {
	out.Constructor = C_AuthSendCode
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthResendCode int64 = 2682713491

type poolAuthResendCode struct {
	pool sync.Pool
}

func (p *poolAuthResendCode) Get() *AuthResendCode {
	x, ok := p.pool.Get().(*AuthResendCode)
	if !ok {
		return &AuthResendCode{}
	}
	x.AppHash = ""
	return x
}

func (p *poolAuthResendCode) Put(x *AuthResendCode) {
	p.pool.Put(x)
}

var PoolAuthResendCode = poolAuthResendCode{}

func ResultAuthResendCode(out *MessageEnvelope, res *AuthResendCode) {
	out.Constructor = C_AuthResendCode
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthRecall int64 = 1172029049

type poolAuthRecall struct {
	pool sync.Pool
}

func (p *poolAuthRecall) Get() *AuthRecall {
	x, ok := p.pool.Get().(*AuthRecall)
	if !ok {
		return &AuthRecall{}
	}
	x.Version = 0
	x.AppVersion = ""
	x.Platform = ""
	x.Vendor = ""
	x.OSVersion = ""
	return x
}

func (p *poolAuthRecall) Put(x *AuthRecall) {
	p.pool.Put(x)
}

var PoolAuthRecall = poolAuthRecall{}

func ResultAuthRecall(out *MessageEnvelope, res *AuthRecall) {
	out.Constructor = C_AuthRecall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthDestroyKey int64 = 3673422656

type poolAuthDestroyKey struct {
	pool sync.Pool
}

func (p *poolAuthDestroyKey) Get() *AuthDestroyKey {
	x, ok := p.pool.Get().(*AuthDestroyKey)
	if !ok {
		return &AuthDestroyKey{}
	}
	return x
}

func (p *poolAuthDestroyKey) Put(x *AuthDestroyKey) {
	p.pool.Put(x)
}

var PoolAuthDestroyKey = poolAuthDestroyKey{}

func ResultAuthDestroyKey(out *MessageEnvelope, res *AuthDestroyKey) {
	out.Constructor = C_AuthDestroyKey
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthPasswordRecovery int64 = 3813475914

type poolAuthPasswordRecovery struct {
	pool sync.Pool
}

func (p *poolAuthPasswordRecovery) Get() *AuthPasswordRecovery {
	x, ok := p.pool.Get().(*AuthPasswordRecovery)
	if !ok {
		return &AuthPasswordRecovery{}
	}
	return x
}

func (p *poolAuthPasswordRecovery) Put(x *AuthPasswordRecovery) {
	p.pool.Put(x)
}

var PoolAuthPasswordRecovery = poolAuthPasswordRecovery{}

func ResultAuthPasswordRecovery(out *MessageEnvelope, res *AuthPasswordRecovery) {
	out.Constructor = C_AuthPasswordRecovery
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthRecalled int64 = 3249025459

type poolAuthRecalled struct {
	pool sync.Pool
}

func (p *poolAuthRecalled) Get() *AuthRecalled {
	x, ok := p.pool.Get().(*AuthRecalled)
	if !ok {
		return &AuthRecalled{}
	}
	x.UpdateID = 0
	x.Available = false
	x.Force = false
	x.CurrentVersion = ""
	return x
}

func (p *poolAuthRecalled) Put(x *AuthRecalled) {
	p.pool.Put(x)
}

var PoolAuthRecalled = poolAuthRecalled{}

func ResultAuthRecalled(out *MessageEnvelope, res *AuthRecalled) {
	out.Constructor = C_AuthRecalled
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthAuthorization int64 = 1140037965

type poolAuthAuthorization struct {
	pool sync.Pool
}

func (p *poolAuthAuthorization) Get() *AuthAuthorization {
	x, ok := p.pool.Get().(*AuthAuthorization)
	if !ok {
		return &AuthAuthorization{}
	}
	return x
}

func (p *poolAuthAuthorization) Put(x *AuthAuthorization) {
	p.pool.Put(x)
}

var PoolAuthAuthorization = poolAuthAuthorization{}

func ResultAuthAuthorization(out *MessageEnvelope, res *AuthAuthorization) {
	out.Constructor = C_AuthAuthorization
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthBotAuthorization int64 = 3304560814

type poolAuthBotAuthorization struct {
	pool sync.Pool
}

func (p *poolAuthBotAuthorization) Get() *AuthBotAuthorization {
	x, ok := p.pool.Get().(*AuthBotAuthorization)
	if !ok {
		return &AuthBotAuthorization{}
	}
	return x
}

func (p *poolAuthBotAuthorization) Put(x *AuthBotAuthorization) {
	p.pool.Put(x)
}

var PoolAuthBotAuthorization = poolAuthBotAuthorization{}

func ResultAuthBotAuthorization(out *MessageEnvelope, res *AuthBotAuthorization) {
	out.Constructor = C_AuthBotAuthorization
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthCheckedPhone int64 = 2236203131

type poolAuthCheckedPhone struct {
	pool sync.Pool
}

func (p *poolAuthCheckedPhone) Get() *AuthCheckedPhone {
	x, ok := p.pool.Get().(*AuthCheckedPhone)
	if !ok {
		return &AuthCheckedPhone{}
	}
	return x
}

func (p *poolAuthCheckedPhone) Put(x *AuthCheckedPhone) {
	p.pool.Put(x)
}

var PoolAuthCheckedPhone = poolAuthCheckedPhone{}

func ResultAuthCheckedPhone(out *MessageEnvelope, res *AuthCheckedPhone) {
	out.Constructor = C_AuthCheckedPhone
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthSentCode int64 = 2375498471

type poolAuthSentCode struct {
	pool sync.Pool
}

func (p *poolAuthSentCode) Get() *AuthSentCode {
	x, ok := p.pool.Get().(*AuthSentCode)
	if !ok {
		return &AuthSentCode{}
	}
	x.SendToPhone = false
	return x
}

func (p *poolAuthSentCode) Put(x *AuthSentCode) {
	p.pool.Put(x)
}

var PoolAuthSentCode = poolAuthSentCode{}

func ResultAuthSentCode(out *MessageEnvelope, res *AuthSentCode) {
	out.Constructor = C_AuthSentCode
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
	ConstructorNames[2228369460] = "AuthRegister"
	ConstructorNames[1579606687] = "AuthBotRegister"
	ConstructorNames[2587620888] = "AuthLogin"
	ConstructorNames[3346962908] = "AuthCheckPassword"
	ConstructorNames[2711231991] = "AuthRecoverPassword"
	ConstructorNames[992431648] = "AuthLogout"
	ConstructorNames[2851553023] = "AuthLoginByToken"
	ConstructorNames[4134648516] = "AuthCheckPhone"
	ConstructorNames[3984043365] = "AuthSendCode"
	ConstructorNames[2682713491] = "AuthResendCode"
	ConstructorNames[1172029049] = "AuthRecall"
	ConstructorNames[3673422656] = "AuthDestroyKey"
	ConstructorNames[3813475914] = "AuthPasswordRecovery"
	ConstructorNames[3249025459] = "AuthRecalled"
	ConstructorNames[1140037965] = "AuthAuthorization"
	ConstructorNames[3304560814] = "AuthBotAuthorization"
	ConstructorNames[2236203131] = "AuthCheckedPhone"
	ConstructorNames[2375498471] = "AuthSentCode"
}
