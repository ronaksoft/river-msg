package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_AccountSetNotifySettings int64 = 2635330403

type poolAccountSetNotifySettings struct {
	pool sync.Pool
}

func (p *poolAccountSetNotifySettings) Get() *AccountSetNotifySettings {
	x, ok := p.pool.Get().(*AccountSetNotifySettings)
	if !ok {
		return &AccountSetNotifySettings{}
	}
	return x
}

func (p *poolAccountSetNotifySettings) Put(x *AccountSetNotifySettings) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	if x.Settings != nil {
		PoolPeerNotifySettings.Put(x.Settings)
		x.Settings = nil
	}
	p.pool.Put(x)
}

var PoolAccountSetNotifySettings = poolAccountSetNotifySettings{}

const C_AccountGetNotifySettings int64 = 4182396369

type poolAccountGetNotifySettings struct {
	pool sync.Pool
}

func (p *poolAccountGetNotifySettings) Get() *AccountGetNotifySettings {
	x, ok := p.pool.Get().(*AccountGetNotifySettings)
	if !ok {
		return &AccountGetNotifySettings{}
	}
	return x
}

func (p *poolAccountGetNotifySettings) Put(x *AccountGetNotifySettings) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolAccountGetNotifySettings = poolAccountGetNotifySettings{}

const C_AccountRegisterDevice int64 = 609377469

type poolAccountRegisterDevice struct {
	pool sync.Pool
}

func (p *poolAccountRegisterDevice) Get() *AccountRegisterDevice {
	x, ok := p.pool.Get().(*AccountRegisterDevice)
	if !ok {
		return &AccountRegisterDevice{}
	}
	return x
}

func (p *poolAccountRegisterDevice) Put(x *AccountRegisterDevice) {
	x.Token = ""
	x.DeviceModel = ""
	x.SystemVersion = ""
	x.AppVersion = ""
	x.LangCode = ""
	x.TokenType = 0
	x.ClientID = ""
	p.pool.Put(x)
}

var PoolAccountRegisterDevice = poolAccountRegisterDevice{}

const C_AccountUnregisterDevice int64 = 472768969

type poolAccountUnregisterDevice struct {
	pool sync.Pool
}

func (p *poolAccountUnregisterDevice) Get() *AccountUnregisterDevice {
	x, ok := p.pool.Get().(*AccountUnregisterDevice)
	if !ok {
		return &AccountUnregisterDevice{}
	}
	return x
}

func (p *poolAccountUnregisterDevice) Put(x *AccountUnregisterDevice) {
	x.TokenType = 0
	x.Token = ""
	p.pool.Put(x)
}

var PoolAccountUnregisterDevice = poolAccountUnregisterDevice{}

const C_AccountUpdateProfile int64 = 2206794686

type poolAccountUpdateProfile struct {
	pool sync.Pool
}

func (p *poolAccountUpdateProfile) Get() *AccountUpdateProfile {
	x, ok := p.pool.Get().(*AccountUpdateProfile)
	if !ok {
		return &AccountUpdateProfile{}
	}
	return x
}

func (p *poolAccountUpdateProfile) Put(x *AccountUpdateProfile) {
	x.FirstName = ""
	x.LastName = ""
	x.Bio = ""
	p.pool.Put(x)
}

var PoolAccountUpdateProfile = poolAccountUpdateProfile{}

const C_AccountCheckUsername int64 = 83507868

type poolAccountCheckUsername struct {
	pool sync.Pool
}

func (p *poolAccountCheckUsername) Get() *AccountCheckUsername {
	x, ok := p.pool.Get().(*AccountCheckUsername)
	if !ok {
		return &AccountCheckUsername{}
	}
	return x
}

func (p *poolAccountCheckUsername) Put(x *AccountCheckUsername) {
	x.Username = ""
	p.pool.Put(x)
}

var PoolAccountCheckUsername = poolAccountCheckUsername{}

const C_AccountUpdateUsername int64 = 1144675268

type poolAccountUpdateUsername struct {
	pool sync.Pool
}

func (p *poolAccountUpdateUsername) Get() *AccountUpdateUsername {
	x, ok := p.pool.Get().(*AccountUpdateUsername)
	if !ok {
		return &AccountUpdateUsername{}
	}
	return x
}

func (p *poolAccountUpdateUsername) Put(x *AccountUpdateUsername) {
	x.Username = ""
	p.pool.Put(x)
}

var PoolAccountUpdateUsername = poolAccountUpdateUsername{}

const C_AccountUploadPhoto int64 = 413055238

type poolAccountUploadPhoto struct {
	pool sync.Pool
}

