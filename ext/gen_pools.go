package msg 
 
import (
    "sync"
)
var (

PoolMessagesSendMedia = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSendMedia)
            return m
        },
    }
PoolPeer = sync.Pool{
        New: func() interface{} {
            m := new(Peer)
            return m
        },
    }
PoolMediaPhoto = sync.Pool{
        New: func() interface{} {
            m := new(MediaPhoto)
            return m
        },
    }
PoolInitUserBound = sync.Pool{
        New: func() interface{} {
            m := new(InitUserBound)
            return m
        },
    }
PoolMediaWebPage = sync.Pool{
        New: func() interface{} {
            m := new(MediaWebPage)
            return m
        },
    }
PoolClientFile = sync.Pool{
        New: func() interface{} {
            m := new(ClientFile)
            return m
        },
    }
PoolGroupsRemovePhoto = sync.Pool{
        New: func() interface{} {
            m := new(GroupsRemovePhoto)
            return m
        },
    }
PoolUpdateUserTyping = sync.Pool{
        New: func() interface{} {
            m := new(UpdateUserTyping)
            return m
        },
    }
PoolLabelsAddToMessage = sync.Pool{
        New: func() interface{} {
            m := new(LabelsAddToMessage)
            return m
        },
    }
PoolInputMediaGeoLocation = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaGeoLocation)
            return m
        },
    }
PoolMessageActionGroupPhotoChanged = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionGroupPhotoChanged)
            return m
        },
    }
PoolGroupFull = sync.Pool{
        New: func() interface{} {
            m := new(GroupFull)
            return m
        },
    }
PoolUpdateDialogPinned = sync.Pool{
        New: func() interface{} {
            m := new(UpdateDialogPinned)
            return m
        },
    }
PoolReplyKeyboardForceReply = sync.Pool{
        New: func() interface{} {
            m := new(ReplyKeyboardForceReply)
            return m
        },
    }
PoolAccountAuthorization = sync.Pool{
        New: func() interface{} {
            m := new(AccountAuthorization)
            return m
        },
    }
PoolUpdateUserPhoto = sync.Pool{
        New: func() interface{} {
            m := new(UpdateUserPhoto)
            return m
        },
    }
PoolDocumentAttributeAudio = sync.Pool{
        New: func() interface{} {
            m := new(DocumentAttributeAudio)
            return m
        },
    }
PoolButtonRequestGeoLocation = sync.Pool{
        New: func() interface{} {
            m := new(ButtonRequestGeoLocation)
            return m
        },
    }
PoolMusicsUnFollow = sync.Pool{
        New: func() interface{} {
            m := new(MusicsUnFollow)
            return m
        },
    }
PoolInputFileLocation = sync.Pool{
        New: func() interface{} {
            m := new(InputFileLocation)
            return m
        },
    }
PoolUpdateGroupPhoto = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGroupPhoto)
            return m
        },
    }
PoolGroupsAddUser = sync.Pool{
        New: func() interface{} {
            m := new(GroupsAddUser)
            return m
        },
    }
PoolAccountUpdatePhoto = sync.Pool{
        New: func() interface{} {
            m := new(AccountUpdatePhoto)
            return m
        },
    }
PoolAck = sync.Pool{
        New: func() interface{} {
            m := new(Ack)
            return m
        },
    }
PoolSystemGetAppUpdate = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetAppUpdate)
            return m
        },
    }
PoolContactUser = sync.Pool{
        New: func() interface{} {
            m := new(ContactUser)
            return m
        },
    }
PoolTestRequest = sync.Pool{
        New: func() interface{} {
            m := new(TestRequest)
            return m
        },
    }
PoolAccountGetNotifySettings = sync.Pool{
        New: func() interface{} {
            m := new(AccountGetNotifySettings)
            return m
        },
    }
PoolUpdateReadHistoryOutbox = sync.Pool{
        New: func() interface{} {
            m := new(UpdateReadHistoryOutbox)
            return m
        },
    }
