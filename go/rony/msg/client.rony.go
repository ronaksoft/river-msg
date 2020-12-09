package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_ClientSendMessageMedia int64 = 1095038539

type poolClientSendMessageMedia struct {
	pool sync.Pool
}

func (p *poolClientSendMessageMedia) Get() *ClientSendMessageMedia {
	x, ok := p.pool.Get().(*ClientSendMessageMedia)
	if !ok {
		return &ClientSendMessageMedia{}
	}
	return x
}

func (p *poolClientSendMessageMedia) Put(x *ClientSendMessageMedia) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MediaType = 0
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
	x.TinyThumb = x.TinyThumb[:0]
	p.pool.Put(x)
}

var PoolClientSendMessageMedia = poolClientSendMessageMedia{}

const C_ClientGlobalSearch int64 = 1742781507

type poolClientGlobalSearch struct {
	pool sync.Pool
}

func (p *poolClientGlobalSearch) Get() *ClientGlobalSearch {
	x, ok := p.pool.Get().(*ClientGlobalSearch)
	if !ok {
		return &ClientGlobalSearch{}
	}
	return x
}

func (p *poolClientGlobalSearch) Put(x *ClientGlobalSearch) {
	x.Text = ""
	x.LabelIDs = x.LabelIDs[:0]
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Limit = 0
	x.SenderID = 0
	p.pool.Put(x)
}

var PoolClientGlobalSearch = poolClientGlobalSearch{}

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
	x.Text = ""
	p.pool.Put(x)
}

var PoolClientContactSearch = poolClientContactSearch{}

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

const C_ClientClearCachedMedia int64 = 1199927718

type poolClientClearCachedMedia struct {
	pool sync.Pool
}

func (p *poolClientClearCachedMedia) Get() *ClientClearCachedMedia {
	x, ok := p.pool.Get().(*ClientClearCachedMedia)
	if !ok {
		return &ClientClearCachedMedia{}
	}
	return x
}

func (p *poolClientClearCachedMedia) Put(x *ClientClearCachedMedia) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.MediaTypes = x.MediaTypes[:0]
	p.pool.Put(x)
}

var PoolClientClearCachedMedia = poolClientClearCachedMedia{}

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
	if x.peer != nil {
		PoolInputPeer.Put(x.peer)
		x.peer = nil
	}
	p.pool.Put(x)
}

var PoolClientGetLastBotKeyboard = poolClientGetLastBotKeyboard{}

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
	x.MediaType = 0
	p.pool.Put(x)
}

var PoolClientGetMediaHistory = poolClientGetMediaHistory{}

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
	x.Limit = 0
	p.pool.Put(x)
}

var PoolClientGetRecentSearch = poolClientGetRecentSearch{}

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
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolClientPutRecentSearch = poolClientPutRecentSearch{}

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
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolClientRemoveRecentSearch = poolClientRemoveRecentSearch{}

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
	x.Extra = false
	p.pool.Put(x)
}

var PoolClientRemoveAllRecentSearches = poolClientRemoveAllRecentSearches{}

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

const C_ClientGetTeamCounters int64 = 3403301140

type poolClientGetTeamCounters struct {
	pool sync.Pool
}

func (p *poolClientGetTeamCounters) Get() *ClientGetTeamCounters {
	x, ok := p.pool.Get().(*ClientGetTeamCounters)
	if !ok {
		return &ClientGetTeamCounters{}
	}
	return x
}

func (p *poolClientGetTeamCounters) Put(x *ClientGetTeamCounters) {
	if x.Team != nil {
		PoolInputTeam.Put(x.Team)
		x.Team = nil
	}
	x.WithMutes = false
	p.pool.Put(x)
}

var PoolClientGetTeamCounters = poolClientGetTeamCounters{}

const C_ClientPendingMessage int64 = 2164891929

type poolClientPendingMessage struct {
	pool sync.Pool
}

func (p *poolClientPendingMessage) Get() *ClientPendingMessage {
	x, ok := p.pool.Get().(*ClientPendingMessage)
	if !ok {
		return &ClientPendingMessage{}
	}
	return x
}

