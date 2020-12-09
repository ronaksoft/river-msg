package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_LabelsCreate int64 = 2138857068

type poolLabelsCreate struct {
	pool sync.Pool
}

func (p *poolLabelsCreate) Get() *LabelsCreate {
	x, ok := p.pool.Get().(*LabelsCreate)
	if !ok {
		return &LabelsCreate{}
	}
	return x
}

func (p *poolLabelsCreate) Put(x *LabelsCreate) {
	x.RandomID = 0
	x.Name = ""
	x.Colour = ""
	p.pool.Put(x)
}

var PoolLabelsCreate = poolLabelsCreate{}

const C_LabelsEdit int64 = 2790466877

type poolLabelsEdit struct {
	pool sync.Pool
}

func (p *poolLabelsEdit) Get() *LabelsEdit {
	x, ok := p.pool.Get().(*LabelsEdit)
	if !ok {
		return &LabelsEdit{}
	}
	return x
}

func (p *poolLabelsEdit) Put(x *LabelsEdit) {
	x.LabelID = 0
	x.Name = ""
	x.Colour = ""
	p.pool.Put(x)
}

var PoolLabelsEdit = poolLabelsEdit{}

const C_LabelsDelete int64 = 3401105936

type poolLabelsDelete struct {
	pool sync.Pool
}

func (p *poolLabelsDelete) Get() *LabelsDelete {
	x, ok := p.pool.Get().(*LabelsDelete)
	if !ok {
		return &LabelsDelete{}
	}
	return x
}

func (p *poolLabelsDelete) Put(x *LabelsDelete) {
	x.LabelIDs = x.LabelIDs[:0]
	p.pool.Put(x)
}

var PoolLabelsDelete = poolLabelsDelete{}

const C_LabelsGet int64 = 2575409921

type poolLabelsGet struct {
	pool sync.Pool
}

func (p *poolLabelsGet) Get() *LabelsGet {
	x, ok := p.pool.Get().(*LabelsGet)
	if !ok {
		return &LabelsGet{}
	}
	return x
}

func (p *poolLabelsGet) Put(x *LabelsGet) {
	p.pool.Put(x)
}

var PoolLabelsGet = poolLabelsGet{}

const C_LabelsAddToMessage int64 = 180144503

type poolLabelsAddToMessage struct {
	pool sync.Pool
}

func (p *poolLabelsAddToMessage) Get() *LabelsAddToMessage {
	x, ok := p.pool.Get().(*LabelsAddToMessage)
	if !ok {
		return &LabelsAddToMessage{}
	}
	return x
}

func (p *poolLabelsAddToMessage) Put(x *LabelsAddToMessage) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]
	p.pool.Put(x)
}

var PoolLabelsAddToMessage = poolLabelsAddToMessage{}

const C_LabelsRemoveFromMessage int64 = 4195197703

type poolLabelsRemoveFromMessage struct {
	pool sync.Pool
}

func (p *poolLabelsRemoveFromMessage) Get() *LabelsRemoveFromMessage {
	x, ok := p.pool.Get().(*LabelsRemoveFromMessage)
	if !ok {
		return &LabelsRemoveFromMessage{}
	}
	return x
}

func (p *poolLabelsRemoveFromMessage) Put(x *LabelsRemoveFromMessage) {
	if x.Peer != nil {
		PoolInputPeer.Put(x.Peer)
		x.Peer = nil
	}
	x.LabelIDs = x.LabelIDs[:0]
	x.MessageIDs = x.MessageIDs[:0]
	p.pool.Put(x)
}

var PoolLabelsRemoveFromMessage = poolLabelsRemoveFromMessage{}

const C_LabelsListItems int64 = 2351763198

type poolLabelsListItems struct {
	pool sync.Pool
}

func (p *poolLabelsListItems) Get() *LabelsListItems {
	x, ok := p.pool.Get().(*LabelsListItems)
	if !ok {
		return &LabelsListItems{}
	}
	return x
}

func (p *poolLabelsListItems) Put(x *LabelsListItems) {
	x.LabelID = 0
	x.MinID = 0
	x.MaxID = 0
	x.Limit = 0
	p.pool.Put(x)
}

var PoolLabelsListItems = poolLabelsListItems{}

const C_LabelItems int64 = 4271841358

type poolLabelItems struct {
	pool sync.Pool
}

func (p *poolLabelItems) Get() *LabelItems {
	x, ok := p.pool.Get().(*LabelItems)
	if !ok {
		return &LabelItems{}
	}
	return x
}

func (p *poolLabelItems) Put(x *LabelItems) {
	x.LabelID = 0
	x.Messages = x.Messages[:0]
	x.Dialogs = x.Dialogs[:0]
	x.Users = x.Users[:0]
	x.Groups = x.Groups[:0]
	p.pool.Put(x)
}

var PoolLabelItems = poolLabelItems{}

func init() {
	registry.RegisterConstructor(2138857068, "LabelsCreate")
	registry.RegisterConstructor(2790466877, "LabelsEdit")
	registry.RegisterConstructor(3401105936, "LabelsDelete")
	registry.RegisterConstructor(2575409921, "LabelsGet")
	registry.RegisterConstructor(180144503, "LabelsAddToMessage")
	registry.RegisterConstructor(4195197703, "LabelsRemoveFromMessage")
	registry.RegisterConstructor(2351763198, "LabelsListItems")
	registry.RegisterConstructor(4271841358, "LabelItems")
}

func (x *LabelsCreate) DeepCopy(z *LabelsCreate) {
	z.RandomID = x.RandomID
	z.Name = x.Name
	z.Colour = x.Colour
}

func (x *LabelsEdit) DeepCopy(z *LabelsEdit) {
	z.LabelID = x.LabelID
	z.Name = x.Name
	z.Colour = x.Colour
}

func (x *LabelsDelete) DeepCopy(z *LabelsDelete) {
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
}

func (x *LabelsGet) DeepCopy(z *LabelsGet) {
}

func (x *LabelsAddToMessage) DeepCopy(z *LabelsAddToMessage) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *LabelsRemoveFromMessage) DeepCopy(z *LabelsRemoveFromMessage) {
	if x.Peer != nil {
		z.Peer = PoolInputPeer.Get()
		x.Peer.DeepCopy(z.Peer)
	}
	z.LabelIDs = append(z.LabelIDs[:0], x.LabelIDs...)
	z.MessageIDs = append(z.MessageIDs[:0], x.MessageIDs...)
}

func (x *LabelsListItems) DeepCopy(z *LabelsListItems) {
	z.LabelID = x.LabelID
	z.MinID = x.MinID
	z.MaxID = x.MaxID
	z.Limit = x.Limit
}

func (x *LabelItems) DeepCopy(z *LabelItems) {
	z.LabelID = x.LabelID
	for idx := range x.Messages {
		if x.Messages[idx] != nil {
			xx := PoolUserMessage.Get()
			x.Messages[idx].DeepCopy(xx)
			z.Messages = append(z.Messages, xx)
		}
	}
	for idx := range x.Dialogs {
		if x.Dialogs[idx] != nil {
			xx := PoolDialog.Get()
			x.Dialogs[idx].DeepCopy(xx)
			z.Dialogs = append(z.Dialogs, xx)
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