PoolDocumentAttributePhoto = sync.Pool{
        New: func() interface{} {
            m := new(DocumentAttributePhoto)
            return m
        },
    }
PoolMessageEnvelope = sync.Pool{
        New: func() interface{} {
            m := new(MessageEnvelope)
            return m
        },
    }
PoolDocument = sync.Pool{
        New: func() interface{} {
            m := new(Document)
            return m
        },
    }
PoolUpdateGetDifference = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGetDifference)
            return m
        },
    }
PoolLabelsAddToDialog = sync.Pool{
        New: func() interface{} {
            m := new(LabelsAddToDialog)
            return m
        },
    }
PoolMessagesSendServiceMessage = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSendServiceMessage)
            return m
        },
    }
PoolInitAuthCompleted = sync.Pool{
        New: func() interface{} {
            m := new(InitAuthCompleted)
            return m
        },
    }
PoolUpdateAccountPrivacy = sync.Pool{
        New: func() interface{} {
            m := new(UpdateAccountPrivacy)
            return m
        },
    }
PoolButtonRequestPhone = sync.Pool{
        New: func() interface{} {
            m := new(ButtonRequestPhone)
            return m
        },
    }
PoolUpdateContainer = sync.Pool{
        New: func() interface{} {
            m := new(UpdateContainer)
            return m
        },
    }
PoolAccountUpdateStatus = sync.Pool{
        New: func() interface{} {
            m := new(AccountUpdateStatus)
            return m
        },
    }
PoolUpdateMessagesDeleted = sync.Pool{
        New: func() interface{} {
            m := new(UpdateMessagesDeleted)
            return m
        },
    }
PoolUpdateGroupAdmins = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGroupAdmins)
            return m
        },
    }
PoolFile = sync.Pool{
        New: func() interface{} {
            m := new(File)
            return m
        },
    }
PoolUser = sync.Pool{
        New: func() interface{} {
            m := new(User)
            return m
        },
    }
PoolUpdateLabelAdded = sync.Pool{
        New: func() interface{} {
            m := new(UpdateLabelAdded)
            return m
        },
    }
PoolUsersMany = sync.Pool{
        New: func() interface{} {
            m := new(UsersMany)
            return m
        },
    }
PoolInputMediaUploadedPhoto = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaUploadedPhoto)
            return m
        },
    }
PoolDraftMessage = sync.Pool{
        New: func() interface{} {
            m := new(DraftMessage)
            return m
        },
    }
PoolInputMediaUploadedDocument = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaUploadedDocument)
            return m
        },
    }
PoolSystemSalts = sync.Pool{
        New: func() interface{} {
            m := new(SystemSalts)
            return m
        },
    }
PoolMessagesSaveDraft = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSaveDraft)
            return m
        },
    }
PoolAccountRegisterDevice = sync.Pool{
        New: func() interface{} {
            m := new(AccountRegisterDevice)
            return m
        },
    }
PoolBot = sync.Pool{
        New: func() interface{} {
            m := new(Bot)
            return m
        },
    }
PoolAuthLogout = sync.Pool{
        New: func() interface{} {
            m := new(AuthLogout)
            return m
        },
    }
PoolButton = sync.Pool{
        New: func() interface{} {
            m := new(Button)
            return m
        },
    }
PoolUsersGet = sync.Pool{
        New: func() interface{} {
            m := new(UsersGet)
            return m
        },
    }
PoolAccountResetAuthorization = sync.Pool{
        New: func() interface{} {
            m := new(AccountResetAuthorization)
            return m
        },
    }
PoolRSAPublicKey = sync.Pool{
        New: func() interface{} {
            m := new(RSAPublicKey)
            return m
        },
    }
PoolMessagesGetDialog = sync.Pool{
        New: func() interface{} {
            m := new(MessagesGetDialog)
            return m
        },
    }
