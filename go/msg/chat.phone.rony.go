package msg

import (
	edge "github.com/ronaksoft/rony/edge"
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

const C_PhoneInitCall int64 = 2975617068

type poolPhoneInitCall struct {
	pool sync.Pool
}

func (p *poolPhoneInitCall) Get() *PhoneInitCall {
	x, ok := p.pool.Get().(*PhoneInitCall)
	if !ok {
		return &PhoneInitCall{}
	}
	return x
}

func (p *poolPhoneInitCall) Put(x *PhoneInitCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	p.pool.Put(x)
}

var PoolPhoneInitCall = poolPhoneInitCall{}

const C_PhoneRequestCall int64 = 907942641

type poolPhoneRequestCall struct {
	pool sync.Pool
}

func (p *poolPhoneRequestCall) Get() *PhoneRequestCall {
	x, ok := p.pool.Get().(*PhoneRequestCall)
	if !ok {
		return &PhoneRequestCall{}
	}
	return x
}

func (p *poolPhoneRequestCall) Put(x *PhoneRequestCall) {
	x.RandomID = 0
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.Initiator = false
	x.Participants = x.Participants[:0]
	x.CallID = 0
	p.pool.Put(x)
}

var PoolPhoneRequestCall = poolPhoneRequestCall{}

const C_PhoneAcceptCall int64 = 4133092858

type poolPhoneAcceptCall struct {
	pool sync.Pool
}

func (p *poolPhoneAcceptCall) Get() *PhoneAcceptCall {
	x, ok := p.pool.Get().(*PhoneAcceptCall)
	if !ok {
		return &PhoneAcceptCall{}
	}
	return x
}

func (p *poolPhoneAcceptCall) Put(x *PhoneAcceptCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneAcceptCall = poolPhoneAcceptCall{}

const C_PhoneDiscardCall int64 = 2712700137

type poolPhoneDiscardCall struct {
	pool sync.Pool
}

func (p *poolPhoneDiscardCall) Get() *PhoneDiscardCall {
	x, ok := p.pool.Get().(*PhoneDiscardCall)
	if !ok {
		return &PhoneDiscardCall{}
	}
	return x
}

func (p *poolPhoneDiscardCall) Put(x *PhoneDiscardCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Duration = 0
	x.Reason = 0
	p.pool.Put(x)
}

var PoolPhoneDiscardCall = poolPhoneDiscardCall{}

const C_PhoneJoinCall int64 = 3019166552

type poolPhoneJoinCall struct {
	pool sync.Pool
}

func (p *poolPhoneJoinCall) Get() *PhoneJoinCall {
	x, ok := p.pool.Get().(*PhoneJoinCall)
	if !ok {
		return &PhoneJoinCall{}
	}
	return x
}

func (p *poolPhoneJoinCall) Put(x *PhoneJoinCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	p.pool.Put(x)
}

var PoolPhoneJoinCall = poolPhoneJoinCall{}

const C_PhoneAddParticipant int64 = 2867411100

type poolPhoneAddParticipant struct {
	pool sync.Pool
}

func (p *poolPhoneAddParticipant) Get() *PhoneAddParticipant {
	x, ok := p.pool.Get().(*PhoneAddParticipant)
	if !ok {
		return &PhoneAddParticipant{}
	}
	return x
}

func (p *poolPhoneAddParticipant) Put(x *PhoneAddParticipant) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneAddParticipant = poolPhoneAddParticipant{}

const C_PhoneRemoveParticipant int64 = 188662172

type poolPhoneRemoveParticipant struct {
	pool sync.Pool
}

func (p *poolPhoneRemoveParticipant) Get() *PhoneRemoveParticipant {
	x, ok := p.pool.Get().(*PhoneRemoveParticipant)
	if !ok {
		return &PhoneRemoveParticipant{}
	}
	return x
}

func (p *poolPhoneRemoveParticipant) Put(x *PhoneRemoveParticipant) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Participants = x.Participants[:0]
	x.Timeout = false
	p.pool.Put(x)
}

var PoolPhoneRemoveParticipant = poolPhoneRemoveParticipant{}

const C_PhoneUpdateCall int64 = 1976202226

type poolPhoneUpdateCall struct {
	pool sync.Pool
}

func (p *poolPhoneUpdateCall) Get() *PhoneUpdateCall {
	x, ok := p.pool.Get().(*PhoneUpdateCall)
	if !ok {
		return &PhoneUpdateCall{}
	}
	return x
}

func (p *poolPhoneUpdateCall) Put(x *PhoneUpdateCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Participants = x.Participants[:0]
	x.Action = 0
	x.ActionData = x.ActionData[:0]
	p.pool.Put(x)
}

var PoolPhoneUpdateCall = poolPhoneUpdateCall{}

const C_PhoneRateCall int64 = 2215486159

type poolPhoneRateCall struct {
	pool sync.Pool
}

func (p *poolPhoneRateCall) Get() *PhoneRateCall {
	x, ok := p.pool.Get().(*PhoneRateCall)
	if !ok {
		return &PhoneRateCall{}
	}
	return x
}

func (p *poolPhoneRateCall) Put(x *PhoneRateCall) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	x.Rate = 0
	x.ReasonType = 0
	x.ReasonData = x.ReasonData[:0]
	p.pool.Put(x)
}

