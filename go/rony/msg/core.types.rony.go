package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

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
	x.ID = 0
	p.pool.Put(x)
}

var PoolPing = poolPing{}

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
	x.ID = 0
	p.pool.Put(x)
}

var PoolPong = poolPong{}

const C_MessageEnvelope int64 = 535232465

type poolMessageEnvelope struct {
	pool sync.Pool
}

func (p *poolMessageEnvelope) Get() *MessageEnvelope {
	x, ok := p.pool.Get().(*MessageEnvelope)
	if !ok {
		return &MessageEnvelope{}
	}
	return x
}

func (p *poolMessageEnvelope) Put(x *MessageEnvelope) {
	x.Constructor = 0
	x.RequestID = 0
	x.Message = x.Message[:0]
	if x.Team != nil {
		PoolInputTeam.Put(x.Team)
		x.Team = nil
	}
	p.pool.Put(x)
}

var PoolMessageEnvelope = poolMessageEnvelope{}

const C_MessageContainer int64 = 1972016308

type poolMessageContainer struct {
	pool sync.Pool
}

func (p *poolMessageContainer) Get() *MessageContainer {
	x, ok := p.pool.Get().(*MessageContainer)
	if !ok {
		return &MessageContainer{}
	}
	return x
}

func (p *poolMessageContainer) Put(x *MessageContainer) {
	x.Length = 0
	x.Envelopes = x.Envelopes[:0]
	p.pool.Put(x)
}

var PoolMessageContainer = poolMessageContainer{}

const C_UpdateEnvelope int64 = 2373884514

type poolUpdateEnvelope struct {
	pool sync.Pool
}

func (p *poolUpdateEnvelope) Get() *UpdateEnvelope {
	x, ok := p.pool.Get().(*UpdateEnvelope)
	if !ok {
		return &UpdateEnvelope{}
	}
	return x
}

func (p *poolUpdateEnvelope) Put(x *UpdateEnvelope) {
	x.Constructor = 0
	x.Update = x.Update[:0]
	x.UCount = 0
	x.UpdateID = 0
	x.Timestamp = 0
	p.pool.Put(x)
}

var PoolUpdateEnvelope = poolUpdateEnvelope{}

const C_UpdateContainer int64 = 661712615

type poolUpdateContainer struct {
	pool sync.Pool
}

func (p *poolUpdateContainer) Get() *UpdateContainer {
	x, ok := p.pool.Get().(*UpdateContainer)
	if !ok {
		return &UpdateContainer{}
	}
	return x
}

func (p *poolUpdateContainer) Put(x *UpdateContainer) {
	x.Length = 0
	x.Updates = x.Updates[:0]
	x.MinUpdateID = 0
	x.MaxUpdateID = 0
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	p.pool.Put(x)
}

var PoolUpdateContainer = poolUpdateContainer{}

const C_ProtoMessage int64 = 2179260159

type poolProtoMessage struct {
	pool sync.Pool
}

func (p *poolProtoMessage) Get() *ProtoMessage {
	x, ok := p.pool.Get().(*ProtoMessage)
	if !ok {
		return &ProtoMessage{}
	}
	return x
}

func (p *poolProtoMessage) Put(x *ProtoMessage) {
	x.AuthID = 0
	x.MessageKey = x.MessageKey[:0]
	x.Payload = x.Payload[:0]
	p.pool.Put(x)
}

var PoolProtoMessage = poolProtoMessage{}

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
	x.ServerSalt = 0
	x.MessageID = 0
	x.SessionID = 0
	if x.Envelope != nil {
		PoolMessageEnvelope.Put(x.Envelope)
		x.Envelope = nil
	}
	p.pool.Put(x)
}

var PoolProtoEncryptedPayload = poolProtoEncryptedPayload{}

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
	x.Code = ""
	x.Items = ""
	p.pool.Put(x)
}

var PoolError = poolError{}

const C_Ack int64 = 447331921

type poolAck struct {
	pool sync.Pool
}

func (p *poolAck) Get() *Ack {
	x, ok := p.pool.Get().(*Ack)
	if !ok {
		return &Ack{}
	}
	return x
}

func (p *poolAck) Put(x *Ack) {
	x.MessageIDs = x.MessageIDs[:0]
	p.pool.Put(x)
}

var PoolAck = poolAck{}

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
	x.Result = false
	p.pool.Put(x)
}

var PoolBool = poolBool{}

const C_Dialog int64 = 1120787796

type poolDialog struct {
	pool sync.Pool
}

func (p *poolDialog) Get() *Dialog {
	x, ok := p.pool.Get().(*Dialog)
	if !ok {
		return &Dialog{}
	}
	return x
}

func (p *poolDialog) Put(x *Dialog) {
	x.TeamID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.TopMessageID = 0
	x.ReadInboxMaxID = 0
	x.ReadOutboxMaxID = 0
	x.UnreadCount = 0
	x.AccessHash = 0
	if x.NotifySettings != nil {
		PoolPeerNotifySettings.Put(x.NotifySettings)
		x.NotifySettings = nil
	}
	x.MentionedCount = 0
	x.Pinned = false
	if x.Draft != nil {
		PoolDraftMessage.Put(x.Draft)
		x.Draft = nil
	}
	x.PinnedMessageID = 0
	p.pool.Put(x)
}

var PoolDialog = poolDialog{}

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
	x.ID = 0
	x.Type = 0
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolInputPeer = poolInputPeer{}

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
	x.ID = 0
	x.Type = 0
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolPeer = poolPeer{}