PoolAccountAuthorizations = sync.Pool{
        New: func() interface{} {
            m := new(AccountAuthorizations)
            return m
        },
    }
PoolClientSendMessageMedia = sync.Pool{
        New: func() interface{} {
            m := new(ClientSendMessageMedia)
            return m
        },
    }
PoolAppUpdate = sync.Pool{
        New: func() interface{} {
            m := new(AppUpdate)
            return m
        },
    }
PoolMusicsAdd = sync.Pool{
        New: func() interface{} {
            m := new(MusicsAdd)
            return m
        },
    }
PoolMusicsGet = sync.Pool{
        New: func() interface{} {
            m := new(MusicsGet)
            return m
        },
    }
PoolDialog = sync.Pool{
        New: func() interface{} {
            m := new(Dialog)
            return m
        },
    }
PoolAuthAuthorization = sync.Pool{
        New: func() interface{} {
            m := new(AuthAuthorization)
            return m
        },
    }
PoolAuthRecall = sync.Pool{
        New: func() interface{} {
            m := new(AuthRecall)
            return m
        },
    }
PoolSystemGetPublicKeys = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetPublicKeys)
            return m
        },
    }
PoolMessageActionGroupDeleteUser = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionGroupDeleteUser)
            return m
        },
    }
PoolAccountUploadPhoto = sync.Pool{
        New: func() interface{} {
            m := new(AccountUploadPhoto)
            return m
        },
    }
PoolMessageActionClearHistory = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionClearHistory)
            return m
        },
    }
PoolGroupsCreate = sync.Pool{
        New: func() interface{} {
            m := new(GroupsCreate)
            return m
        },
    }
PoolMessagesReadHistory = sync.Pool{
        New: func() interface{} {
            m := new(MessagesReadHistory)
            return m
        },
    }
PoolSystemGetServerTime = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetServerTime)
            return m
        },
    }
PoolGroupsUpdateAdmin = sync.Pool{
        New: func() interface{} {
            m := new(GroupsUpdateAdmin)
            return m
        },
    }
PoolMessagesToggleDialogPin = sync.Pool{
        New: func() interface{} {
            m := new(MessagesToggleDialogPin)
            return m
        },
    }
PoolMessagesReorderPinnedDialogs = sync.Pool{
        New: func() interface{} {
            m := new(MessagesReorderPinnedDialogs)
            return m
        },
    }
PoolContactsGet = sync.Pool{
        New: func() interface{} {
            m := new(ContactsGet)
            return m
        },
    }
PoolLabelsMany = sync.Pool{
        New: func() interface{} {
            m := new(LabelsMany)
            return m
        },
    }
PoolMessagesGetDialogs = sync.Pool{
        New: func() interface{} {
            m := new(MessagesGetDialogs)
            return m
        },
    }
PoolUpdateGetState = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGetState)
            return m
        },
    }
PoolAccountUpdateUsername = sync.Pool{
        New: func() interface{} {
            m := new(AccountUpdateUsername)
            return m
        },
    }
PoolSystemGetInfo = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetInfo)
            return m
        },
    }
PoolAccountCheckUsername = sync.Pool{
        New: func() interface{} {
            m := new(AccountCheckUsername)
            return m
        },
    }
PoolUpdateReadHistoryInbox = sync.Pool{
        New: func() interface{} {
            m := new(UpdateReadHistoryInbox)
            return m
        },
    }
PoolUpdateTooLong = sync.Pool{
        New: func() interface{} {
            m := new(UpdateTooLong)
            return m
        },
    }
PoolInputMediaContact = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaContact)
            return m
        },
    }
PoolMessagesSetTyping = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSetTyping)
            return m
        },
    }
PoolUpdateDialogPinnedReorder = sync.Pool{
        New: func() interface{} {
            m := new(UpdateDialogPinnedReorder)
            return m
        },
    }
