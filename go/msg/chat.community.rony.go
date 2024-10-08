// Code generated by Rony's protoc plugin; DO NOT EDIT.
// ProtoC ver. v4.25.3
// Rony ver. v0.12.22
// Source: chat.community.proto

package msg

import (
	bytes "bytes"
	edge "github.com/ronaksoft/rony/edge"
	pools "github.com/ronaksoft/rony/pools"
	registry "github.com/ronaksoft/rony/registry"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

var _ = pools.Imported

const C_CommunitySendMessage int64 = 3506778488

type poolCommunitySendMessage struct {
	pool sync.Pool
}

func (p *poolCommunitySendMessage) Get() *CommunitySendMessage {
	x, ok := p.pool.Get().(*CommunitySendMessage)
	if !ok {
		x = &CommunitySendMessage{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolCommunitySendMessage) Put(x *CommunitySendMessage) {
	if x == nil {
		return
	}

	x.RandomID = 0
	PoolInputPeer.Put(x.Peer)
	x.Body = ""
	for _, z := range x.Entities {
		PoolMessageEntity.Put(z)
	}
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	x.SenderID = 0
	x.SenderMsgID = 0

	p.pool.Put(x)
}

var PoolCommunitySendMessage = poolCommunitySendMessage{}

func (x *CommunitySendMessage) DeepCopy(z *CommunitySendMessage) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.Body = x.Body
	for idx := range x.Entities {
		if x.Entities[idx] == nil {
			continue
		}
		xx := PoolMessageEntity.Get()
		x.Entities[idx].DeepCopy(xx)
		z.Entities = append(z.Entities, xx)
	}
	z.ReplyMarkup = x.ReplyMarkup
	z.ReplyMarkupData = append(z.ReplyMarkupData[:0], x.ReplyMarkupData...)
	z.SenderID = x.SenderID
	z.SenderMsgID = x.SenderMsgID
}

func (x *CommunitySendMessage) Clone() *CommunitySendMessage {
	z := &CommunitySendMessage{}
	x.DeepCopy(z)
	return z
}

func (x *CommunitySendMessage) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunitySendMessage) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunitySendMessage) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunitySendMessage) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunitySendMessage) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySendMessage, x)
}

const C_CommunitySendMedia int64 = 2436824148

type poolCommunitySendMedia struct {
	pool sync.Pool
}

func (p *poolCommunitySendMedia) Get() *CommunitySendMedia {
	x, ok := p.pool.Get().(*CommunitySendMedia)
	if !ok {
		x = &CommunitySendMedia{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolCommunitySendMedia) Put(x *CommunitySendMedia) {
	if x == nil {
		return
	}

	x.RandomID = 0
	PoolInputPeer.Put(x.Peer)
	x.MediaType = 0
	x.MediaData = x.MediaData[:0]
	x.ReplyTo = 0
	x.ClearDraft = false
	x.SenderID = 0
	x.SenderMsgID = 0

	p.pool.Put(x)
}

var PoolCommunitySendMedia = poolCommunitySendMedia{}

func (x *CommunitySendMedia) DeepCopy(z *CommunitySendMedia) {
	z.RandomID = x.RandomID
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.MediaType = x.MediaType
	z.MediaData = append(z.MediaData[:0], x.MediaData...)
	z.ReplyTo = x.ReplyTo
	z.ClearDraft = x.ClearDraft
	z.SenderID = x.SenderID
	z.SenderMsgID = x.SenderMsgID
}

func (x *CommunitySendMedia) Clone() *CommunitySendMedia {
	z := &CommunitySendMedia{}
	x.DeepCopy(z)
	return z
}

func (x *CommunitySendMedia) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunitySendMedia) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunitySendMedia) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunitySendMedia) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunitySendMedia) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySendMedia, x)
}

const C_CommunitySetTyping int64 = 3413516595

type poolCommunitySetTyping struct {
	pool sync.Pool
}

func (p *poolCommunitySetTyping) Get() *CommunitySetTyping {
	x, ok := p.pool.Get().(*CommunitySetTyping)
	if !ok {
		x = &CommunitySetTyping{}
	}

	x.Peer = PoolInputPeer.Get()

	return x
}

func (p *poolCommunitySetTyping) Put(x *CommunitySetTyping) {
	if x == nil {
		return
	}

	PoolInputPeer.Put(x.Peer)
	x.Action = 0
	x.SenderID = 0

	p.pool.Put(x)
}

var PoolCommunitySetTyping = poolCommunitySetTyping{}

func (x *CommunitySetTyping) DeepCopy(z *CommunitySetTyping) {
	if x.Peer != nil {
		if z.Peer == nil {
			z.Peer = PoolInputPeer.Get()
		}
		x.Peer.DeepCopy(z.Peer)
	} else {
		PoolInputPeer.Put(z.Peer)
		z.Peer = nil
	}
	z.Action = x.Action
	z.SenderID = x.SenderID
}

