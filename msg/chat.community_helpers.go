// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.community.proto

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

const C_CommunitySendMessage int64 = 3506778488

type poolCommunitySendMessage struct {
	pool sync.Pool
}

func (p *poolCommunitySendMessage) Get() *CommunitySendMessage {
	x, ok := p.pool.Get().(*CommunitySendMessage)
	if !ok {
		return &CommunitySendMessage{}
	}
	x.Entities = x.Entities[:0]
	x.ReplyMarkup = 0
	x.ReplyMarkupData = nil
	x.ReplyMarkupData = x.ReplyMarkupData[:0]
	return x
}

func (p *poolCommunitySendMessage) Put(x *CommunitySendMessage) {
	p.pool.Put(x)
}

var PoolCommunitySendMessage = poolCommunitySendMessage{}

func ResultCommunitySendMessage(out *MessageEnvelope, res *CommunitySendMessage) {
	out.Constructor = C_CommunitySendMessage
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_CommunitySendMedia int64 = 2436824148

type poolCommunitySendMedia struct {
	pool sync.Pool
}

func (p *poolCommunitySendMedia) Get() *CommunitySendMedia {
	x, ok := p.pool.Get().(*CommunitySendMedia)
	if !ok {
		return &CommunitySendMedia{}
	}
	x.MediaData = x.MediaData[:0]
	x.ReplyTo = 0
	x.ClearDraft = false
	return x
}

func (p *poolCommunitySendMedia) Put(x *CommunitySendMedia) {
	p.pool.Put(x)
}

var PoolCommunitySendMedia = poolCommunitySendMedia{}

func ResultCommunitySendMedia(out *MessageEnvelope, res *CommunitySendMedia) {
	out.Constructor = C_CommunitySendMedia
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_CommunitySetTyping int64 = 3413516595

type poolCommunitySetTyping struct {
	pool sync.Pool
}

func (p *poolCommunitySetTyping) Get() *CommunitySetTyping {
	x, ok := p.pool.Get().(*CommunitySetTyping)
	if !ok {
		return &CommunitySetTyping{}
	}
	return x
}

func (p *poolCommunitySetTyping) Put(x *CommunitySetTyping) {
	p.pool.Put(x)
}

var PoolCommunitySetTyping = poolCommunitySetTyping{}

func ResultCommunitySetTyping(out *MessageEnvelope, res *CommunitySetTyping) {
	out.Constructor = C_CommunitySetTyping
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_CommunityGetUpdates int64 = 2021391963

type poolCommunityGetUpdates struct {
	pool sync.Pool
}

func (p *poolCommunityGetUpdates) Get() *CommunityGetUpdates {
	x, ok := p.pool.Get().(*CommunityGetUpdates)
	if !ok {
		return &CommunityGetUpdates{}
	}
	x.OffsetID = 0
	return x
}

func (p *poolCommunityGetUpdates) Put(x *CommunityGetUpdates) {
	p.pool.Put(x)
}

var PoolCommunityGetUpdates = poolCommunityGetUpdates{}

func ResultCommunityGetUpdates(out *MessageEnvelope, res *CommunityGetUpdates) {
	out.Constructor = C_CommunityGetUpdates
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
	ConstructorNames[3506778488] = "CommunitySendMessage"
	ConstructorNames[2436824148] = "CommunitySendMedia"
	ConstructorNames[3413516595] = "CommunitySetTyping"
	ConstructorNames[2021391963] = "CommunityGetUpdates"
}
