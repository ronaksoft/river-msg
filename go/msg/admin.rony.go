package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_AdminBroadcastMessage int64 = 4051598653

type poolAdminBroadcastMessage struct {
	pool sync.Pool
}

func (p *poolAdminBroadcastMessage) Get() *AdminBroadcastMessage {
	x, ok := p.pool.Get().(*AdminBroadcastMessage)
	if !ok {
		return &AdminBroadcastMessage{}
	}
	return x
}

func (p *poolAdminBroadcastMessage) Put(x *AdminBroadcastMessage) {
	x.Body = ""
	x.ReceiverIDs = x.ReceiverIDs[:0]
	x.Entities = x.Entities[:0]
	x.MediaType = 0
	x.MediaData = x.MediaData[:0]
	p.pool.Put(x)
}

var PoolAdminBroadcastMessage = poolAdminBroadcastMessage{}

const C_AdminSetWelcomeMessage int64 = 4036161171

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
	x.AccessToken = ""
	x.Lang = ""
	x.Template = ""
	p.pool.Put(x)
}

var PoolAdminSetWelcomeMessage = poolAdminSetWelcomeMessage{}

const C_AdminGetWelcomeMessages int64 = 1475548165

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
	x.AccessToken = ""
	p.pool.Put(x)
}

var PoolAdminGetWelcomeMessages = poolAdminGetWelcomeMessages{}

const C_AdminDeleteWelcomeMessage int64 = 1498710407

type poolAdminDeleteWelcomeMessage struct {
	pool sync.Pool
}

func (p *poolAdminDeleteWelcomeMessage) Get() *AdminDeleteWelcomeMessage {
	x, ok := p.pool.Get().(*AdminDeleteWelcomeMessage)
	if !ok {
		return &AdminDeleteWelcomeMessage{}
	}
	return x
}

func (p *poolAdminDeleteWelcomeMessage) Put(x *AdminDeleteWelcomeMessage) {
	x.AccessToken = ""
	x.Lang = ""
	p.pool.Put(x)
}

var PoolAdminDeleteWelcomeMessage = poolAdminDeleteWelcomeMessage{}

const C_AdminSetPushProvider int64 = 894806066

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
	x.AccessToken = ""
	if x.Provider != nil {
		PoolPushProvider.Put(x.Provider)
		x.Provider = nil
	}
	p.pool.Put(x)
}

var PoolAdminSetPushProvider = poolAdminSetPushProvider{}

const C_AdminGetPushProviders int64 = 3791305018

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
	x.AccessToken = ""
	p.pool.Put(x)
}

var PoolAdminGetPushProviders = poolAdminGetPushProviders{}

const C_AdminDeletePushProvider int64 = 2655579321

type poolAdminDeletePushProvider struct {
	pool sync.Pool
}

func (p *poolAdminDeletePushProvider) Get() *AdminDeletePushProvider {
	x, ok := p.pool.Get().(*AdminDeletePushProvider)
	if !ok {
		return &AdminDeletePushProvider{}
	}
	return x
}

func (p *poolAdminDeletePushProvider) Put(x *AdminDeletePushProvider) {
	x.AccessToken = ""
	x.Name = ""
	x.Type = 0
	p.pool.Put(x)
}

var PoolAdminDeletePushProvider = poolAdminDeletePushProvider{}

const C_AdminSetVersion int64 = 1484557854

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
	x.AccessToken = ""
	if x.Version != nil {
		PoolVersion.Put(x.Version)
		x.Version = nil
	}
	p.pool.Put(x)
}

var PoolAdminSetVersion = poolAdminSetVersion{}

const C_AdminGetVersions int64 = 4285938095

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
	x.AccessToken = ""
	p.pool.Put(x)
}

var PoolAdminGetVersions = poolAdminGetVersions{}

const C_AdminSetToken int64 = 583431571

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
	x.AccessToken = ""
	x.Privilege = 0
	p.pool.Put(x)
}

var PoolAdminSetToken = poolAdminSetToken{}