PoolAuthBotRegister = sync.Pool{
        New: func() interface{} {
            m := new(AuthBotRegister)
            return m
        },
    }
PoolGroupsToggleAdmins = sync.Pool{
        New: func() interface{} {
            m := new(GroupsToggleAdmins)
            return m
        },
    }
PoolInitCompleteAuth = sync.Pool{
        New: func() interface{} {
            m := new(InitCompleteAuth)
            return m
        },
    }
PoolAccountSetPrivacy = sync.Pool{
        New: func() interface{} {
            m := new(AccountSetPrivacy)
            return m
        },
    }
PoolUpdateGroupParticipantAdd = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGroupParticipantAdd)
            return m
        },
    }
PoolUserMessage = sync.Pool{
        New: func() interface{} {
            m := new(UserMessage)
            return m
        },
    }
PoolSystemGetSalts = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetSalts)
            return m
        },
    }
PoolMessagesMany = sync.Pool{
        New: func() interface{} {
            m := new(MessagesMany)
            return m
        },
    }
PoolUpdateDifference = sync.Pool{
        New: func() interface{} {
            m := new(UpdateDifference)
            return m
        },
    }
PoolContactsDelete = sync.Pool{
        New: func() interface{} {
            m := new(ContactsDelete)
            return m
        },
    }
PoolMessagesReadContents = sync.Pool{
        New: func() interface{} {
            m := new(MessagesReadContents)
            return m
        },
    }
PoolSystemGetDHGroups = sync.Pool{
        New: func() interface{} {
            m := new(SystemGetDHGroups)
            return m
        },
    }
PoolUpdateGroupParticipantAdmin = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGroupParticipantAdmin)
            return m
        },
    }
PoolPeerMediaInfo = sync.Pool{
        New: func() interface{} {
            m := new(PeerMediaInfo)
            return m
        },
    }
PoolUpdateMessageEdited = sync.Pool{
        New: func() interface{} {
            m := new(UpdateMessageEdited)
            return m
        },
    }
PoolUpdateState = sync.Pool{
        New: func() interface{} {
            m := new(UpdateState)
            return m
        },
    }
PoolUserPhoto = sync.Pool{
        New: func() interface{} {
            m := new(UserPhoto)
            return m
        },
    }
PoolAccountGetPrivacy = sync.Pool{
        New: func() interface{} {
            m := new(AccountGetPrivacy)
            return m
        },
    }
PoolInitBindUser = sync.Pool{
        New: func() interface{} {
            m := new(InitBindUser)
            return m
        },
    }
PoolMessageActionGroupAddUser = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionGroupAddUser)
            return m
        },
    }
PoolMessagesGetPinnedDialogs = sync.Pool{
        New: func() interface{} {
            m := new(MessagesGetPinnedDialogs)
            return m
        },
    }
PoolMessageContainer = sync.Pool{
        New: func() interface{} {
            m := new(MessageContainer)
            return m
        },
    }
PoolMessagesClearHistory = sync.Pool{
        New: func() interface{} {
            m := new(MessagesClearHistory)
            return m
        },
    }
PoolDocumentAttributeVideo = sync.Pool{
        New: func() interface{} {
            m := new(DocumentAttributeVideo)
            return m
        },
    }
PoolTestResponse = sync.Pool{
        New: func() interface{} {
            m := new(TestResponse)
            return m
        },
    }
PoolUpdateDraftMessageCleared = sync.Pool{
        New: func() interface{} {
            m := new(UpdateDraftMessageCleared)
            return m
        },
    }
PoolAccountSetLang = sync.Pool{
        New: func() interface{} {
            m := new(AccountSetLang)
            return m
        },
    }
PoolAccountSetNotifySettings = sync.Pool{
        New: func() interface{} {
            m := new(AccountSetNotifySettings)
            return m
        },
    }
PoolAccountGetAuthorizations = sync.Pool{
        New: func() interface{} {
            m := new(AccountGetAuthorizations)
            return m
        },
    }