var PoolPhoneRateCall = poolPhoneRateCall{}

const C_PhoneGetHistory int64 = 407776572

type poolPhoneGetHistory struct {
	pool sync.Pool
}

func (p *poolPhoneGetHistory) Get() *PhoneGetHistory {
	x, ok := p.pool.Get().(*PhoneGetHistory)
	if !ok {
		return &PhoneGetHistory{}
	}
	return x
}

func (p *poolPhoneGetHistory) Put(x *PhoneGetHistory) {
	x.Limit = 0
	x.After = 0
	p.pool.Put(x)
}

var PoolPhoneGetHistory = poolPhoneGetHistory{}

const C_PhoneCallRecord int64 = 4147150312

type poolPhoneCallRecord struct {
	pool sync.Pool
}

func (p *poolPhoneCallRecord) Get() *PhoneCallRecord {
	x, ok := p.pool.Get().(*PhoneCallRecord)
	if !ok {
		return &PhoneCallRecord{}
	}
	return x
}

func (p *poolPhoneCallRecord) Put(x *PhoneCallRecord) {
	x.UserID = 0
	x.TeamID = 0
	x.CallID = 0
	x.CreatedOn = 0
	x.EndedOn = 0
	x.Incoming = false
	x.PeerID = 0
	x.PeerType = 0
	x.Status = 0
	p.pool.Put(x)
}

var PoolPhoneCallRecord = poolPhoneCallRecord{}

const C_PhoneCallsMany int64 = 1227520020

type poolPhoneCallsMany struct {
	pool sync.Pool
}

func (p *poolPhoneCallsMany) Get() *PhoneCallsMany {
	x, ok := p.pool.Get().(*PhoneCallsMany)
	if !ok {
		return &PhoneCallsMany{}
	}
	return x
}

func (p *poolPhoneCallsMany) Put(x *PhoneCallsMany) {
	x.PhoneCalls = x.PhoneCalls[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	x.Continuous = false
	x.Empty = false
	p.pool.Put(x)
}

var PoolPhoneCallsMany = poolPhoneCallsMany{}

const C_PhoneUpdateAdmin int64 = 442877873

type poolPhoneUpdateAdmin struct {
	pool sync.Pool
}

func (p *poolPhoneUpdateAdmin) Get() *PhoneUpdateAdmin {
	x, ok := p.pool.Get().(*PhoneUpdateAdmin)
	if !ok {
		return &PhoneUpdateAdmin{}
	}
	return x
}

func (p *poolPhoneUpdateAdmin) Put(x *PhoneUpdateAdmin) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.CallID = 0
	if x.User != nil {
		PoolInputUser.Put(x.User)
		x.User = nil
	}
	x.Admin = false
	p.pool.Put(x)
}

var PoolPhoneUpdateAdmin = poolPhoneUpdateAdmin{}

const C_PhoneCall int64 = 3296664529

type poolPhoneCall struct {
	pool sync.Pool
}