const C_AdminDeleteToken int64 = 1959132934

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
	x.AccessToken = ""
	x.DeletedToken = ""
	p.pool.Put(x)
}

var PoolAdminDeleteToken = poolAdminDeleteToken{}

const C_AdminReserveUsername int64 = 697333997

type poolAdminReserveUsername struct {
	pool sync.Pool
}

func (p *poolAdminReserveUsername) Get() *AdminReserveUsername {
	x, ok := p.pool.Get().(*AdminReserveUsername)
	if !ok {
		return &AdminReserveUsername{}
	}
	return x
}

func (p *poolAdminReserveUsername) Put(x *AdminReserveUsername) {
	x.AccessToken = ""
	x.Usernames = x.Usernames[:0]
	x.Delete = false
	p.pool.Put(x)
}

var PoolAdminReserveUsername = poolAdminReserveUsername{}

const C_AdminGetReservedUsernames int64 = 3979011259

type poolAdminGetReservedUsernames struct {
	pool sync.Pool
}

func (p *poolAdminGetReservedUsernames) Get() *AdminGetReservedUsernames {
	x, ok := p.pool.Get().(*AdminGetReservedUsernames)
	if !ok {
		return &AdminGetReservedUsernames{}
	}
	return x
}

func (p *poolAdminGetReservedUsernames) Put(x *AdminGetReservedUsernames) {
	x.AccessToken = ""
	p.pool.Put(x)
}

var PoolAdminGetReservedUsernames = poolAdminGetReservedUsernames{}

const C_AdminTeamCreate int64 = 2968516162

type poolAdminTeamCreate struct {
	pool sync.Pool
}

func (p *poolAdminTeamCreate) Get() *AdminTeamCreate {
	x, ok := p.pool.Get().(*AdminTeamCreate)
	if !ok {
		return &AdminTeamCreate{}
	}
	return x
}

func (p *poolAdminTeamCreate) Put(x *AdminTeamCreate) {
	x.AccessToken = ""
	x.Capacity = 0
	x.ExpireDate = 0
	x.Community = false
	x.Title = ""
	x.CreatorID = 0
	p.pool.Put(x)
}

var PoolAdminTeamCreate = poolAdminTeamCreate{}

const C_AdminToken int64 = 715104459

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
	x.Privilege = 0
	x.Token = ""
	p.pool.Put(x)
}

var PoolAdminToken = poolAdminToken{}

const C_WelcomeMessagesMany int64 = 4147727409

type poolWelcomeMessagesMany struct {
	pool sync.Pool
}

func (p *poolWelcomeMessagesMany) Get() *WelcomeMessagesMany {
	x, ok := p.pool.Get().(*WelcomeMessagesMany)
	if !ok {
		return &WelcomeMessagesMany{}
	}
	return x
}

func (p *poolWelcomeMessagesMany) Put(x *WelcomeMessagesMany) {
	x.Messages = x.Messages[:0]
	x.Count = 0
	p.pool.Put(x)
}

var PoolWelcomeMessagesMany = poolWelcomeMessagesMany{}

const C_VersionsMany int64 = 1851057214

type poolVersionsMany struct {
	pool sync.Pool
}

func (p *poolVersionsMany) Get() *VersionsMany {
	x, ok := p.pool.Get().(*VersionsMany)
	if !ok {
		return &VersionsMany{}
	}
	return x
}

func (p *poolVersionsMany) Put(x *VersionsMany) {
	x.Versions = x.Versions[:0]
	x.Count = 0
	p.pool.Put(x)
}

var PoolVersionsMany = poolVersionsMany{}

const C_PushProvidersMany int64 = 1190715605

type poolPushProvidersMany struct {
	pool sync.Pool
}

func (p *poolPushProvidersMany) Get() *PushProvidersMany {
	x, ok := p.pool.Get().(*PushProvidersMany)
	if !ok {
		return &PushProvidersMany{}
	}
	return x
}

func (p *poolPushProvidersMany) Put(x *PushProvidersMany) {
	x.Providers = x.Providers[:0]
	x.Count = 0
	p.pool.Put(x)
}