PoolUpdateMessageID = sync.Pool{
        New: func() interface{} {
            m := new(UpdateMessageID)
            return m
        },
    }
PoolMessagesGet = sync.Pool{
        New: func() interface{} {
            m := new(MessagesGet)
            return m
        },
    }
PoolContactsImported = sync.Pool{
        New: func() interface{} {
            m := new(ContactsImported)
            return m
        },
    }
PoolMessagesClearDraft = sync.Pool{
        New: func() interface{} {
            m := new(MessagesClearDraft)
            return m
        },
    }
PoolClientPendingMessage = sync.Pool{
        New: func() interface{} {
            m := new(ClientPendingMessage)
            return m
        },
    }
PoolProtoMessage = sync.Pool{
        New: func() interface{} {
            m := new(ProtoMessage)
            return m
        },
    }
PoolLabelsSet = sync.Pool{
        New: func() interface{} {
            m := new(LabelsSet)
            return m
        },
    }
PoolInputMediaPhoto = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaPhoto)
            return m
        },
    }
PoolKeyboardButtonRow = sync.Pool{
        New: func() interface{} {
            m := new(KeyboardButtonRow)
            return m
        },
    }
PoolDocumentAttributeFile = sync.Pool{
        New: func() interface{} {
            m := new(DocumentAttributeFile)
            return m
        },
    }
PoolAuthRegister = sync.Pool{
        New: func() interface{} {
            m := new(AuthRegister)
            return m
        },
    }
PoolAuthCheckedPhone = sync.Pool{
        New: func() interface{} {
            m := new(AuthCheckedPhone)
            return m
        },
    }
PoolMessageActionGroupCreated = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionGroupCreated)
            return m
        },
    }
PoolInputMediaDocument = sync.Pool{
        New: func() interface{} {
            m := new(InputMediaDocument)
            return m
        },
    }
PoolMediaDocument = sync.Pool{
        New: func() interface{} {
            m := new(MediaDocument)
            return m
        },
    }
PoolButtonUrl = sync.Pool{
        New: func() interface{} {
            m := new(ButtonUrl)
            return m
        },
    }
PoolUpdateLabelRemoved = sync.Pool{
        New: func() interface{} {
            m := new(UpdateLabelRemoved)
            return m
        },
    }
PoolInputButtonUrlAuth = sync.Pool{
        New: func() interface{} {
            m := new(InputButtonUrlAuth)
            return m
        },
    }
PoolButtonUrlAuth = sync.Pool{
        New: func() interface{} {
            m := new(ButtonUrlAuth)
            return m
        },
    }
PoolUpdateAuthorizationReset = sync.Pool{
        New: func() interface{} {
            m := new(UpdateAuthorizationReset)
            return m
        },
    }
PoolInitCompleteAuthInternal = sync.Pool{
        New: func() interface{} {
            m := new(InitCompleteAuthInternal)
            return m
        },
    }
PoolUpdateEnvelope = sync.Pool{
        New: func() interface{} {
            m := new(UpdateEnvelope)
            return m
        },
    }
PoolAuthSentCode = sync.Pool{
        New: func() interface{} {
            m := new(AuthSentCode)
            return m
        },
    }
PoolMessageActionContactRegistered = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionContactRegistered)
            return m
        },
    }
PoolMessageActionGroupTitleChanged = sync.Pool{
        New: func() interface{} {
            m := new(MessageActionGroupTitleChanged)
            return m
        },
    }
PoolFileLocation = sync.Pool{
        New: func() interface{} {
            m := new(FileLocation)
            return m
        },
    }
PoolReplyInlineMarkup = sync.Pool{
        New: func() interface{} {
            m := new(ReplyInlineMarkup)
            return m
        },
    }
PoolUpdateGroupParticipantDeleted = sync.Pool{
        New: func() interface{} {
            m := new(UpdateGroupParticipantDeleted)
            return m
        },
    }
