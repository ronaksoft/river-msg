// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: system.proto

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

const C_SystemGetPublicKeys int64 = 1191522796

type poolSystemGetPublicKeys struct {
	pool sync.Pool
}

func (p *poolSystemGetPublicKeys) Get() *SystemGetPublicKeys {
	x, ok := p.pool.Get().(*SystemGetPublicKeys)
	if !ok {
		return &SystemGetPublicKeys{}
	}
	return x
}

func (p *poolSystemGetPublicKeys) Put(x *SystemGetPublicKeys) {
	p.pool.Put(x)
}

var PoolSystemGetPublicKeys = poolSystemGetPublicKeys{}

func ResultSystemGetPublicKeys(out *MessageEnvelope, res *SystemGetPublicKeys) {
	out.Constructor = C_SystemGetPublicKeys
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetDHGroups int64 = 1786665018

type poolSystemGetDHGroups struct {
	pool sync.Pool
}

func (p *poolSystemGetDHGroups) Get() *SystemGetDHGroups {
	x, ok := p.pool.Get().(*SystemGetDHGroups)
	if !ok {
		return &SystemGetDHGroups{}
	}
	return x
}

func (p *poolSystemGetDHGroups) Put(x *SystemGetDHGroups) {
	p.pool.Put(x)
}

var PoolSystemGetDHGroups = poolSystemGetDHGroups{}

func ResultSystemGetDHGroups(out *MessageEnvelope, res *SystemGetDHGroups) {
	out.Constructor = C_SystemGetDHGroups
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetServerTime int64 = 1321179349

type poolSystemGetServerTime struct {
	pool sync.Pool
}

func (p *poolSystemGetServerTime) Get() *SystemGetServerTime {
	x, ok := p.pool.Get().(*SystemGetServerTime)
	if !ok {
		return &SystemGetServerTime{}
	}
	return x
}

func (p *poolSystemGetServerTime) Put(x *SystemGetServerTime) {
	p.pool.Put(x)
}

var PoolSystemGetServerTime = poolSystemGetServerTime{}

func ResultSystemGetServerTime(out *MessageEnvelope, res *SystemGetServerTime) {
	out.Constructor = C_SystemGetServerTime
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetInfo int64 = 1486296237

type poolSystemGetInfo struct {
	pool sync.Pool
}

func (p *poolSystemGetInfo) Get() *SystemGetInfo {
	x, ok := p.pool.Get().(*SystemGetInfo)
	if !ok {
		return &SystemGetInfo{}
	}
	return x
}

func (p *poolSystemGetInfo) Put(x *SystemGetInfo) {
	p.pool.Put(x)
}

var PoolSystemGetInfo = poolSystemGetInfo{}

func ResultSystemGetInfo(out *MessageEnvelope, res *SystemGetInfo) {
	out.Constructor = C_SystemGetInfo
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetSalts int64 = 1705203315

type poolSystemGetSalts struct {
	pool sync.Pool
}

func (p *poolSystemGetSalts) Get() *SystemGetSalts {
	x, ok := p.pool.Get().(*SystemGetSalts)
	if !ok {
		return &SystemGetSalts{}
	}
	return x
}

func (p *poolSystemGetSalts) Put(x *SystemGetSalts) {
	p.pool.Put(x)
}

var PoolSystemGetSalts = poolSystemGetSalts{}

func ResultSystemGetSalts(out *MessageEnvelope, res *SystemGetSalts) {
	out.Constructor = C_SystemGetSalts
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetConfig int64 = 1910333714

type poolSystemGetConfig struct {
	pool sync.Pool
}

func (p *poolSystemGetConfig) Get() *SystemGetConfig {
	x, ok := p.pool.Get().(*SystemGetConfig)
	if !ok {
		return &SystemGetConfig{}
	}
	return x
}

func (p *poolSystemGetConfig) Put(x *SystemGetConfig) {
	p.pool.Put(x)
}

var PoolSystemGetConfig = poolSystemGetConfig{}

func ResultSystemGetConfig(out *MessageEnvelope, res *SystemGetConfig) {
	out.Constructor = C_SystemGetConfig
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemUploadUsage int64 = 3056393082

type poolSystemUploadUsage struct {
	pool sync.Pool
}

func (p *poolSystemUploadUsage) Get() *SystemUploadUsage {
	x, ok := p.pool.Get().(*SystemUploadUsage)
	if !ok {
		return &SystemUploadUsage{}
	}
	return x
}

func (p *poolSystemUploadUsage) Put(x *SystemUploadUsage) {
	x.Usage = x.Usage[:0]
	p.pool.Put(x)
}

var PoolSystemUploadUsage = poolSystemUploadUsage{}

func ResultSystemUploadUsage(out *MessageEnvelope, res *SystemUploadUsage) {
	out.Constructor = C_SystemUploadUsage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemGetResponse int64 = 1415676946

type poolSystemGetResponse struct {
	pool sync.Pool
}

func (p *poolSystemGetResponse) Get() *SystemGetResponse {
	x, ok := p.pool.Get().(*SystemGetResponse)
	if !ok {
		return &SystemGetResponse{}
	}
	return x
}

func (p *poolSystemGetResponse) Put(x *SystemGetResponse) {
	x.RequestIDs = x.RequestIDs[:0]
	p.pool.Put(x)
}

var PoolSystemGetResponse = poolSystemGetResponse{}

func ResultSystemGetResponse(out *MessageEnvelope, res *SystemGetResponse) {
	out.Constructor = C_SystemGetResponse
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientUsage int64 = 453987802

type poolClientUsage struct {
	pool sync.Pool
}

func (p *poolClientUsage) Get() *ClientUsage {
	x, ok := p.pool.Get().(*ClientUsage)
	if !ok {
		return &ClientUsage{}
	}
	return x
}

func (p *poolClientUsage) Put(x *ClientUsage) {
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

func ResultClientUsage(out *MessageEnvelope, res *ClientUsage) {
	out.Constructor = C_ClientUsage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemConfig int64 = 367036084

type poolSystemConfig struct {
	pool sync.Pool
}

func (p *poolSystemConfig) Get() *SystemConfig {
	x, ok := p.pool.Get().(*SystemConfig)
	if !ok {
		return &SystemConfig{}
	}
	return x
}

func (p *poolSystemConfig) Put(x *SystemConfig) {
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

func ResultSystemConfig(out *MessageEnvelope, res *SystemConfig) {
	out.Constructor = C_SystemConfig
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_DataCenter int64 = 3431386561

type poolDataCenter struct {
	pool sync.Pool
}

func (p *poolDataCenter) Get() *DataCenter {
	x, ok := p.pool.Get().(*DataCenter)
	if !ok {
		return &DataCenter{}
	}
	return x
}

func (p *poolDataCenter) Put(x *DataCenter) {
	x.IP = ""
	x.Port = 0
	x.Http = false
	x.Websocket = false
	x.Quic = false
	p.pool.Put(x)
}

var PoolDataCenter = poolDataCenter{}

func ResultDataCenter(out *MessageEnvelope, res *DataCenter) {
	out.Constructor = C_DataCenter
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemSalts int64 = 871116906

type poolSystemSalts struct {
	pool sync.Pool
}

func (p *poolSystemSalts) Get() *SystemSalts {
	x, ok := p.pool.Get().(*SystemSalts)
	if !ok {
		return &SystemSalts{}
	}
	return x
}

func (p *poolSystemSalts) Put(x *SystemSalts) {
	x.Salts = x.Salts[:0]
	x.StartsFrom = 0
	x.Duration = 0
	p.pool.Put(x)
}

var PoolSystemSalts = poolSystemSalts{}

func ResultSystemSalts(out *MessageEnvelope, res *SystemSalts) {
	out.Constructor = C_SystemSalts
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_AppUpdate int64 = 1100207036

type poolAppUpdate struct {
	pool sync.Pool
}

func (p *poolAppUpdate) Get() *AppUpdate {
	x, ok := p.pool.Get().(*AppUpdate)
	if !ok {
		return &AppUpdate{}
	}
	return x
}

func (p *poolAppUpdate) Put(x *AppUpdate) {
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

func ResultAppUpdate(out *MessageEnvelope, res *AppUpdate) {
	out.Constructor = C_AppUpdate
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemInfo int64 = 2754948760

type poolSystemInfo struct {
	pool sync.Pool
}

func (p *poolSystemInfo) Get() *SystemInfo {
	x, ok := p.pool.Get().(*SystemInfo)
	if !ok {
		return &SystemInfo{}
	}
	return x
}

func (p *poolSystemInfo) Put(x *SystemInfo) {
	x.WorkGroupName = ""
	x.BigPhotoUrl = ""
	x.SmallPhotoUrl = ""
	x.StorageUrl = ""
	p.pool.Put(x)
}

var PoolSystemInfo = poolSystemInfo{}

func ResultSystemInfo(out *MessageEnvelope, res *SystemInfo) {
	out.Constructor = C_SystemInfo
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemServerTime int64 = 2854614486

type poolSystemServerTime struct {
	pool sync.Pool
}

func (p *poolSystemServerTime) Get() *SystemServerTime {
	x, ok := p.pool.Get().(*SystemServerTime)
	if !ok {
		return &SystemServerTime{}
	}
	return x
}

func (p *poolSystemServerTime) Put(x *SystemServerTime) {
	x.Timestamp = 0
	p.pool.Put(x)
}

var PoolSystemServerTime = poolSystemServerTime{}

func ResultSystemServerTime(out *MessageEnvelope, res *SystemServerTime) {
	out.Constructor = C_SystemServerTime
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemPublicKeys int64 = 2745130223

type poolSystemPublicKeys struct {
	pool sync.Pool
}

func (p *poolSystemPublicKeys) Get() *SystemPublicKeys {
	x, ok := p.pool.Get().(*SystemPublicKeys)
	if !ok {
		return &SystemPublicKeys{}
	}
	return x
}

func (p *poolSystemPublicKeys) Put(x *SystemPublicKeys) {
	x.RSAPublicKeys = x.RSAPublicKeys[:0]
	p.pool.Put(x)
}

var PoolSystemPublicKeys = poolSystemPublicKeys{}

func ResultSystemPublicKeys(out *MessageEnvelope, res *SystemPublicKeys) {
	out.Constructor = C_SystemPublicKeys
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_SystemDHGroups int64 = 2890748083

type poolSystemDHGroups struct {
	pool sync.Pool
}

func (p *poolSystemDHGroups) Get() *SystemDHGroups {
	x, ok := p.pool.Get().(*SystemDHGroups)
	if !ok {
		return &SystemDHGroups{}
	}
	return x
}

func (p *poolSystemDHGroups) Put(x *SystemDHGroups) {
	x.DHGroups = x.DHGroups[:0]
	p.pool.Put(x)
}

var PoolSystemDHGroups = poolSystemDHGroups{}

func ResultSystemDHGroups(out *MessageEnvelope, res *SystemDHGroups) {
	out.Constructor = C_SystemDHGroups
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
	ConstructorNames[1191522796] = "SystemGetPublicKeys"
	ConstructorNames[1786665018] = "SystemGetDHGroups"
	ConstructorNames[1321179349] = "SystemGetServerTime"
	ConstructorNames[1486296237] = "SystemGetInfo"
	ConstructorNames[1705203315] = "SystemGetSalts"
	ConstructorNames[1910333714] = "SystemGetConfig"
	ConstructorNames[3056393082] = "SystemUploadUsage"
	ConstructorNames[1415676946] = "SystemGetResponse"
	ConstructorNames[453987802] = "ClientUsage"
	ConstructorNames[367036084] = "SystemConfig"
	ConstructorNames[3431386561] = "DataCenter"
	ConstructorNames[871116906] = "SystemSalts"
	ConstructorNames[1100207036] = "AppUpdate"
	ConstructorNames[2754948760] = "SystemInfo"
	ConstructorNames[2854614486] = "SystemServerTime"
	ConstructorNames[2745130223] = "SystemPublicKeys"
	ConstructorNames[2890748083] = "SystemDHGroups"
}