const C_InputPassword int64 = 513021899

type poolInputPassword struct {
	pool sync.Pool
}

func (p *poolInputPassword) Get() *InputPassword {
	x, ok := p.pool.Get().(*InputPassword)
	if !ok {
		return &InputPassword{}
	}
	return x
}

func (p *poolInputPassword) Put(x *InputPassword) {
	x.SrpID = 0
	x.A = x.A[:0]
	x.M1 = x.M1[:0]
	p.pool.Put(x)
}

var PoolInputPassword = poolInputPassword{}

const C_InputFileLocation int64 = 354669666

type poolInputFileLocation struct {
	pool sync.Pool
}

func (p *poolInputFileLocation) Get() *InputFileLocation {
	x, ok := p.pool.Get().(*InputFileLocation)
	if !ok {
		return &InputFileLocation{}
	}
	return x
}

func (p *poolInputFileLocation) Put(x *InputFileLocation) {
	x.ClusterID = 0
	x.FileID = 0
	x.AccessHash = 0
	x.Version = 0
	p.pool.Put(x)
}

var PoolInputFileLocation = poolInputFileLocation{}

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
	x.ClusterID = 0
	x.FileID = 0
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolFileLocation = poolFileLocation{}

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
	if x.PhotoBig != nil {
		PoolFileLocation.Put(x.PhotoBig)
		x.PhotoBig = nil
	}
	if x.PhotoSmall != nil {
		PoolFileLocation.Put(x.PhotoSmall)
		x.PhotoSmall = nil
	}
	x.PhotoID = 0
	p.pool.Put(x)
}

var PoolUserPhoto = poolUserPhoto{}

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
	x.UserID = 0
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolInputUser = poolInputUser{}

const C_User int64 = 765557111

type poolUser struct {
	pool sync.Pool
}

func (p *poolUser) Get() *User {
	x, ok := p.pool.Get().(*User)
	if !ok {
		return &User{}
	}
	return x
}

func (p *poolUser) Put(x *User) {
	x.ID = 0
	x.FirstName = ""
	x.LastName = ""
	x.Username = ""
	x.Status = 0
	x.Restricted = false
	x.AccessHash = 0
	if x.Photo != nil {
		PoolUserPhoto.Put(x.Photo)
		x.Photo = nil
	}
	x.Bio = ""
	x.Phone = ""
	x.LastSeen = 0
	x.PhotoGallery = x.PhotoGallery[:0]
	x.IsBot = false
	x.Deleted = false
	x.Blocked = false
	if x.BotInfo != nil {
		PoolBotInfo.Put(x.BotInfo)
		x.BotInfo = nil
	}
	x.Official = false
	p.pool.Put(x)
}

var PoolUser = poolUser{}

const C_ContactUser int64 = 460099170

type poolContactUser struct {
	pool sync.Pool
}

func (p *poolContactUser) Get() *ContactUser {
	x, ok := p.pool.Get().(*ContactUser)
	if !ok {
		return &ContactUser{}
	}
	return x
}

func (p *poolContactUser) Put(x *ContactUser) {
	x.ID = 0
	x.FirstName = ""
	x.LastName = ""
	x.AccessHash = 0
	x.Phone = ""
	x.Username = ""
	x.ClientID = 0
	if x.Photo != nil {
		PoolUserPhoto.Put(x.Photo)
		x.Photo = nil
	}
	p.pool.Put(x)
}

var PoolContactUser = poolContactUser{}

const C_Bot int64 = 961692401

type poolBot struct {
	pool sync.Pool
}

func (p *poolBot) Get() *Bot {
	x, ok := p.pool.Get().(*Bot)
	if !ok {
		return &Bot{}
	}
	return x
}

func (p *poolBot) Put(x *Bot) {
	x.ID = 0
	x.Name = ""
	x.Username = ""
	x.Bio = ""
	p.pool.Put(x)
}

var PoolBot = poolBot{}

const C_BotCommands int64 = 1852470005

type poolBotCommands struct {
	pool sync.Pool
}

func (p *poolBotCommands) Get() *BotCommands {
	x, ok := p.pool.Get().(*BotCommands)
	if !ok {
		return &BotCommands{}
	}
	return x
}

func (p *poolBotCommands) Put(x *BotCommands) {
	x.Command = ""
	x.Description = ""
	p.pool.Put(x)
}

var PoolBotCommands = poolBotCommands{}

const C_BotInfo int64 = 4059496923

type poolBotInfo struct {
	pool sync.Pool
}

func (p *poolBotInfo) Get() *BotInfo {
	x, ok := p.pool.Get().(*BotInfo)
	if !ok {
		return &BotInfo{}
	}
	return x
}

func (p *poolBotInfo) Put(x *BotInfo) {
	if x.Bot != nil {
		PoolBot.Put(x.Bot)
		x.Bot = nil
	}
	x.UserID = 0
	x.Description = ""
	x.BotCommands = x.BotCommands[:0]
	x.InlineGeo = false
	x.InlinePlaceHolder = ""
	x.InlineQuery = false
	p.pool.Put(x)
}

var PoolBotInfo = poolBotInfo{}

const C_GroupPhoto int64 = 3998516135

type poolGroupPhoto struct {
	pool sync.Pool
}

func (p *poolGroupPhoto) Get() *GroupPhoto {
	x, ok := p.pool.Get().(*GroupPhoto)
	if !ok {
		return &GroupPhoto{}
	}
	return x
}