func (p *poolClientPendingMessage) Put(x *ClientPendingMessage) {
	x.ID = 0
	x.RequestID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.AccessHash = 0
	x.CreatedOn = 0
	x.ReplyTo = 0
	x.Body = ""
	x.SenderID = 0
	x.Entities = x.Entities[:0]
	x.MediaType = 0
	x.Media = x.Media[:0]
	x.ClearDraft = false
	x.FileUploadID = ""
	x.ThumbUploadID = ""
	x.FileID = 0
	x.ThumbID = 0
	x.Sha256 = x.Sha256[:0]
	if x.ServerFile != nil {
		PoolFileLocation.Put(x.ServerFile)
		x.ServerFile = nil
	}
	x.TeamID = 0
	x.TeamAccessHash = 0
	x.TinyThumb = x.TinyThumb[:0]
	p.pool.Put(x)
}

var PoolClientPendingMessage = poolClientPendingMessage{}

const C_ClientSearchResult int64 = 2957647709

type poolClientSearchResult struct {
	pool sync.Pool
}

func (p *poolClientSearchResult) Get() *ClientSearchResult {
	x, ok := p.pool.Get().(*ClientSearchResult)
	if !ok {
		return &ClientSearchResult{}
	}
	return x
}

func (p *poolClientSearchResult) Put(x *ClientSearchResult) {
	x.Messages = x.Messages[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	x.MatchedUsers = x.MatchedUsers[:0]
	x.MatchedGroups = x.MatchedGroups[:0]
	p.pool.Put(x)
}

var PoolClientSearchResult = poolClientSearchResult{}

const C_ClientFilesMany int64 = 1414992553

type poolClientFilesMany struct {
	pool sync.Pool
}

func (p *poolClientFilesMany) Get() *ClientFilesMany {
	x, ok := p.pool.Get().(*ClientFilesMany)
	if !ok {
		return &ClientFilesMany{}
	}
	return x
}

func (p *poolClientFilesMany) Put(x *ClientFilesMany) {
	x.Gifs = x.Gifs[:0]
	x.Total = 0
	p.pool.Put(x)
}

var PoolClientFilesMany = poolClientFilesMany{}

const C_ClientFile int64 = 155127968

type poolClientFile struct {
	pool sync.Pool
}

func (p *poolClientFile) Get() *ClientFile {
	x, ok := p.pool.Get().(*ClientFile)
	if !ok {
		return &ClientFile{}
	}
	return x
}

func (p *poolClientFile) Put(x *ClientFile) {
	x.ClusterID = 0
	x.FileID = 0
	x.AccessHash = 0
	x.Type = 0
	x.MimeType = ""
	x.UserID = 0
	x.GroupID = 0
	x.FileSize = 0
	x.MessageID = 0
	x.PeerID = 0
	x.PeerType = 0
	x.Version = 0
	x.Extension = ""
	x.MD5Checksum = ""
	x.WallpaperID = 0
	x.Attributes = x.Attributes[:0]
	p.pool.Put(x)
}

var PoolClientFile = poolClientFile{}

const C_ClientFileRequest int64 = 3995993899

type poolClientFileRequest struct {
	pool sync.Pool
}

func (p *poolClientFileRequest) Get() *ClientFileRequest {
	x, ok := p.pool.Get().(*ClientFileRequest)
	if !ok {
		return &ClientFileRequest{}
	}
	return x
}

func (p *poolClientFileRequest) Put(x *ClientFileRequest) {
	if x.Next != nil {
		PoolClientFileRequest.Put(x.Next)
		x.Next = nil
	}
	x.PeerID = 0
	x.PeerType = 0
	x.MessageID = 0
	x.ClusterID = 0
	x.FileID = 0
	x.AccessHash = 0
	x.Version = 0
	x.FileSize = 0
	x.ChunkSize = 0
	x.FinishedParts = x.FinishedParts[:0]
	x.TotalParts = 0
	x.SkipDelegateCall = false
	x.FilePath = ""
	x.TempPath = ""
	x.CheckSha256 = false
	x.FileSha256 = x.FileSha256[:0]
	x.IsProfilePhoto = false
	x.GroupID = 0
	x.ThumbID = 0
	x.ThumbPath = ""
	p.pool.Put(x)
}

var PoolClientFileRequest = poolClientFileRequest{}

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
	x.Status = 0
	x.Progress = 0
	x.FilePath = ""
	p.pool.Put(x)
}

var PoolClientFileStatus = poolClientFileStatus{}

const C_ClientCachedMediaInfo int64 = 442767121

type poolClientCachedMediaInfo struct {
	pool sync.Pool
}

func (p *poolClientCachedMediaInfo) Get() *ClientCachedMediaInfo {
	x, ok := p.pool.Get().(*ClientCachedMediaInfo)
	if !ok {
		return &ClientCachedMediaInfo{}
	}
	return x
}