PoolMessagesEdit = sync.Pool{
        New: func() interface{} {
            m := new(MessagesEdit)
            return m
        },
    }
PoolLabelsGet = sync.Pool{
        New: func() interface{} {
            m := new(LabelsGet)
            return m
        },
    }
PoolGroupsEditTitle = sync.Pool{
        New: func() interface{} {
            m := new(GroupsEditTitle)
            return m
        },
    }
PoolAuthLogin = sync.Pool{
        New: func() interface{} {
            m := new(AuthLogin)
            return m
        },
    }
PoolError = sync.Pool{
        New: func() interface{} {
            m := new(Error)
            return m
        },
    }
PoolGroupsUploadPhoto = sync.Pool{
        New: func() interface{} {
            m := new(GroupsUploadPhoto)
            return m
        },
    }
PoolMediaGeoLocation = sync.Pool{
        New: func() interface{} {
            m := new(MediaGeoLocation)
            return m
        },
    }
PoolLabelsRemoveFromDialog = sync.Pool{
        New: func() interface{} {
            m := new(LabelsRemoveFromDialog)
            return m
        },
    }
PoolKeyboardButtonEnvelope = sync.Pool{
        New: func() interface{} {
            m := new(KeyboardButtonEnvelope)
            return m
        },
    }
PoolDBMediaInfo = sync.Pool{
        New: func() interface{} {
            m := new(DBMediaInfo)
            return m
        },
    }
PoolMessagesForward = sync.Pool{
        New: func() interface{} {
            m := new(MessagesForward)
            return m
        },
    }
PoolProtoEncryptedPayload = sync.Pool{
        New: func() interface{} {
            m := new(ProtoEncryptedPayload)
            return m
        },
    }
PoolPhoneContact = sync.Pool{
        New: func() interface{} {
            m := new(PhoneContact)
            return m
        },
    }
PoolAuthResendCode = sync.Pool{
        New: func() interface{} {
            m := new(AuthResendCode)
            return m
        },
    }
PoolUpdateUserStatus = sync.Pool{
        New: func() interface{} {
            m := new(UpdateUserStatus)
            return m
        },
    }
PoolClientFileStatus = sync.Pool{
        New: func() interface{} {
            m := new(ClientFileStatus)
            return m
        },
    }
PoolSystemPublicKeys = sync.Pool{
        New: func() interface{} {
            m := new(SystemPublicKeys)
            return m
        },
    }
PoolDHGroup = sync.Pool{
        New: func() interface{} {
            m := new(DHGroup)
            return m
        },
    }
PoolSystemInfo = sync.Pool{
        New: func() interface{} {
            m := new(SystemInfo)
            return m
        },
    }
PoolInitTestAuth = sync.Pool{
        New: func() interface{} {
            m := new(InitTestAuth)
            return m
        },
    }
PoolAuthLoginByToken = sync.Pool{
        New: func() interface{} {
            m := new(AuthLoginByToken)
            return m
        },
    }
PoolSystemServerTime = sync.Pool{
        New: func() interface{} {
            m := new(SystemServerTime)
            return m
        },
    }
PoolEchoWithDelay = sync.Pool{
        New: func() interface{} {
            m := new(EchoWithDelay)
            return m
        },
    }
PoolGroup = sync.Pool{
        New: func() interface{} {
            m := new(Group)
            return m
        },
    }
PoolSystemDHGroups = sync.Pool{
        New: func() interface{} {
            m := new(SystemDHGroups)
            return m
        },
    }
PoolMessagesSent = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSent)
            return m
        },
    }
PoolClientSearchResult = sync.Pool{
        New: func() interface{} {
            m := new(ClientSearchResult)
            return m
        },
    }
PoolGroupsGetFull = sync.Pool{
        New: func() interface{} {
            m := new(GroupsGetFull)
            return m
        },
    }
