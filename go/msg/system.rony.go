// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: system.proto

package msg

import (
	bytes "bytes"
	edge "github.com/ronaksoft/rony/edge"
	pools "github.com/ronaksoft/rony/pools"
	registry "github.com/ronaksoft/rony/registry"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

var _ = pools.Imported

const C_SystemGetServerKeys int64 = 2510636156

type poolSystemGetServerKeys struct {
	pool sync.Pool
}

func (p *poolSystemGetServerKeys) Get() *SystemGetServerKeys {
	x, ok := p.pool.Get().(*SystemGetServerKeys)
	if !ok {
		x = &SystemGetServerKeys{}
	}

	return x
}

func (p *poolSystemGetServerKeys) Put(x *SystemGetServerKeys) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolSystemGetServerKeys = poolSystemGetServerKeys{}

func (x *SystemGetServerKeys) DeepCopy(z *SystemGetServerKeys) {
}

func (x *SystemGetServerKeys) Clone() *SystemGetServerKeys {
	z := &SystemGetServerKeys{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetServerKeys) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetServerKeys) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetServerKeys) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetServerKeys) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetServerKeys) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetServerKeys, x)
}

const C_SystemGetServerTime int64 = 1321179349

type poolSystemGetServerTime struct {
	pool sync.Pool
}

func (p *poolSystemGetServerTime) Get() *SystemGetServerTime {
	x, ok := p.pool.Get().(*SystemGetServerTime)
	if !ok {
		x = &SystemGetServerTime{}
	}

	return x
}

func (p *poolSystemGetServerTime) Put(x *SystemGetServerTime) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolSystemGetServerTime = poolSystemGetServerTime{}

func (x *SystemGetServerTime) DeepCopy(z *SystemGetServerTime) {
}

func (x *SystemGetServerTime) Clone() *SystemGetServerTime {
	z := &SystemGetServerTime{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetServerTime) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetServerTime) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetServerTime) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetServerTime) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetServerTime) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetServerTime, x)
}

const C_SystemGetInfo int64 = 1486296237

type poolSystemGetInfo struct {
	pool sync.Pool
}

func (p *poolSystemGetInfo) Get() *SystemGetInfo {
	x, ok := p.pool.Get().(*SystemGetInfo)
	if !ok {
		x = &SystemGetInfo{}
	}

	return x
}

func (p *poolSystemGetInfo) Put(x *SystemGetInfo) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolSystemGetInfo = poolSystemGetInfo{}

func (x *SystemGetInfo) DeepCopy(z *SystemGetInfo) {
}

func (x *SystemGetInfo) Clone() *SystemGetInfo {
	z := &SystemGetInfo{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetInfo) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetInfo) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetInfo) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetInfo) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetInfo) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetInfo, x)
}

const C_SystemGetSalts int64 = 1705203315

type poolSystemGetSalts struct {
	pool sync.Pool
}

func (p *poolSystemGetSalts) Get() *SystemGetSalts {
	x, ok := p.pool.Get().(*SystemGetSalts)
	if !ok {
		x = &SystemGetSalts{}
	}

	return x
}

func (p *poolSystemGetSalts) Put(x *SystemGetSalts) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolSystemGetSalts = poolSystemGetSalts{}

func (x *SystemGetSalts) DeepCopy(z *SystemGetSalts) {
}

func (x *SystemGetSalts) Clone() *SystemGetSalts {
	z := &SystemGetSalts{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetSalts) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetSalts) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetSalts) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetSalts) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetSalts) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetSalts, x)
}

const C_SystemGetConfig int64 = 1910333714

type poolSystemGetConfig struct {
	pool sync.Pool
}

func (p *poolSystemGetConfig) Get() *SystemGetConfig {
	x, ok := p.pool.Get().(*SystemGetConfig)
	if !ok {
		x = &SystemGetConfig{}
	}

	return x
}

func (p *poolSystemGetConfig) Put(x *SystemGetConfig) {
	if x == nil {
		return
	}

	p.pool.Put(x)
}

var PoolSystemGetConfig = poolSystemGetConfig{}

func (x *SystemGetConfig) DeepCopy(z *SystemGetConfig) {
}