func (p *poolAccountUploadPhoto) Get() *AccountUploadPhoto {
	x, ok := p.pool.Get().(*AccountUploadPhoto)
	if !ok {
		return &AccountUploadPhoto{}
	}
	return x
}

func (p *poolAccountUploadPhoto) Put(x *AccountUploadPhoto) {
	if x.File != nil {
		PoolInputFile.Put(x.File)
		x.File = nil
	}
	x.ReturnObject = false
	p.pool.Put(x)
}

var PoolAccountUploadPhoto = poolAccountUploadPhoto{}

const C_AccountUpdatePhoto int64 = 1215719904

type poolAccountUpdatePhoto struct {
	pool sync.Pool
}

func (p *poolAccountUpdatePhoto) Get() *AccountUpdatePhoto {
	x, ok := p.pool.Get().(*AccountUpdatePhoto)
	if !ok {
		return &AccountUpdatePhoto{}
	}
	return x
}

func (p *poolAccountUpdatePhoto) Put(x *AccountUpdatePhoto) {
	x.PhotoID = 0
	p.pool.Put(x)
}

var PoolAccountUpdatePhoto = poolAccountUpdatePhoto{}

const C_AccountRemovePhoto int64 = 2390524815

type poolAccountRemovePhoto struct {
	pool sync.Pool
}

func (p *poolAccountRemovePhoto) Get() *AccountRemovePhoto {
	x, ok := p.pool.Get().(*AccountRemovePhoto)
	if !ok {
		return &AccountRemovePhoto{}
	}
	return x
}

func (p *poolAccountRemovePhoto) Put(x *AccountRemovePhoto) {
	x.PhotoID = 0
	p.pool.Put(x)
}

var PoolAccountRemovePhoto = poolAccountRemovePhoto{}

const C_AccountSendChangePhoneCode int64 = 4022476844

type poolAccountSendChangePhoneCode struct {
	pool sync.Pool
}

func (p *poolAccountSendChangePhoneCode) Get() *AccountSendChangePhoneCode {
	x, ok := p.pool.Get().(*AccountSendChangePhoneCode)
	if !ok {
		return &AccountSendChangePhoneCode{}
	}
	return x
}

func (p *poolAccountSendChangePhoneCode) Put(x *AccountSendChangePhoneCode) {
	x.Phone = ""
	x.AppHash = ""
	p.pool.Put(x)
}

var PoolAccountSendChangePhoneCode = poolAccountSendChangePhoneCode{}

const C_AccountResendChangePhoneCode int64 = 523206354

type poolAccountResendChangePhoneCode struct {
	pool sync.Pool
}

func (p *poolAccountResendChangePhoneCode) Get() *AccountResendChangePhoneCode {
	x, ok := p.pool.Get().(*AccountResendChangePhoneCode)
	if !ok {
		return &AccountResendChangePhoneCode{}
	}
	return x
}

func (p *poolAccountResendChangePhoneCode) Put(x *AccountResendChangePhoneCode) {
	x.Phone = ""
	x.PhoneCodeHash = ""
	p.pool.Put(x)
}

var PoolAccountResendChangePhoneCode = poolAccountResendChangePhoneCode{}

const C_AccountChangePhone int64 = 2939511809

type poolAccountChangePhone struct {
	pool sync.Pool
}

func (p *poolAccountChangePhone) Get() *AccountChangePhone {
	x, ok := p.pool.Get().(*AccountChangePhone)
	if !ok {
		return &AccountChangePhone{}
	}
	return x
}

func (p *poolAccountChangePhone) Put(x *AccountChangePhone) {
	x.Phone = ""
	x.PhoneCodeHash = ""
	x.PhoneCode = ""
	if x.Password != nil {
		PoolInputPassword.Put(x.Password)
		x.Password = nil
	}
	p.pool.Put(x)
}

var PoolAccountChangePhone = poolAccountChangePhone{}

const C_AccountSetPrivacy int64 = 435602842

type poolAccountSetPrivacy struct {
	pool sync.Pool
}

func (p *poolAccountSetPrivacy) Get() *AccountSetPrivacy {
	x, ok := p.pool.Get().(*AccountSetPrivacy)
	if !ok {
		return &AccountSetPrivacy{}
	}
	return x
}

func (p *poolAccountSetPrivacy) Put(x *AccountSetPrivacy) {
	x.ChatInvite = x.ChatInvite[:0]
	x.LastSeen = x.LastSeen[:0]
	x.PhoneNumber = x.PhoneNumber[:0]
	x.ProfilePhoto = x.ProfilePhoto[:0]
	x.ForwardedMessage = x.ForwardedMessage[:0]
	x.Call = x.Call[:0]
	p.pool.Put(x)
}

var PoolAccountSetPrivacy = poolAccountSetPrivacy{}

