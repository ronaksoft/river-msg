// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client.proto

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

const C_ClientSendMessageMedia int64 = 1095038539

type poolClientSendMessageMedia struct {
	pool sync.Pool
}

func (p *poolClientSendMessageMedia) Get() *ClientSendMessageMedia {
	x, ok := p.pool.Get().(*ClientSendMessageMedia)
	if !ok {
		return &ClientSendMessageMedia{}
	}
	x.Caption = ""
	x.FileName = ""
	x.FilePath = ""
	x.ThumbFilePath = ""
	x.FileMIME = ""
	x.ThumbMIME = ""
	x.ReplyTo = 0
	x.ClearDraft = false
	x.Attributes = x.Attributes[:0]
	x.FileUploadID = ""
	x.ThumbUploadID = ""
	x.FileID = 0
	x.ThumbID = 0
	x.FileTotalParts = 0
	x.Entities = x.Entities[:0]
	return x
}

func (p *poolClientSendMessageMedia) Put(x *ClientSendMessageMedia) {
	p.pool.Put(x)
}

var PoolClientSendMessageMedia = poolClientSendMessageMedia{}

func ResultClientSendMessageMedia(out *MessageEnvelope, res *ClientSendMessageMedia) {
	out.Constructor = C_ClientSendMessageMedia
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGlobalSearch int64 = 1742781507

type poolClientGlobalSearch struct {
	pool sync.Pool
}

func (p *poolClientGlobalSearch) Get() *ClientGlobalSearch {
	x, ok := p.pool.Get().(*ClientGlobalSearch)
	if !ok {
		return &ClientGlobalSearch{}
	}
	x.Text = ""
	x.LabelIDs = x.LabelIDs[:0]
	x.Peer = nil
	x.SenderID = 0
	return x
}

func (p *poolClientGlobalSearch) Put(x *ClientGlobalSearch) {
	p.pool.Put(x)
}

var PoolClientGlobalSearch = poolClientGlobalSearch{}

func ResultClientGlobalSearch(out *MessageEnvelope, res *ClientGlobalSearch) {
	out.Constructor = C_ClientGlobalSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientContactSearch int64 = 1793449803

type poolClientContactSearch struct {
	pool sync.Pool
}

func (p *poolClientContactSearch) Get() *ClientContactSearch {
	x, ok := p.pool.Get().(*ClientContactSearch)
	if !ok {
		return &ClientContactSearch{}
	}
	return x
}

func (p *poolClientContactSearch) Put(x *ClientContactSearch) {
	p.pool.Put(x)
}

var PoolClientContactSearch = poolClientContactSearch{}

func ResultClientContactSearch(out *MessageEnvelope, res *ClientContactSearch) {
	out.Constructor = C_ClientContactSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGetCachedMedia int64 = 856595701

type poolClientGetCachedMedia struct {
	pool sync.Pool
}

func (p *poolClientGetCachedMedia) Get() *ClientGetCachedMedia {
	x, ok := p.pool.Get().(*ClientGetCachedMedia)
	if !ok {
		return &ClientGetCachedMedia{}
	}
	return x
}

func (p *poolClientGetCachedMedia) Put(x *ClientGetCachedMedia) {
	p.pool.Put(x)
}

var PoolClientGetCachedMedia = poolClientGetCachedMedia{}

func ResultClientGetCachedMedia(out *MessageEnvelope, res *ClientGetCachedMedia) {
	out.Constructor = C_ClientGetCachedMedia
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientClearCachedMedia int64 = 1199927718

type poolClientClearCachedMedia struct {
	pool sync.Pool
}

func (p *poolClientClearCachedMedia) Get() *ClientClearCachedMedia {
	x, ok := p.pool.Get().(*ClientClearCachedMedia)
	if !ok {
		return &ClientClearCachedMedia{}
	}
	x.Peer = nil
	x.MediaTypes = x.MediaTypes[:0]
	return x
}

func (p *poolClientClearCachedMedia) Put(x *ClientClearCachedMedia) {
	p.pool.Put(x)
}

var PoolClientClearCachedMedia = poolClientClearCachedMedia{}

func ResultClientClearCachedMedia(out *MessageEnvelope, res *ClientClearCachedMedia) {
	out.Constructor = C_ClientClearCachedMedia
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGetLastBotKeyboard int64 = 177544569

type poolClientGetLastBotKeyboard struct {
	pool sync.Pool
}

func (p *poolClientGetLastBotKeyboard) Get() *ClientGetLastBotKeyboard {
	x, ok := p.pool.Get().(*ClientGetLastBotKeyboard)
	if !ok {
		return &ClientGetLastBotKeyboard{}
	}
	return x
}

func (p *poolClientGetLastBotKeyboard) Put(x *ClientGetLastBotKeyboard) {
	p.pool.Put(x)
}

var PoolClientGetLastBotKeyboard = poolClientGetLastBotKeyboard{}

func ResultClientGetLastBotKeyboard(out *MessageEnvelope, res *ClientGetLastBotKeyboard) {
	out.Constructor = C_ClientGetLastBotKeyboard
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGetMediaHistory int64 = 1354863379

type poolClientGetMediaHistory struct {
	pool sync.Pool
}

func (p *poolClientGetMediaHistory) Get() *ClientGetMediaHistory {
	x, ok := p.pool.Get().(*ClientGetMediaHistory)
	if !ok {
		return &ClientGetMediaHistory{}
	}
	return x
}

func (p *poolClientGetMediaHistory) Put(x *ClientGetMediaHistory) {
	p.pool.Put(x)
}

var PoolClientGetMediaHistory = poolClientGetMediaHistory{}

func ResultClientGetMediaHistory(out *MessageEnvelope, res *ClientGetMediaHistory) {
	out.Constructor = C_ClientGetMediaHistory
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGetRecentSearch int64 = 2622949116

type poolClientGetRecentSearch struct {
	pool sync.Pool
}

func (p *poolClientGetRecentSearch) Get() *ClientGetRecentSearch {
	x, ok := p.pool.Get().(*ClientGetRecentSearch)
	if !ok {
		return &ClientGetRecentSearch{}
	}
	return x
}

func (p *poolClientGetRecentSearch) Put(x *ClientGetRecentSearch) {
	p.pool.Put(x)
}

var PoolClientGetRecentSearch = poolClientGetRecentSearch{}

func ResultClientGetRecentSearch(out *MessageEnvelope, res *ClientGetRecentSearch) {
	out.Constructor = C_ClientGetRecentSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientPutRecentSearch int64 = 629582533

type poolClientPutRecentSearch struct {
	pool sync.Pool
}

func (p *poolClientPutRecentSearch) Get() *ClientPutRecentSearch {
	x, ok := p.pool.Get().(*ClientPutRecentSearch)
	if !ok {
		return &ClientPutRecentSearch{}
	}
	return x
}

func (p *poolClientPutRecentSearch) Put(x *ClientPutRecentSearch) {
	p.pool.Put(x)
}

var PoolClientPutRecentSearch = poolClientPutRecentSearch{}

func ResultClientPutRecentSearch(out *MessageEnvelope, res *ClientPutRecentSearch) {
	out.Constructor = C_ClientPutRecentSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientRemoveRecentSearch int64 = 1281490259

type poolClientRemoveRecentSearch struct {
	pool sync.Pool
}

func (p *poolClientRemoveRecentSearch) Get() *ClientRemoveRecentSearch {
	x, ok := p.pool.Get().(*ClientRemoveRecentSearch)
	if !ok {
		return &ClientRemoveRecentSearch{}
	}
	return x
}

func (p *poolClientRemoveRecentSearch) Put(x *ClientRemoveRecentSearch) {
	p.pool.Put(x)
}

var PoolClientRemoveRecentSearch = poolClientRemoveRecentSearch{}

func ResultClientRemoveRecentSearch(out *MessageEnvelope, res *ClientRemoveRecentSearch) {
	out.Constructor = C_ClientRemoveRecentSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientRemoveAllRecentSearches int64 = 3599155822

type poolClientRemoveAllRecentSearches struct {
	pool sync.Pool
}

func (p *poolClientRemoveAllRecentSearches) Get() *ClientRemoveAllRecentSearches {
	x, ok := p.pool.Get().(*ClientRemoveAllRecentSearches)
	if !ok {
		return &ClientRemoveAllRecentSearches{}
	}
	return x
}

func (p *poolClientRemoveAllRecentSearches) Put(x *ClientRemoveAllRecentSearches) {
	p.pool.Put(x)
}

var PoolClientRemoveAllRecentSearches = poolClientRemoveAllRecentSearches{}

func ResultClientRemoveAllRecentSearches(out *MessageEnvelope, res *ClientRemoveAllRecentSearches) {
	out.Constructor = C_ClientRemoveAllRecentSearches
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientGetSavedGifs int64 = 3028067090

type poolClientGetSavedGifs struct {
	pool sync.Pool
}

func (p *poolClientGetSavedGifs) Get() *ClientGetSavedGifs {
	x, ok := p.pool.Get().(*ClientGetSavedGifs)
	if !ok {
		return &ClientGetSavedGifs{}
	}
	return x
}

func (p *poolClientGetSavedGifs) Put(x *ClientGetSavedGifs) {
	p.pool.Put(x)
}

var PoolClientGetSavedGifs = poolClientGetSavedGifs{}

func ResultClientGetSavedGifs(out *MessageEnvelope, res *ClientGetSavedGifs) {
	out.Constructor = C_ClientGetSavedGifs
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientPendingMessage int64 = 2164891929

type poolClientPendingMessage struct {
	pool sync.Pool
}

func (p *poolClientPendingMessage) Get() *ClientPendingMessage {
	x, ok := p.pool.Get().(*ClientPendingMessage)
	if !ok {
		return &ClientPendingMessage{}
	}
	x.Entities = x.Entities[:0]
	x.MediaType = 0
	x.Media = nil
	x.Media = x.Media[:0]
	x.ClearDraft = false
	x.FileUploadID = ""
	x.ThumbUploadID = ""
	x.FileID = 0
	x.ThumbID = 0
	x.Sha256 = nil
	x.Sha256 = x.Sha256[:0]
	x.ServerFile = nil
	return x
}

func (p *poolClientPendingMessage) Put(x *ClientPendingMessage) {
	p.pool.Put(x)
}

var PoolClientPendingMessage = poolClientPendingMessage{}

func ResultClientPendingMessage(out *MessageEnvelope, res *ClientPendingMessage) {
	out.Constructor = C_ClientPendingMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientSearchResult int64 = 2957647709

type poolClientSearchResult struct {
	pool sync.Pool
}

func (p *poolClientSearchResult) Get() *ClientSearchResult {
	x, ok := p.pool.Get().(*ClientSearchResult)
	if !ok {
		return &ClientSearchResult{}
	}
	x.Messages = x.Messages[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	x.MatchedUsers = x.MatchedUsers[:0]
	x.MatchedGroups = x.MatchedGroups[:0]
	return x
}

func (p *poolClientSearchResult) Put(x *ClientSearchResult) {
	p.pool.Put(x)
}

var PoolClientSearchResult = poolClientSearchResult{}

func ResultClientSearchResult(out *MessageEnvelope, res *ClientSearchResult) {
	out.Constructor = C_ClientSearchResult
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientFilesMany int64 = 1414992553

type poolClientFilesMany struct {
	pool sync.Pool
}

func (p *poolClientFilesMany) Get() *ClientFilesMany {
	x, ok := p.pool.Get().(*ClientFilesMany)
	if !ok {
		return &ClientFilesMany{}
	}
	x.Gifs = x.Gifs[:0]
	return x
}

func (p *poolClientFilesMany) Put(x *ClientFilesMany) {
	p.pool.Put(x)
}

var PoolClientFilesMany = poolClientFilesMany{}

func ResultClientFilesMany(out *MessageEnvelope, res *ClientFilesMany) {
	out.Constructor = C_ClientFilesMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientFile int64 = 155127968

type poolClientFile struct {
	pool sync.Pool
}

func (p *poolClientFile) Get() *ClientFile {
	x, ok := p.pool.Get().(*ClientFile)
	if !ok {
		return &ClientFile{}
	}
	x.MimeType = ""
	x.UserID = 0
	x.GroupID = 0
	x.MessageID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.Version = 0
	x.Extension = ""
	x.MD5Checksum = ""
	x.WallpaperID = 0
	return x
}

func (p *poolClientFile) Put(x *ClientFile) {
	p.pool.Put(x)
}

var PoolClientFile = poolClientFile{}

func ResultClientFile(out *MessageEnvelope, res *ClientFile) {
	out.Constructor = C_ClientFile
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientFileStatus int64 = 2731095358

type poolClientFileStatus struct {
	pool sync.Pool
}

func (p *poolClientFileStatus) Get() *ClientFileStatus {
	x, ok := p.pool.Get().(*ClientFileStatus)
	if !ok {
		return &ClientFileStatus{}
	}
	return x
}

func (p *poolClientFileStatus) Put(x *ClientFileStatus) {
	p.pool.Put(x)
}

var PoolClientFileStatus = poolClientFileStatus{}

func ResultClientFileStatus(out *MessageEnvelope, res *ClientFileStatus) {
	out.Constructor = C_ClientFileStatus
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientCachedMediaInfo int64 = 442767121

type poolClientCachedMediaInfo struct {
	pool sync.Pool
}

func (p *poolClientCachedMediaInfo) Get() *ClientCachedMediaInfo {
	x, ok := p.pool.Get().(*ClientCachedMediaInfo)
	if !ok {
		return &ClientCachedMediaInfo{}
	}
	x.MediaInfo = x.MediaInfo[:0]
	return x
}

func (p *poolClientCachedMediaInfo) Put(x *ClientCachedMediaInfo) {
	p.pool.Put(x)
}

var PoolClientCachedMediaInfo = poolClientCachedMediaInfo{}

func ResultClientCachedMediaInfo(out *MessageEnvelope, res *ClientCachedMediaInfo) {
	out.Constructor = C_ClientCachedMediaInfo
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientPeerMediaInfo int64 = 2711408875

type poolClientPeerMediaInfo struct {
	pool sync.Pool
}

func (p *poolClientPeerMediaInfo) Get() *ClientPeerMediaInfo {
	x, ok := p.pool.Get().(*ClientPeerMediaInfo)
	if !ok {
		return &ClientPeerMediaInfo{}
	}
	x.Media = x.Media[:0]
	return x
}

func (p *poolClientPeerMediaInfo) Put(x *ClientPeerMediaInfo) {
	p.pool.Put(x)
}

var PoolClientPeerMediaInfo = poolClientPeerMediaInfo{}

func ResultClientPeerMediaInfo(out *MessageEnvelope, res *ClientPeerMediaInfo) {
	out.Constructor = C_ClientPeerMediaInfo
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ClientMediaSize int64 = 1541024203

type poolClientMediaSize struct {
	pool sync.Pool
}

func (p *poolClientMediaSize) Get() *ClientMediaSize {
	x, ok := p.pool.Get().(*ClientMediaSize)
	if !ok {
		return &ClientMediaSize{}
	}
	return x
}

func (p *poolClientMediaSize) Put(x *ClientMediaSize) {
	p.pool.Put(x)
}

var PoolClientMediaSize = poolClientMediaSize{}

func ResultClientMediaSize(out *MessageEnvelope, res *ClientMediaSize) {
	out.Constructor = C_ClientMediaSize
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_RecentSearch int64 = 2045409949

type poolRecentSearch struct {
	pool sync.Pool
}

func (p *poolRecentSearch) Get() *RecentSearch {
	x, ok := p.pool.Get().(*RecentSearch)
	if !ok {
		return &RecentSearch{}
	}
	return x
}

func (p *poolRecentSearch) Put(x *RecentSearch) {
	p.pool.Put(x)
}

var PoolRecentSearch = poolRecentSearch{}

func ResultRecentSearch(out *MessageEnvelope, res *RecentSearch) {
	out.Constructor = C_RecentSearch
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_RecentSearchMany int64 = 1217165147

type poolRecentSearchMany struct {
	pool sync.Pool
}

func (p *poolRecentSearchMany) Get() *RecentSearchMany {
	x, ok := p.pool.Get().(*RecentSearchMany)
	if !ok {
		return &RecentSearchMany{}
	}
	x.RecentSearches = x.RecentSearches[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	return x
}

func (p *poolRecentSearchMany) Put(x *RecentSearchMany) {
	p.pool.Put(x)
}

var PoolRecentSearchMany = poolRecentSearchMany{}

func ResultRecentSearchMany(out *MessageEnvelope, res *RecentSearchMany) {
	out.Constructor = C_RecentSearchMany
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
	ConstructorNames[1095038539] = "ClientSendMessageMedia"
	ConstructorNames[1742781507] = "ClientGlobalSearch"
	ConstructorNames[1793449803] = "ClientContactSearch"
	ConstructorNames[856595701] = "ClientGetCachedMedia"
	ConstructorNames[1199927718] = "ClientClearCachedMedia"
	ConstructorNames[177544569] = "ClientGetLastBotKeyboard"
	ConstructorNames[1354863379] = "ClientGetMediaHistory"
	ConstructorNames[2622949116] = "ClientGetRecentSearch"
	ConstructorNames[629582533] = "ClientPutRecentSearch"
	ConstructorNames[1281490259] = "ClientRemoveRecentSearch"
	ConstructorNames[3599155822] = "ClientRemoveAllRecentSearches"
	ConstructorNames[3028067090] = "ClientGetSavedGifs"
	ConstructorNames[2164891929] = "ClientPendingMessage"
	ConstructorNames[2957647709] = "ClientSearchResult"
	ConstructorNames[1414992553] = "ClientFilesMany"
	ConstructorNames[155127968] = "ClientFile"
	ConstructorNames[2731095358] = "ClientFileStatus"
	ConstructorNames[442767121] = "ClientCachedMediaInfo"
	ConstructorNames[2711408875] = "ClientPeerMediaInfo"
	ConstructorNames[1541024203] = "ClientMediaSize"
	ConstructorNames[2045409949] = "RecentSearch"
	ConstructorNames[1217165147] = "RecentSearchMany"
}
