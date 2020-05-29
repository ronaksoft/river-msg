// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin.api.proto

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

const C_AdminSetWelcomeMessage int64 = 1149591874

type poolAdminSetWelcomeMessage struct {
	pool sync.Pool
}

func (p *poolAdminSetWelcomeMessage) Get() *AdminSetWelcomeMessage {
	x, ok := p.pool.Get().(*AdminSetWelcomeMessage)
	if !ok {
		return &AdminSetWelcomeMessage{}
	}
	return x
}

func (p *poolAdminSetWelcomeMessage) Put(x *AdminSetWelcomeMessage) {
	p.pool.Put(x)
}

var PoolAdminSetWelcomeMessage = poolAdminSetWelcomeMessage{}

func ResultAdminSetWelcomeMessage(out *MessageEnvelope, res *AdminSetWelcomeMessage) {
	out.Constructor = C_AdminSetWelcomeMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminGetWelcomeMessages int64 = 2794709448

type poolAdminGetWelcomeMessages struct {
	pool sync.Pool
}

func (p *poolAdminGetWelcomeMessages) Get() *AdminGetWelcomeMessages {
	x, ok := p.pool.Get().(*AdminGetWelcomeMessages)
	if !ok {
		return &AdminGetWelcomeMessages{}
	}
	return x
}

func (p *poolAdminGetWelcomeMessages) Put(x *AdminGetWelcomeMessages) {
	p.pool.Put(x)
}

var PoolAdminGetWelcomeMessages = poolAdminGetWelcomeMessages{}

func ResultAdminGetWelcomeMessages(out *MessageEnvelope, res *AdminGetWelcomeMessages) {
	out.Constructor = C_AdminGetWelcomeMessages
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminSetPushProvider int64 = 1758606947

type poolAdminSetPushProvider struct {
	pool sync.Pool
}

func (p *poolAdminSetPushProvider) Get() *AdminSetPushProvider {
	x, ok := p.pool.Get().(*AdminSetPushProvider)
	if !ok {
		return &AdminSetPushProvider{}
	}
	return x
}

func (p *poolAdminSetPushProvider) Put(x *AdminSetPushProvider) {
	p.pool.Put(x)
}

var PoolAdminSetPushProvider = poolAdminSetPushProvider{}

func ResultAdminSetPushProvider(out *MessageEnvelope, res *AdminSetPushProvider) {
	out.Constructor = C_AdminSetPushProvider
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminGetPushProviders int64 = 4257963974

type poolAdminGetPushProviders struct {
	pool sync.Pool
}

func (p *poolAdminGetPushProviders) Get() *AdminGetPushProviders {
	x, ok := p.pool.Get().(*AdminGetPushProviders)
	if !ok {
		return &AdminGetPushProviders{}
	}
	return x
}

func (p *poolAdminGetPushProviders) Put(x *AdminGetPushProviders) {
	p.pool.Put(x)
}

var PoolAdminGetPushProviders = poolAdminGetPushProviders{}

func ResultAdminGetPushProviders(out *MessageEnvelope, res *AdminGetPushProviders) {
	out.Constructor = C_AdminGetPushProviders
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminSetVersion int64 = 1311023404

type poolAdminSetVersion struct {
	pool sync.Pool
}

func (p *poolAdminSetVersion) Get() *AdminSetVersion {
	x, ok := p.pool.Get().(*AdminSetVersion)
	if !ok {
		return &AdminSetVersion{}
	}
	return x
}

func (p *poolAdminSetVersion) Put(x *AdminSetVersion) {
	p.pool.Put(x)
}

var PoolAdminSetVersion = poolAdminSetVersion{}

func ResultAdminSetVersion(out *MessageEnvelope, res *AdminSetVersion) {
	out.Constructor = C_AdminSetVersion
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminGetVersions int64 = 934752256

type poolAdminGetVersions struct {
	pool sync.Pool
}

func (p *poolAdminGetVersions) Get() *AdminGetVersions {
	x, ok := p.pool.Get().(*AdminGetVersions)
	if !ok {
		return &AdminGetVersions{}
	}
	return x
}

func (p *poolAdminGetVersions) Put(x *AdminGetVersions) {
	p.pool.Put(x)
}

var PoolAdminGetVersions = poolAdminGetVersions{}

func ResultAdminGetVersions(out *MessageEnvelope, res *AdminGetVersions) {
	out.Constructor = C_AdminGetVersions
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminSetToken int64 = 2892519162

type poolAdminSetToken struct {
	pool sync.Pool
}

func (p *poolAdminSetToken) Get() *AdminSetToken {
	x, ok := p.pool.Get().(*AdminSetToken)
	if !ok {
		return &AdminSetToken{}
	}
	return x
}

func (p *poolAdminSetToken) Put(x *AdminSetToken) {
	p.pool.Put(x)
}

var PoolAdminSetToken = poolAdminSetToken{}

func ResultAdminSetToken(out *MessageEnvelope, res *AdminSetToken) {
	out.Constructor = C_AdminSetToken
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminDeleteToken int64 = 3154441897

type poolAdminDeleteToken struct {
	pool sync.Pool
}

func (p *poolAdminDeleteToken) Get() *AdminDeleteToken {
	x, ok := p.pool.Get().(*AdminDeleteToken)
	if !ok {
		return &AdminDeleteToken{}
	}
	return x
}

func (p *poolAdminDeleteToken) Put(x *AdminDeleteToken) {
	p.pool.Put(x)
}

var PoolAdminDeleteToken = poolAdminDeleteToken{}

func ResultAdminDeleteToken(out *MessageEnvelope, res *AdminDeleteToken) {
	out.Constructor = C_AdminDeleteToken
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AdminToken int64 = 2895609620

type poolAdminToken struct {
	pool sync.Pool
}

func (p *poolAdminToken) Get() *AdminToken {
	x, ok := p.pool.Get().(*AdminToken)
	if !ok {
		return &AdminToken{}
	}
	return x
}

func (p *poolAdminToken) Put(x *AdminToken) {
	p.pool.Put(x)
}

var PoolAdminToken = poolAdminToken{}

func ResultAdminToken(out *MessageEnvelope, res *AdminToken) {
	out.Constructor = C_AdminToken
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_WelcomeMessagesMany int64 = 414982091

type poolWelcomeMessagesMany struct {
	pool sync.Pool
}

func (p *poolWelcomeMessagesMany) Get() *WelcomeMessagesMany {
	x, ok := p.pool.Get().(*WelcomeMessagesMany)
	if !ok {
		return &WelcomeMessagesMany{}
	}
	x.Messages = x.Messages[:0]
	return x
}

func (p *poolWelcomeMessagesMany) Put(x *WelcomeMessagesMany) {
	p.pool.Put(x)
}

var PoolWelcomeMessagesMany = poolWelcomeMessagesMany{}

func ResultWelcomeMessagesMany(out *MessageEnvelope, res *WelcomeMessagesMany) {
	out.Constructor = C_WelcomeMessagesMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_VersionsMany int64 = 2123920547

type poolVersionsMany struct {
	pool sync.Pool
}

func (p *poolVersionsMany) Get() *VersionsMany {
	x, ok := p.pool.Get().(*VersionsMany)
	if !ok {
		return &VersionsMany{}
	}
	x.Versions = x.Versions[:0]
	return x
}

func (p *poolVersionsMany) Put(x *VersionsMany) {
	p.pool.Put(x)
}

var PoolVersionsMany = poolVersionsMany{}

func ResultVersionsMany(out *MessageEnvelope, res *VersionsMany) {
	out.Constructor = C_VersionsMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PushProvidersMany int64 = 5873573

type poolPushProvidersMany struct {
	pool sync.Pool
}

func (p *poolPushProvidersMany) Get() *PushProvidersMany {
	x, ok := p.pool.Get().(*PushProvidersMany)
	if !ok {
		return &PushProvidersMany{}
	}
	x.Providers = x.Providers[:0]
	return x
}

func (p *poolPushProvidersMany) Put(x *PushProvidersMany) {
	p.pool.Put(x)
}

var PoolPushProvidersMany = poolPushProvidersMany{}

func ResultPushProvidersMany(out *MessageEnvelope, res *PushProvidersMany) {
	out.Constructor = C_PushProvidersMany
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
	ConstructorNames[1149591874] = "AdminSetWelcomeMessage"
	ConstructorNames[2794709448] = "AdminGetWelcomeMessages"
	ConstructorNames[1758606947] = "AdminSetPushProvider"
	ConstructorNames[4257963974] = "AdminGetPushProviders"
	ConstructorNames[1311023404] = "AdminSetVersion"
	ConstructorNames[934752256] = "AdminGetVersions"
	ConstructorNames[2892519162] = "AdminSetToken"
	ConstructorNames[3154441897] = "AdminDeleteToken"
	ConstructorNames[2895609620] = "AdminToken"
	ConstructorNames[414982091] = "WelcomeMessagesMany"
	ConstructorNames[2123920547] = "VersionsMany"
	ConstructorNames[5873573] = "PushProvidersMany"
}