const C_AccountGetPrivacy int64 = 934536200

type poolAccountGetPrivacy struct {
	pool sync.Pool
}

func (p *poolAccountGetPrivacy) Get() *AccountGetPrivacy {
	x, ok := p.pool.Get().(*AccountGetPrivacy)
	if !ok {
		return &AccountGetPrivacy{}
	}
	return x
}

func (p *poolAccountGetPrivacy) Put(x *AccountGetPrivacy) {
	x.Key = 0
	p.pool.Put(x)
}

var PoolAccountGetPrivacy = poolAccountGetPrivacy{}

const C_AccountGetAuthorizations int64 = 2563289800

type poolAccountGetAuthorizations struct {
	pool sync.Pool
}

func (p *poolAccountGetAuthorizations) Get() *AccountGetAuthorizations {
	x, ok := p.pool.Get().(*AccountGetAuthorizations)
	if !ok {
		return &AccountGetAuthorizations{}
	}
	return x
}

func (p *poolAccountGetAuthorizations) Put(x *AccountGetAuthorizations) {
	p.pool.Put(x)
}

var PoolAccountGetAuthorizations = poolAccountGetAuthorizations{}

const C_AccountResetAuthorization int64 = 2378769356

type poolAccountResetAuthorization struct {
	pool sync.Pool
}

func (p *poolAccountResetAuthorization) Get() *AccountResetAuthorization {
	x, ok := p.pool.Get().(*AccountResetAuthorization)
	if !ok {
		return &AccountResetAuthorization{}
	}
	return x
}

func (p *poolAccountResetAuthorization) Put(x *AccountResetAuthorization) {
	x.AuthID = 0
	p.pool.Put(x)
}

var PoolAccountResetAuthorization = poolAccountResetAuthorization{}

const C_AccountUpdateStatus int64 = 3359301343

type poolAccountUpdateStatus struct {
	pool sync.Pool
}

func (p *poolAccountUpdateStatus) Get() *AccountUpdateStatus {
	x, ok := p.pool.Get().(*AccountUpdateStatus)
	if !ok {
		return &AccountUpdateStatus{}
	}
	return x
}

func (p *poolAccountUpdateStatus) Put(x *AccountUpdateStatus) {
	x.Online = false
	p.pool.Put(x)
}

var PoolAccountUpdateStatus = poolAccountUpdateStatus{}

const C_AccountSetLang int64 = 1288054277

type poolAccountSetLang struct {
	pool sync.Pool
}

func (p *poolAccountSetLang) Get() *AccountSetLang {
	x, ok := p.pool.Get().(*AccountSetLang)
	if !ok {
		return &AccountSetLang{}
	}
	return x
}

func (p *poolAccountSetLang) Put(x *AccountSetLang) {
	x.LangCode = ""
	p.pool.Put(x)
}

var PoolAccountSetLang = poolAccountSetLang{}

const C_AccountGetPassword int64 = 892760360

type poolAccountGetPassword struct {
	pool sync.Pool
}

func (p *poolAccountGetPassword) Get() *AccountGetPassword {
	x, ok := p.pool.Get().(*AccountGetPassword)
	if !ok {
		return &AccountGetPassword{}
	}
	return x
}

func (p *poolAccountGetPassword) Put(x *AccountGetPassword) {
	p.pool.Put(x)
}

var PoolAccountGetPassword = poolAccountGetPassword{}

const C_AccountGetPasswordSettings int64 = 3344806825

type poolAccountGetPasswordSettings struct {
	pool sync.Pool
}

func (p *poolAccountGetPasswordSettings) Get() *AccountGetPasswordSettings {
	x, ok := p.pool.Get().(*AccountGetPasswordSettings)
	if !ok {
		return &AccountGetPasswordSettings{}
	}
	return x
}

func (p *poolAccountGetPasswordSettings) Put(x *AccountGetPasswordSettings) {
	if x.Password != nil {
		PoolInputPassword.Put(x.Password)
		x.Password = nil
	}
	p.pool.Put(x)
}

var PoolAccountGetPasswordSettings = poolAccountGetPasswordSettings{}

const C_AccountUpdatePasswordSettings int64 = 484279179

type poolAccountUpdatePasswordSettings struct {
	pool sync.Pool
}

func (p *poolAccountUpdatePasswordSettings) Get() *AccountUpdatePasswordSettings {
	x, ok := p.pool.Get().(*AccountUpdatePasswordSettings)
	if !ok {
		return &AccountUpdatePasswordSettings{}
	}
	return x
}