var PoolPushProvidersMany = poolPushProvidersMany{}

const C_WelcomeMessage int64 = 2710083316

type poolWelcomeMessage struct {
	pool sync.Pool
}

func (p *poolWelcomeMessage) Get() *WelcomeMessage {
	x, ok := p.pool.Get().(*WelcomeMessage)
	if !ok {
		return &WelcomeMessage{}
	}
	return x
}

func (p *poolWelcomeMessage) Put(x *WelcomeMessage) {
	x.Lang = ""
	x.Template = ""
	p.pool.Put(x)
}

var PoolWelcomeMessage = poolWelcomeMessage{}

const C_PushProvider int64 = 742531531

type poolPushProvider struct {
	pool sync.Pool
}

func (p *poolPushProvider) Get() *PushProvider {
	x, ok := p.pool.Get().(*PushProvider)
	if !ok {
		return &PushProvider{}
	}
	return x
}

func (p *poolPushProvider) Put(x *PushProvider) {
	x.Name = ""
	x.Type = 0
	x.TestMode = false
	x.Credentials = x.Credentials[:0]
	x.KeyID = ""
	x.TeamID = ""
	x.Topic = ""
	p.pool.Put(x)
}

var PoolPushProvider = poolPushProvider{}

const C_Version int64 = 3565878624

type poolVersion struct {
	pool sync.Pool
}

func (p *poolVersion) Get() *Version {
	x, ok := p.pool.Get().(*Version)
	if !ok {
		return &Version{}
	}
	return x
}

func (p *poolVersion) Put(x *Version) {
	x.Vendor = ""
	x.Stage = ""
	x.OS = ""
	x.MinVersion = ""
	x.CurrentVersion = ""
	x.ForcedVersions = x.ForcedVersions[:0]
	p.pool.Put(x)
}

var PoolVersion = poolVersion{}

const C_ReservedUsernames int64 = 337474487

type poolReservedUsernames struct {
	pool sync.Pool
}

func (p *poolReservedUsernames) Get() *ReservedUsernames {
	x, ok := p.pool.Get().(*ReservedUsernames)
	if !ok {
		return &ReservedUsernames{}
	}
	return x
}

func (p *poolReservedUsernames) Put(x *ReservedUsernames) {
	x.Usernames = x.Usernames[:0]
	x.Count = 0
	p.pool.Put(x)
}

var PoolReservedUsernames = poolReservedUsernames{}

func init() {
	registry.RegisterConstructor(4051598653, "msg.AdminBroadcastMessage")
	registry.RegisterConstructor(4036161171, "msg.AdminSetWelcomeMessage")
	registry.RegisterConstructor(1475548165, "msg.AdminGetWelcomeMessages")
	registry.RegisterConstructor(1498710407, "msg.AdminDeleteWelcomeMessage")
	registry.RegisterConstructor(894806066, "msg.AdminSetPushProvider")
	registry.RegisterConstructor(3791305018, "msg.AdminGetPushProviders")
	registry.RegisterConstructor(2655579321, "msg.AdminDeletePushProvider")
	registry.RegisterConstructor(1484557854, "msg.AdminSetVersion")
	registry.RegisterConstructor(4285938095, "msg.AdminGetVersions")
	registry.RegisterConstructor(583431571, "msg.AdminSetToken")
	registry.RegisterConstructor(1959132934, "msg.AdminDeleteToken")
	registry.RegisterConstructor(697333997, "msg.AdminReserveUsername")
	registry.RegisterConstructor(3979011259, "msg.AdminGetReservedUsernames")
	registry.RegisterConstructor(2968516162, "msg.AdminTeamCreate")
	registry.RegisterConstructor(715104459, "msg.AdminToken")
	registry.RegisterConstructor(4147727409, "msg.WelcomeMessagesMany")
	registry.RegisterConstructor(1851057214, "msg.VersionsMany")
	registry.RegisterConstructor(1190715605, "msg.PushProvidersMany")
	registry.RegisterConstructor(2710083316, "msg.WelcomeMessage")
	registry.RegisterConstructor(742531531, "msg.PushProvider")
	registry.RegisterConstructor(3565878624, "msg.Version")
	registry.RegisterConstructor(337474487, "msg.ReservedUsernames")
}

