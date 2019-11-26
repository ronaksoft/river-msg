package msg 
 
import (
	"github.com/gobwas/pool/pbytes"
)

func ResultMessagesSendMedia(out *MessageEnvelope, res *MessagesSendMedia) {
	out.Constructor = C_MessagesSendMedia
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultPeer(out *MessageEnvelope, res *Peer) {
	out.Constructor = C_Peer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaPhoto(out *MessageEnvelope, res *MediaPhoto) {
	out.Constructor = C_MediaPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitUserBound(out *MessageEnvelope, res *InitUserBound) {
	out.Constructor = C_InitUserBound
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesBroadcast(out *MessageEnvelope, res *MessagesBroadcast) {
	out.Constructor = C_MessagesBroadcast
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaWebPage(out *MessageEnvelope, res *MediaWebPage) {
	out.Constructor = C_MediaWebPage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientFile(out *MessageEnvelope, res *ClientFile) {
	out.Constructor = C_ClientFile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsRemovePhoto(out *MessageEnvelope, res *GroupsRemovePhoto) {
	out.Constructor = C_GroupsRemovePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateUserTyping(out *MessageEnvelope, res *UpdateUserTyping) {
	out.Constructor = C_UpdateUserTyping
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsAddToMessage(out *MessageEnvelope, res *LabelsAddToMessage) {
	out.Constructor = C_LabelsAddToMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaGeoLocation(out *MessageEnvelope, res *InputMediaGeoLocation) {
	out.Constructor = C_InputMediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionGroupPhotoChanged(out *MessageEnvelope, res *MessageActionGroupPhotoChanged) {
	out.Constructor = C_MessageActionGroupPhotoChanged
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupFull(out *MessageEnvelope, res *GroupFull) {
	out.Constructor = C_GroupFull
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateDialogPinned(out *MessageEnvelope, res *UpdateDialogPinned) {
	out.Constructor = C_UpdateDialogPinned
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultReplyKeyboardForceReply(out *MessageEnvelope, res *ReplyKeyboardForceReply) {
	out.Constructor = C_ReplyKeyboardForceReply
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountAuthorization(out *MessageEnvelope, res *AccountAuthorization) {
	out.Constructor = C_AccountAuthorization
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateUserPhoto(out *MessageEnvelope, res *UpdateUserPhoto) {
	out.Constructor = C_UpdateUserPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocumentAttributeAudio(out *MessageEnvelope, res *DocumentAttributeAudio) {
	out.Constructor = C_DocumentAttributeAudio
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonRequestGeoLocation(out *MessageEnvelope, res *ButtonRequestGeoLocation) {
	out.Constructor = C_ButtonRequestGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsUnFollow(out *MessageEnvelope, res *MusicsUnFollow) {
	out.Constructor = C_MusicsUnFollow
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputFileLocation(out *MessageEnvelope, res *InputFileLocation) {
	out.Constructor = C_InputFileLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGroupPhoto(out *MessageEnvelope, res *UpdateGroupPhoto) {
	out.Constructor = C_UpdateGroupPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsAddUser(out *MessageEnvelope, res *GroupsAddUser) {
	out.Constructor = C_GroupsAddUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUpdatePhoto(out *MessageEnvelope, res *AccountUpdatePhoto) {
	out.Constructor = C_AccountUpdatePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAck(out *MessageEnvelope, res *Ack) {
	out.Constructor = C_Ack
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetAppUpdate(out *MessageEnvelope, res *SystemGetAppUpdate) {
	out.Constructor = C_SystemGetAppUpdate
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactUser(out *MessageEnvelope, res *ContactUser) {
	out.Constructor = C_ContactUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultTestRequest(out *MessageEnvelope, res *TestRequest) {
	out.Constructor = C_TestRequest
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountGetNotifySettings(out *MessageEnvelope, res *AccountGetNotifySettings) {
	out.Constructor = C_AccountGetNotifySettings
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateReadHistoryOutbox(out *MessageEnvelope, res *UpdateReadHistoryOutbox) {
	out.Constructor = C_UpdateReadHistoryOutbox
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocumentAttributePhoto(out *MessageEnvelope, res *DocumentAttributePhoto) {
	out.Constructor = C_DocumentAttributePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageEnvelope(out *MessageEnvelope, res *MessageEnvelope) {
	out.Constructor = C_MessageEnvelope
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocument(out *MessageEnvelope, res *Document) {
	out.Constructor = C_Document
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGetDifference(out *MessageEnvelope, res *UpdateGetDifference) {
	out.Constructor = C_UpdateGetDifference
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsAddToDialog(out *MessageEnvelope, res *LabelsAddToDialog) {
	out.Constructor = C_LabelsAddToDialog
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitAuthCompleted(out *MessageEnvelope, res *InitAuthCompleted) {
	out.Constructor = C_InitAuthCompleted
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateAccountPrivacy(out *MessageEnvelope, res *UpdateAccountPrivacy) {
	out.Constructor = C_UpdateAccountPrivacy
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonRequestPhone(out *MessageEnvelope, res *ButtonRequestPhone) {
	out.Constructor = C_ButtonRequestPhone
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateContainer(out *MessageEnvelope, res *UpdateContainer) {
	out.Constructor = C_UpdateContainer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUpdateStatus(out *MessageEnvelope, res *AccountUpdateStatus) {
	out.Constructor = C_AccountUpdateStatus
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateMessagesDeleted(out *MessageEnvelope, res *UpdateMessagesDeleted) {
	out.Constructor = C_UpdateMessagesDeleted
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGroupAdmins(out *MessageEnvelope, res *UpdateGroupAdmins) {
	out.Constructor = C_UpdateGroupAdmins
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFile(out *MessageEnvelope, res *File) {
	out.Constructor = C_File
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUser(out *MessageEnvelope, res *User) {
	out.Constructor = C_User
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesBroadcastProgress(out *MessageEnvelope, res *MessagesBroadcastProgress) {
	out.Constructor = C_MessagesBroadcastProgress
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateLabelAdded(out *MessageEnvelope, res *UpdateLabelAdded) {
	out.Constructor = C_UpdateLabelAdded
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUsersMany(out *MessageEnvelope, res *UsersMany) {
	out.Constructor = C_UsersMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaUploadedPhoto(out *MessageEnvelope, res *InputMediaUploadedPhoto) {
	out.Constructor = C_InputMediaUploadedPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDraftMessage(out *MessageEnvelope, res *DraftMessage) {
	out.Constructor = C_DraftMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaUploadedDocument(out *MessageEnvelope, res *InputMediaUploadedDocument) {
	out.Constructor = C_InputMediaUploadedDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemSalts(out *MessageEnvelope, res *SystemSalts) {
	out.Constructor = C_SystemSalts
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesSaveDraft(out *MessageEnvelope, res *MessagesSaveDraft) {
	out.Constructor = C_MessagesSaveDraft
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountRegisterDevice(out *MessageEnvelope, res *AccountRegisterDevice) {
	out.Constructor = C_AccountRegisterDevice
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultBot(out *MessageEnvelope, res *Bot) {
	out.Constructor = C_Bot
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthLogout(out *MessageEnvelope, res *AuthLogout) {
	out.Constructor = C_AuthLogout
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButton(out *MessageEnvelope, res *Button) {
	out.Constructor = C_Button
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUsersGet(out *MessageEnvelope, res *UsersGet) {
	out.Constructor = C_UsersGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountResetAuthorization(out *MessageEnvelope, res *AccountResetAuthorization) {
	out.Constructor = C_AccountResetAuthorization
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultRSAPublicKey(out *MessageEnvelope, res *RSAPublicKey) {
	out.Constructor = C_RSAPublicKey
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesGetDialog(out *MessageEnvelope, res *MessagesGetDialog) {
	out.Constructor = C_MessagesGetDialog
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountAuthorizations(out *MessageEnvelope, res *AccountAuthorizations) {
	out.Constructor = C_AccountAuthorizations
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientSendMessageMedia(out *MessageEnvelope, res *ClientSendMessageMedia) {
	out.Constructor = C_ClientSendMessageMedia
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAppUpdate(out *MessageEnvelope, res *AppUpdate) {
	out.Constructor = C_AppUpdate
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsAdd(out *MessageEnvelope, res *MusicsAdd) {
	out.Constructor = C_MusicsAdd
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsGet(out *MessageEnvelope, res *MusicsGet) {
	out.Constructor = C_MusicsGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDialog(out *MessageEnvelope, res *Dialog) {
	out.Constructor = C_Dialog
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthAuthorization(out *MessageEnvelope, res *AuthAuthorization) {
	out.Constructor = C_AuthAuthorization
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthRecall(out *MessageEnvelope, res *AuthRecall) {
	out.Constructor = C_AuthRecall
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetPublicKeys(out *MessageEnvelope, res *SystemGetPublicKeys) {
	out.Constructor = C_SystemGetPublicKeys
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionGroupDeleteUser(out *MessageEnvelope, res *MessageActionGroupDeleteUser) {
	out.Constructor = C_MessageActionGroupDeleteUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUploadPhoto(out *MessageEnvelope, res *AccountUploadPhoto) {
	out.Constructor = C_AccountUploadPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionClearHistory(out *MessageEnvelope, res *MessageActionClearHistory) {
	out.Constructor = C_MessageActionClearHistory
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsCreate(out *MessageEnvelope, res *GroupsCreate) {
	out.Constructor = C_GroupsCreate
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesReadHistory(out *MessageEnvelope, res *MessagesReadHistory) {
	out.Constructor = C_MessagesReadHistory
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetServerTime(out *MessageEnvelope, res *SystemGetServerTime) {
	out.Constructor = C_SystemGetServerTime
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsUpdateAdmin(out *MessageEnvelope, res *GroupsUpdateAdmin) {
	out.Constructor = C_GroupsUpdateAdmin
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesToggleDialogPin(out *MessageEnvelope, res *MessagesToggleDialogPin) {
	out.Constructor = C_MessagesToggleDialogPin
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesReorderPinnedDialogs(out *MessageEnvelope, res *MessagesReorderPinnedDialogs) {
	out.Constructor = C_MessagesReorderPinnedDialogs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactsGet(out *MessageEnvelope, res *ContactsGet) {
	out.Constructor = C_ContactsGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsMany(out *MessageEnvelope, res *LabelsMany) {
	out.Constructor = C_LabelsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesGetDialogs(out *MessageEnvelope, res *MessagesGetDialogs) {
	out.Constructor = C_MessagesGetDialogs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGetState(out *MessageEnvelope, res *UpdateGetState) {
	out.Constructor = C_UpdateGetState
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUpdateUsername(out *MessageEnvelope, res *AccountUpdateUsername) {
	out.Constructor = C_AccountUpdateUsername
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetInfo(out *MessageEnvelope, res *SystemGetInfo) {
	out.Constructor = C_SystemGetInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountCheckUsername(out *MessageEnvelope, res *AccountCheckUsername) {
	out.Constructor = C_AccountCheckUsername
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateReadHistoryInbox(out *MessageEnvelope, res *UpdateReadHistoryInbox) {
	out.Constructor = C_UpdateReadHistoryInbox
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateTooLong(out *MessageEnvelope, res *UpdateTooLong) {
	out.Constructor = C_UpdateTooLong
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaContact(out *MessageEnvelope, res *InputMediaContact) {
	out.Constructor = C_InputMediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesSetTyping(out *MessageEnvelope, res *MessagesSetTyping) {
	out.Constructor = C_MessagesSetTyping
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateDialogPinnedReorder(out *MessageEnvelope, res *UpdateDialogPinnedReorder) {
	out.Constructor = C_UpdateDialogPinnedReorder
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthBotRegister(out *MessageEnvelope, res *AuthBotRegister) {
	out.Constructor = C_AuthBotRegister
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsToggleAdmins(out *MessageEnvelope, res *GroupsToggleAdmins) {
	out.Constructor = C_GroupsToggleAdmins
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitCompleteAuth(out *MessageEnvelope, res *InitCompleteAuth) {
	out.Constructor = C_InitCompleteAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountSetPrivacy(out *MessageEnvelope, res *AccountSetPrivacy) {
	out.Constructor = C_AccountSetPrivacy
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGroupParticipantAdd(out *MessageEnvelope, res *UpdateGroupParticipantAdd) {
	out.Constructor = C_UpdateGroupParticipantAdd
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUserMessage(out *MessageEnvelope, res *UserMessage) {
	out.Constructor = C_UserMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetSalts(out *MessageEnvelope, res *SystemGetSalts) {
	out.Constructor = C_SystemGetSalts
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesMany(out *MessageEnvelope, res *MessagesMany) {
	out.Constructor = C_MessagesMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateDifference(out *MessageEnvelope, res *UpdateDifference) {
	out.Constructor = C_UpdateDifference
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactsDelete(out *MessageEnvelope, res *ContactsDelete) {
	out.Constructor = C_ContactsDelete
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesReadContents(out *MessageEnvelope, res *MessagesReadContents) {
	out.Constructor = C_MessagesReadContents
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemGetDHGroups(out *MessageEnvelope, res *SystemGetDHGroups) {
	out.Constructor = C_SystemGetDHGroups
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGroupParticipantAdmin(out *MessageEnvelope, res *UpdateGroupParticipantAdmin) {
	out.Constructor = C_UpdateGroupParticipantAdmin
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultPeerMediaInfo(out *MessageEnvelope, res *PeerMediaInfo) {
	out.Constructor = C_PeerMediaInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateMessageEdited(out *MessageEnvelope, res *UpdateMessageEdited) {
	out.Constructor = C_UpdateMessageEdited
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateState(out *MessageEnvelope, res *UpdateState) {
	out.Constructor = C_UpdateState
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUserPhoto(out *MessageEnvelope, res *UserPhoto) {
	out.Constructor = C_UserPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountGetPrivacy(out *MessageEnvelope, res *AccountGetPrivacy) {
	out.Constructor = C_AccountGetPrivacy
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitBindUser(out *MessageEnvelope, res *InitBindUser) {
	out.Constructor = C_InitBindUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionGroupAddUser(out *MessageEnvelope, res *MessageActionGroupAddUser) {
	out.Constructor = C_MessageActionGroupAddUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesGetPinnedDialogs(out *MessageEnvelope, res *MessagesGetPinnedDialogs) {
	out.Constructor = C_MessagesGetPinnedDialogs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageContainer(out *MessageEnvelope, res *MessageContainer) {
	out.Constructor = C_MessageContainer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesClearHistory(out *MessageEnvelope, res *MessagesClearHistory) {
	out.Constructor = C_MessagesClearHistory
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocumentAttributeVideo(out *MessageEnvelope, res *DocumentAttributeVideo) {
	out.Constructor = C_DocumentAttributeVideo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultTestResponse(out *MessageEnvelope, res *TestResponse) {
	out.Constructor = C_TestResponse
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateDraftMessageCleared(out *MessageEnvelope, res *UpdateDraftMessageCleared) {
	out.Constructor = C_UpdateDraftMessageCleared
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountSetLang(out *MessageEnvelope, res *AccountSetLang) {
	out.Constructor = C_AccountSetLang
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountSetNotifySettings(out *MessageEnvelope, res *AccountSetNotifySettings) {
	out.Constructor = C_AccountSetNotifySettings
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFileGetMany(out *MessageEnvelope, res *FileGetMany) {
	out.Constructor = C_FileGetMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountGetAuthorizations(out *MessageEnvelope, res *AccountGetAuthorizations) {
	out.Constructor = C_AccountGetAuthorizations
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateMessageID(out *MessageEnvelope, res *UpdateMessageID) {
	out.Constructor = C_UpdateMessageID
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesGet(out *MessageEnvelope, res *MessagesGet) {
	out.Constructor = C_MessagesGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactsImported(out *MessageEnvelope, res *ContactsImported) {
	out.Constructor = C_ContactsImported
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesClearDraft(out *MessageEnvelope, res *MessagesClearDraft) {
	out.Constructor = C_MessagesClearDraft
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientPendingMessage(out *MessageEnvelope, res *ClientPendingMessage) {
	out.Constructor = C_ClientPendingMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultProtoMessage(out *MessageEnvelope, res *ProtoMessage) {
	out.Constructor = C_ProtoMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsSet(out *MessageEnvelope, res *LabelsSet) {
	out.Constructor = C_LabelsSet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaPhoto(out *MessageEnvelope, res *InputMediaPhoto) {
	out.Constructor = C_InputMediaPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultKeyboardButtonRow(out *MessageEnvelope, res *KeyboardButtonRow) {
	out.Constructor = C_KeyboardButtonRow
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocumentAttributeFile(out *MessageEnvelope, res *DocumentAttributeFile) {
	out.Constructor = C_DocumentAttributeFile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthRegister(out *MessageEnvelope, res *AuthRegister) {
	out.Constructor = C_AuthRegister
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthCheckedPhone(out *MessageEnvelope, res *AuthCheckedPhone) {
	out.Constructor = C_AuthCheckedPhone
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionGroupCreated(out *MessageEnvelope, res *MessageActionGroupCreated) {
	out.Constructor = C_MessageActionGroupCreated
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputMediaDocument(out *MessageEnvelope, res *InputMediaDocument) {
	out.Constructor = C_InputMediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaDocument(out *MessageEnvelope, res *MediaDocument) {
	out.Constructor = C_MediaDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonUrl(out *MessageEnvelope, res *ButtonUrl) {
	out.Constructor = C_ButtonUrl
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateLabelRemoved(out *MessageEnvelope, res *UpdateLabelRemoved) {
	out.Constructor = C_UpdateLabelRemoved
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputButtonUrlAuth(out *MessageEnvelope, res *InputButtonUrlAuth) {
	out.Constructor = C_InputButtonUrlAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonUrlAuth(out *MessageEnvelope, res *ButtonUrlAuth) {
	out.Constructor = C_ButtonUrlAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateAuthorizationReset(out *MessageEnvelope, res *UpdateAuthorizationReset) {
	out.Constructor = C_UpdateAuthorizationReset
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitCompleteAuthInternal(out *MessageEnvelope, res *InitCompleteAuthInternal) {
	out.Constructor = C_InitCompleteAuthInternal
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateEnvelope(out *MessageEnvelope, res *UpdateEnvelope) {
	out.Constructor = C_UpdateEnvelope
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthSentCode(out *MessageEnvelope, res *AuthSentCode) {
	out.Constructor = C_AuthSentCode
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionContactRegistered(out *MessageEnvelope, res *MessageActionContactRegistered) {
	out.Constructor = C_MessageActionContactRegistered
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionGroupTitleChanged(out *MessageEnvelope, res *MessageActionGroupTitleChanged) {
	out.Constructor = C_MessageActionGroupTitleChanged
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFileLocation(out *MessageEnvelope, res *FileLocation) {
	out.Constructor = C_FileLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultReplyInlineMarkup(out *MessageEnvelope, res *ReplyInlineMarkup) {
	out.Constructor = C_ReplyInlineMarkup
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateGroupParticipantDeleted(out *MessageEnvelope, res *UpdateGroupParticipantDeleted) {
	out.Constructor = C_UpdateGroupParticipantDeleted
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesEdit(out *MessageEnvelope, res *MessagesEdit) {
	out.Constructor = C_MessagesEdit
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsGet(out *MessageEnvelope, res *LabelsGet) {
	out.Constructor = C_LabelsGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsEditTitle(out *MessageEnvelope, res *GroupsEditTitle) {
	out.Constructor = C_GroupsEditTitle
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthLogin(out *MessageEnvelope, res *AuthLogin) {
	out.Constructor = C_AuthLogin
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultError(out *MessageEnvelope, res *Error) {
	out.Constructor = C_Error
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsUploadPhoto(out *MessageEnvelope, res *GroupsUploadPhoto) {
	out.Constructor = C_GroupsUploadPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaGeoLocation(out *MessageEnvelope, res *MediaGeoLocation) {
	out.Constructor = C_MediaGeoLocation
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsRemoveFromDialog(out *MessageEnvelope, res *LabelsRemoveFromDialog) {
	out.Constructor = C_LabelsRemoveFromDialog
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageActionScreenShotTaken(out *MessageEnvelope, res *MessageActionScreenShotTaken) {
	out.Constructor = C_MessageActionScreenShotTaken
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultKeyboardButtonEnvelope(out *MessageEnvelope, res *KeyboardButtonEnvelope) {
	out.Constructor = C_KeyboardButtonEnvelope
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDBMediaInfo(out *MessageEnvelope, res *DBMediaInfo) {
	out.Constructor = C_DBMediaInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesForward(out *MessageEnvelope, res *MessagesForward) {
	out.Constructor = C_MessagesForward
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultProtoEncryptedPayload(out *MessageEnvelope, res *ProtoEncryptedPayload) {
	out.Constructor = C_ProtoEncryptedPayload
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultPhoneContact(out *MessageEnvelope, res *PhoneContact) {
	out.Constructor = C_PhoneContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthResendCode(out *MessageEnvelope, res *AuthResendCode) {
	out.Constructor = C_AuthResendCode
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateUserStatus(out *MessageEnvelope, res *UpdateUserStatus) {
	out.Constructor = C_UpdateUserStatus
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientFileStatus(out *MessageEnvelope, res *ClientFileStatus) {
	out.Constructor = C_ClientFileStatus
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemPublicKeys(out *MessageEnvelope, res *SystemPublicKeys) {
	out.Constructor = C_SystemPublicKeys
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDHGroup(out *MessageEnvelope, res *DHGroup) {
	out.Constructor = C_DHGroup
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemInfo(out *MessageEnvelope, res *SystemInfo) {
	out.Constructor = C_SystemInfo
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitTestAuth(out *MessageEnvelope, res *InitTestAuth) {
	out.Constructor = C_InitTestAuth
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthLoginByToken(out *MessageEnvelope, res *AuthLoginByToken) {
	out.Constructor = C_AuthLoginByToken
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemServerTime(out *MessageEnvelope, res *SystemServerTime) {
	out.Constructor = C_SystemServerTime
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultEchoWithDelay(out *MessageEnvelope, res *EchoWithDelay) {
	out.Constructor = C_EchoWithDelay
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroup(out *MessageEnvelope, res *Group) {
	out.Constructor = C_Group
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultSystemDHGroups(out *MessageEnvelope, res *SystemDHGroups) {
	out.Constructor = C_SystemDHGroups
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesSent(out *MessageEnvelope, res *MessagesSent) {
	out.Constructor = C_MessagesSent
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientSearchResult(out *MessageEnvelope, res *ClientSearchResult) {
	out.Constructor = C_ClientSearchResult
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsGetFull(out *MessageEnvelope, res *GroupsGetFull) {
	out.Constructor = C_GroupsGetFull
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateReadMessagesContents(out *MessageEnvelope, res *UpdateReadMessagesContents) {
	out.Constructor = C_UpdateReadMessagesContents
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonBuy(out *MessageEnvelope, res *ButtonBuy) {
	out.Constructor = C_ButtonBuy
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesSend(out *MessageEnvelope, res *MessagesSend) {
	out.Constructor = C_MessagesSend
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientUpdateMessagesDeleted(out *MessageEnvelope, res *ClientUpdateMessagesDeleted) {
	out.Constructor = C_ClientUpdateMessagesDeleted
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsDeleteUser(out *MessageEnvelope, res *GroupsDeleteUser) {
	out.Constructor = C_GroupsDeleteUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateNotifySettings(out *MessageEnvelope, res *UpdateNotifySettings) {
	out.Constructor = C_UpdateNotifySettings
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitConnectTest(out *MessageEnvelope, res *InitConnectTest) {
	out.Constructor = C_InitConnectTest
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultReplyKeyboardMarkup(out *MessageEnvelope, res *ReplyKeyboardMarkup) {
	out.Constructor = C_ReplyKeyboardMarkup
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthRecalled(out *MessageEnvelope, res *AuthRecalled) {
	out.Constructor = C_AuthRecalled
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesDialogs(out *MessageEnvelope, res *MessagesDialogs) {
	out.Constructor = C_MessagesDialogs
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultStartBot(out *MessageEnvelope, res *StartBot) {
	out.Constructor = C_StartBot
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthBotAuthorization(out *MessageEnvelope, res *AuthBotAuthorization) {
	out.Constructor = C_AuthBotAuthorization
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUsersGetFull(out *MessageEnvelope, res *UsersGetFull) {
	out.Constructor = C_UsersGetFull
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFileMany{(out *MessageEnvelope, res *FileMany{) {
	out.Constructor = C_FileMany{
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputPeer(out *MessageEnvelope, res *InputPeer) {
	out.Constructor = C_InputPeer
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesGetHistory(out *MessageEnvelope, res *MessagesGetHistory) {
	out.Constructor = C_MessagesGetHistory
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsDelete(out *MessageEnvelope, res *LabelsDelete) {
	out.Constructor = C_LabelsDelete
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateNewMessage(out *MessageEnvelope, res *UpdateNewMessage) {
	out.Constructor = C_UpdateNewMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupsUpdatePhoto(out *MessageEnvelope, res *GroupsUpdatePhoto) {
	out.Constructor = C_GroupsUpdatePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateDraftMessage(out *MessageEnvelope, res *UpdateDraftMessage) {
	out.Constructor = C_UpdateDraftMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactsImport(out *MessageEnvelope, res *ContactsImport) {
	out.Constructor = C_ContactsImport
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultPeerNotifySettings(out *MessageEnvelope, res *PeerNotifySettings) {
	out.Constructor = C_PeerNotifySettings
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessageEntity(out *MessageEnvelope, res *MessageEntity) {
	out.Constructor = C_MessageEntity
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabel(out *MessageEnvelope, res *Label) {
	out.Constructor = C_Label
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesDelete(out *MessageEnvelope, res *MessagesDelete) {
	out.Constructor = C_MessagesDelete
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaSize(out *MessageEnvelope, res *MediaSize) {
	out.Constructor = C_MediaSize
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthDestroyKey(out *MessageEnvelope, res *AuthDestroyKey) {
	out.Constructor = C_AuthDestroyKey
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMessagesSendScreenShotNotification(out *MessageEnvelope, res *MessagesSendScreenShotNotification) {
	out.Constructor = C_MessagesSendScreenShotNotification
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUpdateProfile(out *MessageEnvelope, res *AccountUpdateProfile) {
	out.Constructor = C_AccountUpdateProfile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountRemovePhoto(out *MessageEnvelope, res *AccountRemovePhoto) {
	out.Constructor = C_AccountRemovePhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMediaContact(out *MessageEnvelope, res *MediaContact) {
	out.Constructor = C_MediaContact
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFileSavePart(out *MessageEnvelope, res *FileSavePart) {
	out.Constructor = C_FileSavePart
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountPrivacyRules(out *MessageEnvelope, res *AccountPrivacyRules) {
	out.Constructor = C_AccountPrivacyRules
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientUpdatePendingMessageDelivery(out *MessageEnvelope, res *ClientUpdatePendingMessageDelivery) {
	out.Constructor = C_ClientUpdatePendingMessageDelivery
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonSwitchInline(out *MessageEnvelope, res *ButtonSwitchInline) {
	out.Constructor = C_ButtonSwitchInline
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputUser(out *MessageEnvelope, res *InputUser) {
	out.Constructor = C_InputUser
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputFile(out *MessageEnvelope, res *InputFile) {
	out.Constructor = C_InputFile
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultContactsMany(out *MessageEnvelope, res *ContactsMany) {
	out.Constructor = C_ContactsMany
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsFollow(out *MessageEnvelope, res *MusicsFollow) {
	out.Constructor = C_MusicsFollow
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsRemove(out *MessageEnvelope, res *MusicsRemove) {
	out.Constructor = C_MusicsRemove
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultPrivacyRule(out *MessageEnvelope, res *PrivacyRule) {
	out.Constructor = C_PrivacyRule
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountUnregisterDevice(out *MessageEnvelope, res *AccountUnregisterDevice) {
	out.Constructor = C_AccountUnregisterDevice
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthSendCode(out *MessageEnvelope, res *AuthSendCode) {
	out.Constructor = C_AuthSendCode
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupPhoto(out *MessageEnvelope, res *GroupPhoto) {
	out.Constructor = C_GroupPhoto
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultButtonCallback(out *MessageEnvelope, res *ButtonCallback) {
	out.Constructor = C_ButtonCallback
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultMusicsSetStatus(out *MessageEnvelope, res *MusicsSetStatus) {
	out.Constructor = C_MusicsSetStatus
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultGroupParticipant(out *MessageEnvelope, res *GroupParticipant) {
	out.Constructor = C_GroupParticipant
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInputDocument(out *MessageEnvelope, res *InputDocument) {
	out.Constructor = C_InputDocument
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultBool(out *MessageEnvelope, res *Bool) {
	out.Constructor = C_Bool
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitResponse(out *MessageEnvelope, res *InitResponse) {
	out.Constructor = C_InitResponse
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAuthCheckPhone(out *MessageEnvelope, res *AuthCheckPhone) {
	out.Constructor = C_AuthCheckPhone
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultDocumentAttribute(out *MessageEnvelope, res *DocumentAttribute) {
	out.Constructor = C_DocumentAttribute
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultInitConnect(out *MessageEnvelope, res *InitConnect) {
	out.Constructor = C_InitConnect
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultLabelsRemoveFromMessage(out *MessageEnvelope, res *LabelsRemoveFromMessage) {
	out.Constructor = C_LabelsRemoveFromMessage
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultClientUpdateSynced(out *MessageEnvelope, res *ClientUpdateSynced) {
	out.Constructor = C_ClientUpdateSynced
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultFileGet(out *MessageEnvelope, res *FileGet) {
	out.Constructor = C_FileGet
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultAccountChangePhone(out *MessageEnvelope, res *AccountChangePhone) {
	out.Constructor = C_AccountChangePhone
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
func ResultUpdateUsername(out *MessageEnvelope, res *UpdateUsername) {
	out.Constructor = C_UpdateUsername
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}