func (p *poolAccountUpdatePasswordSettings) Put(x *AccountUpdatePasswordSettings) {
	if x.Password != nil {
		PoolInputPassword.Put(x.Password)
		x.Password = nil
	}
	x.PasswordHash = x.PasswordHash[:0]
	x.Algorithm = 0
	x.AlgorithmData = x.AlgorithmData[:0]
	x.Hint = ""
	x.Questions = x.Questions[:0]
	p.pool.Put(x)
}

var PoolAccountUpdatePasswordSettings = poolAccountUpdatePasswordSettings{}

const C_AccountRecoverPassword int64 = 4107350339

type poolAccountRecoverPassword struct {
	pool sync.Pool
}

func (p *poolAccountRecoverPassword) Get() *AccountRecoverPassword {
	x, ok := p.pool.Get().(*AccountRecoverPassword)
	if !ok {
		return &AccountRecoverPassword{}
	}
	return x
}

func (p *poolAccountRecoverPassword) Put(x *AccountRecoverPassword) {
	x.Answers = x.Answers[:0]
	x.Algorithm = 0
	x.AlgorithmData = x.AlgorithmData[:0]
	x.SrpID = 0
	p.pool.Put(x)
}

var PoolAccountRecoverPassword = poolAccountRecoverPassword{}

const C_AccountGetTeams int64 = 3180865232

type poolAccountGetTeams struct {
	pool sync.Pool
}

func (p *poolAccountGetTeams) Get() *AccountGetTeams {
	x, ok := p.pool.Get().(*AccountGetTeams)
	if !ok {
		return &AccountGetTeams{}
	}
	return x
}

func (p *poolAccountGetTeams) Put(x *AccountGetTeams) {
	p.pool.Put(x)
}

var PoolAccountGetTeams = poolAccountGetTeams{}

const C_AccountPasswordSettings int64 = 957483519

type poolAccountPasswordSettings struct {
	pool sync.Pool
}

func (p *poolAccountPasswordSettings) Get() *AccountPasswordSettings {
	x, ok := p.pool.Get().(*AccountPasswordSettings)
	if !ok {
		return &AccountPasswordSettings{}
	}
	return x
}

func (p *poolAccountPasswordSettings) Put(x *AccountPasswordSettings) {
	x.Hint = ""
	x.Questions = x.Questions[:0]
	p.pool.Put(x)
}

var PoolAccountPasswordSettings = poolAccountPasswordSettings{}

const C_SecurityQuestions int64 = 763648334

type poolSecurityQuestions struct {
	pool sync.Pool
}

func (p *poolSecurityQuestions) Get() *SecurityQuestions {
	x, ok := p.pool.Get().(*SecurityQuestions)
	if !ok {
		return &SecurityQuestions{}
	}
	return x
}

func (p *poolSecurityQuestions) Put(x *SecurityQuestions) {
	x.Questions = x.Questions[:0]
	p.pool.Put(x)
}

var PoolSecurityQuestions = poolSecurityQuestions{}

const C_RecoveryQuestion int64 = 2918071096

type poolRecoveryQuestion struct {
	pool sync.Pool
}

func (p *poolRecoveryQuestion) Get() *RecoveryQuestion {
	x, ok := p.pool.Get().(*RecoveryQuestion)
	if !ok {
		return &RecoveryQuestion{}
	}
	return x
}

func (p *poolRecoveryQuestion) Put(x *RecoveryQuestion) {
	x.ID = 0
	x.Text = ""
	p.pool.Put(x)
}

var PoolRecoveryQuestion = poolRecoveryQuestion{}

const C_SecurityQuestion int64 = 2312942506

type poolSecurityQuestion struct {
	pool sync.Pool
}

func (p *poolSecurityQuestion) Get() *SecurityQuestion {
	x, ok := p.pool.Get().(*SecurityQuestion)
	if !ok {
		return &SecurityQuestion{}
	}
	return x
}

func (p *poolSecurityQuestion) Put(x *SecurityQuestion) {
	x.ID = 0
	x.Text = ""
	x.Answer = ""
	p.pool.Put(x)
}

var PoolSecurityQuestion = poolSecurityQuestion{}

const C_SecurityAnswer int64 = 1114722082

type poolSecurityAnswer struct {
	pool sync.Pool
}

func (p *poolSecurityAnswer) Get() *SecurityAnswer {
	x, ok := p.pool.Get().(*SecurityAnswer)
	if !ok {
		return &SecurityAnswer{}
	}
	return x
}

func (p *poolSecurityAnswer) Put(x *SecurityAnswer) {
	x.QuestionID = 0
	x.Answer = ""
	p.pool.Put(x)
}

var PoolSecurityAnswer = poolSecurityAnswer{}

const C_AccountPassword int64 = 4014653466

type poolAccountPassword struct {
	pool sync.Pool
}