PoolUpdateReadMessagesContents = sync.Pool{
        New: func() interface{} {
            m := new(UpdateReadMessagesContents)
            return m
        },
    }
PoolButtonBuy = sync.Pool{
        New: func() interface{} {
            m := new(ButtonBuy)
            return m
        },
    }
PoolMessagesSend = sync.Pool{
        New: func() interface{} {
            m := new(MessagesSend)
            return m
        },
    }
PoolClientUpdateMessagesDeleted = sync.Pool{
        New: func() interface{} {
            m := new(ClientUpdateMessagesDeleted)
            return m
        },
    }
PoolGroupsDeleteUser = sync.Pool{
        New: func() interface{} {
            m := new(GroupsDeleteUser)
            return m
        },
    }
PoolUpdateNotifySettings = sync.Pool{
        New: func() interface{} {
            m := new(UpdateNotifySettings)
            return m
        },
    }
PoolInitConnectTest = sync.Pool{
        New: func() interface{} {
            m := new(InitConnectTest)
            return m
        },
    }
PoolReplyKeyboardMarkup = sync.Pool{
        New: func() interface{} {
            m := new(ReplyKeyboardMarkup)
            return m
        },
    }
PoolAuthRecalled = sync.Pool{
        New: func() interface{} {
            m := new(AuthRecalled)
            return m
        },
    }
PoolMessagesDialogs = sync.Pool{
        New: func() interface{} {
            m := new(MessagesDialogs)
            return m
        },
    }
PoolStartBot = sync.Pool{
        New: func() interface{} {
            m := new(StartBot)
            return m
        },
    }
PoolAuthBotAuthorization = sync.Pool{
        New: func() interface{} {
            m := new(AuthBotAuthorization)
            return m
        },
    }
PoolUsersGetFull = sync.Pool{
        New: func() interface{} {
            m := new(UsersGetFull)
            return m
        },
    }
PoolInputPeer = sync.Pool{
        New: func() interface{} {
            m := new(InputPeer)
            return m
        },
    }
PoolMessagesGetHistory = sync.Pool{
        New: func() interface{} {
            m := new(MessagesGetHistory)
            return m
        },
    }
PoolLabelsDelete = sync.Pool{
        New: func() interface{} {
            m := new(LabelsDelete)
            return m
        },
    }
PoolUpdateNewMessage = sync.Pool{
        New: func() interface{} {
            m := new(UpdateNewMessage)
            return m
        },
    }
PoolGroupsUpdatePhoto = sync.Pool{
        New: func() interface{} {
            m := new(GroupsUpdatePhoto)
            return m
        },
    }
PoolUpdateDraftMessage = sync.Pool{
        New: func() interface{} {
            m := new(UpdateDraftMessage)
            return m
        },
    }
PoolContactsImport = sync.Pool{
        New: func() interface{} {
            m := new(ContactsImport)
            return m
        },
    }
PoolPeerNotifySettings = sync.Pool{
        New: func() interface{} {
            m := new(PeerNotifySettings)
            return m
        },
    }
PoolMessageEntity = sync.Pool{
        New: func() interface{} {
            m := new(MessageEntity)
            return m
        },
    }
PoolLabel = sync.Pool{
        New: func() interface{} {
            m := new(Label)
            return m
        },
    }
PoolMessagesDelete = sync.Pool{
        New: func() interface{} {
            m := new(MessagesDelete)
            return m
        },
    }
PoolMediaSize = sync.Pool{
        New: func() interface{} {
            m := new(MediaSize)
            return m
        },
    }
PoolAuthDestroyKey = sync.Pool{
        New: func() interface{} {
            m := new(AuthDestroyKey)
            return m
        },
    }
PoolAccountUpdateProfile = sync.Pool{
        New: func() interface{} {
            m := new(AccountUpdateProfile)
            return m
        },
    }
PoolAccountRemovePhoto = sync.Pool{
        New: func() interface{} {
            m := new(AccountRemovePhoto)
            return m
        },
    }