func (x *CommunitySetTyping) Clone() *CommunitySetTyping {
	z := &CommunitySetTyping{}
	x.DeepCopy(z)
	return z
}

func (x *CommunitySetTyping) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunitySetTyping) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunitySetTyping) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunitySetTyping) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunitySetTyping) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunitySetTyping, x)
}

const C_CommunityGetUpdates int64 = 2021391963

type poolCommunityGetUpdates struct {
	pool sync.Pool
}

func (p *poolCommunityGetUpdates) Get() *CommunityGetUpdates {
	x, ok := p.pool.Get().(*CommunityGetUpdates)
	if !ok {
		x = &CommunityGetUpdates{}
	}

	return x
}

func (p *poolCommunityGetUpdates) Put(x *CommunityGetUpdates) {
	if x == nil {
		return
	}

	x.WaitAfterInMS = 0
	x.WaitMaxInMS = 0
	x.SizeLimit = 0
	x.OffsetID = 0

	p.pool.Put(x)
}

var PoolCommunityGetUpdates = poolCommunityGetUpdates{}

func (x *CommunityGetUpdates) DeepCopy(z *CommunityGetUpdates) {
	z.WaitAfterInMS = x.WaitAfterInMS
	z.WaitMaxInMS = x.WaitMaxInMS
	z.SizeLimit = x.SizeLimit
	z.OffsetID = x.OffsetID
}

func (x *CommunityGetUpdates) Clone() *CommunityGetUpdates {
	z := &CommunityGetUpdates{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityGetUpdates) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityGetUpdates) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityGetUpdates) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityGetUpdates) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityGetUpdates) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityGetUpdates, x)
}

const C_CommunityGetMembers int64 = 2022915988

type poolCommunityGetMembers struct {
	pool sync.Pool
}

func (p *poolCommunityGetMembers) Get() *CommunityGetMembers {
	x, ok := p.pool.Get().(*CommunityGetMembers)
	if !ok {
		x = &CommunityGetMembers{}
	}

	return x
}

func (p *poolCommunityGetMembers) Put(x *CommunityGetMembers) {
	if x == nil {
		return
	}

	x.Offset = 0
	x.Limit = 0

	p.pool.Put(x)
}

var PoolCommunityGetMembers = poolCommunityGetMembers{}

func (x *CommunityGetMembers) DeepCopy(z *CommunityGetMembers) {
	z.Offset = x.Offset
	z.Limit = x.Limit
}

func (x *CommunityGetMembers) Clone() *CommunityGetMembers {
	z := &CommunityGetMembers{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityGetMembers) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityGetMembers) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityGetMembers) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityGetMembers) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityGetMembers) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityGetMembers, x)
}

const C_CommunityRecall int64 = 890349574

type poolCommunityRecall struct {
	pool sync.Pool
}

func (p *poolCommunityRecall) Get() *CommunityRecall {
	x, ok := p.pool.Get().(*CommunityRecall)
	if !ok {
		x = &CommunityRecall{}
	}

	return x
}

func (p *poolCommunityRecall) Put(x *CommunityRecall) {
	if x == nil {
		return
	}

	x.TeamID = 0
	x.AccessKey = x.AccessKey[:0]

	p.pool.Put(x)
}

var PoolCommunityRecall = poolCommunityRecall{}

func (x *CommunityRecall) DeepCopy(z *CommunityRecall) {
	z.TeamID = x.TeamID
	z.AccessKey = append(z.AccessKey[:0], x.AccessKey...)
}

func (x *CommunityRecall) Clone() *CommunityRecall {
	z := &CommunityRecall{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityRecall) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityRecall) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityRecall) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityRecall) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityRecall) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityRecall, x)
}

const C_CommunityAuthorizeUser int64 = 1452766231

type poolCommunityAuthorizeUser struct {
	pool sync.Pool
}

func (p *poolCommunityAuthorizeUser) Get() *CommunityAuthorizeUser {
	x, ok := p.pool.Get().(*CommunityAuthorizeUser)
	if !ok {
		x = &CommunityAuthorizeUser{}
	}

	return x
}

func (p *poolCommunityAuthorizeUser) Put(x *CommunityAuthorizeUser) {
	if x == nil {
		return
	}

	x.Phone = ""
	x.FirstName = ""
	x.LastName = ""
	x.Provider = ""

	p.pool.Put(x)
}

var PoolCommunityAuthorizeUser = poolCommunityAuthorizeUser{}

func (x *CommunityAuthorizeUser) DeepCopy(z *CommunityAuthorizeUser) {
	z.Phone = x.Phone
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Provider = x.Provider
}

func (x *CommunityAuthorizeUser) Clone() *CommunityAuthorizeUser {
	z := &CommunityAuthorizeUser{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityAuthorizeUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityAuthorizeUser) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityAuthorizeUser) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityAuthorizeUser) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityAuthorizeUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityAuthorizeUser, x)
}

const C_CommunityUser int64 = 3812837958