func (p *poolPhoneCall) Get() *PhoneCall {
	x, ok := p.pool.Get().(*PhoneCall)
	if !ok {
		return &PhoneCall{}
	}
	return x
}

func (p *poolPhoneCall) Put(x *PhoneCall) {
	x.ID = 0
	x.CreatedOn = 0
	p.pool.Put(x)
}

var PoolPhoneCall = poolPhoneCall{}

const C_PhoneInit int64 = 3464876187

type poolPhoneInit struct {
	pool sync.Pool
}

func (p *poolPhoneInit) Get() *PhoneInit {
	x, ok := p.pool.Get().(*PhoneInit)
	if !ok {
		return &PhoneInit{}
	}
	return x
}

func (p *poolPhoneInit) Put(x *PhoneInit) {
	x.IceServers = x.IceServers[:0]
	p.pool.Put(x)
}

var PoolPhoneInit = poolPhoneInit{}

const C_PhoneParticipants int64 = 2567653219

type poolPhoneParticipants struct {
	pool sync.Pool
}

func (p *poolPhoneParticipants) Get() *PhoneParticipants {
	x, ok := p.pool.Get().(*PhoneParticipants)
	if !ok {
		return &PhoneParticipants{}
	}
	return x
}

func (p *poolPhoneParticipants) Put(x *PhoneParticipants) {
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneParticipants = poolPhoneParticipants{}

const C_IceServer int64 = 4291892363

type poolIceServer struct {
	pool sync.Pool
}

func (p *poolIceServer) Get() *IceServer {
	x, ok := p.pool.Get().(*IceServer)
	if !ok {
		return &IceServer{}
	}
	return x
}

func (p *poolIceServer) Put(x *IceServer) {
	x.Urls = x.Urls[:0]
	x.Username = ""
	x.Credential = ""
	p.pool.Put(x)
}

var PoolIceServer = poolIceServer{}

const C_PhoneParticipant int64 = 226273622

type poolPhoneParticipant struct {
	pool sync.Pool
}

func (p *poolPhoneParticipant) Get() *PhoneParticipant {
	x, ok := p.pool.Get().(*PhoneParticipant)
	if !ok {
		return &PhoneParticipant{}
	}
	return x
}

func (p *poolPhoneParticipant) Put(x *PhoneParticipant) {
	x.ConnectionId = 0
	if x.Peer != nil {
		PoolInputUser.Put(x.Peer)
		x.Peer = nil
	}
	x.Initiator = false
	x.Admin = false
	p.pool.Put(x)
}

var PoolPhoneParticipant = poolPhoneParticipant{}

const C_PhoneParticipantSDP int64 = 545454774

type poolPhoneParticipantSDP struct {
	pool sync.Pool
}

func (p *poolPhoneParticipantSDP) Get() *PhoneParticipantSDP {
	x, ok := p.pool.Get().(*PhoneParticipantSDP)
	if !ok {
		return &PhoneParticipantSDP{}
	}
	return x
}

func (p *poolPhoneParticipantSDP) Put(x *PhoneParticipantSDP) {
	x.ConnectionId = 0
	if x.Peer != nil {
		PoolInputUser.Put(x.Peer)
		x.Peer = nil
	}
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneParticipantSDP = poolPhoneParticipantSDP{}

const C_PhoneActionCallEmpty int64 = 1073285997

type poolPhoneActionCallEmpty struct {
	pool sync.Pool
}

func (p *poolPhoneActionCallEmpty) Get() *PhoneActionCallEmpty {
	x, ok := p.pool.Get().(*PhoneActionCallEmpty)
	if !ok {
		return &PhoneActionCallEmpty{}
	}
	return x
}

func (p *poolPhoneActionCallEmpty) Put(x *PhoneActionCallEmpty) {
	x.Empty = false
	p.pool.Put(x)
}

var PoolPhoneActionCallEmpty = poolPhoneActionCallEmpty{}

const C_PhoneActionAccepted int64 = 2493210645

type poolPhoneActionAccepted struct {
	pool sync.Pool
}

func (p *poolPhoneActionAccepted) Get() *PhoneActionAccepted {
	x, ok := p.pool.Get().(*PhoneActionAccepted)
	if !ok {
		return &PhoneActionAccepted{}
	}
	return x
}

func (p *poolPhoneActionAccepted) Put(x *PhoneActionAccepted) {
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneActionAccepted = poolPhoneActionAccepted{}

const C_PhoneActionRequested int64 = 1678316869

type poolPhoneActionRequested struct {
	pool sync.Pool
}

func (p *poolPhoneActionRequested) Get() *PhoneActionRequested {
	x, ok := p.pool.Get().(*PhoneActionRequested)
	if !ok {
		return &PhoneActionRequested{}
	}
	return x
}

func (p *poolPhoneActionRequested) Put(x *PhoneActionRequested) {
	x.SDP = ""
	x.Type = ""
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneActionRequested = poolPhoneActionRequested{}

const C_PhoneActionCallWaiting int64 = 3634710697

type poolPhoneActionCallWaiting struct {
	pool sync.Pool
}

func (p *poolPhoneActionCallWaiting) Get() *PhoneActionCallWaiting {
	x, ok := p.pool.Get().(*PhoneActionCallWaiting)
	if !ok {
		return &PhoneActionCallWaiting{}
	}
	return x
}

func (p *poolPhoneActionCallWaiting) Put(x *PhoneActionCallWaiting) {
	x.Empty = false
	p.pool.Put(x)
}

var PoolPhoneActionCallWaiting = poolPhoneActionCallWaiting{}

const C_PhoneActionDiscarded int64 = 4285966731

type poolPhoneActionDiscarded struct {
	pool sync.Pool
}

func (p *poolPhoneActionDiscarded) Get() *PhoneActionDiscarded {
	x, ok := p.pool.Get().(*PhoneActionDiscarded)
	if !ok {
		return &PhoneActionDiscarded{}
	}
	return x
}

func (p *poolPhoneActionDiscarded) Put(x *PhoneActionDiscarded) {
	x.Duration = 0
	x.Video = false
	x.Reason = 0
	x.Terminate = false
	p.pool.Put(x)
}

var PoolPhoneActionDiscarded = poolPhoneActionDiscarded{}

const C_PhoneActionIceExchange int64 = 1618781621

type poolPhoneActionIceExchange struct {
	pool sync.Pool
}

func (p *poolPhoneActionIceExchange) Get() *PhoneActionIceExchange {
	x, ok := p.pool.Get().(*PhoneActionIceExchange)
	if !ok {
		return &PhoneActionIceExchange{}
	}
	return x
}

func (p *poolPhoneActionIceExchange) Put(x *PhoneActionIceExchange) {
	x.Candidate = ""
	x.SdpMLineIndex = 0
	x.SdpMid = ""
	x.UsernameFragment = ""
	p.pool.Put(x)
}

var PoolPhoneActionIceExchange = poolPhoneActionIceExchange{}

const C_PhoneActionAck int64 = 1221076803

type poolPhoneActionAck struct {
	pool sync.Pool
}

func (p *poolPhoneActionAck) Get() *PhoneActionAck {
	x, ok := p.pool.Get().(*PhoneActionAck)
	if !ok {
		return &PhoneActionAck{}
	}
	return x
}

func (p *poolPhoneActionAck) Put(x *PhoneActionAck) {
	p.pool.Put(x)
}

var PoolPhoneActionAck = poolPhoneActionAck{}

const C_PhoneActionParticipantAdded int64 = 2638615078

type poolPhoneActionParticipantAdded struct {
	pool sync.Pool
}

func (p *poolPhoneActionParticipantAdded) Get() *PhoneActionParticipantAdded {
	x, ok := p.pool.Get().(*PhoneActionParticipantAdded)
	if !ok {
		return &PhoneActionParticipantAdded{}
	}
	return x
}

func (p *poolPhoneActionParticipantAdded) Put(x *PhoneActionParticipantAdded) {
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneActionParticipantAdded = poolPhoneActionParticipantAdded{}

const C_PhoneActionParticipantRemoved int64 = 3280922507

type poolPhoneActionParticipantRemoved struct {
	pool sync.Pool
}

func (p *poolPhoneActionParticipantRemoved) Get() *PhoneActionParticipantRemoved {
	x, ok := p.pool.Get().(*PhoneActionParticipantRemoved)
	if !ok {
		return &PhoneActionParticipantRemoved{}
	}
	return x
}

func (p *poolPhoneActionParticipantRemoved) Put(x *PhoneActionParticipantRemoved) {
	x.UserIDs = x.UserIDs[:0]
	x.Timeout = false
	p.pool.Put(x)
}

var PoolPhoneActionParticipantRemoved = poolPhoneActionParticipantRemoved{}

const C_PhoneActionJoinRequested int64 = 656125601

type poolPhoneActionJoinRequested struct {
	pool sync.Pool
}

func (p *poolPhoneActionJoinRequested) Get() *PhoneActionJoinRequested {
	x, ok := p.pool.Get().(*PhoneActionJoinRequested)
	if !ok {
		return &PhoneActionJoinRequested{}
	}
	return x
}

func (p *poolPhoneActionJoinRequested) Put(x *PhoneActionJoinRequested) {
	x.UserIDs = x.UserIDs[:0]
	p.pool.Put(x)
}

var PoolPhoneActionJoinRequested = poolPhoneActionJoinRequested{}

const C_PhoneActionMediaSettingsUpdated int64 = 2310335221

type poolPhoneActionMediaSettingsUpdated struct {
	pool sync.Pool
}

func (p *poolPhoneActionMediaSettingsUpdated) Get() *PhoneActionMediaSettingsUpdated {
	x, ok := p.pool.Get().(*PhoneActionMediaSettingsUpdated)
	if !ok {
		return &PhoneActionMediaSettingsUpdated{}
	}
	return x
}

func (p *poolPhoneActionMediaSettingsUpdated) Put(x *PhoneActionMediaSettingsUpdated) {
	x.Video = false
	x.Audio = false
	p.pool.Put(x)
}

var PoolPhoneActionMediaSettingsUpdated = poolPhoneActionMediaSettingsUpdated{}

const C_PhoneActionReactionSet int64 = 2047679815

type poolPhoneActionReactionSet struct {
	pool sync.Pool
}

func (p *poolPhoneActionReactionSet) Get() *PhoneActionReactionSet {
	x, ok := p.pool.Get().(*PhoneActionReactionSet)
	if !ok {
		return &PhoneActionReactionSet{}
	}
	return x
}

func (p *poolPhoneActionReactionSet) Put(x *PhoneActionReactionSet) {
	x.Reaction = ""
	p.pool.Put(x)
}

var PoolPhoneActionReactionSet = poolPhoneActionReactionSet{}

const C_PhoneActionSDPOffer int64 = 931453435

type poolPhoneActionSDPOffer struct {
	pool sync.Pool
}

func (p *poolPhoneActionSDPOffer) Get() *PhoneActionSDPOffer {
	x, ok := p.pool.Get().(*PhoneActionSDPOffer)
	if !ok {
		return &PhoneActionSDPOffer{}
	}
	return x
}

func (p *poolPhoneActionSDPOffer) Put(x *PhoneActionSDPOffer) {
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneActionSDPOffer = poolPhoneActionSDPOffer{}

const C_PhoneActionSDPAnswer int64 = 835530308

type poolPhoneActionSDPAnswer struct {
	pool sync.Pool
}

func (p *poolPhoneActionSDPAnswer) Get() *PhoneActionSDPAnswer {
	x, ok := p.pool.Get().(*PhoneActionSDPAnswer)
	if !ok {
		return &PhoneActionSDPAnswer{}
	}
	return x
}

func (p *poolPhoneActionSDPAnswer) Put(x *PhoneActionSDPAnswer) {
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneActionSDPAnswer = poolPhoneActionSDPAnswer{}

func init() {
	registry.RegisterConstructor(2975617068, "PhoneInitCall")
	registry.RegisterConstructor(907942641, "PhoneRequestCall")
	registry.RegisterConstructor(4133092858, "PhoneAcceptCall")
	registry.RegisterConstructor(2712700137, "PhoneDiscardCall")
	registry.RegisterConstructor(3019166552, "PhoneJoinCall")
	registry.RegisterConstructor(2867411100, "PhoneAddParticipant")
	registry.RegisterConstructor(188662172, "PhoneRemoveParticipant")
	registry.RegisterConstructor(1976202226, "PhoneUpdateCall")
	registry.RegisterConstructor(2215486159, "PhoneRateCall")
	registry.RegisterConstructor(407776572, "PhoneGetHistory")
	registry.RegisterConstructor(4147150312, "PhoneCallRecord")
	registry.RegisterConstructor(1227520020, "PhoneCallsMany")
	registry.RegisterConstructor(442877873, "PhoneUpdateAdmin")
	registry.RegisterConstructor(3296664529, "PhoneCall")
	registry.RegisterConstructor(3464876187, "PhoneInit")
	registry.RegisterConstructor(2567653219, "PhoneParticipants")
	registry.RegisterConstructor(4291892363, "IceServer")
	registry.RegisterConstructor(226273622, "PhoneParticipant")
	registry.RegisterConstructor(545454774, "PhoneParticipantSDP")
	registry.RegisterConstructor(1073285997, "PhoneActionCallEmpty")
	registry.RegisterConstructor(2493210645, "PhoneActionAccepted")
	registry.RegisterConstructor(1678316869, "PhoneActionRequested")
	registry.RegisterConstructor(3634710697, "PhoneActionCallWaiting")
	registry.RegisterConstructor(4285966731, "PhoneActionDiscarded")
	registry.RegisterConstructor(1618781621, "PhoneActionIceExchange")
	registry.RegisterConstructor(1221076803, "PhoneActionAck")
	registry.RegisterConstructor(2638615078, "PhoneActionParticipantAdded")
	registry.RegisterConstructor(3280922507, "PhoneActionParticipantRemoved")
	registry.RegisterConstructor(656125601, "PhoneActionJoinRequested")
	registry.RegisterConstructor(2310335221, "PhoneActionMediaSettingsUpdated")
	registry.RegisterConstructor(2047679815, "PhoneActionReactionSet")
	registry.RegisterConstructor(931453435, "PhoneActionSDPOffer")
	registry.RegisterConstructor(835530308, "PhoneActionSDPAnswer")
}

func (x *PhoneInitCall) DeepCopy(z *PhoneInitCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
}

func (x *PhoneRequestCall) DeepCopy(z *PhoneRequestCall) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Initiator = x.Initiator
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolPhoneParticipantSDP.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
	z.CallID = x.CallID
}

func (x *PhoneAcceptCall) DeepCopy(z *PhoneAcceptCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolPhoneParticipantSDP.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
}

func (x *PhoneDiscardCall) DeepCopy(z *PhoneDiscardCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	z.Duration = x.Duration
	z.Reason = x.Reason
}

func (x *PhoneJoinCall) DeepCopy(z *PhoneJoinCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
}

func (x *PhoneAddParticipant) DeepCopy(z *PhoneAddParticipant) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolInputUser.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
}

func (x *PhoneRemoveParticipant) DeepCopy(z *PhoneRemoveParticipant) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolInputUser.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
	z.Timeout = x.Timeout
}

func (x *PhoneUpdateCall) DeepCopy(z *PhoneUpdateCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolInputUser.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
	z.Action = x.Action
	z.ActionData = append(z.ActionData[:0], x.ActionData...)
}

func (x *PhoneRateCall) DeepCopy(z *PhoneRateCall) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	z.Rate = x.Rate
	z.ReasonType = x.ReasonType
	z.ReasonData = append(z.ReasonData[:0], x.ReasonData...)
}

func (x *PhoneGetHistory) DeepCopy(z *PhoneGetHistory) {
	z.Limit = x.Limit
	z.After = x.After
}

func (x *PhoneCallRecord) DeepCopy(z *PhoneCallRecord) {
	z.UserID = x.UserID
	z.TeamID = x.TeamID
	z.CallID = x.CallID
	z.CreatedOn = x.CreatedOn
	z.EndedOn = x.EndedOn
	z.Incoming = x.Incoming
	z.PeerID = x.PeerID
	z.PeerType = x.PeerType
	z.Status = x.Status
}

func (x *PhoneCallsMany) DeepCopy(z *PhoneCallsMany) {
	for idx := range x.PhoneCalls {
		if x.PhoneCalls[idx] != nil {
			xx := PoolPhoneCallRecord.Get()
			x.PhoneCalls[idx].DeepCopy(xx)
			z.PhoneCalls = append(z.PhoneCalls, xx)
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
	z.Continuous = x.Continuous
	z.Empty = x.Empty
}

func (x *PhoneUpdateAdmin) DeepCopy(z *PhoneUpdateAdmin) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.CallID = x.CallID
	if x.User != nil {
		z.User = PoolInputUser.Get()
		x.User.DeepCopy(z.User)
	}
	z.Admin = x.Admin
}

func (x *PhoneCall) DeepCopy(z *PhoneCall) {
	z.ID = x.ID
	z.CreatedOn = x.CreatedOn
}

func (x *PhoneInit) DeepCopy(z *PhoneInit) {
	for idx := range x.IceServers {
		if x.IceServers[idx] != nil {
			xx := PoolIceServer.Get()
			x.IceServers[idx].DeepCopy(xx)
			z.IceServers = append(z.IceServers, xx)
		}
	}
}

func (x *PhoneParticipants) DeepCopy(z *PhoneParticipants) {
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolPhoneParticipant.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
}

func (x *IceServer) DeepCopy(z *IceServer) {
	z.Urls = append(z.Urls[:0], x.Urls...)
	z.Username = x.Username
	z.Credential = x.Credential
}

func (x *PhoneParticipant) DeepCopy(z *PhoneParticipant) {
	z.ConnectionId = x.ConnectionId
	if x.Peer != nil {
		z.Peer = PoolInputUser.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.Initiator = x.Initiator
	z.Admin = x.Admin
}

func (x *PhoneParticipantSDP) DeepCopy(z *PhoneParticipantSDP) {
	z.ConnectionId = x.ConnectionId
	if x.Peer != nil {
		z.Peer = PoolInputUser.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.SDP = x.SDP
	z.Type = x.Type
}

func (x *PhoneActionCallEmpty) DeepCopy(z *PhoneActionCallEmpty) {
	z.Empty = x.Empty
}

func (x *PhoneActionAccepted) DeepCopy(z *PhoneActionAccepted) {
	z.SDP = x.SDP
	z.Type = x.Type
}

func (x *PhoneActionRequested) DeepCopy(z *PhoneActionRequested) {
	z.SDP = x.SDP
	z.Type = x.Type
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolPhoneParticipant.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
}

func (x *PhoneActionCallWaiting) DeepCopy(z *PhoneActionCallWaiting) {
	z.Empty = x.Empty
}

func (x *PhoneActionDiscarded) DeepCopy(z *PhoneActionDiscarded) {
	z.Duration = x.Duration
	z.Video = x.Video
	z.Reason = x.Reason
	z.Terminate = x.Terminate
}

func (x *PhoneActionIceExchange) DeepCopy(z *PhoneActionIceExchange) {
	z.Candidate = x.Candidate
	z.SdpMLineIndex = x.SdpMLineIndex
	z.SdpMid = x.SdpMid
	z.UsernameFragment = x.UsernameFragment
}

func (x *PhoneActionAck) DeepCopy(z *PhoneActionAck) {
}

func (x *PhoneActionParticipantAdded) DeepCopy(z *PhoneActionParticipantAdded) {
	for idx := range x.Participants {
		if x.Participants[idx] != nil {
			xx := PoolPhoneParticipant.Get()
			x.Participants[idx].DeepCopy(xx)
			z.Participants = append(z.Participants, xx)
		}
	}
}

func (x *PhoneActionParticipantRemoved) DeepCopy(z *PhoneActionParticipantRemoved) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
	z.Timeout = x.Timeout
}

func (x *PhoneActionJoinRequested) DeepCopy(z *PhoneActionJoinRequested) {
	z.UserIDs = append(z.UserIDs[:0], x.UserIDs...)
}

func (x *PhoneActionMediaSettingsUpdated) DeepCopy(z *PhoneActionMediaSettingsUpdated) {
	z.Video = x.Video
	z.Audio = x.Audio
}

func (x *PhoneActionReactionSet) DeepCopy(z *PhoneActionReactionSet) {
	z.Reaction = x.Reaction
}

func (x *PhoneActionSDPOffer) DeepCopy(z *PhoneActionSDPOffer) {
	z.SDP = x.SDP
	z.Type = x.Type
}

func (x *PhoneActionSDPAnswer) DeepCopy(z *PhoneActionSDPAnswer) {
	z.SDP = x.SDP
	z.Type = x.Type
}

func (x *PhoneInitCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneInitCall, x)
}

func (x *PhoneRequestCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneRequestCall, x)
}

func (x *PhoneAcceptCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneAcceptCall, x)
}

func (x *PhoneDiscardCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneDiscardCall, x)
}

func (x *PhoneJoinCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneJoinCall, x)
}