func (x *SystemGetConfig) Clone() *SystemGetConfig {
	z := &SystemGetConfig{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetConfig) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetConfig) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetConfig) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetConfig) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetConfig) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetConfig, x)
}

const C_SystemUploadUsage int64 = 3056393082

type poolSystemUploadUsage struct {
	pool sync.Pool
}

func (p *poolSystemUploadUsage) Get() *SystemUploadUsage {
	x, ok := p.pool.Get().(*SystemUploadUsage)
	if !ok {
		x = &SystemUploadUsage{}
	}

	return x
}

func (p *poolSystemUploadUsage) Put(x *SystemUploadUsage) {
	if x == nil {
		return
	}

	for _, z := range x.Usage {
		PoolClientUsage.Put(z)
	}
	x.Usage = x.Usage[:0]

	p.pool.Put(x)
}

var PoolSystemUploadUsage = poolSystemUploadUsage{}

func (x *SystemUploadUsage) DeepCopy(z *SystemUploadUsage) {
	for idx := range x.Usage {
		if x.Usage[idx] == nil {
			continue
		}
		xx := PoolClientUsage.Get()
		x.Usage[idx].DeepCopy(xx)
		z.Usage = append(z.Usage, xx)
	}
}

func (x *SystemUploadUsage) Clone() *SystemUploadUsage {
	z := &SystemUploadUsage{}
	x.DeepCopy(z)
	return z
}

func (x *SystemUploadUsage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemUploadUsage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemUploadUsage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemUploadUsage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemUploadUsage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemUploadUsage, x)
}

const C_SystemGetResponse int64 = 1415676946

type poolSystemGetResponse struct {
	pool sync.Pool
}

func (p *poolSystemGetResponse) Get() *SystemGetResponse {
	x, ok := p.pool.Get().(*SystemGetResponse)
	if !ok {
		x = &SystemGetResponse{}
	}

	return x
}

func (p *poolSystemGetResponse) Put(x *SystemGetResponse) {
	if x == nil {
		return
	}

	x.RequestIDs = x.RequestIDs[:0]

	p.pool.Put(x)
}

var PoolSystemGetResponse = poolSystemGetResponse{}

func (x *SystemGetResponse) DeepCopy(z *SystemGetResponse) {
	z.RequestIDs = append(z.RequestIDs[:0], x.RequestIDs...)
}

func (x *SystemGetResponse) Clone() *SystemGetResponse {
	z := &SystemGetResponse{}
	x.DeepCopy(z)
	return z
}

func (x *SystemGetResponse) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemGetResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemGetResponse) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemGetResponse) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemGetResponse) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemGetResponse, x)
}

const C_ClientUsage int64 = 453987802

type poolClientUsage struct {
	pool sync.Pool
}

func (p *poolClientUsage) Get() *ClientUsage {
	x, ok := p.pool.Get().(*ClientUsage)
	if !ok {
		x = &ClientUsage{}
	}

	return x
}

func (p *poolClientUsage) Put(x *ClientUsage) {
	if x == nil {
		return
	}

	x.Year = 0
	x.Month = 0
	x.Day = 0
	x.UserID = 0
	x.ForegroundTime = 0
	x.AvgResponseTime = 0
	x.TotalRequests = 0
	x.ReceivedMessages = 0
	x.SentMessages = 0
	x.ReceivedMedia = 0
	x.SentMedia = 0
	x.UploadBytes = 0
	x.DownloadBytes = 0

	p.pool.Put(x)
}

var PoolClientUsage = poolClientUsage{}

func (x *ClientUsage) DeepCopy(z *ClientUsage) {
	z.Year = x.Year
	z.Month = x.Month
	z.Day = x.Day
	z.UserID = x.UserID
	z.ForegroundTime = x.ForegroundTime
	z.AvgResponseTime = x.AvgResponseTime
	z.TotalRequests = x.TotalRequests
	z.ReceivedMessages = x.ReceivedMessages
	z.SentMessages = x.SentMessages
	z.ReceivedMedia = x.ReceivedMedia
	z.SentMedia = x.SentMedia
	z.UploadBytes = x.UploadBytes
	z.DownloadBytes = x.DownloadBytes
}

