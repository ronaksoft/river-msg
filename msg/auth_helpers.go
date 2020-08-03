// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

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

const C_InitConnect int64 = 4150793517

type poolInitConnect struct {
	pool sync.Pool
}

func (p *poolInitConnect) Get() *InitConnect {
	x, ok := p.pool.Get().(*InitConnect)
	if !ok {
		return &InitConnect{}
	}
	return x
}

func (p *poolInitConnect) Put(x *InitConnect) {
	p.pool.Put(x)
}

var PoolInitConnect = poolInitConnect{}

func ResultInitConnect(out *MessageEnvelope, res *InitConnect) {
	out.Constructor = C_InitConnect
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitCompleteAuth int64 = 1583178320

type poolInitCompleteAuth struct {
	pool sync.Pool
}

func (p *poolInitCompleteAuth) Get() *InitCompleteAuth {
	x, ok := p.pool.Get().(*InitCompleteAuth)
	if !ok {
		return &InitCompleteAuth{}
	}
	return x
}

func (p *poolInitCompleteAuth) Put(x *InitCompleteAuth) {
	x.ClientDHPubKey = x.ClientDHPubKey[:0]
	x.EncryptedPayload = x.EncryptedPayload[:0]
	p.pool.Put(x)
}

var PoolInitCompleteAuth = poolInitCompleteAuth{}

func ResultInitCompleteAuth(out *MessageEnvelope, res *InitCompleteAuth) {
	out.Constructor = C_InitCompleteAuth
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitConnectTest int64 = 3188015450

type poolInitConnectTest struct {
	pool sync.Pool
}

func (p *poolInitConnectTest) Get() *InitConnectTest {
	x, ok := p.pool.Get().(*InitConnectTest)
	if !ok {
		return &InitConnectTest{}
	}
	return x
}

func (p *poolInitConnectTest) Put(x *InitConnectTest) {
	p.pool.Put(x)
}

var PoolInitConnectTest = poolInitConnectTest{}

func ResultInitConnectTest(out *MessageEnvelope, res *InitConnectTest) {
	out.Constructor = C_InitConnectTest
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitBindUser int64 = 1933549113

type poolInitBindUser struct {
	pool sync.Pool
}

func (p *poolInitBindUser) Get() *InitBindUser {
	x, ok := p.pool.Get().(*InitBindUser)
	if !ok {
		return &InitBindUser{}
	}
	return x
}

func (p *poolInitBindUser) Put(x *InitBindUser) {
	p.pool.Put(x)
}

var PoolInitBindUser = poolInitBindUser{}

func ResultInitBindUser(out *MessageEnvelope, res *InitBindUser) {
	out.Constructor = C_InitBindUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AuthRegister int64 = 2228369460

type poolAuthRegister struct {
	pool sync.Pool
}

func (p *poolAuthRegister) Get() *AuthRegister {
	x, ok := p.pool.Get().(*AuthRegister)
	if !ok {
		return &AuthRegister{}
	}
	return x
}

func (p *poolAuthRegister) Put(x *AuthRegister) {
	x.LangCode = ""
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
	return x
}

func (p *poolAuthLogout) Put(x *AuthLogout) {
	x.AuthIDs = x.AuthIDs[:0]
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
	return x
}

func (p *poolAuthSendCode) Put(x *AuthSendCode) {
	x.AppHash = ""
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
	return x
}

func (p *poolAuthResendCode) Put(x *AuthResendCode) {
	x.AppHash = ""
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
	return x
}

func (p *poolAuthRecall) Put(x *AuthRecall) {
	x.Version = 0
	x.AppVersion = ""
	x.Platform = ""
	x.Vendor = ""
	x.OSVersion = ""
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

const C_InitTestAuth int64 = 2762878006

type poolInitTestAuth struct {
	pool sync.Pool
}

func (p *poolInitTestAuth) Get() *InitTestAuth {
	x, ok := p.pool.Get().(*InitTestAuth)
	if !ok {
		return &InitTestAuth{}
	}
	return x
}

func (p *poolInitTestAuth) Put(x *InitTestAuth) {
	x.AuthKey = x.AuthKey[:0]
	p.pool.Put(x)
}

var PoolInitTestAuth = poolInitTestAuth{}

func ResultInitTestAuth(out *MessageEnvelope, res *InitTestAuth) {
	out.Constructor = C_InitTestAuth
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitResponse int64 = 4130340247

type poolInitResponse struct {
	pool sync.Pool
}

func (p *poolInitResponse) Get() *InitResponse {
	x, ok := p.pool.Get().(*InitResponse)
	if !ok {
		return &InitResponse{}
	}
	return x
}

func (p *poolInitResponse) Put(x *InitResponse) {
	p.pool.Put(x)
}

var PoolInitResponse = poolInitResponse{}

func ResultInitResponse(out *MessageEnvelope, res *InitResponse) {
	out.Constructor = C_InitResponse
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitCompleteAuthInternal int64 = 2360982492

type poolInitCompleteAuthInternal struct {
	pool sync.Pool
}

func (p *poolInitCompleteAuthInternal) Get() *InitCompleteAuthInternal {
	x, ok := p.pool.Get().(*InitCompleteAuthInternal)
	if !ok {
		return &InitCompleteAuthInternal{}
	}
	return x
}

func (p *poolInitCompleteAuthInternal) Put(x *InitCompleteAuthInternal) {
	x.SecretNonce = x.SecretNonce[:0]
	p.pool.Put(x)
}

var PoolInitCompleteAuthInternal = poolInitCompleteAuthInternal{}

func ResultInitCompleteAuthInternal(out *MessageEnvelope, res *InitCompleteAuthInternal) {
	out.Constructor = C_InitCompleteAuthInternal
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitAuthCompleted int64 = 627708982

type poolInitAuthCompleted struct {
	pool sync.Pool
}

func (p *poolInitAuthCompleted) Get() *InitAuthCompleted {
	x, ok := p.pool.Get().(*InitAuthCompleted)
	if !ok {
		return &InitAuthCompleted{}
	}
	return x
}

func (p *poolInitAuthCompleted) Put(x *InitAuthCompleted) {
	x.ServerDHPubKey = x.ServerDHPubKey[:0]
	p.pool.Put(x)
}

var PoolInitAuthCompleted = poolInitAuthCompleted{}

func ResultInitAuthCompleted(out *MessageEnvelope, res *InitAuthCompleted) {
	out.Constructor = C_InitAuthCompleted
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InitUserBound int64 = 128391141

type poolInitUserBound struct {
	pool sync.Pool
}

func (p *poolInitUserBound) Get() *InitUserBound {
	x, ok := p.pool.Get().(*InitUserBound)
	if !ok {
		return &InitUserBound{}
	}
	return x
}

func (p *poolInitUserBound) Put(x *InitUserBound) {
	p.pool.Put(x)
}

var PoolInitUserBound = poolInitUserBound{}

func ResultInitUserBound(out *MessageEnvelope, res *InitUserBound) {
	out.Constructor = C_InitUserBound
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
	return x
}

func (p *poolAuthRecalled) Put(x *AuthRecalled) {
	x.UpdateID = 0
	x.Available = false
	x.Force = false
	x.CurrentVersion = ""
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
	x.ActiveSessions = 0
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
	x.AuthKey = x.AuthKey[:0]
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
	return x
}

func (p *poolAuthSentCode) Put(x *AuthSentCode) {
	x.SendToPhone = false
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
	ConstructorNames[4150793517] = "InitConnect"
	ConstructorNames[1583178320] = "InitCompleteAuth"
	ConstructorNames[3188015450] = "InitConnectTest"
	ConstructorNames[1933549113] = "InitBindUser"
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
	ConstructorNames[2762878006] = "InitTestAuth"
	ConstructorNames[4130340247] = "InitResponse"
	ConstructorNames[2360982492] = "InitCompleteAuthInternal"
	ConstructorNames[627708982] = "InitAuthCompleted"
	ConstructorNames[128391141] = "InitUserBound"
	ConstructorNames[3813475914] = "AuthPasswordRecovery"
	ConstructorNames[3249025459] = "AuthRecalled"
	ConstructorNames[1140037965] = "AuthAuthorization"
	ConstructorNames[3304560814] = "AuthBotAuthorization"
	ConstructorNames[2236203131] = "AuthCheckedPhone"
	ConstructorNames[2375498471] = "AuthSentCode"
}