func (p *poolGroupPhoto) Put(x *GroupPhoto) {
	if x.PhotoBig != nil {
		PoolFileLocation.Put(x.PhotoBig)
		x.PhotoBig = nil
	}
	if x.PhotoSmall != nil {
		PoolFileLocation.Put(x.PhotoSmall)
		x.PhotoSmall = nil
	}
	x.PhotoID = 0
	p.pool.Put(x)
}

var PoolGroupPhoto = poolGroupPhoto{}

const C_Group int64 = 2885774273

type poolGroup struct {
	pool sync.Pool
}

func (p *poolGroup) Get() *Group {
	x, ok := p.pool.Get().(*Group)
	if !ok {
		return &Group{}
	}
	return x
}

func (p *poolGroup) Put(x *Group) {
	x.TeamID = 0
	x.ID = 0
	x.Title = ""
	x.CreatedOn = 0
	x.Participants = 0
	x.EditedOn = 0
	x.Flags = x.Flags[:0]
	if x.Photo != nil {
		PoolGroupPhoto.Put(x.Photo)
		x.Photo = nil
	}
	p.pool.Put(x)
}

var PoolGroup = poolGroup{}

const C_GroupFull int64 = 205850814

type poolGroupFull struct {
	pool sync.Pool
}

func (p *poolGroupFull) Get() *GroupFull {
	x, ok := p.pool.Get().(*GroupFull)
	if !ok {
		return &GroupFull{}
	}
	return x
}

func (p *poolGroupFull) Put(x *GroupFull) {
	if x.Group != nil {
		PoolGroup.Put(x.Group)
		x.Group = nil
	}
	x.Users = x.Users[:0]
	x.Participants = x.Participants[:0]
	if x.NotifySettings != nil {
		PoolPeerNotifySettings.Put(x.NotifySettings)
		x.NotifySettings = nil
	}
	x.PhotoGallery = x.PhotoGallery[:0]
	p.pool.Put(x)
}

var PoolGroupFull = poolGroupFull{}

const C_GroupParticipant int64 = 4072279665

type poolGroupParticipant struct {
	pool sync.Pool
}

func (p *poolGroupParticipant) Get() *GroupParticipant {
	x, ok := p.pool.Get().(*GroupParticipant)
	if !ok {
		return &GroupParticipant{}
	}
	return x
}

func (p *poolGroupParticipant) Put(x *GroupParticipant) {
	x.UserID = 0
	x.FirstName = ""
	x.LastName = ""
	x.Type = 0
	x.AccessHash = 0
	x.Username = ""
	if x.Photo != nil {
		PoolUserPhoto.Put(x.Photo)
		x.Photo = nil
	}
	p.pool.Put(x)
}

var PoolGroupParticipant = poolGroupParticipant{}

const C_UserMessage int64 = 1677556362

type poolUserMessage struct {
	pool sync.Pool
}

func (p *poolUserMessage) Get() *UserMessage {
	x, ok := p.pool.Get().(*UserMessage)
	if !ok {
		return &UserMessage{}
	}
	return x
}

func (p *poolUserMessage) Put(x *UserMessage) {
	x.TeamID = 0
	x.ID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.CreatedOn = 0
	x.EditedOn = 0
	x.Fwd = false
	x.FwdSenderID = 0
	x.FwdChannelID = 0
	x.FwdChannelMessageID = 0
	x.Flags = 0
	x.MessageType = 0
	x.Body = ""
	x.SenderID = 0
	x.ContentRead = false
	x.Inbox = false
	x.ReplyTo = 0
	x.MessageAction = 0
	x.MessageActionData = x.MessageActionData[:0]
	x.Entities = x.Entities[:0]
	x.MediaType = 0
	x.Media = x.Media[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	x.LabelIDs = x.LabelIDs[:0]
	x.ViaBotID = 0
	x.Reactions = x.Reactions[:0]
	x.YourReactions = x.YourReactions[:0]
	p.pool.Put(x)
}

var PoolUserMessage = poolUserMessage{}

const C_ReactionCounter int64 = 3277490830

type poolReactionCounter struct {
	pool sync.Pool
}

func (p *poolReactionCounter) Get() *ReactionCounter {
	x, ok := p.pool.Get().(*ReactionCounter)
	if !ok {
		return &ReactionCounter{}
	}
	return x
}

func (p *poolReactionCounter) Put(x *ReactionCounter) {
	x.Reaction = ""
	x.Total = 0
	p.pool.Put(x)
}

var PoolReactionCounter = poolReactionCounter{}

const C_DraftMessage int64 = 869564229

type poolDraftMessage struct {
	pool sync.Pool
}

func (p *poolDraftMessage) Get() *DraftMessage {
	x, ok := p.pool.Get().(*DraftMessage)
	if !ok {
		return &DraftMessage{}
	}
	return x
}

func (p *poolDraftMessage) Put(x *DraftMessage) {
	x.TeamID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.Date = 0
	x.Body = ""
	x.Entities = x.Entities[:0]
	x.ReplyTo = 0
	x.EditedID = 0
	p.pool.Put(x)
}

var PoolDraftMessage = poolDraftMessage{}

const C_MessageEntity int64 = 3479443932

type poolMessageEntity struct {
	pool sync.Pool
}

func (p *poolMessageEntity) Get() *MessageEntity {
	x, ok := p.pool.Get().(*MessageEntity)
	if !ok {
		return &MessageEntity{}
	}
	return x
}

func (p *poolMessageEntity) Put(x *MessageEntity) {
	x.Type = 0
	x.Offset = 0
	x.Length = 0
	x.UserID = 0
	p.pool.Put(x)
}

var PoolMessageEntity = poolMessageEntity{}

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
	x.FingerPrint = 0
	x.N = ""
	x.E = 0
	p.pool.Put(x)
}