func (x *ClientUsage) Clone() *ClientUsage {
	z := &ClientUsage{}
	x.DeepCopy(z)
	return z
}

func (x *ClientUsage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ClientUsage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *ClientUsage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *ClientUsage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *ClientUsage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ClientUsage, x)
}

const C_SystemConfig int64 = 367036084

type poolSystemConfig struct {
	pool sync.Pool
}

func (p *poolSystemConfig) Get() *SystemConfig {
	x, ok := p.pool.Get().(*SystemConfig)
	if !ok {
		x = &SystemConfig{}
	}

	return x
}

func (p *poolSystemConfig) Put(x *SystemConfig) {
	if x == nil {
		return
	}

	x.GifBot = ""
	x.WikiBot = ""
	x.TestMode = false
	x.PhoneCallEnabled = false
	x.ExpireOn = 0
	x.GroupMaxSize = 0
	x.ForwardedMaxCount = 0
	x.OnlineUpdatePeriodInSec = 0
	x.EditTimeLimitInSec = 0
	x.RevokeTimeLimitInSec = 0
	x.PinnedDialogsMaxCount = 0
	x.UrlPrefix = 0
	x.MessageMaxLength = 0
	x.CaptionMaxLength = 0
	for _, z := range x.DCs {
		PoolDataCenter.Put(z)
	}
	x.DCs = x.DCs[:0]
	x.MaxLabels = 0
	x.TopPeerDecayRate = 0
	x.TopPeerMaxStep = 0
	x.MaxActiveSessions = 0
	x.Reactions = x.Reactions[:0]
	x.MaxUploadSize = 0
	x.MaxUploadPartSize = 0
	x.MaxUploadParts = 0

	p.pool.Put(x)
}

var PoolSystemConfig = poolSystemConfig{}

func (x *SystemConfig) DeepCopy(z *SystemConfig) {
	z.GifBot = x.GifBot
	z.WikiBot = x.WikiBot
	z.TestMode = x.TestMode
	z.PhoneCallEnabled = x.PhoneCallEnabled
	z.ExpireOn = x.ExpireOn
	z.GroupMaxSize = x.GroupMaxSize
	z.ForwardedMaxCount = x.ForwardedMaxCount
	z.OnlineUpdatePeriodInSec = x.OnlineUpdatePeriodInSec
	z.EditTimeLimitInSec = x.EditTimeLimitInSec
	z.RevokeTimeLimitInSec = x.RevokeTimeLimitInSec
	z.PinnedDialogsMaxCount = x.PinnedDialogsMaxCount
	z.UrlPrefix = x.UrlPrefix
	z.MessageMaxLength = x.MessageMaxLength
	z.CaptionMaxLength = x.CaptionMaxLength
	for idx := range x.DCs {
		if x.DCs[idx] == nil {
			continue
		}
		xx := PoolDataCenter.Get()
		x.DCs[idx].DeepCopy(xx)
		z.DCs = append(z.DCs, xx)
	}
	z.MaxLabels = x.MaxLabels
	z.TopPeerDecayRate = x.TopPeerDecayRate
	z.TopPeerMaxStep = x.TopPeerMaxStep
	z.MaxActiveSessions = x.MaxActiveSessions
	z.Reactions = append(z.Reactions[:0], x.Reactions...)
	z.MaxUploadSize = x.MaxUploadSize
	z.MaxUploadPartSize = x.MaxUploadPartSize
	z.MaxUploadParts = x.MaxUploadParts
}

func (x *SystemConfig) Clone() *SystemConfig {
	z := &SystemConfig{}
	x.DeepCopy(z)
	return z
}

func (x *SystemConfig) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemConfig) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemConfig) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemConfig) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemConfig) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemConfig, x)
}

const C_DataCenter int64 = 3431386561

type poolDataCenter struct {
	pool sync.Pool
}

func (p *poolDataCenter) Get() *DataCenter {
	x, ok := p.pool.Get().(*DataCenter)
	if !ok {
		x = &DataCenter{}
	}

	return x
}