func (p *poolClientCachedMediaInfo) Put(x *ClientCachedMediaInfo) {
	x.MediaInfo = x.MediaInfo[:0]
	p.pool.Put(x)
}

var PoolClientCachedMediaInfo = poolClientCachedMediaInfo{}

const C_ClientPeerMediaInfo int64 = 2711408875

type poolClientPeerMediaInfo struct {
	pool sync.Pool
}

func (p *poolClientPeerMediaInfo) Get() *ClientPeerMediaInfo {
	x, ok := p.pool.Get().(*ClientPeerMediaInfo)
	if !ok {
		return &ClientPeerMediaInfo{}
	}
	return x
}

func (p *poolClientPeerMediaInfo) Put(x *ClientPeerMediaInfo) {
	x.PeerID = 0
	x.PeerType = 0
	x.Media = x.Media[:0]
	p.pool.Put(x)
}

var PoolClientPeerMediaInfo = poolClientPeerMediaInfo{}

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
	x.MediaType = 0
	x.TotalSize = 0
	p.pool.Put(x)
}

var PoolClientMediaSize = poolClientMediaSize{}

const C_ClientRecentSearch int64 = 2069517672

type poolClientRecentSearch struct {
	pool sync.Pool
}

func (p *poolClientRecentSearch) Get() *ClientRecentSearch {
	x, ok := p.pool.Get().(*ClientRecentSearch)
	if !ok {
		return &ClientRecentSearch{}
	}
	return x
}