func (p *poolAccountPassword) Get() *AccountPassword {
	x, ok := p.pool.Get().(*AccountPassword)
	if !ok {
		return &AccountPassword{}
	}
	return x
}

func (p *poolAccountPassword) Put(x *AccountPassword) {
	x.HasPassword = false
	x.Hint = ""
	x.Algorithm = 0
	x.AlgorithmData = x.AlgorithmData[:0]
	x.SrpB = x.SrpB[:0]
	x.RandomData = x.RandomData[:0]
	x.SrpID = 0
	x.Questions = x.Questions[:0]
	p.pool.Put(x)
}

var PoolAccountPassword = poolAccountPassword{}

const C_AccountAuthorizations int64 = 1563073032

type poolAccountAuthorizations struct {
	pool sync.Pool
}

func (p *poolAccountAuthorizations) Get() *AccountAuthorizations {
	x, ok := p.pool.Get().(*AccountAuthorizations)
	if !ok {
		return &AccountAuthorizations{}
	}
	return x
}

func (p *poolAccountAuthorizations) Put(x *AccountAuthorizations) {
	x.Authorizations = x.Authorizations[:0]
	p.pool.Put(x)
}

var PoolAccountAuthorizations = poolAccountAuthorizations{}

const C_AccountAuthorization int64 = 1307277999

type poolAccountAuthorization struct {
	pool sync.Pool
}

func (p *poolAccountAuthorization) Get() *AccountAuthorization {
	x, ok := p.pool.Get().(*AccountAuthorization)
	if !ok {
		return &AccountAuthorization{}
	}
	return x
}

func (p *poolAccountAuthorization) Put(x *AccountAuthorization) {
	x.AuthID = 0
	x.Model = ""
	x.AppVersion = ""
	x.SystemVersion = ""
	x.LangCode = ""
	x.CreatedAt = 0
	x.ActiveAt = 0
	x.ClientIP = ""
	x.LastAccess = 0
	p.pool.Put(x)
}

var PoolAccountAuthorization = poolAccountAuthorization{}

const C_AccountPrivacyRules int64 = 219898582

type poolAccountPrivacyRules struct {
	pool sync.Pool
}

func (p *poolAccountPrivacyRules) Get() *AccountPrivacyRules {
	x, ok := p.pool.Get().(*AccountPrivacyRules)
	if !ok {
		return &AccountPrivacyRules{}
	}
	return x
}

func (p *poolAccountPrivacyRules) Put(x *AccountPrivacyRules) {
	x.Rules = x.Rules[:0]
	p.pool.Put(x)
}

var PoolAccountPrivacyRules = poolAccountPrivacyRules{}

func init() {
	registry.RegisterConstructor(2635330403, "msg.AccountSetNotifySettings")
	registry.RegisterConstructor(4182396369, "msg.AccountGetNotifySettings")
	registry.RegisterConstructor(609377469, "msg.AccountRegisterDevice")
	registry.RegisterConstructor(472768969, "msg.AccountUnregisterDevice")
	registry.RegisterConstructor(2206794686, "msg.AccountUpdateProfile")
	registry.RegisterConstructor(83507868, "msg.AccountCheckUsername")
	registry.RegisterConstructor(1144675268, "msg.AccountUpdateUsername")
	registry.RegisterConstructor(413055238, "msg.AccountUploadPhoto")
	registry.RegisterConstructor(1215719904, "msg.AccountUpdatePhoto")
	registry.RegisterConstructor(2390524815, "msg.AccountRemovePhoto")
	registry.RegisterConstructor(4022476844, "msg.AccountSendChangePhoneCode")
	registry.RegisterConstructor(523206354, "msg.AccountResendChangePhoneCode")
	registry.RegisterConstructor(2939511809, "msg.AccountChangePhone")
	registry.RegisterConstructor(435602842, "msg.AccountSetPrivacy")
	registry.RegisterConstructor(934536200, "msg.AccountGetPrivacy")
	registry.RegisterConstructor(2563289800, "msg.AccountGetAuthorizations")
	registry.RegisterConstructor(2378769356, "msg.AccountResetAuthorization")
	registry.RegisterConstructor(3359301343, "msg.AccountUpdateStatus")
	registry.RegisterConstructor(1288054277, "msg.AccountSetLang")
	registry.RegisterConstructor(892760360, "msg.AccountGetPassword")
	registry.RegisterConstructor(3344806825, "msg.AccountGetPasswordSettings")
	registry.RegisterConstructor(484279179, "msg.AccountUpdatePasswordSettings")
	registry.RegisterConstructor(4107350339, "msg.AccountRecoverPassword")
	registry.RegisterConstructor(3180865232, "msg.AccountGetTeams")
	registry.RegisterConstructor(957483519, "msg.AccountPasswordSettings")
	registry.RegisterConstructor(763648334, "msg.SecurityQuestions")
	registry.RegisterConstructor(2918071096, "msg.RecoveryQuestion")
	registry.RegisterConstructor(2312942506, "msg.SecurityQuestion")
	registry.RegisterConstructor(1114722082, "msg.SecurityAnswer")
	registry.RegisterConstructor(4014653466, "msg.AccountPassword")
	registry.RegisterConstructor(1563073032, "msg.AccountAuthorizations")
	registry.RegisterConstructor(1307277999, "msg.AccountAuthorization")
	registry.RegisterConstructor(219898582, "msg.AccountPrivacyRules")
}