var PoolRSAPublicKey = poolRSAPublicKey{}

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
	x.FingerPrint = 0
	x.Prime = ""
	x.Gen = 0
	p.pool.Put(x)
}

var PoolDHGroup = poolDHGroup{}

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
	x.ClientID = 0
	x.FirstName = ""
	x.LastName = ""
	x.Phone = ""
	p.pool.Put(x)
}

var PoolPhoneContact = poolPhoneContact{}

const C_PeerNotifySettings int64 = 3475030132

type poolPeerNotifySettings struct {
	pool sync.Pool
}

func (p *poolPeerNotifySettings) Get() *PeerNotifySettings {
	x, ok := p.pool.Get().(*PeerNotifySettings)
	if !ok {
		return &PeerNotifySettings{}
	}
	return x
}

func (p *poolPeerNotifySettings) Put(x *PeerNotifySettings) {
	x.Flags = 0
	x.MuteUntil = 0
	x.Sound = ""
	p.pool.Put(x)
}

var PoolPeerNotifySettings = poolPeerNotifySettings{}

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
	x.FileID = 0
	x.TotalParts = 0
	x.FileName = ""
	x.MD5Checksum = ""
	p.pool.Put(x)
}

var PoolInputFile = poolInputFile{}

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
	x.ID = 0
	x.AccessHash = 0
	x.ClusterID = 0
	p.pool.Put(x)
}

var PoolInputDocument = poolInputDocument{}

const C_PrivacyRule int64 = 3954700912

type poolPrivacyRule struct {
	pool sync.Pool
}

func (p *poolPrivacyRule) Get() *PrivacyRule {
	x, ok := p.pool.Get().(*PrivacyRule)
	if !ok {
		return &PrivacyRule{}
	}
	return x
}

func (p *poolPrivacyRule) Put(x *PrivacyRule) {
	x.PrivacyType = 0
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolPrivacyRule = poolPrivacyRule{}

const C_Label int64 = 3479601132

type poolLabel struct {
	pool sync.Pool
}

func (p *poolLabel) Get() *Label {
	x, ok := p.pool.Get().(*Label)
	if !ok {
		return &Label{}
	}
	return x
}

func (p *poolLabel) Put(x *Label) {
	x.ID = 0
	x.Name = ""
	x.Colour = ""
	x.Count = 0
	p.pool.Put(x)
}

var PoolLabel = poolLabel{}

const C_LabelsMany int64 = 1423713603

type poolLabelsMany struct {
	pool sync.Pool
}

func (p *poolLabelsMany) Get() *LabelsMany {
	x, ok := p.pool.Get().(*LabelsMany)
	if !ok {
		return &LabelsMany{}
	}
	return x
}

func (p *poolLabelsMany) Put(x *LabelsMany) {
	x.Labels = x.Labels[:0]
	x.Empty = false
	p.pool.Put(x)
}

var PoolLabelsMany = poolLabelsMany{}

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
	x.Lat = 0
	x.Long = 0
	p.pool.Put(x)
}

var PoolInputGeoLocation = poolInputGeoLocation{}

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
	x.Lat = 0
	x.Long = 0
	p.pool.Put(x)
}

var PoolGeoLocation = poolGeoLocation{}

const C_InputTeam int64 = 2937769744

type poolInputTeam struct {
	pool sync.Pool
}

func (p *poolInputTeam) Get() *InputTeam {
	x, ok := p.pool.Get().(*InputTeam)
	if !ok {
		return &InputTeam{}
	}
	return x
}

func (p *poolInputTeam) Put(x *InputTeam) {
	x.ID = 0
	x.AccessHash = 0
	p.pool.Put(x)
}

var PoolInputTeam = poolInputTeam{}

const C_TeamPhoto int64 = 3678645727

type poolTeamPhoto struct {
	pool sync.Pool
}

func (p *poolTeamPhoto) Get() *TeamPhoto {
	x, ok := p.pool.Get().(*TeamPhoto)
	if !ok {
		return &TeamPhoto{}
	}
	return x
}

func (p *poolTeamPhoto) Put(x *TeamPhoto) {
	if x.PhotoBig != nil {
		PoolFileLocation.Put(x.PhotoBig)
		x.PhotoBig = nil
	}
	if x.PhotoSmall != nil {
		PoolFileLocation.Put(x.PhotoSmall)
		x.PhotoSmall = nil
	}
	p.pool.Put(x)
}

var PoolTeamPhoto = poolTeamPhoto{}

const C_Team int64 = 1691486497

type poolTeam struct {
	pool sync.Pool
}

func (p *poolTeam) Get() *Team {
	x, ok := p.pool.Get().(*Team)
	if !ok {
		return &Team{}
	}
	return x
}

func (p *poolTeam) Put(x *Team) {
	x.ID = 0
	x.Name = ""
	x.CreatorID = 0
	x.AccessHash = 0
	x.Flags = x.Flags[:0]
	x.Capacity = 0
	x.Community = false
	if x.Photo != nil {
		PoolTeamPhoto.Put(x.Photo)
		x.Photo = nil
	}
	p.pool.Put(x)
}

var PoolTeam = poolTeam{}

