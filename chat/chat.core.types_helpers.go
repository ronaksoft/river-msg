// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.core.types.proto

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

const C_Ping int64 = 2246546115

type poolPing struct {
	pool sync.Pool
}

func (p *poolPing) Get() *Ping {
	x, ok := p.pool.Get().(*Ping)
	if !ok {
		return &Ping{}
	}
	return x
}

func (p *poolPing) Put(x *Ping) {
	p.pool.Put(x)
}

var PoolPing = poolPing{}

func ResultPing(out *MessageEnvelope, res *Ping) {
	out.Constructor = C_Ping
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Pong int64 = 2171268721

type poolPong struct {
	pool sync.Pool
}

func (p *poolPong) Get() *Pong {
	x, ok := p.pool.Get().(*Pong)
	if !ok {
		return &Pong{}
	}
	return x
}

func (p *poolPong) Put(x *Pong) {
	p.pool.Put(x)
}

var PoolPong = poolPong{}

func ResultPong(out *MessageEnvelope, res *Pong) {
	out.Constructor = C_Pong
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageEnvelope int64 = 535232465

type poolMessageEnvelope struct {
	pool sync.Pool
}

func (p *poolMessageEnvelope) Get() *MessageEnvelope {
	x, ok := p.pool.Get().(*MessageEnvelope)
	if !ok {
		return &MessageEnvelope{}
	}
	x.Message = x.Message[:0]
	return x
}

func (p *poolMessageEnvelope) Put(x *MessageEnvelope) {
	p.pool.Put(x)
}

var PoolMessageEnvelope = poolMessageEnvelope{}

func ResultMessageEnvelope(out *MessageEnvelope, res *MessageEnvelope) {
	out.Constructor = C_MessageEnvelope
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageContainer int64 = 1972016308

type poolMessageContainer struct {
	pool sync.Pool
}

func (p *poolMessageContainer) Get() *MessageContainer {
	x, ok := p.pool.Get().(*MessageContainer)
	if !ok {
		return &MessageContainer{}
	}
	x.Envelopes = x.Envelopes[:0]
	return x
}

func (p *poolMessageContainer) Put(x *MessageContainer) {
	p.pool.Put(x)
}

var PoolMessageContainer = poolMessageContainer{}

func ResultMessageContainer(out *MessageEnvelope, res *MessageContainer) {
	out.Constructor = C_MessageContainer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateEnvelope int64 = 2373884514

type poolUpdateEnvelope struct {
	pool sync.Pool
}

func (p *poolUpdateEnvelope) Get() *UpdateEnvelope {
	x, ok := p.pool.Get().(*UpdateEnvelope)
	if !ok {
		return &UpdateEnvelope{}
	}
	x.Update = x.Update[:0]
	return x
}

func (p *poolUpdateEnvelope) Put(x *UpdateEnvelope) {
	p.pool.Put(x)
}

var PoolUpdateEnvelope = poolUpdateEnvelope{}

func ResultUpdateEnvelope(out *MessageEnvelope, res *UpdateEnvelope) {
	out.Constructor = C_UpdateEnvelope
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UpdateContainer int64 = 661712615

type poolUpdateContainer struct {
	pool sync.Pool
}

func (p *poolUpdateContainer) Get() *UpdateContainer {
	x, ok := p.pool.Get().(*UpdateContainer)
	if !ok {
		return &UpdateContainer{}
	}
	x.Updates = x.Updates[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	return x
}

func (p *poolUpdateContainer) Put(x *UpdateContainer) {
	p.pool.Put(x)
}

var PoolUpdateContainer = poolUpdateContainer{}

func ResultUpdateContainer(out *MessageEnvelope, res *UpdateContainer) {
	out.Constructor = C_UpdateContainer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ProtoMessage int64 = 2179260159

type poolProtoMessage struct {
	pool sync.Pool
}

func (p *poolProtoMessage) Get() *ProtoMessage {
	x, ok := p.pool.Get().(*ProtoMessage)
	if !ok {
		return &ProtoMessage{}
	}
	x.AuthID = 0
	x.MessageKey = nil
	x.MessageKey = x.MessageKey[:0]
	x.Payload = x.Payload[:0]
	return x
}

func (p *poolProtoMessage) Put(x *ProtoMessage) {
	p.pool.Put(x)
}

var PoolProtoMessage = poolProtoMessage{}

func ResultProtoMessage(out *MessageEnvelope, res *ProtoMessage) {
	out.Constructor = C_ProtoMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ProtoEncryptedPayload int64 = 2668405547

type poolProtoEncryptedPayload struct {
	pool sync.Pool
}

func (p *poolProtoEncryptedPayload) Get() *ProtoEncryptedPayload {
	x, ok := p.pool.Get().(*ProtoEncryptedPayload)
	if !ok {
		return &ProtoEncryptedPayload{}
	}
	return x
}

func (p *poolProtoEncryptedPayload) Put(x *ProtoEncryptedPayload) {
	p.pool.Put(x)
}

var PoolProtoEncryptedPayload = poolProtoEncryptedPayload{}

func ResultProtoEncryptedPayload(out *MessageEnvelope, res *ProtoEncryptedPayload) {
	out.Constructor = C_ProtoEncryptedPayload
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Error int64 = 2619118453

type poolError struct {
	pool sync.Pool
}

func (p *poolError) Get() *Error {
	x, ok := p.pool.Get().(*Error)
	if !ok {
		return &Error{}
	}
	return x
}

func (p *poolError) Put(x *Error) {
	p.pool.Put(x)
}

var PoolError = poolError{}

func ResultError(out *MessageEnvelope, res *Error) {
	out.Constructor = C_Error
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Ack int64 = 447331921

type poolAck struct {
	pool sync.Pool
}

func (p *poolAck) Get() *Ack {
	x, ok := p.pool.Get().(*Ack)
	if !ok {
		return &Ack{}
	}
	x.MessageIDs = x.MessageIDs[:0]
	return x
}

func (p *poolAck) Put(x *Ack) {
	p.pool.Put(x)
}

var PoolAck = poolAck{}

func ResultAck(out *MessageEnvelope, res *Ack) {
	out.Constructor = C_Ack
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Bool int64 = 4122188204

type poolBool struct {
	pool sync.Pool
}

func (p *poolBool) Get() *Bool {
	x, ok := p.pool.Get().(*Bool)
	if !ok {
		return &Bool{}
	}
	return x
}

func (p *poolBool) Put(x *Bool) {
	p.pool.Put(x)
}

var PoolBool = poolBool{}

func ResultBool(out *MessageEnvelope, res *Bool) {
	out.Constructor = C_Bool
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Dialog int64 = 1120787796

type poolDialog struct {
	pool sync.Pool
}

func (p *poolDialog) Get() *Dialog {
	x, ok := p.pool.Get().(*Dialog)
	if !ok {
		return &Dialog{}
	}
	x.NotifySettings = nil
	x.MentionedCount = 0
	x.Draft = nil
	return x
}

func (p *poolDialog) Put(x *Dialog) {
	p.pool.Put(x)
}

var PoolDialog = poolDialog{}

func ResultDialog(out *MessageEnvelope, res *Dialog) {
	out.Constructor = C_Dialog
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputPeer int64 = 3374092470

type poolInputPeer struct {
	pool sync.Pool
}

func (p *poolInputPeer) Get() *InputPeer {
	x, ok := p.pool.Get().(*InputPeer)
	if !ok {
		return &InputPeer{}
	}
	return x
}

func (p *poolInputPeer) Put(x *InputPeer) {
	p.pool.Put(x)
}

var PoolInputPeer = poolInputPeer{}

func ResultInputPeer(out *MessageEnvelope, res *InputPeer) {
	out.Constructor = C_InputPeer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Peer int64 = 47470215

type poolPeer struct {
	pool sync.Pool
}

func (p *poolPeer) Get() *Peer {
	x, ok := p.pool.Get().(*Peer)
	if !ok {
		return &Peer{}
	}
	return x
}

func (p *poolPeer) Put(x *Peer) {
	p.pool.Put(x)
}

var PoolPeer = poolPeer{}

func ResultPeer(out *MessageEnvelope, res *Peer) {
	out.Constructor = C_Peer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputUser int64 = 3865689926

type poolInputUser struct {
	pool sync.Pool
}

func (p *poolInputUser) Get() *InputUser {
	x, ok := p.pool.Get().(*InputUser)
	if !ok {
		return &InputUser{}
	}
	return x
}

func (p *poolInputUser) Put(x *InputUser) {
	p.pool.Put(x)
}

var PoolInputUser = poolInputUser{}

func ResultInputUser(out *MessageEnvelope, res *InputUser) {
	out.Constructor = C_InputUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputPassword int64 = 513021899

type poolInputPassword struct {
	pool sync.Pool
}

func (p *poolInputPassword) Get() *InputPassword {
	x, ok := p.pool.Get().(*InputPassword)
	if !ok {
		return &InputPassword{}
	}
	x.A = x.A[:0]
	x.M1 = x.M1[:0]
	return x
}

func (p *poolInputPassword) Put(x *InputPassword) {
	p.pool.Put(x)
}

var PoolInputPassword = poolInputPassword{}

func ResultInputPassword(out *MessageEnvelope, res *InputPassword) {
	out.Constructor = C_InputPassword
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputFileLocation int64 = 354669666

type poolInputFileLocation struct {
	pool sync.Pool
}

func (p *poolInputFileLocation) Get() *InputFileLocation {
	x, ok := p.pool.Get().(*InputFileLocation)
	if !ok {
		return &InputFileLocation{}
	}
	x.Version = 0
	return x
}

func (p *poolInputFileLocation) Put(x *InputFileLocation) {
	p.pool.Put(x)
}

var PoolInputFileLocation = poolInputFileLocation{}

func ResultInputFileLocation(out *MessageEnvelope, res *InputFileLocation) {
	out.Constructor = C_InputFileLocation
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_FileLocation int64 = 2432133155

type poolFileLocation struct {
	pool sync.Pool
}

func (p *poolFileLocation) Get() *FileLocation {
	x, ok := p.pool.Get().(*FileLocation)
	if !ok {
		return &FileLocation{}
	}
	return x
}

func (p *poolFileLocation) Put(x *FileLocation) {
	p.pool.Put(x)
}

var PoolFileLocation = poolFileLocation{}

func ResultFileLocation(out *MessageEnvelope, res *FileLocation) {
	out.Constructor = C_FileLocation
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UserPhoto int64 = 1881347437

type poolUserPhoto struct {
	pool sync.Pool
}

func (p *poolUserPhoto) Get() *UserPhoto {
	x, ok := p.pool.Get().(*UserPhoto)
	if !ok {
		return &UserPhoto{}
	}
	return x
}

func (p *poolUserPhoto) Put(x *UserPhoto) {
	p.pool.Put(x)
}

var PoolUserPhoto = poolUserPhoto{}

func ResultUserPhoto(out *MessageEnvelope, res *UserPhoto) {
	out.Constructor = C_UserPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_User int64 = 765557111

type poolUser struct {
	pool sync.Pool
}

func (p *poolUser) Get() *User {
	x, ok := p.pool.Get().(*User)
	if !ok {
		return &User{}
	}
	x.Username = ""
	x.Photo = nil
	x.Bio = ""
	x.Phone = ""
	x.LastSeen = 0
	x.PhotoGallery = x.PhotoGallery[:0]
	x.IsBot = false
	x.Deleted = false
	x.Blocked = false
	x.BotInfo = nil
	x.Official = false
	return x
}

func (p *poolUser) Put(x *User) {
	p.pool.Put(x)
}

var PoolUser = poolUser{}

func ResultUser(out *MessageEnvelope, res *User) {
	out.Constructor = C_User
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Bot int64 = 961692401

type poolBot struct {
	pool sync.Pool
}

func (p *poolBot) Get() *Bot {
	x, ok := p.pool.Get().(*Bot)
	if !ok {
		return &Bot{}
	}
	x.Bio = ""
	return x
}

func (p *poolBot) Put(x *Bot) {
	p.pool.Put(x)
}

var PoolBot = poolBot{}

func ResultBot(out *MessageEnvelope, res *Bot) {
	out.Constructor = C_Bot
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_BotCommands int64 = 1852470005

type poolBotCommands struct {
	pool sync.Pool
}

func (p *poolBotCommands) Get() *BotCommands {
	x, ok := p.pool.Get().(*BotCommands)
	if !ok {
		return &BotCommands{}
	}
	x.Description = ""
	return x
}

func (p *poolBotCommands) Put(x *BotCommands) {
	p.pool.Put(x)
}

var PoolBotCommands = poolBotCommands{}

func ResultBotCommands(out *MessageEnvelope, res *BotCommands) {
	out.Constructor = C_BotCommands
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_BotInfo int64 = 4059496923

type poolBotInfo struct {
	pool sync.Pool
}

func (p *poolBotInfo) Get() *BotInfo {
	x, ok := p.pool.Get().(*BotInfo)
	if !ok {
		return &BotInfo{}
	}
	x.Description = ""
	x.BotCommands = x.BotCommands[:0]
	x.InlineGeo = false
	x.InlinePlaceHolder = ""
	x.InlineQuery = false
	return x
}

func (p *poolBotInfo) Put(x *BotInfo) {
	p.pool.Put(x)
}

var PoolBotInfo = poolBotInfo{}

func ResultBotInfo(out *MessageEnvelope, res *BotInfo) {
	out.Constructor = C_BotInfo
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_ContactUser int64 = 460099170

type poolContactUser struct {
	pool sync.Pool
}

func (p *poolContactUser) Get() *ContactUser {
	x, ok := p.pool.Get().(*ContactUser)
	if !ok {
		return &ContactUser{}
	}
	x.Photo = nil
	return x
}

func (p *poolContactUser) Put(x *ContactUser) {
	p.pool.Put(x)
}

var PoolContactUser = poolContactUser{}

func ResultContactUser(out *MessageEnvelope, res *ContactUser) {
	out.Constructor = C_ContactUser
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_UserMessage int64 = 1677556362

type poolUserMessage struct {
	pool sync.Pool
}

func (p *poolUserMessage) Get() *UserMessage {
	x, ok := p.pool.Get().(*UserMessage)
	if !ok {
		return &UserMessage{}
	}
	x.MessageAction = 0
	x.MessageActionData = nil
	x.MessageActionData = x.MessageActionData[:0]
	x.Entities = x.Entities[:0]
	x.MediaType = 0
	x.Media = nil
	x.Media = x.Media[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	x.LabelIDs = x.LabelIDs[:0]
	x.ViaBotID = 0
	return x
}

func (p *poolUserMessage) Put(x *UserMessage) {
	p.pool.Put(x)
}

var PoolUserMessage = poolUserMessage{}

func ResultUserMessage(out *MessageEnvelope, res *UserMessage) {
	out.Constructor = C_UserMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_DraftMessage int64 = 869564229

type poolDraftMessage struct {
	pool sync.Pool
}

func (p *poolDraftMessage) Get() *DraftMessage {
	x, ok := p.pool.Get().(*DraftMessage)
	if !ok {
		return &DraftMessage{}
	}
	x.Entities = x.Entities[:0]
	x.EditedID = 0
	return x
}

func (p *poolDraftMessage) Put(x *DraftMessage) {
	p.pool.Put(x)
}

var PoolDraftMessage = poolDraftMessage{}

func ResultDraftMessage(out *MessageEnvelope, res *DraftMessage) {
	out.Constructor = C_DraftMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_MessageEntity int64 = 3479443932

type poolMessageEntity struct {
	pool sync.Pool
}

func (p *poolMessageEntity) Get() *MessageEntity {
	x, ok := p.pool.Get().(*MessageEntity)
	if !ok {
		return &MessageEntity{}
	}
	x.UserID = 0
	return x
}

func (p *poolMessageEntity) Put(x *MessageEntity) {
	p.pool.Put(x)
}

var PoolMessageEntity = poolMessageEntity{}

func ResultMessageEntity(out *MessageEnvelope, res *MessageEntity) {
	out.Constructor = C_MessageEntity
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_RSAPublicKey int64 = 1046601890

type poolRSAPublicKey struct {
	pool sync.Pool
}

func (p *poolRSAPublicKey) Get() *RSAPublicKey {
	x, ok := p.pool.Get().(*RSAPublicKey)
	if !ok {
		return &RSAPublicKey{}
	}
	return x
}

func (p *poolRSAPublicKey) Put(x *RSAPublicKey) {
	p.pool.Put(x)
}

var PoolRSAPublicKey = poolRSAPublicKey{}

func ResultRSAPublicKey(out *MessageEnvelope, res *RSAPublicKey) {
	out.Constructor = C_RSAPublicKey
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_DHGroup int64 = 2751503049

type poolDHGroup struct {
	pool sync.Pool
}

func (p *poolDHGroup) Get() *DHGroup {
	x, ok := p.pool.Get().(*DHGroup)
	if !ok {
		return &DHGroup{}
	}
	return x
}

func (p *poolDHGroup) Put(x *DHGroup) {
	p.pool.Put(x)
}

var PoolDHGroup = poolDHGroup{}

func ResultDHGroup(out *MessageEnvelope, res *DHGroup) {
	out.Constructor = C_DHGroup
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PhoneContact int64 = 2672574672

type poolPhoneContact struct {
	pool sync.Pool
}

func (p *poolPhoneContact) Get() *PhoneContact {
	x, ok := p.pool.Get().(*PhoneContact)
	if !ok {
		return &PhoneContact{}
	}
	return x
}

func (p *poolPhoneContact) Put(x *PhoneContact) {
	p.pool.Put(x)
}

var PoolPhoneContact = poolPhoneContact{}

func ResultPhoneContact(out *MessageEnvelope, res *PhoneContact) {
	out.Constructor = C_PhoneContact
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PeerNotifySettings int64 = 3475030132

type poolPeerNotifySettings struct {
	pool sync.Pool
}

func (p *poolPeerNotifySettings) Get() *PeerNotifySettings {
	x, ok := p.pool.Get().(*PeerNotifySettings)
	if !ok {
		return &PeerNotifySettings{}
	}
	x.MuteUntil = 0
	x.Sound = ""
	return x
}

func (p *poolPeerNotifySettings) Put(x *PeerNotifySettings) {
	p.pool.Put(x)
}

var PoolPeerNotifySettings = poolPeerNotifySettings{}

func ResultPeerNotifySettings(out *MessageEnvelope, res *PeerNotifySettings) {
	out.Constructor = C_PeerNotifySettings
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputFile int64 = 3882180383

type poolInputFile struct {
	pool sync.Pool
}

func (p *poolInputFile) Get() *InputFile {
	x, ok := p.pool.Get().(*InputFile)
	if !ok {
		return &InputFile{}
	}
	return x
}

func (p *poolInputFile) Put(x *InputFile) {
	p.pool.Put(x)
}

var PoolInputFile = poolInputFile{}

func ResultInputFile(out *MessageEnvelope, res *InputFile) {
	out.Constructor = C_InputFile
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputDocument int64 = 4081048424

type poolInputDocument struct {
	pool sync.Pool
}

func (p *poolInputDocument) Get() *InputDocument {
	x, ok := p.pool.Get().(*InputDocument)
	if !ok {
		return &InputDocument{}
	}
	return x
}

func (p *poolInputDocument) Put(x *InputDocument) {
	p.pool.Put(x)
}

var PoolInputDocument = poolInputDocument{}

func ResultInputDocument(out *MessageEnvelope, res *InputDocument) {
	out.Constructor = C_InputDocument
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupPhoto int64 = 3998516135

type poolGroupPhoto struct {
	pool sync.Pool
}

func (p *poolGroupPhoto) Get() *GroupPhoto {
	x, ok := p.pool.Get().(*GroupPhoto)
	if !ok {
		return &GroupPhoto{}
	}
	x.PhotoID = 0
	return x
}

func (p *poolGroupPhoto) Put(x *GroupPhoto) {
	p.pool.Put(x)
}

var PoolGroupPhoto = poolGroupPhoto{}

func ResultGroupPhoto(out *MessageEnvelope, res *GroupPhoto) {
	out.Constructor = C_GroupPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Group int64 = 2885774273

type poolGroup struct {
	pool sync.Pool
}

func (p *poolGroup) Get() *Group {
	x, ok := p.pool.Get().(*Group)
	if !ok {
		return &Group{}
	}
	x.EditedOn = 0
	x.Flags = x.Flags[:0]
	x.Photo = nil
	return x
}

func (p *poolGroup) Put(x *Group) {
	p.pool.Put(x)
}

var PoolGroup = poolGroup{}

func ResultGroup(out *MessageEnvelope, res *Group) {
	out.Constructor = C_Group
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupFull int64 = 205850814

type poolGroupFull struct {
	pool sync.Pool
}

func (p *poolGroupFull) Get() *GroupFull {
	x, ok := p.pool.Get().(*GroupFull)
	if !ok {
		return &GroupFull{}
	}
	x.Users = x.Users[:0]
	x.Participants = x.Participants[:0]
	x.PhotoGallery = x.PhotoGallery[:0]
	return x
}

func (p *poolGroupFull) Put(x *GroupFull) {
	p.pool.Put(x)
}

var PoolGroupFull = poolGroupFull{}

func ResultGroupFull(out *MessageEnvelope, res *GroupFull) {
	out.Constructor = C_GroupFull
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GroupParticipant int64 = 4072279665

type poolGroupParticipant struct {
	pool sync.Pool
}

func (p *poolGroupParticipant) Get() *GroupParticipant {
	x, ok := p.pool.Get().(*GroupParticipant)
	if !ok {
		return &GroupParticipant{}
	}
	x.Photo = nil
	return x
}

func (p *poolGroupParticipant) Put(x *GroupParticipant) {
	p.pool.Put(x)
}

var PoolGroupParticipant = poolGroupParticipant{}

func ResultGroupParticipant(out *MessageEnvelope, res *GroupParticipant) {
	out.Constructor = C_GroupParticipant
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PrivacyRule int64 = 3954700912

type poolPrivacyRule struct {
	pool sync.Pool
}

func (p *poolPrivacyRule) Get() *PrivacyRule {
	x, ok := p.pool.Get().(*PrivacyRule)
	if !ok {
		return &PrivacyRule{}
	}
	x.UserIDs = x.UserIDs[:0]
	return x
}

func (p *poolPrivacyRule) Put(x *PrivacyRule) {
	p.pool.Put(x)
}

var PoolPrivacyRule = poolPrivacyRule{}

func ResultPrivacyRule(out *MessageEnvelope, res *PrivacyRule) {
	out.Constructor = C_PrivacyRule
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_Label int64 = 3479601132

type poolLabel struct {
	pool sync.Pool
}

func (p *poolLabel) Get() *Label {
	x, ok := p.pool.Get().(*Label)
	if !ok {
		return &Label{}
	}
	x.Count = 0
	return x
}

func (p *poolLabel) Put(x *Label) {
	p.pool.Put(x)
}

var PoolLabel = poolLabel{}

func ResultLabel(out *MessageEnvelope, res *Label) {
	out.Constructor = C_Label
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_LabelsMany int64 = 1423713603

type poolLabelsMany struct {
	pool sync.Pool
}

func (p *poolLabelsMany) Get() *LabelsMany {
	x, ok := p.pool.Get().(*LabelsMany)
	if !ok {
		return &LabelsMany{}
	}
	x.Labels = x.Labels[:0]
	x.Empty = false
	return x
}

func (p *poolLabelsMany) Put(x *LabelsMany) {
	p.pool.Put(x)
}

var PoolLabelsMany = poolLabelsMany{}

func ResultLabelsMany(out *MessageEnvelope, res *LabelsMany) {
	out.Constructor = C_LabelsMany
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_InputGeoLocation int64 = 1403425127

type poolInputGeoLocation struct {
	pool sync.Pool
}

func (p *poolInputGeoLocation) Get() *InputGeoLocation {
	x, ok := p.pool.Get().(*InputGeoLocation)
	if !ok {
		return &InputGeoLocation{}
	}
	return x
}

func (p *poolInputGeoLocation) Put(x *InputGeoLocation) {
	p.pool.Put(x)
}

var PoolInputGeoLocation = poolInputGeoLocation{}

func ResultInputGeoLocation(out *MessageEnvelope, res *InputGeoLocation) {
	out.Constructor = C_InputGeoLocation
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_GeoLocation int64 = 3794405429

type poolGeoLocation struct {
	pool sync.Pool
}

func (p *poolGeoLocation) Get() *GeoLocation {
	x, ok := p.pool.Get().(*GeoLocation)
	if !ok {
		return &GeoLocation{}
	}
	return x
}

func (p *poolGeoLocation) Put(x *GeoLocation) {
	p.pool.Put(x)
}

var PoolGeoLocation = poolGeoLocation{}

func ResultGeoLocation(out *MessageEnvelope, res *GeoLocation) {
	out.Constructor = C_GeoLocation
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
	ConstructorNames[2246546115] = "Ping"
	ConstructorNames[2171268721] = "Pong"
	ConstructorNames[535232465] = "MessageEnvelope"
	ConstructorNames[1972016308] = "MessageContainer"
	ConstructorNames[2373884514] = "UpdateEnvelope"
	ConstructorNames[661712615] = "UpdateContainer"
	ConstructorNames[2179260159] = "ProtoMessage"
	ConstructorNames[2668405547] = "ProtoEncryptedPayload"
	ConstructorNames[2619118453] = "Error"
	ConstructorNames[447331921] = "Ack"
	ConstructorNames[4122188204] = "Bool"
	ConstructorNames[1120787796] = "Dialog"
	ConstructorNames[3374092470] = "InputPeer"
	ConstructorNames[47470215] = "Peer"
	ConstructorNames[3865689926] = "InputUser"
	ConstructorNames[513021899] = "InputPassword"
	ConstructorNames[354669666] = "InputFileLocation"
	ConstructorNames[2432133155] = "FileLocation"
	ConstructorNames[1881347437] = "UserPhoto"
	ConstructorNames[765557111] = "User"
	ConstructorNames[961692401] = "Bot"
	ConstructorNames[1852470005] = "BotCommands"
	ConstructorNames[4059496923] = "BotInfo"
	ConstructorNames[460099170] = "ContactUser"
	ConstructorNames[1677556362] = "UserMessage"
	ConstructorNames[869564229] = "DraftMessage"
	ConstructorNames[3479443932] = "MessageEntity"
	ConstructorNames[1046601890] = "RSAPublicKey"
	ConstructorNames[2751503049] = "DHGroup"
	ConstructorNames[2672574672] = "PhoneContact"
	ConstructorNames[3475030132] = "PeerNotifySettings"
	ConstructorNames[3882180383] = "InputFile"
	ConstructorNames[4081048424] = "InputDocument"
	ConstructorNames[3998516135] = "GroupPhoto"
	ConstructorNames[2885774273] = "Group"
	ConstructorNames[205850814] = "GroupFull"
	ConstructorNames[4072279665] = "GroupParticipant"
	ConstructorNames[3954700912] = "PrivacyRule"
	ConstructorNames[3479601132] = "Label"
	ConstructorNames[1423713603] = "LabelsMany"
	ConstructorNames[1403425127] = "InputGeoLocation"
	ConstructorNames[3794405429] = "GeoLocation"
}