func (x *AccountSetNotifySettings) DeepCopy(z *AccountSetNotifySettings) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	if x.Settings != nil {
		z.Settings = PoolPeerNotifySettings.Get()
		x.Settings.DeepCopy(z.Settings)
	}
}

func (x *AccountGetNotifySettings) DeepCopy(z *AccountGetNotifySettings) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *AccountRegisterDevice) DeepCopy(z *AccountRegisterDevice) {
	z.Token = x.Token
	z.DeviceModel = x.DeviceModel
	z.SystemVersion = x.SystemVersion
	z.AppVersion = x.AppVersion
	z.LangCode = x.LangCode
	z.TokenType = x.TokenType
	z.ClientID = x.ClientID
}

func (x *AccountUnregisterDevice) DeepCopy(z *AccountUnregisterDevice) {
	z.TokenType = x.TokenType
	z.Token = x.Token
}

func (x *AccountUpdateProfile) DeepCopy(z *AccountUpdateProfile) {
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Bio = x.Bio
}

func (x *AccountCheckUsername) DeepCopy(z *AccountCheckUsername) {
	z.Username = x.Username
}

func (x *AccountUpdateUsername) DeepCopy(z *AccountUpdateUsername) {
	z.Username = x.Username
}

func (x *AccountUploadPhoto) DeepCopy(z *AccountUploadPhoto) {
	if x.File != nil {
		z.File = PoolInputFile.Get()
		x.File.DeepCopy(z.File)
	}
	z.ReturnObject = x.ReturnObject
}

func (x *AccountUpdatePhoto) DeepCopy(z *AccountUpdatePhoto) {
	z.PhotoID = x.PhotoID
}

func (x *AccountRemovePhoto) DeepCopy(z *AccountRemovePhoto) {
	z.PhotoID = x.PhotoID
}

func (x *AccountSendChangePhoneCode) DeepCopy(z *AccountSendChangePhoneCode) {
	z.Phone = x.Phone
	z.AppHash = x.AppHash
}

func (x *AccountResendChangePhoneCode) DeepCopy(z *AccountResendChangePhoneCode) {
	z.Phone = x.Phone
	z.PhoneCodeHash = x.PhoneCodeHash
}

func (x *AccountChangePhone) DeepCopy(z *AccountChangePhone) {
	z.Phone = x.Phone
	z.PhoneCodeHash = x.PhoneCodeHash
	z.PhoneCode = x.PhoneCode
	if x.Password != nil {
		z.Password = PoolInputPassword.Get()
		x.Password.DeepCopy(z.Password)
	}
}

func (x *AccountSetPrivacy) DeepCopy(z *AccountSetPrivacy) {
	for idx := range x.ChatInvite {
		if x.ChatInvite[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.ChatInvite[idx].DeepCopy(xx)
			z.ChatInvite = append(z.ChatInvite, xx)
		}
	}
	for idx := range x.LastSeen {
		if x.LastSeen[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.LastSeen[idx].DeepCopy(xx)
			z.LastSeen = append(z.LastSeen, xx)
		}
	}
	for idx := range x.PhoneNumber {
		if x.PhoneNumber[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.PhoneNumber[idx].DeepCopy(xx)
			z.PhoneNumber = append(z.PhoneNumber, xx)
		}
	}
	for idx := range x.ProfilePhoto {
		if x.ProfilePhoto[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.ProfilePhoto[idx].DeepCopy(xx)
			z.ProfilePhoto = append(z.ProfilePhoto, xx)
		}
	}
	for idx := range x.ForwardedMessage {
		if x.ForwardedMessage[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.ForwardedMessage[idx].DeepCopy(xx)
			z.ForwardedMessage = append(z.ForwardedMessage, xx)
		}
	}
	for idx := range x.Call {
		if x.Call[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.Call[idx].DeepCopy(xx)
			z.Call = append(z.Call, xx)
		}
	}
}

func (x *AccountGetPrivacy) DeepCopy(z *AccountGetPrivacy) {
	z.Key = x.Key
}