type poolCommunityUser struct {
	pool sync.Pool
}

func (p *poolCommunityUser) Get() *CommunityUser {
	x, ok := p.pool.Get().(*CommunityUser)
	if !ok {
		x = &CommunityUser{}
	}

	return x
}

func (p *poolCommunityUser) Put(x *CommunityUser) {
	if x == nil {
		return
	}

	x.UserID = 0
	x.FirstName = ""
	x.LastName = ""
	x.Phone = ""

	p.pool.Put(x)
}

var PoolCommunityUser = poolCommunityUser{}

func (x *CommunityUser) DeepCopy(z *CommunityUser) {
	z.UserID = x.UserID
	z.FirstName = x.FirstName
	z.LastName = x.LastName
	z.Phone = x.Phone
}

func (x *CommunityUser) Clone() *CommunityUser {
	z := &CommunityUser{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityUser) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityUser) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityUser) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityUser) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityUser) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUser, x)
}

const C_CommunityUpdateEnvelope int64 = 1076119993

type poolCommunityUpdateEnvelope struct {
	pool sync.Pool
}

func (p *poolCommunityUpdateEnvelope) Get() *CommunityUpdateEnvelope {
	x, ok := p.pool.Get().(*CommunityUpdateEnvelope)
	if !ok {
		x = &CommunityUpdateEnvelope{}
	}

	return x
}

func (p *poolCommunityUpdateEnvelope) Put(x *CommunityUpdateEnvelope) {
	if x == nil {
		return
	}

	x.OffsetID = 0
	x.PartitionID = 0
	x.Constructor = 0
	x.Update = x.Update[:0]

	p.pool.Put(x)
}

var PoolCommunityUpdateEnvelope = poolCommunityUpdateEnvelope{}

func (x *CommunityUpdateEnvelope) DeepCopy(z *CommunityUpdateEnvelope) {
	z.OffsetID = x.OffsetID
	z.PartitionID = x.PartitionID
	z.Constructor = x.Constructor
	z.Update = append(z.Update[:0], x.Update...)
}

func (x *CommunityUpdateEnvelope) Clone() *CommunityUpdateEnvelope {
	z := &CommunityUpdateEnvelope{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityUpdateEnvelope) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityUpdateEnvelope) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityUpdateEnvelope) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityUpdateEnvelope) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityUpdateEnvelope) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUpdateEnvelope, x)
}

const C_CommunityUpdateContainer int64 = 918339432

type poolCommunityUpdateContainer struct {
	pool sync.Pool
}

func (p *poolCommunityUpdateContainer) Get() *CommunityUpdateContainer {
	x, ok := p.pool.Get().(*CommunityUpdateContainer)
	if !ok {
		x = &CommunityUpdateContainer{}
	}

	return x
}

func (p *poolCommunityUpdateContainer) Put(x *CommunityUpdateContainer) {
	if x == nil {
		return
	}

	for _, z := range x.Updates {
		PoolCommunityUpdateEnvelope.Put(z)
	}
	x.Updates = x.Updates[:0]
	x.Empty = false

	p.pool.Put(x)
}

var PoolCommunityUpdateContainer = poolCommunityUpdateContainer{}

func (x *CommunityUpdateContainer) DeepCopy(z *CommunityUpdateContainer) {
	for idx := range x.Updates {
		if x.Updates[idx] == nil {
			continue
		}
		xx := PoolCommunityUpdateEnvelope.Get()
		x.Updates[idx].DeepCopy(xx)
		z.Updates = append(z.Updates, xx)
	}
	z.Empty = x.Empty
}

func (x *CommunityUpdateContainer) Clone() *CommunityUpdateContainer {
	z := &CommunityUpdateContainer{}
	x.DeepCopy(z)
	return z
}

func (x *CommunityUpdateContainer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{Merge: true}.Unmarshal(b, x)
}

func (x *CommunityUpdateContainer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CommunityUpdateContainer) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *CommunityUpdateContainer) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *CommunityUpdateContainer) PushToContext(ctx *edge.RequestCtx) {
	ctx.PushMessage(C_CommunityUpdateContainer, x)
}

func init() {
	registry.RegisterConstructor(3506778488, "CommunitySendMessage")
	registry.RegisterConstructor(2436824148, "CommunitySendMedia")
	registry.RegisterConstructor(3413516595, "CommunitySetTyping")
	registry.RegisterConstructor(2021391963, "CommunityGetUpdates")
	registry.RegisterConstructor(2022915988, "CommunityGetMembers")
	registry.RegisterConstructor(890349574, "CommunityRecall")
	registry.RegisterConstructor(1452766231, "CommunityAuthorizeUser")
	registry.RegisterConstructor(3812837958, "CommunityUser")
	registry.RegisterConstructor(1076119993, "CommunityUpdateEnvelope")
	registry.RegisterConstructor(918339432, "CommunityUpdateContainer")
}

var _ = bytes.MinRead