PoolMediaContact = sync.Pool{
        New: func() interface{} {
            m := new(MediaContact)
            return m
        },
    }
PoolFileSavePart = sync.Pool{
        New: func() interface{} {
            m := new(FileSavePart)
            return m
        },
    }
PoolAccountPrivacyRules = sync.Pool{
        New: func() interface{} {
            m := new(AccountPrivacyRules)
            return m
        },
    }
PoolClientUpdatePendingMessageDelivery = sync.Pool{
        New: func() interface{} {
            m := new(ClientUpdatePendingMessageDelivery)
            return m
        },
    }
PoolButtonSwitchInline = sync.Pool{
        New: func() interface{} {
            m := new(ButtonSwitchInline)
            return m
        },
    }
PoolInputUser = sync.Pool{
        New: func() interface{} {
            m := new(InputUser)
            return m
        },
    }
PoolInputFile = sync.Pool{
        New: func() interface{} {
            m := new(InputFile)
            return m
        },
    }
PoolContactsMany = sync.Pool{
        New: func() interface{} {
            m := new(ContactsMany)
            return m
        },
    }
PoolMusicsFollow = sync.Pool{
        New: func() interface{} {
            m := new(MusicsFollow)
            return m
        },
    }
PoolMusicsRemove = sync.Pool{
        New: func() interface{} {
            m := new(MusicsRemove)
            return m
        },
    }
PoolPrivacyRule = sync.Pool{
        New: func() interface{} {
            m := new(PrivacyRule)
            return m
        },
    }
PoolAccountUnregisterDevice = sync.Pool{
        New: func() interface{} {
            m := new(AccountUnregisterDevice)
            return m
        },
    }
PoolAuthSendCode = sync.Pool{
        New: func() interface{} {
            m := new(AuthSendCode)
            return m
        },
    }
PoolGroupPhoto = sync.Pool{
        New: func() interface{} {
            m := new(GroupPhoto)
            return m
        },
    }
PoolButtonCallback = sync.Pool{
        New: func() interface{} {
            m := new(ButtonCallback)
            return m
        },
    }
PoolMusicsSetStatus = sync.Pool{
        New: func() interface{} {
            m := new(MusicsSetStatus)
            return m
        },
    }
PoolGroupParticipant = sync.Pool{
        New: func() interface{} {
            m := new(GroupParticipant)
            return m
        },
    }
PoolInputDocument = sync.Pool{
        New: func() interface{} {
            m := new(InputDocument)
            return m
        },
    }
PoolBool = sync.Pool{
        New: func() interface{} {
            m := new(Bool)
            return m
        },
    }
PoolInitResponse = sync.Pool{
        New: func() interface{} {
            m := new(InitResponse)
            return m
        },
    }
PoolAuthCheckPhone = sync.Pool{
        New: func() interface{} {
            m := new(AuthCheckPhone)
            return m
        },
    }
PoolDocumentAttribute = sync.Pool{
        New: func() interface{} {
            m := new(DocumentAttribute)
            return m
        },
    }
PoolInitConnect = sync.Pool{
        New: func() interface{} {
            m := new(InitConnect)
            return m
        },
    }
PoolLabelsRemoveFromMessage = sync.Pool{
        New: func() interface{} {
            m := new(LabelsRemoveFromMessage)
            return m
        },
    }
PoolClientUpdateSynced = sync.Pool{
        New: func() interface{} {
            m := new(ClientUpdateSynced)
            return m
        },
    }
PoolFileGet = sync.Pool{
        New: func() interface{} {
            m := new(FileGet)
            return m
        },
    }
PoolAccountChangePhone = sync.Pool{
        New: func() interface{} {
            m := new(AccountChangePhone)
            return m
        },
    }
PoolUpdateUsername = sync.Pool{
        New: func() interface{} {
            m := new(UpdateUsername)
            return m
        },
    }
)