func (x *AccountGetAuthorizations) DeepCopy(z *AccountGetAuthorizations) {
}

func (x *AccountResetAuthorization) DeepCopy(z *AccountResetAuthorization) {
	z.AuthID = x.AuthID
}

func (x *AccountUpdateStatus) DeepCopy(z *AccountUpdateStatus) {
	z.Online = x.Online
}

func (x *AccountSetLang) DeepCopy(z *AccountSetLang) {
	z.LangCode = x.LangCode
}

func (x *AccountGetPassword) DeepCopy(z *AccountGetPassword) {
}

func (x *AccountGetPasswordSettings) DeepCopy(z *AccountGetPasswordSettings) {
	if x.Password != nil {
		z.Password = PoolInputPassword.Get()
		x.Password.DeepCopy(z.Password)
	}
}

func (x *AccountUpdatePasswordSettings) DeepCopy(z *AccountUpdatePasswordSettings) {
	if x.Password != nil {
		z.Password = PoolInputPassword.Get()
		x.Password.DeepCopy(z.Password)
	}
	z.PasswordHash = append(z.PasswordHash[:0], x.PasswordHash...)
	z.Algorithm = x.Algorithm
	z.AlgorithmData = append(z.AlgorithmData[:0], x.AlgorithmData...)
	z.Hint = x.Hint
	for idx := range x.Questions {
		if x.Questions[idx] != nil {
			xx := PoolSecurityQuestion.Get()
			x.Questions[idx].DeepCopy(xx)
			z.Questions = append(z.Questions, xx)
		}
	}
}

func (x *AccountRecoverPassword) DeepCopy(z *AccountRecoverPassword) {
	for idx := range x.Answers {
		if x.Answers[idx] != nil {
			xx := PoolSecurityAnswer.Get()
			x.Answers[idx].DeepCopy(xx)
			z.Answers = append(z.Answers, xx)
		}
	}
	z.Algorithm = x.Algorithm
	z.AlgorithmData = append(z.AlgorithmData[:0], x.AlgorithmData...)
	z.SrpID = x.SrpID
}

func (x *AccountGetTeams) DeepCopy(z *AccountGetTeams) {
}

func (x *AccountPasswordSettings) DeepCopy(z *AccountPasswordSettings) {
	z.Hint = x.Hint
	for idx := range x.Questions {
		if x.Questions[idx] != nil {
			xx := PoolRecoveryQuestion.Get()
			x.Questions[idx].DeepCopy(xx)
			z.Questions = append(z.Questions, xx)
		}
	}
}

func (x *SecurityQuestions) DeepCopy(z *SecurityQuestions) {
	for idx := range x.Questions {
		if x.Questions[idx] != nil {
			xx := PoolSecurityQuestion.Get()
			x.Questions[idx].DeepCopy(xx)
			z.Questions = append(z.Questions, xx)
		}
	}
}

func (x *RecoveryQuestion) DeepCopy(z *RecoveryQuestion) {
	z.ID = x.ID
	z.Text = x.Text
}

func (x *SecurityQuestion) DeepCopy(z *SecurityQuestion) {
	z.ID = x.ID
	z.Text = x.Text
	z.Answer = x.Answer
}

func (x *SecurityAnswer) DeepCopy(z *SecurityAnswer) {
	z.QuestionID = x.QuestionID
	z.Answer = x.Answer
}

func (x *AccountPassword) DeepCopy(z *AccountPassword) {
	z.HasPassword = x.HasPassword
	z.Hint = x.Hint
	z.Algorithm = x.Algorithm
	z.AlgorithmData = append(z.AlgorithmData[:0], x.AlgorithmData...)
	z.SrpB = append(z.SrpB[:0], x.SrpB...)
	z.RandomData = append(z.RandomData[:0], x.RandomData...)
	z.SrpID = x.SrpID
	for idx := range x.Questions {
		if x.Questions[idx] != nil {
			xx := PoolRecoveryQuestion.Get()
			x.Questions[idx].DeepCopy(xx)
			z.Questions = append(z.Questions, xx)
		}
	}
}

func (x *AccountAuthorizations) DeepCopy(z *AccountAuthorizations) {
	for idx := range x.Authorizations {
		if x.Authorizations[idx] != nil {
			xx := PoolAccountAuthorization.Get()
			x.Authorizations[idx].DeepCopy(xx)
			z.Authorizations = append(z.Authorizations, xx)
		}
	}
}

func (x *AccountAuthorization) DeepCopy(z *AccountAuthorization) {
	z.AuthID = x.AuthID
	z.Model = x.Model
	z.AppVersion = x.AppVersion
	z.SystemVersion = x.SystemVersion
	z.LangCode = x.LangCode
	z.CreatedAt = x.CreatedAt
	z.ActiveAt = x.ActiveAt
	z.ClientIP = x.ClientIP
	z.LastAccess = x.LastAccess
}