func (x *AdminBroadcastMessage) DeepCopy(z *AdminBroadcastMessage) {
	z.Body = x.Body
	z.ReceiverIDs = append(z.ReceiverIDs[:0], x.ReceiverIDs...)
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.MediaType = x.MediaType
	z.MediaData = append(z.MediaData[:0], x.MediaData...)
}

func (x *AdminSetWelcomeMessage) DeepCopy(z *AdminSetWelcomeMessage) {
	z.AccessToken = x.AccessToken
	z.Lang = x.Lang
	z.Template = x.Template
}

func (x *AdminGetWelcomeMessages) DeepCopy(z *AdminGetWelcomeMessages) {
	z.AccessToken = x.AccessToken
}

func (x *AdminDeleteWelcomeMessage) DeepCopy(z *AdminDeleteWelcomeMessage) {
	z.AccessToken = x.AccessToken
	z.Lang = x.Lang
}

func (x *AdminSetPushProvider) DeepCopy(z *AdminSetPushProvider) {
	z.AccessToken = x.AccessToken
	if x.Provider != nil {
		z.Provider = PoolPushProvider.Get()
		x.Provider.DeepCopy(z.Provider)
	}
}

func (x *AdminGetPushProviders) DeepCopy(z *AdminGetPushProviders) {
	z.AccessToken = x.AccessToken
}

func (x *AdminDeletePushProvider) DeepCopy(z *AdminDeletePushProvider) {
	z.AccessToken = x.AccessToken
	z.Name = x.Name
	z.Type = x.Type
}

func (x *AdminSetVersion) DeepCopy(z *AdminSetVersion) {
	z.AccessToken = x.AccessToken
	if x.Version != nil {
		z.Version = PoolVersion.Get()
		x.Version.DeepCopy(z.Version)
	}
}

func (x *AdminGetVersions) DeepCopy(z *AdminGetVersions) {
	z.AccessToken = x.AccessToken
}

func (x *AdminSetToken) DeepCopy(z *AdminSetToken) {
	z.AccessToken = x.AccessToken
	z.Privilege = x.Privilege
}

func (x *AdminDeleteToken) DeepCopy(z *AdminDeleteToken) {
	z.AccessToken = x.AccessToken
	z.DeletedToken = x.DeletedToken
}

func (x *AdminReserveUsername) DeepCopy(z *AdminReserveUsername) {
	z.AccessToken = x.AccessToken
	z.Usernames = append(z.Usernames[:0], x.Usernames...)
	z.Delete = x.Delete
}

func (x *AdminGetReservedUsernames) DeepCopy(z *AdminGetReservedUsernames) {
	z.AccessToken = x.AccessToken
}

func (x *AdminTeamCreate) DeepCopy(z *AdminTeamCreate) {
	z.AccessToken = x.AccessToken
	z.Capacity = x.Capacity
	z.ExpireDate = x.ExpireDate
	z.Community = x.Community
	z.Title = x.Title
	z.CreatorID = x.CreatorID
}

func (x *AdminToken) DeepCopy(z *AdminToken) {
	z.Privilege = x.Privilege
	z.Token = x.Token
}

func (x *WelcomeMessagesMany) DeepCopy(z *WelcomeMessagesMany) {
	for idx := range x.Messages {
		if x.Messages[idx] != nil {
			xx := PoolWelcomeMessage.Get()
			x.Messages[idx].DeepCopy(xx)
			z.Messages = append(z.Messages, xx)
		}
	}
	z.Count = x.Count
}

func (x *VersionsMany) DeepCopy(z *VersionsMany) {
	for idx := range x.Versions {
		if x.Versions[idx] != nil {
			xx := PoolVersion.Get()
			x.Versions[idx].DeepCopy(xx)
			z.Versions = append(z.Versions, xx)
		}
	}
	z.Count = x.Count
}