func (x *PhoneAddParticipant) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneAddParticipant, x)
}

func (x *PhoneRemoveParticipant) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneRemoveParticipant, x)
}

func (x *PhoneUpdateCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneUpdateCall, x)
}

func (x *PhoneRateCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneRateCall, x)
}

func (x *PhoneGetHistory) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneGetHistory, x)
}

func (x *PhoneCallRecord) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneCallRecord, x)
}

func (x *PhoneCallsMany) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneCallsMany, x)
}

func (x *PhoneUpdateAdmin) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneUpdateAdmin, x)
}

func (x *PhoneCall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneCall, x)
}

func (x *PhoneInit) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneInit, x)
}

func (x *PhoneParticipants) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneParticipants, x)
}

func (x *IceServer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_IceServer, x)
}

func (x *PhoneParticipant) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneParticipant, x)
}

func (x *PhoneParticipantSDP) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneParticipantSDP, x)
}

func (x *PhoneActionCallEmpty) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionCallEmpty, x)
}

func (x *PhoneActionAccepted) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionAccepted, x)
}

func (x *PhoneActionRequested) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionRequested, x)
}

func (x *PhoneActionCallWaiting) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionCallWaiting, x)
}

func (x *PhoneActionDiscarded) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionDiscarded, x)
}

