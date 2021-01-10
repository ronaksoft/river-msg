package rony

import (
	registry "github.com/ronaksoft/rony/registry"
	proto "google.golang.org/protobuf/proto"
	sync "sync"
)

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
	x.Auth = x.Auth[:0]
	x.Header = x.Header[:0]
	p.pool.Put(x)
}

var PoolMessageEnvelope = poolMessageEnvelope{}

const C_KeyValue int64 = 4276272820

type poolKeyValue struct {
	pool sync.Pool
}

func (p *poolKeyValue) Get() *KeyValue {
	x, ok := p.pool.Get().(*KeyValue)
	if !ok {
		return &KeyValue{}
	}
	return x
}

func (p *poolKeyValue) Put(x *KeyValue) {
	x.Key = ""
	x.Value = ""
	p.pool.Put(x)
}

var PoolKeyValue = poolKeyValue{}

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
	x.Template = ""
	x.TemplateItems = x.TemplateItems[:0]
	x.LocalTemplate = ""
	x.LocalTemplateItems = x.LocalTemplateItems[:0]
	p.pool.Put(x)
}

var PoolError = poolError{}

const C_Redirect int64 = 981138557

type poolRedirect struct {
	pool sync.Pool
}

func (p *poolRedirect) Get() *Redirect {
	x, ok := p.pool.Get().(*Redirect)
	if !ok {
		return &Redirect{}
	}
	return x
}

func (p *poolRedirect) Put(x *Redirect) {
	x.Reason = 0
	if x.Leader != nil {
		PoolNodeInfo.Put(x.Leader)
		x.Leader = nil
	}
	x.Followers = x.Followers[:0]
	x.WaitInSec = 0
	p.pool.Put(x)
}

var PoolRedirect = poolRedirect{}

const C_NodeInfo int64 = 2894317502

type poolNodeInfo struct {
	pool sync.Pool
}

func (p *poolNodeInfo) Get() *NodeInfo {
	x, ok := p.pool.Get().(*NodeInfo)
	if !ok {
		return &NodeInfo{}
	}
	return x
}

func (p *poolNodeInfo) Put(x *NodeInfo) {
	x.ReplicaSet = 0
	x.ServerID = ""
	x.HostPorts = x.HostPorts[:0]
	x.Leader = false
	p.pool.Put(x)
}

var PoolNodeInfo = poolNodeInfo{}

const C_GetNodes int64 = 362407405

type poolGetNodes struct {
	pool sync.Pool
}

func (p *poolGetNodes) Get() *GetNodes {
	x, ok := p.pool.Get().(*GetNodes)
	if !ok {
		return &GetNodes{}
	}
	return x
}

func (p *poolGetNodes) Put(x *GetNodes) {
	x.ReplicaSet = x.ReplicaSet[:0]
	p.pool.Put(x)
}

var PoolGetNodes = poolGetNodes{}

const C_NodeInfoMany int64 = 1963715709

type poolNodeInfoMany struct {
	pool sync.Pool
}

func (p *poolNodeInfoMany) Get() *NodeInfoMany {
	x, ok := p.pool.Get().(*NodeInfoMany)
	if !ok {
		return &NodeInfoMany{}
	}
	return x
}

func (p *poolNodeInfoMany) Put(x *NodeInfoMany) {
	x.Nodes = x.Nodes[:0]
	p.pool.Put(x)
}

var PoolNodeInfoMany = poolNodeInfoMany{}

func init() {
	registry.RegisterConstructor(535232465, "MessageEnvelope")
	registry.RegisterConstructor(4276272820, "KeyValue")
	registry.RegisterConstructor(1972016308, "MessageContainer")
	registry.RegisterConstructor(2619118453, "Error")
	registry.RegisterConstructor(981138557, "Redirect")
	registry.RegisterConstructor(2894317502, "NodeInfo")
	registry.RegisterConstructor(362407405, "GetNodes")
	registry.RegisterConstructor(1963715709, "NodeInfoMany")
}

func (x *MessageEnvelope) DeepCopy(z *MessageEnvelope) {
	z.Constructor = x.Constructor
	z.RequestID = x.RequestID
	z.Message = append(z.Message[:0], x.Message...)
	z.Auth = append(z.Auth[:0], x.Auth...)
	for idx := range x.Header {
		if x.Header[idx] != nil {
			xx := PoolKeyValue.Get()
			x.Header[idx].DeepCopy(xx)
			z.Header = append(z.Header, xx)
		}
	}
}

func (x *KeyValue) DeepCopy(z *KeyValue) {
	z.Key = x.Key
	z.Value = x.Value
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

func (x *Error) DeepCopy(z *Error) {
	z.Code = x.Code
	z.Items = x.Items
	z.Template = x.Template
	z.TemplateItems = append(z.TemplateItems[:0], x.TemplateItems...)
	z.LocalTemplate = x.LocalTemplate
	z.LocalTemplateItems = append(z.LocalTemplateItems[:0], x.LocalTemplateItems...)
}

func (x *Redirect) DeepCopy(z *Redirect) {
	z.Reason = x.Reason
	if x.Leader != nil {
		z.Leader = PoolNodeInfo.Get()
		x.Leader.DeepCopy(z.Leader)
	}
	for idx := range x.Followers {
		if x.Followers[idx] != nil {
			xx := PoolNodeInfo.Get()
			x.Followers[idx].DeepCopy(xx)
			z.Followers = append(z.Followers, xx)
		}
	}
	z.WaitInSec = x.WaitInSec
}

func (x *NodeInfo) DeepCopy(z *NodeInfo) {
	z.ReplicaSet = x.ReplicaSet
	z.ServerID = x.ServerID
	z.HostPorts = append(z.HostPorts[:0], x.HostPorts...)
	z.Leader = x.Leader
}

func (x *GetNodes) DeepCopy(z *GetNodes) {
	z.ReplicaSet = append(z.ReplicaSet[:0], x.ReplicaSet...)
}

func (x *NodeInfoMany) DeepCopy(z *NodeInfoMany) {
	for idx := range x.Nodes {
		if x.Nodes[idx] != nil {
			xx := PoolNodeInfo.Get()
			x.Nodes[idx].DeepCopy(xx)
			z.Nodes = append(z.Nodes, xx)
		}
	}
}

func (x *MessageEnvelope) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *KeyValue) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageContainer) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Error) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Redirect) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *NodeInfo) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *GetNodes) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *NodeInfoMany) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *MessageEnvelope) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *KeyValue) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *MessageContainer) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *Error) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *Redirect) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *NodeInfo) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *GetNodes) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}

func (x *NodeInfoMany) Unmarshal(b []byte) error {
	return proto.UnmarshalOptions{}.Unmarshal(b, x)
}