func (x *PushProvidersMany) DeepCopy(z *PushProvidersMany) {
	for idx := range x.Providers {
		if x.Providers[idx] != nil {
			xx := PoolPushProvider.Get()
			x.Providers[idx].DeepCopy(xx)
			z.Providers = append(z.Providers, xx)
		}
	}
	z.Count = x.Count
}

func (x *WelcomeMessage) DeepCopy(z *WelcomeMessage) {
	z.Lang = x.Lang
	z.Template = x.Template
}

func (x *PushProvider) DeepCopy(z *PushProvider) {
	z.Name = x.Name
	z.Type = x.Type
	z.TestMode = x.TestMode
	z.Credentials = append(z.Credentials[:0], x.Credentials...)
	z.KeyID = x.KeyID
	z.TeamID = x.TeamID
	z.Topic = x.Topic
}

func (x *Version) DeepCopy(z *Version) {
	z.Vendor = x.Vendor
	z.Stage = x.Stage
	z.OS = x.OS
	z.MinVersion = x.MinVersion
	z.CurrentVersion = x.CurrentVersion
	z.ForcedVersions = append(z.ForcedVersions[:0], x.ForcedVersions...)
}

func (x *ReservedUsernames) DeepCopy(z *ReservedUsernames) {
	z.Usernames = append(z.Usernames[:0], x.Usernames...)
	z.Count = x.Count
}

func (x *AdminBroadcastMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminBroadcastMessage, x)
}

func (x *AdminSetWelcomeMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminSetWelcomeMessage, x)
}

func (x *AdminGetWelcomeMessages) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminGetWelcomeMessages, x)
}

func (x *AdminDeleteWelcomeMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminDeleteWelcomeMessage, x)
}

func (x *AdminSetPushProvider) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminSetPushProvider, x)
}

func (x *AdminGetPushProviders) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminGetPushProviders, x)
}

func (x *AdminDeletePushProvider) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminDeletePushProvider, x)
}

func (x *AdminSetVersion) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminSetVersion, x)
}

func (x *AdminGetVersions) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminGetVersions, x)
}

func (x *AdminSetToken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminSetToken, x)
}

func (x *AdminDeleteToken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminDeleteToken, x)
}

func (x *AdminReserveUsername) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminReserveUsername, x)
}

func (x *AdminGetReservedUsernames) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminGetReservedUsernames, x)
}

func (x *AdminTeamCreate) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminTeamCreate, x)
}

func (x *AdminToken) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AdminToken, x)
}

func (x *WelcomeMessagesMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_WelcomeMessagesMany, x)
}

func (x *VersionsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_VersionsMany, x)
}

func (x *PushProvidersMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PushProvidersMany, x)
}

func (x *WelcomeMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_WelcomeMessage, x)
}

func (x *PushProvider) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PushProvider, x)
}

func (x *Version) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Version, x)
}

func (x *ReservedUsernames) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ReservedUsernames, x)
}

func (x *AdminBroadcastMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminSetWelcomeMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminGetWelcomeMessages) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminDeleteWelcomeMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminSetPushProvider) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminGetPushProviders) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminDeletePushProvider) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminSetVersion) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminGetVersions) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminSetToken) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminDeleteToken) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminReserveUsername) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminGetReservedUsernames) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminTeamCreate) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminToken) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *WelcomeMessagesMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *VersionsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *PushProvidersMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *WelcomeMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *PushProvider) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Version) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ReservedUsernames) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AdminBroadcastMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminSetWelcomeMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminGetWelcomeMessages) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminDeleteWelcomeMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminSetPushProvider) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminGetPushProviders) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminDeletePushProvider) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminSetVersion) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminGetVersions) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminSetToken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminDeleteToken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminReserveUsername) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminGetReservedUsernames) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminTeamCreate) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AdminToken) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *WelcomeMessagesMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *VersionsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PushProvidersMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *WelcomeMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PushProvider) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Version) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ReservedUsernames) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}