func (x *PhoneActionIceExchange) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionIceExchange, x)
}

func (x *PhoneActionAck) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionAck, x)
}

func (x *PhoneActionParticipantAdded) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionParticipantAdded, x)
}

func (x *PhoneActionParticipantRemoved) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionParticipantRemoved, x)
}

func (x *PhoneActionJoinRequested) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionJoinRequested, x)
}

func (x *PhoneActionMediaSettingsUpdated) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionMediaSettingsUpdated, x)
}

func (x *PhoneActionReactionSet) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionReactionSet, x)
}

func (x *PhoneActionSDPOffer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionSDPOffer, x)
}

func (x *PhoneActionSDPAnswer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_PhoneActionSDPAnswer, x)
}

func (x *PhoneInitCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneRequestCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneAcceptCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneDiscardCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneJoinCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneAddParticipant) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneRemoveParticipant) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneUpdateCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneRateCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneGetHistory) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneCallRecord) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneCallsMany) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneUpdateAdmin) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneCall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneInit) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneParticipants) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *IceServer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneParticipant) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneParticipantSDP) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionCallEmpty) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionAccepted) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionRequested) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionCallWaiting) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionDiscarded) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionIceExchange) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionAck) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionParticipantAdded) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionParticipantRemoved) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionJoinRequested) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionMediaSettingsUpdated) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionReactionSet) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionSDPOffer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneActionSDPAnswer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PhoneInitCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneRequestCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneAcceptCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneDiscardCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneJoinCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneAddParticipant) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneRemoveParticipant) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneUpdateCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneRateCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneGetHistory) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneCallRecord) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneCallsMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneUpdateAdmin) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneCall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneInit) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneParticipants) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *IceServer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneParticipant) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneParticipantSDP) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionCallEmpty) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionAccepted) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionRequested) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionCallWaiting) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionDiscarded) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionIceExchange) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionAck) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionParticipantAdded) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionParticipantRemoved) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionJoinRequested) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionMediaSettingsUpdated) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionReactionSet) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionSDPOffer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *PhoneActionSDPAnswer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}