func init() {
	registry.RegisterConstructor(2246546115, "Ping")
	registry.RegisterConstructor(2171268721, "Pong")
	registry.RegisterConstructor(535232465, "MessageEnvelope")
	registry.RegisterConstructor(1972016308, "MessageContainer")
	registry.RegisterConstructor(2373884514, "UpdateEnvelope")
	registry.RegisterConstructor(661712615, "UpdateContainer")
	registry.RegisterConstructor(2179260159, "ProtoMessage")
	registry.RegisterConstructor(2668405547, "ProtoEncryptedPayload")
	registry.RegisterConstructor(2619118453, "Error")
	registry.RegisterConstructor(447331921, "Ack")
	registry.RegisterConstructor(4122188204, "Bool")
	registry.RegisterConstructor(1120787796, "Dialog")
	registry.RegisterConstructor(3374092470, "InputPeer")
	registry.RegisterConstructor(47470215, "Peer")
	registry.RegisterConstructor(513021899, "InputPassword")
	registry.RegisterConstructor(354669666, "InputFileLocation")
	registry.RegisterConstructor(2432133155, "FileLocation")
	registry.RegisterConstructor(1881347437, "UserPhoto")
	registry.RegisterConstructor(3865689926, "InputUser")
	registry.RegisterConstructor(765557111, "User")
	registry.RegisterConstructor(460099170, "ContactUser")
	registry.RegisterConstructor(961692401, "Bot")
	registry.RegisterConstructor(1852470005, "BotCommands")
	registry.RegisterConstructor(4059496923, "BotInfo")
	registry.RegisterConstructor(3998516135, "GroupPhoto")
	registry.RegisterConstructor(2885774273, "Group")
	registry.RegisterConstructor(205850814, "GroupFull")
	registry.RegisterConstructor(4072279665, "GroupParticipant")
	registry.RegisterConstructor(1677556362, "UserMessage")
	registry.RegisterConstructor(3277490830, "ReactionCounter")
	registry.RegisterConstructor(869564229, "DraftMessage")
	registry.RegisterConstructor(3479443932, "MessageEntity")
	registry.RegisterConstructor(1046601890, "RSAPublicKey")
	registry.RegisterConstructor(2751503049, "DHGroup")
	registry.RegisterConstructor(2672574672, "PhoneContact")
	registry.RegisterConstructor(3475030132, "PeerNotifySettings")
	registry.RegisterConstructor(3882180383, "InputFile")
	registry.RegisterConstructor(4081048424, "InputDocument")
	registry.RegisterConstructor(3954700912, "PrivacyRule")
	registry.RegisterConstructor(3479601132, "Label")
	registry.RegisterConstructor(1423713603, "LabelsMany")
	registry.RegisterConstructor(1403425127, "InputGeoLocation")
	registry.RegisterConstructor(3794405429, "GeoLocation")
	registry.RegisterConstructor(2937769744, "InputTeam")
	registry.RegisterConstructor(3678645727, "TeamPhoto")
	registry.RegisterConstructor(1691486497, "Team")
}

func (x *Ping) DeepCopy(z *Ping) {
	z.ID = x.ID
}

func (x *Pong) DeepCopy(z *Pong) {
	z.ID = x.ID
}

func (x *MessageEnvelope) DeepCopy(z *MessageEnvelope) {
	z.Constructor = x.Constructor
	z.RequestID = x.RequestID
	z.Message = append(z.Message[:0], x.Message...)
	if x.Team != nil {
		z.Team = PoolInputTeam.Get()
		x.Team.DeepCopy(z.Team)
	}
}

func (x *MessageContainer) DeepCopy(z *MessageContainer) {
	z.Length = x.Length
	for idx := range x.Envelopes {
		if x.Envelopes[idx] != nil {
			xx := PoolMessageEnvelope.Get()
			x.Envelopes[idx].DeepCopy(xx)
			z.Envelopes = append(z.Envelopes, xx)
		}
	}
}

func (x *UpdateEnvelope) DeepCopy(z *UpdateEnvelope) {
	z.Constructor = x.Constructor
	z.Update = append(z.Update[:0], x.Update...)
	z.UCount = x.UCount
	z.UpdateID = x.UpdateID
	z.Timestamp = x.Timestamp
}

func (x *UpdateContainer) DeepCopy(z *UpdateContainer) {
	z.Length = x.Length
	for idx := range x.Updates {
		if x.Updates[idx] != nil {
			xx := PoolUpdateEnvelope.Get()
			x.Updates[idx].DeepCopy(xx)
			z.Updates = append(z.Updates, xx)
		}
	}
	z.MinUpdateID = x.MinUpdateID
	z.MaxUpdateID = x.MaxUpdateID
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	for idx := range x.Groups {
		if x.Groups[idx] != nil {
			xx := PoolGroup.Get()
			x.Groups[idx].DeepCopy(xx)
			z.Groups = append(z.Groups, xx)
		}
	}
}

func (x *ProtoMessage) DeepCopy(z *ProtoMessage) {
	z.AuthID = x.AuthID
	z.MessageKey = append(z.MessageKey[:0], x.MessageKey...)
	z.Payload = append(z.Payload[:0], x.Payload...)
}

func (x *ProtoEncryptedPayload) DeepCopy(z *ProtoEncryptedPayload) {
	z.ServerSalt = x.ServerSalt
	z.MessageID = x.MessageID
	z.SessionID = x.SessionID
	if x.Envelope != nil {
		z.Envelope = PoolMessageEnvelope.Get()
		x.Envelope.DeepCopy(z.Envelope)
	}
}

