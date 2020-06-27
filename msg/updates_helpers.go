// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: updates.proto

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

const C_UpdateGetState int64 = 1437250230

type poolUpdateGetState struct {
	pool sync.Pool
}

func (p *poolUpdateGetState) Get() *UpdateGetState {
	x, ok := p.pool.Get().(*UpdateGetState)
	if !ok {
		return &UpdateGetState{}
	}
	return x
}

func (p *poolUpdateGetState) Put(x *UpdateGetState) {
	p.pool.Put(x)
}

var PoolUpdateGetState = poolUpdateGetState{}

func ResultUpdateGetState(out *MessageEnvelope, res *UpdateGetState) {
	out.Constructor = C_UpdateGetState
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGetDifference int64 = 556775761

type poolUpdateGetDifference struct {
	pool sync.Pool
}

func (p *poolUpdateGetDifference) Get() *UpdateGetDifference {
	x, ok := p.pool.Get().(*UpdateGetDifference)
	if !ok {
		return &UpdateGetDifference{}
	}
	return x
}

func (p *poolUpdateGetDifference) Put(x *UpdateGetDifference) {
	p.pool.Put(x)
}

var PoolUpdateGetDifference = poolUpdateGetDifference{}

func ResultUpdateGetDifference(out *MessageEnvelope, res *UpdateGetDifference) {
	out.Constructor = C_UpdateGetDifference
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateDifference int64 = 1742546619

type poolUpdateDifference struct {
	pool sync.Pool
}

func (p *poolUpdateDifference) Get() *UpdateDifference {
	x, ok := p.pool.Get().(*UpdateDifference)
	if !ok {
		return &UpdateDifference{}
	}
	x.Updates = x.Updates[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	x.CurrentUpdateID = 0
	return x
}

func (p *poolUpdateDifference) Put(x *UpdateDifference) {
	p.pool.Put(x)
}

var PoolUpdateDifference = poolUpdateDifference{}

func ResultUpdateDifference(out *MessageEnvelope, res *UpdateDifference) {
	out.Constructor = C_UpdateDifference
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateTooLong int64 = 1531755547

type poolUpdateTooLong struct {
	pool sync.Pool
}

func (p *poolUpdateTooLong) Get() *UpdateTooLong {
	x, ok := p.pool.Get().(*UpdateTooLong)
	if !ok {
		return &UpdateTooLong{}
	}
	return x
}

func (p *poolUpdateTooLong) Put(x *UpdateTooLong) {
	p.pool.Put(x)
}

var PoolUpdateTooLong = poolUpdateTooLong{}

func ResultUpdateTooLong(out *MessageEnvelope, res *UpdateTooLong) {
	out.Constructor = C_UpdateTooLong
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateState int64 = 1837585836

type poolUpdateState struct {
	pool sync.Pool
}

func (p *poolUpdateState) Get() *UpdateState {
	x, ok := p.pool.Get().(*UpdateState)
	if !ok {
		return &UpdateState{}
	}
	return x
}

func (p *poolUpdateState) Put(x *UpdateState) {
	p.pool.Put(x)
}

var PoolUpdateState = poolUpdateState{}

func ResultUpdateState(out *MessageEnvelope, res *UpdateState) {
	out.Constructor = C_UpdateState
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateMessageID int64 = 2139063022

type poolUpdateMessageID struct {
	pool sync.Pool
}

func (p *poolUpdateMessageID) Get() *UpdateMessageID {
	x, ok := p.pool.Get().(*UpdateMessageID)
	if !ok {
		return &UpdateMessageID{}
	}
	return x
}

func (p *poolUpdateMessageID) Put(x *UpdateMessageID) {
	p.pool.Put(x)
}

var PoolUpdateMessageID = poolUpdateMessageID{}

func ResultUpdateMessageID(out *MessageEnvelope, res *UpdateMessageID) {
	out.Constructor = C_UpdateMessageID
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateNewMessage int64 = 3426925183

type poolUpdateNewMessage struct {
	pool sync.Pool
}

func (p *poolUpdateNewMessage) Get() *UpdateNewMessage {
	x, ok := p.pool.Get().(*UpdateNewMessage)
	if !ok {
		return &UpdateNewMessage{}
	}
	x.AccessHash = 0
	x.SenderRefID = 0
	return x
}

func (p *poolUpdateNewMessage) Put(x *UpdateNewMessage) {
	p.pool.Put(x)
}

var PoolUpdateNewMessage = poolUpdateNewMessage{}

func ResultUpdateNewMessage(out *MessageEnvelope, res *UpdateNewMessage) {
	out.Constructor = C_UpdateNewMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateMessageEdited int64 = 1825079988

type poolUpdateMessageEdited struct {
	pool sync.Pool
}

func (p *poolUpdateMessageEdited) Get() *UpdateMessageEdited {
	x, ok := p.pool.Get().(*UpdateMessageEdited)
	if !ok {
		return &UpdateMessageEdited{}
	}
	return x
}

func (p *poolUpdateMessageEdited) Put(x *UpdateMessageEdited) {
	p.pool.Put(x)
}

var PoolUpdateMessageEdited = poolUpdateMessageEdited{}

func ResultUpdateMessageEdited(out *MessageEnvelope, res *UpdateMessageEdited) {
	out.Constructor = C_UpdateMessageEdited
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateMessagesDeleted int64 = 670568714

type poolUpdateMessagesDeleted struct {
	pool sync.Pool
}

func (p *poolUpdateMessagesDeleted) Get() *UpdateMessagesDeleted {
	x, ok := p.pool.Get().(*UpdateMessagesDeleted)
	if !ok {
		return &UpdateMessagesDeleted{}
	}
	x.TeamID = 0
	x.MessageIDs = x.MessageIDs[:0]
	x.Peer = nil
	return x
}

func (p *poolUpdateMessagesDeleted) Put(x *UpdateMessagesDeleted) {
	p.pool.Put(x)
}

var PoolUpdateMessagesDeleted = poolUpdateMessagesDeleted{}

func ResultUpdateMessagesDeleted(out *MessageEnvelope, res *UpdateMessagesDeleted) {
	out.Constructor = C_UpdateMessagesDeleted
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateReadHistoryInbox int64 = 1529128378

type poolUpdateReadHistoryInbox struct {
	pool sync.Pool
}

func (p *poolUpdateReadHistoryInbox) Get() *UpdateReadHistoryInbox {
	x, ok := p.pool.Get().(*UpdateReadHistoryInbox)
	if !ok {
		return &UpdateReadHistoryInbox{}
	}
	x.TeamID = 0
	return x
}

func (p *poolUpdateReadHistoryInbox) Put(x *UpdateReadHistoryInbox) {
	p.pool.Put(x)
}

var PoolUpdateReadHistoryInbox = poolUpdateReadHistoryInbox{}

func ResultUpdateReadHistoryInbox(out *MessageEnvelope, res *UpdateReadHistoryInbox) {
	out.Constructor = C_UpdateReadHistoryInbox
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateReadHistoryOutbox int64 = 510866108

type poolUpdateReadHistoryOutbox struct {
	pool sync.Pool
}

func (p *poolUpdateReadHistoryOutbox) Get() *UpdateReadHistoryOutbox {
	x, ok := p.pool.Get().(*UpdateReadHistoryOutbox)
	if !ok {
		return &UpdateReadHistoryOutbox{}
	}
	x.TeamID = 0
	return x
}

func (p *poolUpdateReadHistoryOutbox) Put(x *UpdateReadHistoryOutbox) {
	p.pool.Put(x)
}

var PoolUpdateReadHistoryOutbox = poolUpdateReadHistoryOutbox{}

func ResultUpdateReadHistoryOutbox(out *MessageEnvelope, res *UpdateReadHistoryOutbox) {
	out.Constructor = C_UpdateReadHistoryOutbox
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateUserTyping int64 = 178254060

type poolUpdateUserTyping struct {
	pool sync.Pool
}

func (p *poolUpdateUserTyping) Get() *UpdateUserTyping {
	x, ok := p.pool.Get().(*UpdateUserTyping)
	if !ok {
		return &UpdateUserTyping{}
	}
	x.TeamID = 0
	return x
}

func (p *poolUpdateUserTyping) Put(x *UpdateUserTyping) {
	p.pool.Put(x)
}

var PoolUpdateUserTyping = poolUpdateUserTyping{}

func ResultUpdateUserTyping(out *MessageEnvelope, res *UpdateUserTyping) {
	out.Constructor = C_UpdateUserTyping
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateUserStatus int64 = 2696747995

type poolUpdateUserStatus struct {
	pool sync.Pool
}

func (p *poolUpdateUserStatus) Get() *UpdateUserStatus {
	x, ok := p.pool.Get().(*UpdateUserStatus)
	if !ok {
		return &UpdateUserStatus{}
	}
	return x
}

func (p *poolUpdateUserStatus) Put(x *UpdateUserStatus) {
	p.pool.Put(x)
}

var PoolUpdateUserStatus = poolUpdateUserStatus{}

func ResultUpdateUserStatus(out *MessageEnvelope, res *UpdateUserStatus) {
	out.Constructor = C_UpdateUserStatus
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateUsername int64 = 4290110589

type poolUpdateUsername struct {
	pool sync.Pool
}

func (p *poolUpdateUsername) Get() *UpdateUsername {
	x, ok := p.pool.Get().(*UpdateUsername)
	if !ok {
		return &UpdateUsername{}
	}
	x.Phone = ""
	return x
}

func (p *poolUpdateUsername) Put(x *UpdateUsername) {
	p.pool.Put(x)
}

var PoolUpdateUsername = poolUpdateUsername{}

func ResultUpdateUsername(out *MessageEnvelope, res *UpdateUsername) {
	out.Constructor = C_UpdateUsername
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateUserPhoto int64 = 302028082

type poolUpdateUserPhoto struct {
	pool sync.Pool
}

func (p *poolUpdateUserPhoto) Get() *UpdateUserPhoto {
	x, ok := p.pool.Get().(*UpdateUserPhoto)
	if !ok {
		return &UpdateUserPhoto{}
	}
	x.Photo = nil
	x.PhotoID = 0
	x.DeletedPhotoIDs = x.DeletedPhotoIDs[:0]
	return x
}

func (p *poolUpdateUserPhoto) Put(x *UpdateUserPhoto) {
	p.pool.Put(x)
}

var PoolUpdateUserPhoto = poolUpdateUserPhoto{}

func ResultUpdateUserPhoto(out *MessageEnvelope, res *UpdateUserPhoto) {
	out.Constructor = C_UpdateUserPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateNotifySettings int64 = 3187524885

type poolUpdateNotifySettings struct {
	pool sync.Pool
}

func (p *poolUpdateNotifySettings) Get() *UpdateNotifySettings {
	x, ok := p.pool.Get().(*UpdateNotifySettings)
	if !ok {
		return &UpdateNotifySettings{}
	}
	return x
}

func (p *poolUpdateNotifySettings) Put(x *UpdateNotifySettings) {
	p.pool.Put(x)
}

var PoolUpdateNotifySettings = poolUpdateNotifySettings{}

func ResultUpdateNotifySettings(out *MessageEnvelope, res *UpdateNotifySettings) {
	out.Constructor = C_UpdateNotifySettings
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGroupParticipantAdd int64 = 1623827837

type poolUpdateGroupParticipantAdd struct {
	pool sync.Pool
}

func (p *poolUpdateGroupParticipantAdd) Get() *UpdateGroupParticipantAdd {
	x, ok := p.pool.Get().(*UpdateGroupParticipantAdd)
	if !ok {
		return &UpdateGroupParticipantAdd{}
	}
	return x
}

func (p *poolUpdateGroupParticipantAdd) Put(x *UpdateGroupParticipantAdd) {
	p.pool.Put(x)
}

var PoolUpdateGroupParticipantAdd = poolUpdateGroupParticipantAdd{}

func ResultUpdateGroupParticipantAdd(out *MessageEnvelope, res *UpdateGroupParticipantAdd) {
	out.Constructor = C_UpdateGroupParticipantAdd
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGroupParticipantDeleted int64 = 2489941844

type poolUpdateGroupParticipantDeleted struct {
	pool sync.Pool
}

func (p *poolUpdateGroupParticipantDeleted) Get() *UpdateGroupParticipantDeleted {
	x, ok := p.pool.Get().(*UpdateGroupParticipantDeleted)
	if !ok {
		return &UpdateGroupParticipantDeleted{}
	}
	return x
}

func (p *poolUpdateGroupParticipantDeleted) Put(x *UpdateGroupParticipantDeleted) {
	p.pool.Put(x)
}

var PoolUpdateGroupParticipantDeleted = poolUpdateGroupParticipantDeleted{}

func ResultUpdateGroupParticipantDeleted(out *MessageEnvelope, res *UpdateGroupParticipantDeleted) {
	out.Constructor = C_UpdateGroupParticipantDeleted
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGroupParticipantAdmin int64 = 1813022164

type poolUpdateGroupParticipantAdmin struct {
	pool sync.Pool
}

func (p *poolUpdateGroupParticipantAdmin) Get() *UpdateGroupParticipantAdmin {
	x, ok := p.pool.Get().(*UpdateGroupParticipantAdmin)
	if !ok {
		return &UpdateGroupParticipantAdmin{}
	}
	return x
}

func (p *poolUpdateGroupParticipantAdmin) Put(x *UpdateGroupParticipantAdmin) {
	p.pool.Put(x)
}

var PoolUpdateGroupParticipantAdmin = poolUpdateGroupParticipantAdmin{}

func ResultUpdateGroupParticipantAdmin(out *MessageEnvelope, res *UpdateGroupParticipantAdmin) {
	out.Constructor = C_UpdateGroupParticipantAdmin
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGroupAdmins int64 = 694155405

type poolUpdateGroupAdmins struct {
	pool sync.Pool
}

func (p *poolUpdateGroupAdmins) Get() *UpdateGroupAdmins {
	x, ok := p.pool.Get().(*UpdateGroupAdmins)
	if !ok {
		return &UpdateGroupAdmins{}
	}
	return x
}

func (p *poolUpdateGroupAdmins) Put(x *UpdateGroupAdmins) {
	p.pool.Put(x)
}

var PoolUpdateGroupAdmins = poolUpdateGroupAdmins{}

func ResultUpdateGroupAdmins(out *MessageEnvelope, res *UpdateGroupAdmins) {
	out.Constructor = C_UpdateGroupAdmins
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateGroupPhoto int64 = 367193154

type poolUpdateGroupPhoto struct {
	pool sync.Pool
}

func (p *poolUpdateGroupPhoto) Get() *UpdateGroupPhoto {
	x, ok := p.pool.Get().(*UpdateGroupPhoto)
	if !ok {
		return &UpdateGroupPhoto{}
	}
	x.Photo = nil
	x.PhotoID = 0
	return x
}

func (p *poolUpdateGroupPhoto) Put(x *UpdateGroupPhoto) {
	p.pool.Put(x)
}

var PoolUpdateGroupPhoto = poolUpdateGroupPhoto{}

func ResultUpdateGroupPhoto(out *MessageEnvelope, res *UpdateGroupPhoto) {
	out.Constructor = C_UpdateGroupPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateReadMessagesContents int64 = 2991403048

type poolUpdateReadMessagesContents struct {
	pool sync.Pool
}

func (p *poolUpdateReadMessagesContents) Get() *UpdateReadMessagesContents {
	x, ok := p.pool.Get().(*UpdateReadMessagesContents)
	if !ok {
		return &UpdateReadMessagesContents{}
	}
	x.MessageIDs = x.MessageIDs[:0]
	return x
}

func (p *poolUpdateReadMessagesContents) Put(x *UpdateReadMessagesContents) {
	p.pool.Put(x)
}

var PoolUpdateReadMessagesContents = poolUpdateReadMessagesContents{}

func ResultUpdateReadMessagesContents(out *MessageEnvelope, res *UpdateReadMessagesContents) {
	out.Constructor = C_UpdateReadMessagesContents
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateAuthorizationReset int64 = 2359297647

type poolUpdateAuthorizationReset struct {
	pool sync.Pool
}

func (p *poolUpdateAuthorizationReset) Get() *UpdateAuthorizationReset {
	x, ok := p.pool.Get().(*UpdateAuthorizationReset)
	if !ok {
		return &UpdateAuthorizationReset{}
	}
	return x
}

func (p *poolUpdateAuthorizationReset) Put(x *UpdateAuthorizationReset) {
	p.pool.Put(x)
}

var PoolUpdateAuthorizationReset = poolUpdateAuthorizationReset{}

func ResultUpdateAuthorizationReset(out *MessageEnvelope, res *UpdateAuthorizationReset) {
	out.Constructor = C_UpdateAuthorizationReset
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateDraftMessage int64 = 3453026195

type poolUpdateDraftMessage struct {
	pool sync.Pool
}

func (p *poolUpdateDraftMessage) Get() *UpdateDraftMessage {
	x, ok := p.pool.Get().(*UpdateDraftMessage)
	if !ok {
		return &UpdateDraftMessage{}
	}
	x.UpdateID = 0
	return x
}

func (p *poolUpdateDraftMessage) Put(x *UpdateDraftMessage) {
	p.pool.Put(x)
}

var PoolUpdateDraftMessage = poolUpdateDraftMessage{}

func ResultUpdateDraftMessage(out *MessageEnvelope, res *UpdateDraftMessage) {
	out.Constructor = C_UpdateDraftMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateDraftMessageCleared int64 = 2011635602

type poolUpdateDraftMessageCleared struct {
	pool sync.Pool
}

func (p *poolUpdateDraftMessageCleared) Get() *UpdateDraftMessageCleared {
	x, ok := p.pool.Get().(*UpdateDraftMessageCleared)
	if !ok {
		return &UpdateDraftMessageCleared{}
	}
	x.UpdateID = 0
	return x
}

func (p *poolUpdateDraftMessageCleared) Put(x *UpdateDraftMessageCleared) {
	p.pool.Put(x)
}

var PoolUpdateDraftMessageCleared = poolUpdateDraftMessageCleared{}

func ResultUpdateDraftMessageCleared(out *MessageEnvelope, res *UpdateDraftMessageCleared) {
	out.Constructor = C_UpdateDraftMessageCleared
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateDialogPinned int64 = 231538299

type poolUpdateDialogPinned struct {
	pool sync.Pool
}

func (p *poolUpdateDialogPinned) Get() *UpdateDialogPinned {
	x, ok := p.pool.Get().(*UpdateDialogPinned)
	if !ok {
		return &UpdateDialogPinned{}
	}
	return x
}

func (p *poolUpdateDialogPinned) Put(x *UpdateDialogPinned) {
	p.pool.Put(x)
}

var PoolUpdateDialogPinned = poolUpdateDialogPinned{}

func ResultUpdateDialogPinned(out *MessageEnvelope, res *UpdateDialogPinned) {
	out.Constructor = C_UpdateDialogPinned
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateDialogPinnedReorder int64 = 1567423539

type poolUpdateDialogPinnedReorder struct {
	pool sync.Pool
}

func (p *poolUpdateDialogPinnedReorder) Get() *UpdateDialogPinnedReorder {
	x, ok := p.pool.Get().(*UpdateDialogPinnedReorder)
	if !ok {
		return &UpdateDialogPinnedReorder{}
	}
	x.Peer = x.Peer[:0]
	return x
}

func (p *poolUpdateDialogPinnedReorder) Put(x *UpdateDialogPinnedReorder) {
	p.pool.Put(x)
}

var PoolUpdateDialogPinnedReorder = poolUpdateDialogPinnedReorder{}

func ResultUpdateDialogPinnedReorder(out *MessageEnvelope, res *UpdateDialogPinnedReorder) {
	out.Constructor = C_UpdateDialogPinnedReorder
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateAccountPrivacy int64 = 629173761

type poolUpdateAccountPrivacy struct {
	pool sync.Pool
}

func (p *poolUpdateAccountPrivacy) Get() *UpdateAccountPrivacy {
	x, ok := p.pool.Get().(*UpdateAccountPrivacy)
	if !ok {
		return &UpdateAccountPrivacy{}
	}
	x.ChatInvite = x.ChatInvite[:0]
	x.LastSeen = x.LastSeen[:0]
	x.PhoneNumber = x.PhoneNumber[:0]
	x.ProfilePhoto = x.ProfilePhoto[:0]
	x.ForwardedMessage = x.ForwardedMessage[:0]
	x.Call = x.Call[:0]
	return x
}

func (p *poolUpdateAccountPrivacy) Put(x *UpdateAccountPrivacy) {
	p.pool.Put(x)
}

var PoolUpdateAccountPrivacy = poolUpdateAccountPrivacy{}

func ResultUpdateAccountPrivacy(out *MessageEnvelope, res *UpdateAccountPrivacy) {
	out.Constructor = C_UpdateAccountPrivacy
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateLabelItemsAdded int64 = 2216022057

type poolUpdateLabelItemsAdded struct {
	pool sync.Pool
}

func (p *poolUpdateLabelItemsAdded) Get() *UpdateLabelItemsAdded {
	x, ok := p.pool.Get().(*UpdateLabelItemsAdded)
	if !ok {
		return &UpdateLabelItemsAdded{}
	}
	x.MessageIDs = x.MessageIDs[:0]
	x.LabelIDs = x.LabelIDs[:0]
	x.Labels = x.Labels[:0]
	return x
}

func (p *poolUpdateLabelItemsAdded) Put(x *UpdateLabelItemsAdded) {
	p.pool.Put(x)
}

var PoolUpdateLabelItemsAdded = poolUpdateLabelItemsAdded{}

func ResultUpdateLabelItemsAdded(out *MessageEnvelope, res *UpdateLabelItemsAdded) {
	out.Constructor = C_UpdateLabelItemsAdded
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateLabelItemsRemoved int64 = 830226827

type poolUpdateLabelItemsRemoved struct {
	pool sync.Pool
}

func (p *poolUpdateLabelItemsRemoved) Get() *UpdateLabelItemsRemoved {
	x, ok := p.pool.Get().(*UpdateLabelItemsRemoved)
	if !ok {
		return &UpdateLabelItemsRemoved{}
	}
	x.MessageIDs = x.MessageIDs[:0]
	x.LabelIDs = x.LabelIDs[:0]
	x.Labels = x.Labels[:0]
	return x
}

func (p *poolUpdateLabelItemsRemoved) Put(x *UpdateLabelItemsRemoved) {
	p.pool.Put(x)
}

var PoolUpdateLabelItemsRemoved = poolUpdateLabelItemsRemoved{}

func ResultUpdateLabelItemsRemoved(out *MessageEnvelope, res *UpdateLabelItemsRemoved) {
	out.Constructor = C_UpdateLabelItemsRemoved
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateLabelSet int64 = 2353687359

type poolUpdateLabelSet struct {
	pool sync.Pool
}

func (p *poolUpdateLabelSet) Get() *UpdateLabelSet {
	x, ok := p.pool.Get().(*UpdateLabelSet)
	if !ok {
		return &UpdateLabelSet{}
	}
	x.Labels = x.Labels[:0]
	return x
}

func (p *poolUpdateLabelSet) Put(x *UpdateLabelSet) {
	p.pool.Put(x)
}

var PoolUpdateLabelSet = poolUpdateLabelSet{}

func ResultUpdateLabelSet(out *MessageEnvelope, res *UpdateLabelSet) {
	out.Constructor = C_UpdateLabelSet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateLabelDeleted int64 = 3702192307

type poolUpdateLabelDeleted struct {
	pool sync.Pool
}

func (p *poolUpdateLabelDeleted) Get() *UpdateLabelDeleted {
	x, ok := p.pool.Get().(*UpdateLabelDeleted)
	if !ok {
		return &UpdateLabelDeleted{}
	}
	x.LabelIDs = x.LabelIDs[:0]
	return x
}

func (p *poolUpdateLabelDeleted) Put(x *UpdateLabelDeleted) {
	p.pool.Put(x)
}

var PoolUpdateLabelDeleted = poolUpdateLabelDeleted{}

func ResultUpdateLabelDeleted(out *MessageEnvelope, res *UpdateLabelDeleted) {
	out.Constructor = C_UpdateLabelDeleted
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateUserBlocked int64 = 3750625773

type poolUpdateUserBlocked struct {
	pool sync.Pool
}

func (p *poolUpdateUserBlocked) Get() *UpdateUserBlocked {
	x, ok := p.pool.Get().(*UpdateUserBlocked)
	if !ok {
		return &UpdateUserBlocked{}
	}
	return x
}

func (p *poolUpdateUserBlocked) Put(x *UpdateUserBlocked) {
	p.pool.Put(x)
}

var PoolUpdateUserBlocked = poolUpdateUserBlocked{}

func ResultUpdateUserBlocked(out *MessageEnvelope, res *UpdateUserBlocked) {
	out.Constructor = C_UpdateUserBlocked
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateMessagePoll int64 = 383248674

type poolUpdateMessagePoll struct {
	pool sync.Pool
}

func (p *poolUpdateMessagePoll) Get() *UpdateMessagePoll {
	x, ok := p.pool.Get().(*UpdateMessagePoll)
	if !ok {
		return &UpdateMessagePoll{}
	}
	x.Poll = nil
	return x
}

func (p *poolUpdateMessagePoll) Put(x *UpdateMessagePoll) {
	p.pool.Put(x)
}

var PoolUpdateMessagePoll = poolUpdateMessagePoll{}

func ResultUpdateMessagePoll(out *MessageEnvelope, res *UpdateMessagePoll) {
	out.Constructor = C_UpdateMessagePoll
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateBotCallbackQuery int64 = 3408999713

type poolUpdateBotCallbackQuery struct {
	pool sync.Pool
}

func (p *poolUpdateBotCallbackQuery) Get() *UpdateBotCallbackQuery {
	x, ok := p.pool.Get().(*UpdateBotCallbackQuery)
	if !ok {
		return &UpdateBotCallbackQuery{}
	}
	x.MessageID = 0
	x.Data = nil
	x.Data = x.Data[:0]
	return x
}

func (p *poolUpdateBotCallbackQuery) Put(x *UpdateBotCallbackQuery) {
	p.pool.Put(x)
}

var PoolUpdateBotCallbackQuery = poolUpdateBotCallbackQuery{}

func ResultUpdateBotCallbackQuery(out *MessageEnvelope, res *UpdateBotCallbackQuery) {
	out.Constructor = C_UpdateBotCallbackQuery
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateBotInlineQuery int64 = 4065328876

type poolUpdateBotInlineQuery struct {
	pool sync.Pool
}

func (p *poolUpdateBotInlineQuery) Get() *UpdateBotInlineQuery {
	x, ok := p.pool.Get().(*UpdateBotInlineQuery)
	if !ok {
		return &UpdateBotInlineQuery{}
	}
	x.Offset = ""
	x.Geo = nil
	return x
}

func (p *poolUpdateBotInlineQuery) Put(x *UpdateBotInlineQuery) {
	p.pool.Put(x)
}

var PoolUpdateBotInlineQuery = poolUpdateBotInlineQuery{}

func ResultUpdateBotInlineQuery(out *MessageEnvelope, res *UpdateBotInlineQuery) {
	out.Constructor = C_UpdateBotInlineQuery
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateBotInlineSend int64 = 2208028013

type poolUpdateBotInlineSend struct {
	pool sync.Pool
}

func (p *poolUpdateBotInlineSend) Get() *UpdateBotInlineSend {
	x, ok := p.pool.Get().(*UpdateBotInlineSend)
	if !ok {
		return &UpdateBotInlineSend{}
	}
	x.Geo = nil
	return x
}

func (p *poolUpdateBotInlineSend) Put(x *UpdateBotInlineSend) {
	p.pool.Put(x)
}

var PoolUpdateBotInlineSend = poolUpdateBotInlineSend{}

func ResultUpdateBotInlineSend(out *MessageEnvelope, res *UpdateBotInlineSend) {
	out.Constructor = C_UpdateBotInlineSend
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateTeamCreated int64 = 2762932529

type poolUpdateTeamCreated struct {
	pool sync.Pool
}

func (p *poolUpdateTeamCreated) Get() *UpdateTeamCreated {
	x, ok := p.pool.Get().(*UpdateTeamCreated)
	if !ok {
		return &UpdateTeamCreated{}
	}
	return x
}

func (p *poolUpdateTeamCreated) Put(x *UpdateTeamCreated) {
	p.pool.Put(x)
}

var PoolUpdateTeamCreated = poolUpdateTeamCreated{}

func ResultUpdateTeamCreated(out *MessageEnvelope, res *UpdateTeamCreated) {
	out.Constructor = C_UpdateTeamCreated
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateTeamMemberAdded int64 = 1307755890

type poolUpdateTeamMemberAdded struct {
	pool sync.Pool
}

func (p *poolUpdateTeamMemberAdded) Get() *UpdateTeamMemberAdded {
	x, ok := p.pool.Get().(*UpdateTeamMemberAdded)
	if !ok {
		return &UpdateTeamMemberAdded{}
	}
	return x
}

func (p *poolUpdateTeamMemberAdded) Put(x *UpdateTeamMemberAdded) {
	p.pool.Put(x)
}

var PoolUpdateTeamMemberAdded = poolUpdateTeamMemberAdded{}

func ResultUpdateTeamMemberAdded(out *MessageEnvelope, res *UpdateTeamMemberAdded) {
	out.Constructor = C_UpdateTeamMemberAdded
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateTeamMemberRemoved int64 = 99543064

type poolUpdateTeamMemberRemoved struct {
	pool sync.Pool
}

func (p *poolUpdateTeamMemberRemoved) Get() *UpdateTeamMemberRemoved {
	x, ok := p.pool.Get().(*UpdateTeamMemberRemoved)
	if !ok {
		return &UpdateTeamMemberRemoved{}
	}
	return x
}

func (p *poolUpdateTeamMemberRemoved) Put(x *UpdateTeamMemberRemoved) {
	p.pool.Put(x)
}

var PoolUpdateTeamMemberRemoved = poolUpdateTeamMemberRemoved{}

func ResultUpdateTeamMemberRemoved(out *MessageEnvelope, res *UpdateTeamMemberRemoved) {
	out.Constructor = C_UpdateTeamMemberRemoved
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateTeamPhotoChanged int64 = 2787282465

type poolUpdateTeamPhotoChanged struct {
	pool sync.Pool
}

func (p *poolUpdateTeamPhotoChanged) Get() *UpdateTeamPhotoChanged {
	x, ok := p.pool.Get().(*UpdateTeamPhotoChanged)
	if !ok {
		return &UpdateTeamPhotoChanged{}
	}
	return x
}

func (p *poolUpdateTeamPhotoChanged) Put(x *UpdateTeamPhotoChanged) {
	p.pool.Put(x)
}

var PoolUpdateTeamPhotoChanged = poolUpdateTeamPhotoChanged{}

func ResultUpdateTeamPhotoChanged(out *MessageEnvelope, res *UpdateTeamPhotoChanged) {
	out.Constructor = C_UpdateTeamPhotoChanged
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateCalendarEventAdded int64 = 297964741

type poolUpdateCalendarEventAdded struct {
	pool sync.Pool
}

func (p *poolUpdateCalendarEventAdded) Get() *UpdateCalendarEventAdded {
	x, ok := p.pool.Get().(*UpdateCalendarEventAdded)
	if !ok {
		return &UpdateCalendarEventAdded{}
	}
	return x
}

func (p *poolUpdateCalendarEventAdded) Put(x *UpdateCalendarEventAdded) {
	p.pool.Put(x)
}

var PoolUpdateCalendarEventAdded = poolUpdateCalendarEventAdded{}

func ResultUpdateCalendarEventAdded(out *MessageEnvelope, res *UpdateCalendarEventAdded) {
	out.Constructor = C_UpdateCalendarEventAdded
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateCalendarEventRemoved int64 = 2986798389

type poolUpdateCalendarEventRemoved struct {
	pool sync.Pool
}

func (p *poolUpdateCalendarEventRemoved) Get() *UpdateCalendarEventRemoved {
	x, ok := p.pool.Get().(*UpdateCalendarEventRemoved)
	if !ok {
		return &UpdateCalendarEventRemoved{}
	}
	return x
}

func (p *poolUpdateCalendarEventRemoved) Put(x *UpdateCalendarEventRemoved) {
	p.pool.Put(x)
}

var PoolUpdateCalendarEventRemoved = poolUpdateCalendarEventRemoved{}

func ResultUpdateCalendarEventRemoved(out *MessageEnvelope, res *UpdateCalendarEventRemoved) {
	out.Constructor = C_UpdateCalendarEventRemoved
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateCalendarEventEdited int64 = 516349098

type poolUpdateCalendarEventEdited struct {
	pool sync.Pool
}

func (p *poolUpdateCalendarEventEdited) Get() *UpdateCalendarEventEdited {
	x, ok := p.pool.Get().(*UpdateCalendarEventEdited)
	if !ok {
		return &UpdateCalendarEventEdited{}
	}
	return x
}

func (p *poolUpdateCalendarEventEdited) Put(x *UpdateCalendarEventEdited) {
	p.pool.Put(x)
}

var PoolUpdateCalendarEventEdited = poolUpdateCalendarEventEdited{}

func ResultUpdateCalendarEventEdited(out *MessageEnvelope, res *UpdateCalendarEventEdited) {
	out.Constructor = C_UpdateCalendarEventEdited
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
	ConstructorNames[1437250230] = "UpdateGetState"
	ConstructorNames[556775761] = "UpdateGetDifference"
	ConstructorNames[1742546619] = "UpdateDifference"
	ConstructorNames[1531755547] = "UpdateTooLong"
	ConstructorNames[1837585836] = "UpdateState"
	ConstructorNames[2139063022] = "UpdateMessageID"
	ConstructorNames[3426925183] = "UpdateNewMessage"
	ConstructorNames[1825079988] = "UpdateMessageEdited"
	ConstructorNames[670568714] = "UpdateMessagesDeleted"
	ConstructorNames[1529128378] = "UpdateReadHistoryInbox"
	ConstructorNames[510866108] = "UpdateReadHistoryOutbox"
	ConstructorNames[178254060] = "UpdateUserTyping"
	ConstructorNames[2696747995] = "UpdateUserStatus"
	ConstructorNames[4290110589] = "UpdateUsername"
	ConstructorNames[302028082] = "UpdateUserPhoto"
	ConstructorNames[3187524885] = "UpdateNotifySettings"
	ConstructorNames[1623827837] = "UpdateGroupParticipantAdd"
	ConstructorNames[2489941844] = "UpdateGroupParticipantDeleted"
	ConstructorNames[1813022164] = "UpdateGroupParticipantAdmin"
	ConstructorNames[694155405] = "UpdateGroupAdmins"
	ConstructorNames[367193154] = "UpdateGroupPhoto"
	ConstructorNames[2991403048] = "UpdateReadMessagesContents"
	ConstructorNames[2359297647] = "UpdateAuthorizationReset"
	ConstructorNames[3453026195] = "UpdateDraftMessage"
	ConstructorNames[2011635602] = "UpdateDraftMessageCleared"
	ConstructorNames[231538299] = "UpdateDialogPinned"
	ConstructorNames[1567423539] = "UpdateDialogPinnedReorder"
	ConstructorNames[629173761] = "UpdateAccountPrivacy"
	ConstructorNames[2216022057] = "UpdateLabelItemsAdded"
	ConstructorNames[830226827] = "UpdateLabelItemsRemoved"
	ConstructorNames[2353687359] = "UpdateLabelSet"
	ConstructorNames[3702192307] = "UpdateLabelDeleted"
	ConstructorNames[3750625773] = "UpdateUserBlocked"
	ConstructorNames[383248674] = "UpdateMessagePoll"
	ConstructorNames[3408999713] = "UpdateBotCallbackQuery"
	ConstructorNames[4065328876] = "UpdateBotInlineQuery"
	ConstructorNames[2208028013] = "UpdateBotInlineSend"
	ConstructorNames[2762932529] = "UpdateTeamCreated"
	ConstructorNames[1307755890] = "UpdateTeamMemberAdded"
	ConstructorNames[99543064] = "UpdateTeamMemberRemoved"
	ConstructorNames[2787282465] = "UpdateTeamPhotoChanged"
	ConstructorNames[297964741] = "UpdateCalendarEventAdded"
	ConstructorNames[2986798389] = "UpdateCalendarEventRemoved"
	ConstructorNames[516349098] = "UpdateCalendarEventEdited"
}