func (p *poolDataCenter) Put(x *DataCenter) {
	if x == nil {
		return
	}

	x.IP = ""
	x.Port = 0
	x.Http = false
	x.Websocket = false
	x.Quic = false

	p.pool.Put(x)
}

var PoolDataCenter = poolDataCenter{}

func (x *DataCenter) DeepCopy(z *DataCenter) {
	z.IP = x.IP
	z.Port = x.Port
	z.Http = x.Http
	z.Websocket = x.Websocket
	z.Quic = x.Quic
}

func (x *DataCenter) Clone() *DataCenter {
	z := &DataCenter{}
	x.DeepCopy(z)
	return z
}

func (x *DataCenter) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *DataCenter) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DataCenter) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *DataCenter) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *DataCenter) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_DataCenter, x)
}

const C_SystemSalts int64 = 871116906

type poolSystemSalts struct {
	pool sync.Pool
}

func (p *poolSystemSalts) Get() *SystemSalts {
	x, ok := p.pool.Get().(*SystemSalts)
	if !ok {
		x = &SystemSalts{}
	}

	return x
}

func (p *poolSystemSalts) Put(x *SystemSalts) {
	if x == nil {
		return
	}

	x.Salts = x.Salts[:0]
	x.StartsFrom = 0
	x.Duration = 0

	p.pool.Put(x)
}

var PoolSystemSalts = poolSystemSalts{}

func (x *SystemSalts) DeepCopy(z *SystemSalts) {
	z.Salts = append(z.Salts[:0], x.Salts...)
	z.StartsFrom = x.StartsFrom
	z.Duration = x.Duration
}

func (x *SystemSalts) Clone() *SystemSalts {
	z := &SystemSalts{}
	x.DeepCopy(z)
	return z
}

func (x *SystemSalts) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemSalts) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemSalts) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemSalts) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemSalts) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemSalts, x)
}

const C_AppUpdate int64 = 1100207036

type poolAppUpdate struct {
	pool sync.Pool
}

func (p *poolAppUpdate) Get() *AppUpdate {
	x, ok := p.pool.Get().(*AppUpdate)
	if !ok {
		x = &AppUpdate{}
	}

	return x
}

func (p *poolAppUpdate) Put(x *AppUpdate) {
	if x == nil {
		return
	}

	x.Available = false
	x.Mandatory = false
	x.Identifier = ""
	x.VersionName = ""
	x.DownloadUrl = ""
	x.Description = ""
	x.DisplayInterval = 0

	p.pool.Put(x)
}

var PoolAppUpdate = poolAppUpdate{}

func (x *AppUpdate) DeepCopy(z *AppUpdate) {
	z.Available = x.Available
	z.Mandatory = x.Mandatory
	z.Identifier = x.Identifier
	z.VersionName = x.VersionName
	z.DownloadUrl = x.DownloadUrl
	z.Description = x.Description
	z.DisplayInterval = x.DisplayInterval
}

func (x *AppUpdate) Clone() *AppUpdate {
	z := &AppUpdate{}
	x.DeepCopy(z)
	return z
}

func (x *AppUpdate) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AppUpdate) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *AppUpdate) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *AppUpdate) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *AppUpdate) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AppUpdate, x)
}

const C_SystemInfo int64 = 2754948760

type poolSystemInfo struct {
	pool sync.Pool
}

func (p *poolSystemInfo) Get() *SystemInfo {
	x, ok := p.pool.Get().(*SystemInfo)
	if !ok {
		x = &SystemInfo{}
	}

	return x
}

func (p *poolSystemInfo) Put(x *SystemInfo) {
	if x == nil {
		return
	}

	x.WorkGroupName = ""
	x.BigPhotoUrl = ""
	x.SmallPhotoUrl = ""
	x.StorageUrl = ""

	p.pool.Put(x)
}

var PoolSystemInfo = poolSystemInfo{}

func (x *SystemInfo) DeepCopy(z *SystemInfo) {
	z.WorkGroupName = x.WorkGroupName
	z.BigPhotoUrl = x.BigPhotoUrl
	z.SmallPhotoUrl = x.SmallPhotoUrl
	z.StorageUrl = x.StorageUrl
}