func (x *Error) DeepCopy(z *Error) {
	z.Code = x.Code
	z.Items = x.Items
}

func (x *Ack) DeepCopy(z *Ack) {
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *Bool) DeepCopy(z *Bool) {
	z.Result = x.Result
}

func (x *Dialog) DeepCopy(z *Dialog) {
	z.TeamID = x.TeamID
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.TopMessageID = x.TopMessageID
	z.ReadInboxMaxID = x.ReadInboxMaxID
	z.ReadOutboxMaxID = x.ReadOutboxMaxID
	z.UnreadCount = x.UnreadCount
	z.AccessHash = x.AccessHash
	if x.NotifySettings != nil {
		z.NotifySettings = PoolPeerNotifySettings.Get()
		x.NotifySettings.DeepCopy(z.NotifySettings)
	}
	z.MentionedCount = x.MentionedCount
	z.Pinned = x.Pinned
	if x.Draft != nil {
		z.Draft = PoolDraftMessage.Get()
		x.Draft.DeepCopy(z.Draft)
	}
	z.PinnedMessageID = x.PinnedMessageID
}

func (x *InputPeer) DeepCopy(z *InputPeer) {
	z.ID = x.ID
	z.Type = x.Type
	z.AccessHash = x.AccessHash
}

func (x *Peer) DeepCopy(z *Peer) {
	z.ID = x.ID
	z.Type = x.Type
	z.AccessHash = x.AccessHash
}

func (x *InputPassword) DeepCopy(z *InputPassword) {
	z.SrpID = x.SrpID
	z.A = append(z.A[:0], x.A...)
	z.M1 = append(z.M1[:0], x.M1...)
}

func (x *InputFileLocation) DeepCopy(z *InputFileLocation) {
	z.ClusterID = x.ClusterID
	z.FileID = x.FileID
	z.AccessHash = x.AccessHash
	z.Version = x.Version
}

func (x *FileLocation) DeepCopy(z *FileLocation) {
	z.ClusterID = x.ClusterID
	z.FileID = x.FileID
	z.AccessHash = x.AccessHash
}

func (x *UserPhoto) DeepCopy(z *UserPhoto) {
	if x.PhotoBig != nil {
		z.PhotoBig = PoolFileLocation.Get()
		x.PhotoBig.DeepCopy(z.PhotoBig)
	}
	if x.PhotoSmall != nil {
		z.PhotoSmall = PoolFileLocation.Get()
		x.PhotoSmall.DeepCopy(z.PhotoSmall)
	}
	z.PhotoID = x.PhotoID
}

func (x *InputUser) DeepCopy(z *InputUser) {
	z.UserID = x.UserID
	z.AccessHash = x.AccessHash
}

func (x *User) DeepCopy(z *User) {
	z.ID = x.ID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Username = x.Username
	z.Status = x.Status
	z.Restricted = x.Restricted
	z.AccessHash = x.AccessHash
	if x.Photo != nil {
		z.Photo = PoolUserPhoto.Get()
		x.Photo.DeepCopy(z.Photo)
	}
	z.Bio = x.Bio
	z.Phone = x.Phone
	z.LastSeen = x.LastSeen
	for idx := range x.PhotoGallery {
		if x.PhotoGallery[idx] != nil {
			xx := PoolUserPhoto.Get()
			x.PhotoGallery[idx].DeepCopy(xx)
			z.PhotoGallery = append(z.PhotoGallery, xx)
		}
	}
	z.IsBot = x.IsBot
	z.Deleted = x.Deleted
	z.Blocked = x.Blocked
	if x.BotInfo != nil {
		z.BotInfo = PoolBotInfo.Get()
		x.BotInfo.DeepCopy(z.BotInfo)
	}
	z.Official = x.Official
}

func (x *ContactUser) DeepCopy(z *ContactUser) {
	z.ID = x.ID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.AccessHash = x.AccessHash
	z.Phone = x.Phone
	z.Username = x.Username
	z.ClientID = x.ClientID
	if x.Photo != nil {
		z.Photo = PoolUserPhoto.Get()
		x.Photo.DeepCopy(z.Photo)
	}
}

func (x *Bot) DeepCopy(z *Bot) {
	z.ID = x.ID
	z.Name = x.Name
	z.Username = x.Username
	z.Bio = x.Bio
}

func (x *BotCommands) DeepCopy(z *BotCommands) {
	z.Command = x.Command
	z.Description = x.Description
}

func (x *BotInfo) DeepCopy(z *BotInfo) {
	if x.Bot != nil {
		z.Bot = PoolBot.Get()
		x.Bot.DeepCopy(z.Bot)
	}
	z.UserID = x.UserID
	z.Description = x.Description
	for idx := range x.BotCommands {
		if x.BotCommands[idx] != nil {
			xx := PoolBotCommands.Get()
			x.BotCommands[idx].DeepCopy(xx)
			z.BotCommands = append(z.BotCommands, xx)
		}
	}
	z.InlineGeo = x.InlineGeo
	z.InlinePlaceHolder = x.InlinePlaceHolder
	z.InlineQuery = x.InlineQuery
}

func (x *GroupPhoto) DeepCopy(z *GroupPhoto) {
	if x.PhotoBig != nil {
		z.PhotoBig = PoolFileLocation.Get()
		x.PhotoBig.DeepCopy(z.PhotoBig)
	}
	if x.PhotoSmall != nil {
		z.PhotoSmall = PoolFileLocation.Get()
		x.PhotoSmall.DeepCopy(z.PhotoSmall)
	}
	z.PhotoID = x.PhotoID
}

