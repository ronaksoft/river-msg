// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.phone.proto

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
		*x.Peer = InputPeer{}
	}

	p.pool.Put(x)
}

var PoolPhoneInitCall = poolPhoneInitCall{}

func ResultPhoneInitCall(out *MessageEnvelope, res *PhoneInitCall) {
	out.Constructor = C_PhoneInitCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputPeer{}
	}

	x.Initiator = false
	x.Participants = x.Participants[:0]
	x.CallID = 0
	p.pool.Put(x)
}

var PoolPhoneRequestCall = poolPhoneRequestCall{}

func ResultPhoneRequestCall(out *MessageEnvelope, res *PhoneRequestCall) {
	out.Constructor = C_PhoneRequestCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputPeer{}
	}

	x.CallID = 0
	x.Participants = x.Participants[:0]
	p.pool.Put(x)
}

var PoolPhoneAcceptCall = poolPhoneAcceptCall{}

func ResultPhoneAcceptCall(out *MessageEnvelope, res *PhoneAcceptCall) {
	out.Constructor = C_PhoneAcceptCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputPeer{}
	}

	x.CallID = 0
	x.Participants = x.Participants[:0]
	x.Duration = 0
	x.Reason = 0
	p.pool.Put(x)
}

var PoolPhoneDiscardCall = poolPhoneDiscardCall{}

func ResultPhoneDiscardCall(out *MessageEnvelope, res *PhoneDiscardCall) {
	out.Constructor = C_PhoneDiscardCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputPeer{}
	}

	x.CallID = 0
	x.Participants = x.Participants[:0]
	x.Action = 0
	x.ActionData = x.ActionData[:0]
	p.pool.Put(x)
}

var PoolPhoneUpdateCall = poolPhoneUpdateCall{}

func ResultPhoneUpdateCall(out *MessageEnvelope, res *PhoneUpdateCall) {
	out.Constructor = C_PhoneUpdateCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputPeer{}
	}

	x.CallID = 0
	x.Rate = 0
	x.Comment = ""
	p.pool.Put(x)
}

var PoolPhoneRateCall = poolPhoneRateCall{}

func ResultPhoneRateCall(out *MessageEnvelope, res *PhoneRateCall) {
	out.Constructor = C_PhoneRateCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
	x.Date = 0
	p.pool.Put(x)
}

var PoolPhoneCall = poolPhoneCall{}