func (p *poolClientRecentSearch) Put(x *ClientRecentSearch) {
	if x.Peer != nil {
		PoolPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Date = 0
	p.pool.Put(x)
}

var PoolClientRecentSearch = poolClientRecentSearch{}

const C_ClientRecentSearchMany int64 = 3236847495

type poolClientRecentSearchMany struct {
	pool sync.Pool
}

func (p *poolClientRecentSearchMany) Get() *ClientRecentSearchMany {
	x, ok := p.pool.Get().(*ClientRecentSearchMany)
	if !ok {
		return &ClientRecentSearchMany{}
	}
	return x
}

func (p *poolClientRecentSearchMany) Put(x *ClientRecentSearchMany) {
	x.RecentSearches = x.RecentSearches[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	p.pool.Put(x)
}

var PoolClientRecentSearchMany = poolClientRecentSearchMany{}

const C_ClientTeamCounters int64 = 769069696

type poolClientTeamCounters struct {
	pool sync.Pool
}

func (p *poolClientTeamCounters) Get() *ClientTeamCounters {
	x, ok := p.pool.Get().(*ClientTeamCounters)
	if !ok {
		return &ClientTeamCounters{}
	}
	return x
}

func (p *poolClientTeamCounters) Put(x *ClientTeamCounters) {
	x.UnreadCount = 0
	x.MentionCount = 0
	p.pool.Put(x)
}

var PoolClientTeamCounters = poolClientTeamCounters{}

const C_ClientGetFrequentlyReactions int64 = 2316768708

type poolClientGetFrequentlyReactions struct {
	pool sync.Pool
}

func (p *poolClientGetFrequentlyReactions) Get() *ClientGetFrequentlyReactions {
	x, ok := p.pool.Get().(*ClientGetFrequentlyReactions)
	if !ok {
		return &ClientGetFrequentlyReactions{}
	}
	return x
}

func (p *poolClientGetFrequentlyReactions) Put(x *ClientGetFrequentlyReactions) {
	p.pool.Put(x)
}

var PoolClientGetFrequentlyReactions = poolClientGetFrequentlyReactions{}

const C_ClientFrequentlyReactions int64 = 2830270594

type poolClientFrequentlyReactions struct {
	pool sync.Pool
}

func (p *poolClientFrequentlyReactions) Get() *ClientFrequentlyReactions {
	x, ok := p.pool.Get().(*ClientFrequentlyReactions)
	if !ok {
		return &ClientFrequentlyReactions{}
	}
	return x
}

func (p *poolClientFrequentlyReactions) Put(x *ClientFrequentlyReactions) {
	x.Reactions = x.Reactions[:0]
	p.pool.Put(x)
}

var PoolClientFrequentlyReactions = poolClientFrequentlyReactions{}

const C_ClientDismissNotification int64 = 1698398006

type poolClientDismissNotification struct {
	pool sync.Pool
}

func (p *poolClientDismissNotification) Get() *ClientDismissNotification {
	x, ok := p.pool.Get().(*ClientDismissNotification)
	if !ok {
		return &ClientDismissNotification{}
	}
	return x
}

func (p *poolClientDismissNotification) Put(x *ClientDismissNotification) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Ts = 0
	p.pool.Put(x)
}

var PoolClientDismissNotification = poolClientDismissNotification{}

const C_ClientGetNotificationDismissTime int64 = 4106535811

type poolClientGetNotificationDismissTime struct {
	pool sync.Pool
}

func (p *poolClientGetNotificationDismissTime) Get() *ClientGetNotificationDismissTime {
	x, ok := p.pool.Get().(*ClientGetNotificationDismissTime)
	if !ok {
		return &ClientGetNotificationDismissTime{}
	}
	return x
}

func (p *poolClientGetNotificationDismissTime) Put(x *ClientGetNotificationDismissTime) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolClientGetNotificationDismissTime = poolClientGetNotificationDismissTime{}

const C_ClientNotificationDismissTime int64 = 3077814065

type poolClientNotificationDismissTime struct {
	pool sync.Pool
}

func (p *poolClientNotificationDismissTime) Get() *ClientNotificationDismissTime {
	x, ok := p.pool.Get().(*ClientNotificationDismissTime)
	if !ok {
		return &ClientNotificationDismissTime{}
	}
	return x
}

func (p *poolClientNotificationDismissTime) Put(x *ClientNotificationDismissTime) {
	x.Ts = 0
	p.pool.Put(x)
}

var PoolClientNotificationDismissTime = poolClientNotificationDismissTime{}

func init() {
	registry.RegisterConstructor(1095038539, "ClientSendMessageMedia")
	registry.RegisterConstructor(1742781507, "ClientGlobalSearch")
	registry.RegisterConstructor(1793449803, "ClientContactSearch")
	registry.RegisterConstructor(856595701, "ClientGetCachedMedia")
	registry.RegisterConstructor(1199927718, "ClientClearCachedMedia")
	registry.RegisterConstructor(177544569, "ClientGetLastBotKeyboard")
	registry.RegisterConstructor(1354863379, "ClientGetMediaHistory")
	registry.RegisterConstructor(2622949116, "ClientGetRecentSearch")
	registry.RegisterConstructor(629582533, "ClientPutRecentSearch")
	registry.RegisterConstructor(1281490259, "ClientRemoveRecentSearch")
	registry.RegisterConstructor(3599155822, "ClientRemoveAllRecentSearches")
	registry.RegisterConstructor(3028067090, "ClientGetSavedGifs")
	registry.RegisterConstructor(3403301140, "ClientGetTeamCounters")
	registry.RegisterConstructor(2164891929, "ClientPendingMessage")
	registry.RegisterConstructor(2957647709, "ClientSearchResult")
	registry.RegisterConstructor(1414992553, "ClientFilesMany")
	registry.RegisterConstructor(155127968, "ClientFile")
	registry.RegisterConstructor(3995993899, "ClientFileRequest")
	registry.RegisterConstructor(2731095358, "ClientFileStatus")
	registry.RegisterConstructor(442767121, "ClientCachedMediaInfo")
	registry.RegisterConstructor(2711408875, "ClientPeerMediaInfo")
	registry.RegisterConstructor(1541024203, "ClientMediaSize")
	registry.RegisterConstructor(2069517672, "ClientRecentSearch")
	registry.RegisterConstructor(3236847495, "ClientRecentSearchMany")
	registry.RegisterConstructor(769069696, "ClientTeamCounters")
	registry.RegisterConstructor(2316768708, "ClientGetFrequentlyReactions")
	registry.RegisterConstructor(2830270594, "ClientFrequentlyReactions")
	registry.RegisterConstructor(1698398006, "ClientDismissNotification")
	registry.RegisterConstructor(4106535811, "ClientGetNotificationDismissTime")
	registry.RegisterConstructor(3077814065, "ClientNotificationDismissTime")
}

func (x *ClientSendMessageMedia) DeepCopy(z *ClientSendMessageMedia) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MediaType = x.MediaType
	z.Caption = x.Caption
	z.FileName = x.FileName
	z.FilePath = x.FilePath
	z.ThumbFilePath = x.ThumbFilePath
	z.FileMIME = x.FileMIME
	z.ThumbMIME = x.ThumbMIME
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
	z.FileUploadID = x.FileUploadID
	z.ThumbUploadID = x.ThumbUploadID
	z.FileID = x.FileID
	z.ThumbID = x.ThumbID
	z.FileTotalParts = x.FileTotalParts
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.TinyThumb = append(z.TinyThumb[:0], x.TinyThumb...)
}