func (x *AccountPrivacyRules) DeepCopy(z *AccountPrivacyRules) {
	for idx := range x.Rules {
		if x.Rules[idx] != nil {
			xx := PoolPrivacyRule.Get()
			x.Rules[idx].DeepCopy(xx)
			z.Rules = append(z.Rules, xx)
		}
	}
}

func (x *AccountSetNotifySettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountSetNotifySettings, x)
}

func (x *AccountGetNotifySettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetNotifySettings, x)
}

func (x *AccountRegisterDevice) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountRegisterDevice, x)
}

func (x *AccountUnregisterDevice) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUnregisterDevice, x)
}

func (x *AccountUpdateProfile) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUpdateProfile, x)
}

func (x *AccountCheckUsername) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountCheckUsername, x)
}

func (x *AccountUpdateUsername) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUpdateUsername, x)
}

func (x *AccountUploadPhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUploadPhoto, x)
}

func (x *AccountUpdatePhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUpdatePhoto, x)
}

func (x *AccountRemovePhoto) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountRemovePhoto, x)
}

func (x *AccountSendChangePhoneCode) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountSendChangePhoneCode, x)
}

func (x *AccountResendChangePhoneCode) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountResendChangePhoneCode, x)
}

func (x *AccountChangePhone) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountChangePhone, x)
}

func (x *AccountSetPrivacy) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountSetPrivacy, x)
}

func (x *AccountGetPrivacy) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetPrivacy, x)
}

func (x *AccountGetAuthorizations) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetAuthorizations, x)
}

func (x *AccountResetAuthorization) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountResetAuthorization, x)
}

func (x *AccountUpdateStatus) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUpdateStatus, x)
}

func (x *AccountSetLang) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountSetLang, x)
}

func (x *AccountGetPassword) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetPassword, x)
}

func (x *AccountGetPasswordSettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetPasswordSettings, x)
}

func (x *AccountUpdatePasswordSettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountUpdatePasswordSettings, x)
}

func (x *AccountRecoverPassword) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountRecoverPassword, x)
}

func (x *AccountGetTeams) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountGetTeams, x)
}

func (x *AccountPasswordSettings) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountPasswordSettings, x)
}

func (x *SecurityQuestions) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SecurityQuestions, x)
}

func (x *RecoveryQuestion) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_RecoveryQuestion, x)
}

func (x *SecurityQuestion) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SecurityQuestion, x)
}

func (x *SecurityAnswer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_SecurityAnswer, x)
}

func (x *AccountPassword) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountPassword, x)
}

func (x *AccountAuthorizations) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountAuthorizations, x)
}

func (x *AccountAuthorization) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountAuthorization, x)
}

func (x *AccountPrivacyRules) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_AccountPrivacyRules, x)
}

func (x *AccountSetNotifySettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetNotifySettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountRegisterDevice) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUnregisterDevice) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUpdateProfile) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountCheckUsername) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUpdateUsername) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUploadPhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUpdatePhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountRemovePhoto) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountSendChangePhoneCode) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountResendChangePhoneCode) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountChangePhone) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountSetPrivacy) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetPrivacy) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetAuthorizations) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountResetAuthorization) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUpdateStatus) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountSetLang) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetPassword) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetPasswordSettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountUpdatePasswordSettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountRecoverPassword) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountGetTeams) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountPasswordSettings) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *SecurityQuestions) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *RecoveryQuestion) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *SecurityQuestion) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *SecurityAnswer) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountPassword) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountAuthorizations) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountAuthorization) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountPrivacyRules) MarshalTo(b []byte) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(b, x)
}

func (x *AccountSetNotifySettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetNotifySettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountRegisterDevice) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUnregisterDevice) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUpdateProfile) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountCheckUsername) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUpdateUsername) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUploadPhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUpdatePhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountRemovePhoto) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountSendChangePhoneCode) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountResendChangePhoneCode) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountChangePhone) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountSetPrivacy) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetPrivacy) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetAuthorizations) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountResetAuthorization) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUpdateStatus) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountSetLang) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetPassword) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetPasswordSettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountUpdatePasswordSettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountRecoverPassword) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountGetTeams) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountPasswordSettings) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SecurityQuestions) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *RecoveryQuestion) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SecurityQuestion) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *SecurityAnswer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountPassword) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountAuthorizations) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountAuthorization) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *AccountPrivacyRules) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}