func ResultPhoneCall(out *MessageEnvelope, res *PhoneCall) {
	out.Constructor = C_PhoneCall
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneInit(out *MessageEnvelope, res *PhoneInit) {
	out.Constructor = C_PhoneInit
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultIceServer(out *MessageEnvelope, res *IceServer) {
	out.Constructor = C_IceServer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputUser{}
	}

	x.Initiator = false
	x.Admin = false
	p.pool.Put(x)
}

var PoolPhoneParticipant = poolPhoneParticipant{}

func ResultPhoneParticipant(out *MessageEnvelope, res *PhoneParticipant) {
	out.Constructor = C_PhoneParticipant
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
		*x.Peer = InputUser{}
	}

	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneParticipantSDP = poolPhoneParticipantSDP{}

func ResultPhoneParticipantSDP(out *MessageEnvelope, res *PhoneParticipantSDP) {
	out.Constructor = C_PhoneParticipantSDP
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneActionCallEmpty(out *MessageEnvelope, res *PhoneActionCallEmpty) {
	out.Constructor = C_PhoneActionCallEmpty
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneActionAccepted(out *MessageEnvelope, res *PhoneActionAccepted) {
	out.Constructor = C_PhoneActionAccepted
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneActionRequested(out *MessageEnvelope, res *PhoneActionRequested) {
	out.Constructor = C_PhoneActionRequested
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneActionCallWaiting(out *MessageEnvelope, res *PhoneActionCallWaiting) {
	out.Constructor = C_PhoneActionCallWaiting
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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
	p.pool.Put(x)
}

var PoolPhoneActionDiscarded = poolPhoneActionDiscarded{}

func ResultPhoneActionDiscarded(out *MessageEnvelope, res *PhoneActionDiscarded) {
	out.Constructor = C_PhoneActionDiscarded
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

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

func ResultPhoneActionIceExchange(out *MessageEnvelope, res *PhoneActionIceExchange) {
	out.Constructor = C_PhoneActionIceExchange
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PhoneMediaSettingsUpdated int64 = 163140236

type poolPhoneMediaSettingsUpdated struct {
	pool sync.Pool
}

func (p *poolPhoneMediaSettingsUpdated) Get() *PhoneMediaSettingsUpdated {
	x, ok := p.pool.Get().(*PhoneMediaSettingsUpdated)
	if !ok {
		return &PhoneMediaSettingsUpdated{}
	}
	return x
}

func (p *poolPhoneMediaSettingsUpdated) Put(x *PhoneMediaSettingsUpdated) {
	x.Video = false
	x.Audio = false
	p.pool.Put(x)
}

var PoolPhoneMediaSettingsUpdated = poolPhoneMediaSettingsUpdated{}

func ResultPhoneMediaSettingsUpdated(out *MessageEnvelope, res *PhoneMediaSettingsUpdated) {
	out.Constructor = C_PhoneMediaSettingsUpdated
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PhoneReactionSet int64 = 3821475130

type poolPhoneReactionSet struct {
	pool sync.Pool
}

func (p *poolPhoneReactionSet) Get() *PhoneReactionSet {
	x, ok := p.pool.Get().(*PhoneReactionSet)
	if !ok {
		return &PhoneReactionSet{}
	}
	return x
}

func (p *poolPhoneReactionSet) Put(x *PhoneReactionSet) {
	x.Reaction = ""
	p.pool.Put(x)
}

var PoolPhoneReactionSet = poolPhoneReactionSet{}

func ResultPhoneReactionSet(out *MessageEnvelope, res *PhoneReactionSet) {
	out.Constructor = C_PhoneReactionSet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PhoneSDPOffer int64 = 2063600460

type poolPhoneSDPOffer struct {
	pool sync.Pool
}

func (p *poolPhoneSDPOffer) Get() *PhoneSDPOffer {
	x, ok := p.pool.Get().(*PhoneSDPOffer)
	if !ok {
		return &PhoneSDPOffer{}
	}
	return x
}

func (p *poolPhoneSDPOffer) Put(x *PhoneSDPOffer) {
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneSDPOffer = poolPhoneSDPOffer{}

func ResultPhoneSDPOffer(out *MessageEnvelope, res *PhoneSDPOffer) {
	out.Constructor = C_PhoneSDPOffer
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_PhoneSDPAnswer int64 = 1686408377

type poolPhoneSDPAnswer struct {
	pool sync.Pool
}

func (p *poolPhoneSDPAnswer) Get() *PhoneSDPAnswer {
	x, ok := p.pool.Get().(*PhoneSDPAnswer)
	if !ok {
		return &PhoneSDPAnswer{}
	}
	return x
}

func (p *poolPhoneSDPAnswer) Put(x *PhoneSDPAnswer) {
	x.SDP = ""
	x.Type = ""
	p.pool.Put(x)
}

var PoolPhoneSDPAnswer = poolPhoneSDPAnswer{}

func ResultPhoneSDPAnswer(out *MessageEnvelope, res *PhoneSDPAnswer) {
	out.Constructor = C_PhoneSDPAnswer
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
	ConstructorNames[2975617068] = "PhoneInitCall"
	ConstructorNames[907942641] = "PhoneRequestCall"
	ConstructorNames[4133092858] = "PhoneAcceptCall"
	ConstructorNames[2712700137] = "PhoneDiscardCall"
	ConstructorNames[1976202226] = "PhoneUpdateCall"
	ConstructorNames[2215486159] = "PhoneRateCall"
	ConstructorNames[3296664529] = "PhoneCall"
	ConstructorNames[3464876187] = "PhoneInit"
	ConstructorNames[4291892363] = "IceServer"
	ConstructorNames[226273622] = "PhoneParticipant"
	ConstructorNames[545454774] = "PhoneParticipantSDP"
	ConstructorNames[1073285997] = "PhoneActionCallEmpty"
	ConstructorNames[2493210645] = "PhoneActionAccepted"
	ConstructorNames[1678316869] = "PhoneActionRequested"
	ConstructorNames[3634710697] = "PhoneActionCallWaiting"
	ConstructorNames[4285966731] = "PhoneActionDiscarded"
	ConstructorNames[1618781621] = "PhoneActionIceExchange"
	ConstructorNames[163140236] = "PhoneMediaSettingsUpdated"
	ConstructorNames[3821475130] = "PhoneReactionSet"
	ConstructorNames[2063600460] = "PhoneSDPOffer"
	ConstructorNames[1686408377] = "PhoneSDPAnswer"
}