func (x *SystemInfo) Clone() *SystemInfo {
	z := &SystemInfo{}
	x.DeepCopy(z)
	return z
}

func (x *SystemInfo) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemInfo) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemInfo) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemInfo) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemInfo) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemInfo, x)
}

const C_SystemServerTime int64 = 2854614486

type poolSystemServerTime struct {
	pool sync.Pool
}

func (p *poolSystemServerTime) Get() *SystemServerTime {
	x, ok := p.pool.Get().(*SystemServerTime)
	if !ok {
		x = &SystemServerTime{}
	}

	return x
}

func (p *poolSystemServerTime) Put(x *SystemServerTime) {
	if x == nil {
		return
	}

	x.Timestamp = 0

	p.pool.Put(x)
}

var PoolSystemServerTime = poolSystemServerTime{}

func (x *SystemServerTime) DeepCopy(z *SystemServerTime) {
	z.Timestamp = x.Timestamp
}

func (x *SystemServerTime) Clone() *SystemServerTime {
	z := &SystemServerTime{}
	x.DeepCopy(z)
	return z
}

func (x *SystemServerTime) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemServerTime) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemServerTime) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemServerTime) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemServerTime) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemServerTime, x)
}

const C_SystemKeys int64 = 3677510435

type poolSystemKeys struct {
	pool sync.Pool
}

func (p *poolSystemKeys) Get() *SystemKeys {
	x, ok := p.pool.Get().(*SystemKeys)
	if !ok {
		x = &SystemKeys{}
	}

	return x
}

func (p *poolSystemKeys) Put(x *SystemKeys) {
	if x == nil {
		return
	}

	for _, z := range x.RSAPublicKeys {
		PoolRSAPublicKey.Put(z)
	}
	x.RSAPublicKeys = x.RSAPublicKeys[:0]
	for _, z := range x.DHGroups {
		PoolDHGroup.Put(z)
	}
	x.DHGroups = x.DHGroups[:0]

	p.pool.Put(x)
}

var PoolSystemKeys = poolSystemKeys{}

func (x *SystemKeys) DeepCopy(z *SystemKeys) {
	for idx := range x.RSAPublicKeys {
		if x.RSAPublicKeys[idx] == nil {
			continue
		}
		xx := PoolRSAPublicKey.Get()
		x.RSAPublicKeys[idx].DeepCopy(xx)
		z.RSAPublicKeys = append(z.RSAPublicKeys, xx)
	}
	for idx := range x.DHGroups {
		if x.DHGroups[idx] == nil {
			continue
		}
		xx := PoolDHGroup.Get()
		x.DHGroups[idx].DeepCopy(xx)
		z.DHGroups = append(z.DHGroups, xx)
	}
}

func (x *SystemKeys) Clone() *SystemKeys {
	z := &SystemKeys{}
	x.DeepCopy(z)
	return z
}

func (x *SystemKeys) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SystemKeys) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *SystemKeys) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *SystemKeys) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SystemKeys) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SystemKeys, x)
}

func init() {
	registry.RegisterConstructor(2510636156, "SystemGetServerKeys")
	registry.RegisterConstructor(1321179349, "SystemGetServerTime")
	registry.RegisterConstructor(1486296237, "SystemGetInfo")
	registry.RegisterConstructor(1705203315, "SystemGetSalts")
	registry.RegisterConstructor(1910333714, "SystemGetConfig")
	registry.RegisterConstructor(3056393082, "SystemUploadUsage")
	registry.RegisterConstructor(1415676946, "SystemGetResponse")
	registry.RegisterConstructor(453987802, "ClientUsage")
	registry.RegisterConstructor(367036084, "SystemConfig")
	registry.RegisterConstructor(3431386561, "DataCenter")
	registry.RegisterConstructor(871116906, "SystemSalts")
	registry.RegisterConstructor(1100207036, "AppUpdate")
	registry.RegisterConstructor(2754948760, "SystemInfo")
	registry.RegisterConstructor(2854614486, "SystemServerTime")
	registry.RegisterConstructor(3677510435, "SystemKeys")
}

var _ = bytes.MinRead