func (x *Group) DeepCopy(z *Group) {
	z.TeamID = x.TeamID
	z.ID = x.ID
	z.Title = x.Title
	z.CreatedOn = x.CreatedOn
	z.Participants = x.Participants
	z.EditedOn = x.EditedOn
	z.Flags = append(z.Flags[:0], x.Flags...)
	if x.Photo != nil {
		z.Photo = PoolGroupPhoto.Get()
		x.Photo.DeepCopy(z.Photo)
	}
}

func (x *GroupFull) DeepCopy(z *GroupFull) {
	if x.Group != nil {
		z.Group = PoolGroup.Get()
		x.Group.DeepCopy(z.Group)
	}
	for idx := range x.Users {
		if x.Users[idx] != nil {
			xx := PoolUser.Get()
			x.Users[idx].DeepCopy(xx)
			z.Users = append(z.Users, xx)
		}
	}
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolGroupParticipant.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
	if x.NotifySettings != nil {
		z.NotifySettings = PoolPeerNotifySettings.Get()
		x.NotifySettings.DeepCopy(z.NotifySettings)
	}
	for idx := range x.PhotoGallery {
		if x.PhotoGallery[idx] != nil {
			xx := PoolGroupPhoto.Get()
			x.PhotoGallery[idx].DeepCopy(xx)
			z.PhotoGallery = append(z.PhotoGallery, xx)
		}
	}
}

func (x *GroupParticipant) DeepCopy(z *GroupParticipant) {
	z.UserID = x.UserID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Type = x.Type
	z.AccessHash = x.AccessHash
	z.Username = x.Username
	if x.Photo != nil {
		z.Photo = PoolUserPhoto.Get()
		x.Photo.DeepCopy(z.Photo)
	}
}

func (x *UserMessage) DeepCopy(z *UserMessage) {
	z.TeamID = x.TeamID
	z.ID = x.ID
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.CreatedOn = x.CreatedOn
	z.EditedOn = x.EditedOn
	z.Fwd = x.Fwd
	z.FwdSenderID = x.FwdSenderID
	z.FwdChannelID = x.FwdChannelID
	z.FwdChannelMessageID = x.FwdChannelMessageID
	z.Flags = x.Flags
	z.MessageType = x.MessageType
	z.Body = x.Body
	z.SenderID = x.SenderID
	z.ContentRead = x.ContentRead
	z.Inbox = x.Inbox
	z.ReplyTo = x.ReplyTo
	z.MessageAction = x.MessageAction
	z.MessageActionData = append(z.MessageActionData[:0], x.MessageActionData...)
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.MediaType = x.MediaType
	z.Media = append(z.Media[:0], x.Media...)
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	z.ViaBotID = x.ViaBotID
	for idx := range x.Reactions {
		if x.Reactions[idx] != nil {
			xx := PoolReactionCounter.Get()
			x.Reactions[idx].DeepCopy(xx)
			z.Reactions = append(z.Reactions, xx)
		}
	}
	z.YourReactions = append(z.YourReactions[:0], x.YourReactions...)
}

func (x *ReactionCounter) DeepCopy(z *ReactionCounter) {
	z.Reaction = x.Reaction
	z.Total = x.Total
}

func (x *DraftMessage) DeepCopy(z *DraftMessage) {
	z.TeamID = x.TeamID
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.Date = x.Date
	z.Body = x.Body
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.ReplyTo = x.ReplyTo
	z.EditedID = x.EditedID
}

func (x *MessageEntity) DeepCopy(z *MessageEntity) {
	z.Type = x.Type
	z.Offset = x.Offset
	z.Length = x.Length
	z.UserID = x.UserID
}

func (x *RSAPublicKey) DeepCopy(z *RSAPublicKey) {
	z.FingerPrint = x.FingerPrint
	z.N = x.N
	z.E = x.E
}

func (x *DHGroup) DeepCopy(z *DHGroup) {
	z.FingerPrint = x.FingerPrint
	z.Prime = x.Prime
	z.Gen = x.Gen
}

func (x *PhoneContact) DeepCopy(z *PhoneContact) {
	z.ClientID = x.ClientID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Phone = x.Phone
}

func (x *PeerNotifySettings) DeepCopy(z *PeerNotifySettings) {
	z.Flags = x.Flags
	z.MuteUntil = x.MuteUntil
	z.Sound = x.Sound
}

func (x *InputFile) DeepCopy(z *InputFile) {
	z.FileID = x.FileID
	z.TotalParts = x.TotalParts
	z.FileName = x.FileName
	z.MD5Checksum = x.MD5Checksum
}

func (x *InputDocument) DeepCopy(z *InputDocument) {
	z.ID = x.ID
	z.AccessHash = x.AccessHash
	z.ClusterID = x.ClusterID
}

func (x *PrivacyRule) DeepCopy(z *PrivacyRule) {
	z.PrivacyType = x.PrivacyType
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *Label) DeepCopy(z *Label) {
	z.ID = x.ID
	z.Name = x.Name
	z.Colour = x.Colour
	z.Count = x.Count
}

func (x *LabelsMany) DeepCopy(z *LabelsMany) {
	for idx := range x.Labels {
		if x.Labels[idx] != nil {
			xx := PoolLabel.Get()
			x.Labels[idx].DeepCopy(xx)
			z.Labels = append(z.Labels, xx)
		}
	}
	z.Empty = x.Empty
}

