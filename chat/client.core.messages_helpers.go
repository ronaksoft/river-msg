// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client.core.messages.proto

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
	x.ClearDraft = false
	x.FileUploadID = ""
	x.ThumbUploadID = ""
	x.FileID = 0
	x.ThumbID = 0
	return x
}

func (p *poolClientPendingMessage) Put(x *ClientPendingMessage) {
	p.pool.Put(x)
}

var PoolClientPendingMessage = poolClientPendingMessage{}

func ResultClientPendingMessage(out *MessageEnvelope, res *ClientPendingMessage) {
	out.Constructor = C_ClientPendingMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

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
	return x
}

func (p *poolClientSendMessageMedia) Put(x *ClientSendMessageMedia) {
	p.pool.Put(x)
}

var PoolClientSendMessageMedia = poolClientSendMessageMedia{}

func ResultClientSendMessageMedia(out *MessageEnvelope, res *ClientSendMessageMedia) {
	out.Constructor = C_ClientSendMessageMedia
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	return x
}

func (p *poolClientGlobalSearch) Put(x *ClientGlobalSearch) {
	p.pool.Put(x)
}

var PoolClientGlobalSearch = poolClientGlobalSearch{}

func ResultClientGlobalSearch(out *MessageEnvelope, res *ClientGlobalSearch) {
	out.Constructor = C_ClientGlobalSearch
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	return x
}

func (p *poolClientFile) Put(x *ClientFile) {
	p.pool.Put(x)
}

var PoolClientFile = poolClientFile{}

func ResultClientFile(out *MessageEnvelope, res *ClientFile) {
	out.Constructor = C_ClientFile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
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
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_DBMediaInfo int64 = 2652925823

type poolDBMediaInfo struct {
	pool sync.Pool
}

func (p *poolDBMediaInfo) Get() *DBMediaInfo {
	x, ok := p.pool.Get().(*DBMediaInfo)
	if !ok {
		return &DBMediaInfo{}
	}
	x.MediaInfo = x.MediaInfo[:0]
	return x
}

func (p *poolDBMediaInfo) Put(x *DBMediaInfo) {
	p.pool.Put(x)
}

var PoolDBMediaInfo = poolDBMediaInfo{}

func ResultDBMediaInfo(out *MessageEnvelope, res *DBMediaInfo) {
	out.Constructor = C_DBMediaInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_PeerMediaInfo int64 = 1816749655

type poolPeerMediaInfo struct {
	pool sync.Pool
}

func (p *poolPeerMediaInfo) Get() *PeerMediaInfo {
	x, ok := p.pool.Get().(*PeerMediaInfo)
	if !ok {
		return &PeerMediaInfo{}
	}
	x.Media = x.Media[:0]
	return x
}

func (p *poolPeerMediaInfo) Put(x *PeerMediaInfo) {
	p.pool.Put(x)
}

var PoolPeerMediaInfo = poolPeerMediaInfo{}

func ResultPeerMediaInfo(out *MessageEnvelope, res *PeerMediaInfo) {
	out.Constructor = C_PeerMediaInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

const C_MediaSize int64 = 3543753456

type poolMediaSize struct {
	pool sync.Pool
}

func (p *poolMediaSize) Get() *MediaSize {
	x, ok := p.pool.Get().(*MediaSize)
	if !ok {
		return &MediaSize{}
	}
	return x
}

func (p *poolMediaSize) Put(x *MediaSize) {
	p.pool.Put(x)
}

var PoolMediaSize = poolMediaSize{}

func ResultMediaSize(out *MessageEnvelope, res *MediaSize) {
	out.Constructor = C_MediaSize
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

func init() {
	ConstructorNames[2164891929] = "ClientPendingMessage"
	ConstructorNames[1095038539] = "ClientSendMessageMedia"
	ConstructorNames[1742781507] = "ClientGlobalSearch"
	ConstructorNames[1793449803] = "ClientContactSearch"
	ConstructorNames[2957647709] = "ClientSearchResult"
	ConstructorNames[155127968] = "ClientFile"
	ConstructorNames[2731095358] = "ClientFileStatus"
	ConstructorNames[2652925823] = "DBMediaInfo"
	ConstructorNames[1816749655] = "PeerMediaInfo"
	ConstructorNames[3543753456] = "MediaSize"
}