func (x *ClientGlobalSearch) DeepCopy(z *ClientGlobalSearch) {
	z.Text = x.Text
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Limit = x.Limit
	z.SenderID = x.SenderID
}

func (x *ClientContactSearch) DeepCopy(z *ClientContactSearch) {
	z.Text = x.Text
}

func (x *ClientGetCachedMedia) DeepCopy(z *ClientGetCachedMedia) {
}

func (x *ClientClearCachedMedia) DeepCopy(z *ClientClearCachedMedia) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.MediaTypes = append(z.MediaTypes[:0], x.MediaTypes...)
}

func (x *ClientGetLastBotKeyboard) DeepCopy(z *ClientGetLastBotKeyboard) {
	if x.peer != nil {
		z.peer = PoolInputPeer.Get()
		x.peer.DeepCopy(z.peer)
	}
}

func (x *ClientGetMediaHistory) DeepCopy(z *ClientGetMediaHistory) {
	z.MediaType = x.MediaType
}

func (x *ClientGetRecentSearch) DeepCopy(z *ClientGetRecentSearch) {
	z.Limit = x.Limit
}

func (x *ClientPutRecentSearch) DeepCopy(z *ClientPutRecentSearch) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *ClientRemoveRecentSearch) DeepCopy(z *ClientRemoveRecentSearch) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *ClientRemoveAllRecentSearches) DeepCopy(z *ClientRemoveAllRecentSearches) {
	z.Extra = x.Extra
}

func (x *ClientGetSavedGifs) DeepCopy(z *ClientGetSavedGifs) {
}

func (x *ClientGetTeamCounters) DeepCopy(z *ClientGetTeamCounters) {
	if x.Team != nil {
		z.Team = PoolInputTeam.Get()
		x.Team.DeepCopy(z.Team)
	}
	z.WithMutes = x.WithMutes
}

func (x *ClientPendingMessage) DeepCopy(z *ClientPendingMessage) {
	z.ID = x.ID
	z.RequestID = x.RequestID
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.AccessHash = x.AccessHash
	z.CreatedOn = x.CreatedOn
	z.ReplyTo = x.ReplyTo
	z.Body = x.Body
	z.SenderID = x.SenderID
	for idx := range x.Entities {
		if x.Entities[idx] != nil {
			xx := PoolMessageEntity.Get()
			x.Entities[idx].DeepCopy(xx)
			z.Entities = append(z.Entities, xx)
		}
	}
	z.MediaType = x.MediaType
	z.Media = append(z.Media[:0], x.Media...)
	z.ClearDraft = x.ClearDraft
	z.FileUploadID = x.FileUploadID
	z.ThumbUploadID = x.ThumbUploadID
	z.FileID = x.FileID
	z.ThumbID = x.ThumbID
	z.Sha256 = append(z.Sha256[:0], x.Sha256...)
	if x.ServerFile != nil {
		z.ServerFile = PoolFileLocation.Get()
		x.ServerFile.DeepCopy(z.ServerFile)
	}
	z.TeamID = x.TeamID
	z.TeamAccessHash = x.TeamAccessHash
	z.TinyThumb = append(z.TinyThumb[:0], x.TinyThumb...)
}

func (x *ClientSearchResult) DeepCopy(z *ClientSearchResult) {
	for idx := range x.Messages {
		if x.Messages[idx] != nil {
			xx := PoolUserMessage.Get()
			x.Messages[idx].DeepCopy(xx)
			z.Messages = append(z.Messages, xx)
		}
	}
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
	for idx := range x.MatchedUsers {
		if x.MatchedUsers[idx] != nil {
			xx := PoolUser.Get()
			x.MatchedUsers[idx].DeepCopy(xx)
			z.MatchedUsers = append(z.MatchedUsers, xx)
		}
	}
	for idx := range x.MatchedGroups {
		if x.MatchedGroups[idx] != nil {
			xx := PoolGroup.Get()
			x.MatchedGroups[idx].DeepCopy(xx)
			z.MatchedGroups = append(z.MatchedGroups, xx)
		}
	}
}

func (x *ClientFilesMany) DeepCopy(z *ClientFilesMany) {
	for idx := range x.Gifs {
		if x.Gifs[idx] != nil {
			xx := PoolClientFile.Get()
			x.Gifs[idx].DeepCopy(xx)
			z.Gifs = append(z.Gifs, xx)
		}
	}
	z.Total = x.Total
}

