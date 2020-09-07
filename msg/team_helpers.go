// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team.proto

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

const C_TeamGet int64 = 1172720786

type poolTeamGet struct {
	pool sync.Pool
}

func (p *poolTeamGet) Get() *TeamGet {
	x, ok := p.pool.Get().(*TeamGet)
	if !ok {
		return &TeamGet{}
	}
	return x
}

func (p *poolTeamGet) Put(x *TeamGet) {
	p.pool.Put(x)
}

var PoolTeamGet = poolTeamGet{}

func ResultTeamGet(out *MessageEnvelope, res *TeamGet) {
	out.Constructor = C_TeamGet
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamAddMember int64 = 3889056091

type poolTeamAddMember struct {
	pool sync.Pool
}

func (p *poolTeamAddMember) Get() *TeamAddMember {
	x, ok := p.pool.Get().(*TeamAddMember)
	if !ok {
		return &TeamAddMember{}
	}
	return x
}

func (p *poolTeamAddMember) Put(x *TeamAddMember) {
	p.pool.Put(x)
}

var PoolTeamAddMember = poolTeamAddMember{}

func ResultTeamAddMember(out *MessageEnvelope, res *TeamAddMember) {
	out.Constructor = C_TeamAddMember
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamRemoveMember int64 = 4200364613

type poolTeamRemoveMember struct {
	pool sync.Pool
}

func (p *poolTeamRemoveMember) Get() *TeamRemoveMember {
	x, ok := p.pool.Get().(*TeamRemoveMember)
	if !ok {
		return &TeamRemoveMember{}
	}
	return x
}

func (p *poolTeamRemoveMember) Put(x *TeamRemoveMember) {
	p.pool.Put(x)
}

var PoolTeamRemoveMember = poolTeamRemoveMember{}

func ResultTeamRemoveMember(out *MessageEnvelope, res *TeamRemoveMember) {
	out.Constructor = C_TeamRemoveMember
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamPromote int64 = 382328820

type poolTeamPromote struct {
	pool sync.Pool
}

func (p *poolTeamPromote) Get() *TeamPromote {
	x, ok := p.pool.Get().(*TeamPromote)
	if !ok {
		return &TeamPromote{}
	}
	return x
}

func (p *poolTeamPromote) Put(x *TeamPromote) {
	p.pool.Put(x)
}

var PoolTeamPromote = poolTeamPromote{}

func ResultTeamPromote(out *MessageEnvelope, res *TeamPromote) {
	out.Constructor = C_TeamPromote
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamDemote int64 = 2331393294

type poolTeamDemote struct {
	pool sync.Pool
}

func (p *poolTeamDemote) Get() *TeamDemote {
	x, ok := p.pool.Get().(*TeamDemote)
	if !ok {
		return &TeamDemote{}
	}
	return x
}

func (p *poolTeamDemote) Put(x *TeamDemote) {
	p.pool.Put(x)
}

var PoolTeamDemote = poolTeamDemote{}

func ResultTeamDemote(out *MessageEnvelope, res *TeamDemote) {
	out.Constructor = C_TeamDemote
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamLeave int64 = 1413785879

type poolTeamLeave struct {
	pool sync.Pool
}

func (p *poolTeamLeave) Get() *TeamLeave {
	x, ok := p.pool.Get().(*TeamLeave)
	if !ok {
		return &TeamLeave{}
	}
	return x
}

func (p *poolTeamLeave) Put(x *TeamLeave) {
	p.pool.Put(x)
}

var PoolTeamLeave = poolTeamLeave{}

func ResultTeamLeave(out *MessageEnvelope, res *TeamLeave) {
	out.Constructor = C_TeamLeave
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamJoin int64 = 1725794017

type poolTeamJoin struct {
	pool sync.Pool
}

func (p *poolTeamJoin) Get() *TeamJoin {
	x, ok := p.pool.Get().(*TeamJoin)
	if !ok {
		return &TeamJoin{}
	}
	return x
}

func (p *poolTeamJoin) Put(x *TeamJoin) {
	x.Token = x.Token[:0]
	p.pool.Put(x)
}

var PoolTeamJoin = poolTeamJoin{}

func ResultTeamJoin(out *MessageEnvelope, res *TeamJoin) {
	out.Constructor = C_TeamJoin
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamListMembers int64 = 3107323194

type poolTeamListMembers struct {
	pool sync.Pool
}

func (p *poolTeamListMembers) Get() *TeamListMembers {
	x, ok := p.pool.Get().(*TeamListMembers)
	if !ok {
		return &TeamListMembers{}
	}
	return x
}

func (p *poolTeamListMembers) Put(x *TeamListMembers) {
	p.pool.Put(x)
}

var PoolTeamListMembers = poolTeamListMembers{}

func ResultTeamListMembers(out *MessageEnvelope, res *TeamListMembers) {
	out.Constructor = C_TeamListMembers
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamEdit int64 = 3481894956

type poolTeamEdit struct {
	pool sync.Pool
}

func (p *poolTeamEdit) Get() *TeamEdit {
	x, ok := p.pool.Get().(*TeamEdit)
	if !ok {
		return &TeamEdit{}
	}
	return x
}

func (p *poolTeamEdit) Put(x *TeamEdit) {
	p.pool.Put(x)
}

var PoolTeamEdit = poolTeamEdit{}

func ResultTeamEdit(out *MessageEnvelope, res *TeamEdit) {
	out.Constructor = C_TeamEdit
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamUploadPhoto int64 = 1595699082

type poolTeamUploadPhoto struct {
	pool sync.Pool
}

func (p *poolTeamUploadPhoto) Get() *TeamUploadPhoto {
	x, ok := p.pool.Get().(*TeamUploadPhoto)
	if !ok {
		return &TeamUploadPhoto{}
	}
	return x
}

func (p *poolTeamUploadPhoto) Put(x *TeamUploadPhoto) {
	p.pool.Put(x)
}

var PoolTeamUploadPhoto = poolTeamUploadPhoto{}

func ResultTeamUploadPhoto(out *MessageEnvelope, res *TeamUploadPhoto) {
	out.Constructor = C_TeamUploadPhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamRemovePhoto int64 = 3388888323

type poolTeamRemovePhoto struct {
	pool sync.Pool
}

func (p *poolTeamRemovePhoto) Get() *TeamRemovePhoto {
	x, ok := p.pool.Get().(*TeamRemovePhoto)
	if !ok {
		return &TeamRemovePhoto{}
	}
	return x
}

func (p *poolTeamRemovePhoto) Put(x *TeamRemovePhoto) {
	p.pool.Put(x)
}

var PoolTeamRemovePhoto = poolTeamRemovePhoto{}

func ResultTeamRemovePhoto(out *MessageEnvelope, res *TeamRemovePhoto) {
	out.Constructor = C_TeamRemovePhoto
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamMembers int64 = 2208941294

type poolTeamMembers struct {
	pool sync.Pool
}

func (p *poolTeamMembers) Get() *TeamMembers {
	x, ok := p.pool.Get().(*TeamMembers)
	if !ok {
		return &TeamMembers{}
	}
	return x
}

func (p *poolTeamMembers) Put(x *TeamMembers) {
	x.Members = x.Members[:0]
	x.Users = x.Users[:0]
	p.pool.Put(x)
}

var PoolTeamMembers = poolTeamMembers{}

func ResultTeamMembers(out *MessageEnvelope, res *TeamMembers) {
	out.Constructor = C_TeamMembers
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamMember int64 = 1965775170

type poolTeamMember struct {
	pool sync.Pool
}

func (p *poolTeamMember) Get() *TeamMember {
	x, ok := p.pool.Get().(*TeamMember)
	if !ok {
		return &TeamMember{}
	}
	return x
}

func (p *poolTeamMember) Put(x *TeamMember) {
	p.pool.Put(x)
}

var PoolTeamMember = poolTeamMember{}

func ResultTeamMember(out *MessageEnvelope, res *TeamMember) {
	out.Constructor = C_TeamMember
	protoSize := res.Size()
	if protoSize > cap(out.Message) {
		pbytes.Put(out.Message)
		out.Message = pbytes.GetLen(protoSize)
	} else {
		out.Message = out.Message[:protoSize]
	}
	res.MarshalToSizedBuffer(out.Message)
}

const C_TeamsMany int64 = 2225718663

type poolTeamsMany struct {
	pool sync.Pool
}

func (p *poolTeamsMany) Get() *TeamsMany {
	x, ok := p.pool.Get().(*TeamsMany)
	if !ok {
		return &TeamsMany{}
	}
	return x
}

func (p *poolTeamsMany) Put(x *TeamsMany) {
	x.Teams = x.Teams[:0]
	x.Users = x.Users[:0]
	p.pool.Put(x)
}

var PoolTeamsMany = poolTeamsMany{}

func ResultTeamsMany(out *MessageEnvelope, res *TeamsMany) {
	out.Constructor = C_TeamsMany
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
	ConstructorNames[1172720786] = "TeamGet"
	ConstructorNames[3889056091] = "TeamAddMember"
	ConstructorNames[4200364613] = "TeamRemoveMember"
	ConstructorNames[382328820] = "TeamPromote"
	ConstructorNames[2331393294] = "TeamDemote"
	ConstructorNames[1413785879] = "TeamLeave"
	ConstructorNames[1725794017] = "TeamJoin"
	ConstructorNames[3107323194] = "TeamListMembers"
	ConstructorNames[3481894956] = "TeamEdit"
	ConstructorNames[1595699082] = "TeamUploadPhoto"
	ConstructorNames[3388888323] = "TeamRemovePhoto"
	ConstructorNames[2208941294] = "TeamMembers"
	ConstructorNames[1965775170] = "TeamMember"
	ConstructorNames[2225718663] = "TeamsMany"
}