func (x *InputGeoLocation) DeepCopy(z *InputGeoLocation) {
	z.Lat = x.Lat
	z.Long = x.Long
}

func (x *GeoLocation) DeepCopy(z *GeoLocation) {
	z.Lat = x.Lat
	z.Long = x.Long
}

func (x *InputTeam) DeepCopy(z *InputTeam) {
	z.ID = x.ID
	z.AccessHash = x.AccessHash
}

func (x *TeamPhoto) DeepCopy(z *TeamPhoto) {
	if x.PhotoBig != nil {
		z.PhotoBig = PoolFileLocation.Get()
		x.PhotoBig.DeepCopy(z.PhotoBig)
	}
	if x.PhotoSmall != nil {
		z.PhotoSmall = PoolFileLocation.Get()
		x.PhotoSmall.DeepCopy(z.PhotoSmall)
	}
}

func (x *Team) DeepCopy(z *Team) {
	z.ID = x.ID
	z.Name = x.Name
	z.CreatorID = x.CreatorID
	z.AccessHash = x.AccessHash
	z.Flags = append(z.Flags[:0], x.Flags...)
	z.Capacity = x.Capacity
	z.Community = x.Community
	if x.Photo != nil {
		z.Photo = PoolTeamPhoto.Get()
		x.Photo.DeepCopy(z.Photo)
	}
}

func (x *Ping) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Ping, x)
}

func (x *Pong) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Pong, x)
}

func (x *MessageEnvelope) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageEnvelope, x)
}

func (x *MessageContainer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageContainer, x)
}

func (x *UpdateEnvelope) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UpdateEnvelope, x)
}

func (x *UpdateContainer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UpdateContainer, x)
}

func (x *ProtoMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ProtoMessage, x)
}

func (x *ProtoEncryptedPayload) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ProtoEncryptedPayload, x)
}

func (x *Error) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Error, x)
}

func (x *Ack) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Ack, x)
}

func (x *Bool) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Bool, x)
}

func (x *Dialog) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Dialog, x)
}

func (x *InputPeer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputPeer, x)
}

func (x *Peer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Peer, x)
}

func (x *InputPassword) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputPassword, x)
}

func (x *InputFileLocation) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputFileLocation, x)
}

func (x *FileLocation) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_FileLocation, x)
}

func (x *UserPhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UserPhoto, x)
}

func (x *InputUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputUser, x)
}

func (x *User) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_User, x)
}

func (x *ContactUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ContactUser, x)
}

func (x *Bot) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Bot, x)
}

func (x *BotCommands) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotCommands, x)
}

func (x *BotInfo) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_BotInfo, x)
}

func (x *GroupPhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_GroupPhoto, x)
}

func (x *Group) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Group, x)
}

func (x *GroupFull) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_GroupFull, x)
}

func (x *GroupParticipant) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_GroupParticipant, x)
}

func (x *UserMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_UserMessage, x)
}

func (x *ReactionCounter) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_ReactionCounter, x)
}

func (x *DraftMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_DraftMessage, x)
}

func (x *MessageEntity) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_MessageEntity, x)
}

func (x *RSAPublicKey) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_RSAPublicKey, x)
}

func (x *DHGroup) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_DHGroup, x)
}

func (x *PhoneContact) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneContact, x)
}

func (x *PeerNotifySettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PeerNotifySettings, x)
}

func (x *InputFile) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputFile, x)
}

func (x *InputDocument) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputDocument, x)
}

func (x *PrivacyRule) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PrivacyRule, x)
}

func (x *Label) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Label, x)
}

func (x *LabelsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_LabelsMany, x)
}

func (x *InputGeoLocation) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputGeoLocation, x)
}

func (x *GeoLocation) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_GeoLocation, x)
}

func (x *InputTeam) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_InputTeam, x)
}

func (x *TeamPhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_TeamPhoto, x)
}

func (x *Team) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_Team, x)
}

func (x *Ping) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Pong) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *MessageEnvelope) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *MessageContainer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UpdateEnvelope) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UpdateContainer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ProtoMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ProtoEncryptedPayload) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Error) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Ack) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Bool) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Dialog) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputPeer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Peer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputPassword) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputFileLocation) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *FileLocation) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UserPhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputUser) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *User) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ContactUser) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Bot) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotCommands) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *BotInfo) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *GroupPhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Group) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *GroupFull) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *GroupParticipant) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *UserMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *ReactionCounter) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *DraftMessage) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *MessageEntity) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *RSAPublicKey) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *DHGroup) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *PhoneContact) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *PeerNotifySettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputFile) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputDocument) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *PrivacyRule) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Label) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *LabelsMany) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputGeoLocation) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *GeoLocation) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *InputTeam) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *TeamPhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Team) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *Ping) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Pong) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *MessageEnvelope) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *MessageContainer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UpdateEnvelope) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UpdateContainer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ProtoMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ProtoEncryptedPayload) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Error) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Ack) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Bool) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Dialog) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputPeer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Peer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputPassword) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputFileLocation) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *FileLocation) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UserPhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *User) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ContactUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Bot) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BotCommands) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *BotInfo) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *GroupPhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Group) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *GroupFull) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *GroupParticipant) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *UserMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *ReactionCounter) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *DraftMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *MessageEntity) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *RSAPublicKey) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *DHGroup) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PhoneContact) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PeerNotifySettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputFile) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputDocument) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *PrivacyRule) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Label) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *LabelsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputGeoLocation) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *GeoLocation) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *InputTeam) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *TeamPhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *Team) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}