func (x *ClientFile) DeepCopy(z *ClientFile) {
	z.ClusterID = x.ClusterID
	z.FileID = x.FileID
	z.AccessHash = x.AccessHash
	z.Type = x.Type
	z.MimeType = x.MimeType
	z.UserID = x.UserID
	z.GroupID = x.GroupID
	z.FileSize = x.FileSize
	z.MessageID = x.MessageID
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.Version = x.Version
	z.Extension = x.Extension
	z.MD5Checksum = x.MD5Checksum
	z.WallpaperID = x.WallpaperID
	for idx := range x.Attributes {
		if x.Attributes[idx] != nil {
			xx := PoolDocumentAttribute.Get()
			x.Attributes[idx].DeepCopy(xx)
			z.Attributes = append(z.Attributes, xx)
		}
	}
}

func (x *ClientFileRequest) DeepCopy(z *ClientFileRequest) {
	if x.Next != nil {
		z.Next = PoolClientFileRequest.Get()
		x.Next.DeepCopy(z.Next)
	}
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.MessageID = x.MessageID
	z.ClusterID = x.ClusterID
	z.FileID = x.FileID
	z.AccessHash = x.AccessHash
	z.Version = x.Version
	z.FileSize = x.FileSize
	z.ChunkSize = x.ChunkSize
	z.FinishedParts = append(z.FinishedParts[:0], x.FinishedParts...)
	z.TotalParts = x.TotalParts
	z.SkipDelegateCall = x.SkipDelegateCall
	z.FilePath = x.FilePath
	z.TempPath = x.TempPath
	z.CheckSha256 = x.CheckSha256
	z.FileSha256 = append(z.FileSha256[:0], x.FileSha256...)
	z.IsProfilePhoto = x.IsProfilePhoto
	z.GroupID = x.GroupID
	z.ThumbID = x.ThumbID
	z.ThumbPath = x.ThumbPath
}

func (x *ClientFileStatus) DeepCopy(z *ClientFileStatus) {
	z.Status = x.Status
	z.Progress = x.Progress
	z.FilePath = x.FilePath
}

func (x *ClientCachedMediaInfo) DeepCopy(z *ClientCachedMediaInfo) {
	for idx := range x.MediaInfo {
		if x.MediaInfo[idx] != nil {
			xx := PoolClientPeerMediaInfo.Get()
			x.MediaInfo[idx].DeepCopy(xx)
			z.MediaInfo = append(z.MediaInfo, xx)
		}
	}
}

func (x *ClientPeerMediaInfo) DeepCopy(z *ClientPeerMediaInfo) {
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	for idx := range x.Media {
		if x.Media[idx] != nil {
			xx := PoolClientMediaSize.Get()
			x.Media[idx].DeepCopy(xx)
			z.Media = append(z.Media, xx)
		}
	}
}

func (x *ClientMediaSize) DeepCopy(z *ClientMediaSize) {
	z.MediaType = x.MediaType
	z.TotalSize = x.TotalSize
}

func (x *ClientRecentSearch) DeepCopy(z *ClientRecentSearch) {
	if x.Peer != nil {
		z.Peer = PoolPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Date = x.Date
}

func (x *ClientRecentSearchMany) DeepCopy(z *ClientRecentSearchMany) {
	for idx := range x.RecentSearches {
		if x.RecentSearches[idx] != nil {
			xx := PoolClientRecentSearch.Get()
			x.RecentSearches[idx].DeepCopy(xx)
			z.RecentSearches = append(z.RecentSearches, xx)
		}
	}
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

func (x *ClientTeamCounters) DeepCopy(z *ClientTeamCounters) {
	z.UnreadCount = x.UnreadCount
	z.MentionCount = x.MentionCount
}

func (x *ClientGetFrequentlyReactions) DeepCopy(z *ClientGetFrequentlyReactions) {
}

func (x *ClientFrequentlyReactions) DeepCopy(z *ClientFrequentlyReactions) {
	z.Reactions = append(z.Reactions[:0], x.Reactions...)
}

func (x *ClientDismissNotification) DeepCopy(z *ClientDismissNotification) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Ts = x.Ts
}

func (x *ClientGetNotificationDismissTime) DeepCopy(z *ClientGetNotificationDismissTime) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *ClientNotificationDismissTime) DeepCopy(z *ClientNotificationDismissTime) {
	z.Ts = x.Ts
}
