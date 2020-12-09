package msg

import (
	registry "github.com/ronaksoft/rony/registry"
	sync "sync"
)

const C_ReplyKeyboardMarkup int64 = 3207405102

type poolReplyKeyboardMarkup struct {
	pool sync.Pool
}

func (p *poolReplyKeyboardMarkup) Get() *ReplyKeyboardMarkup {
	x, ok := p.pool.Get().(*ReplyKeyboardMarkup)
	if !ok {
		return &ReplyKeyboardMarkup{}
	}
	return x
}

func (p *poolReplyKeyboardMarkup) Put(x *ReplyKeyboardMarkup) {
	x.SingleUse = false
	x.Selective = false
	x.Resize = false
	x.Rows = x.Rows[:0]
	p.pool.Put(x)
}

var PoolReplyKeyboardMarkup = poolReplyKeyboardMarkup{}

const C_ReplyInlineMarkup int64 = 2436413989

type poolReplyInlineMarkup struct {
	pool sync.Pool
}

func (p *poolReplyInlineMarkup) Get() *ReplyInlineMarkup {
	x, ok := p.pool.Get().(*ReplyInlineMarkup)
	if !ok {
		return &ReplyInlineMarkup{}
	}
	return x
}

func (p *poolReplyInlineMarkup) Put(x *ReplyInlineMarkup) {
	x.Rows = x.Rows[:0]
	p.pool.Put(x)
}

var PoolReplyInlineMarkup = poolReplyInlineMarkup{}

const C_ReplyKeyboardHide int64 = 3134153306

type poolReplyKeyboardHide struct {
	pool sync.Pool
}

func (p *poolReplyKeyboardHide) Get() *ReplyKeyboardHide {
	x, ok := p.pool.Get().(*ReplyKeyboardHide)
	if !ok {
		return &ReplyKeyboardHide{}
	}
	return x
}

func (p *poolReplyKeyboardHide) Put(x *ReplyKeyboardHide) {
	x.Selective = false
	p.pool.Put(x)
}

var PoolReplyKeyboardHide = poolReplyKeyboardHide{}

const C_ReplyKeyboardForceReply int64 = 258469686

type poolReplyKeyboardForceReply struct {
	pool sync.Pool
}

func (p *poolReplyKeyboardForceReply) Get() *ReplyKeyboardForceReply {
	x, ok := p.pool.Get().(*ReplyKeyboardForceReply)
	if !ok {
		return &ReplyKeyboardForceReply{}
	}
	return x
}

func (p *poolReplyKeyboardForceReply) Put(x *ReplyKeyboardForceReply) {
	x.SingleUse = false
	x.Selective = false
	p.pool.Put(x)
}

var PoolReplyKeyboardForceReply = poolReplyKeyboardForceReply{}

const C_KeyboardButtonRow int64 = 2222403758

type poolKeyboardButtonRow struct {
	pool sync.Pool
}

func (p *poolKeyboardButtonRow) Get() *KeyboardButtonRow {
	x, ok := p.pool.Get().(*KeyboardButtonRow)
	if !ok {
		return &KeyboardButtonRow{}
	}
	return x
}

func (p *poolKeyboardButtonRow) Put(x *KeyboardButtonRow) {
	x.Buttons = x.Buttons[:0]
	p.pool.Put(x)
}

var PoolKeyboardButtonRow = poolKeyboardButtonRow{}

const C_KeyboardButtonEnvelope int64 = 2639543624

type poolKeyboardButtonEnvelope struct {
	pool sync.Pool
}

func (p *poolKeyboardButtonEnvelope) Get() *KeyboardButtonEnvelope {
	x, ok := p.pool.Get().(*KeyboardButtonEnvelope)
	if !ok {
		return &KeyboardButtonEnvelope{}
	}
	return x
}

func (p *poolKeyboardButtonEnvelope) Put(x *KeyboardButtonEnvelope) {
	x.Constructor = 0
	x.Data = x.Data[:0]
	p.pool.Put(x)
}

var PoolKeyboardButtonEnvelope = poolKeyboardButtonEnvelope{}

const C_Button int64 = 1034594571

type poolButton struct {
	pool sync.Pool
}

func (p *poolButton) Get() *Button {
	x, ok := p.pool.Get().(*Button)
	if !ok {
		return &Button{}
	}
	return x
}

func (p *poolButton) Put(x *Button) {
	x.Text = ""
	p.pool.Put(x)
}

var PoolButton = poolButton{}

const C_ButtonUrl int64 = 2309530052

type poolButtonUrl struct {
	pool sync.Pool
}

func (p *poolButtonUrl) Get() *ButtonUrl {
	x, ok := p.pool.Get().(*ButtonUrl)
	if !ok {
		return &ButtonUrl{}
	}
	return x
}

func (p *poolButtonUrl) Put(x *ButtonUrl) {
	x.Text = ""
	x.Url = ""
	p.pool.Put(x)
}

var PoolButtonUrl = poolButtonUrl{}

const C_ButtonCallback int64 = 4007711268

type poolButtonCallback struct {
	pool sync.Pool
}

func (p *poolButtonCallback) Get() *ButtonCallback {
	x, ok := p.pool.Get().(*ButtonCallback)
	if !ok {
		return &ButtonCallback{}
	}
	return x
}

func (p *poolButtonCallback) Put(x *ButtonCallback) {
	x.Text = ""
	x.Data = x.Data[:0]
	p.pool.Put(x)
}

var PoolButtonCallback = poolButtonCallback{}

const C_ButtonRequestPhone int64 = 630958494

type poolButtonRequestPhone struct {
	pool sync.Pool
}

func (p *poolButtonRequestPhone) Get() *ButtonRequestPhone {
	x, ok := p.pool.Get().(*ButtonRequestPhone)
	if !ok {
		return &ButtonRequestPhone{}
	}
	return x
}

func (p *poolButtonRequestPhone) Put(x *ButtonRequestPhone) {
	x.Text = ""
	p.pool.Put(x)
}

var PoolButtonRequestPhone = poolButtonRequestPhone{}

const C_ButtonRequestGeoLocation int64 = 323515934

type poolButtonRequestGeoLocation struct {
	pool sync.Pool
}

func (p *poolButtonRequestGeoLocation) Get() *ButtonRequestGeoLocation {
	x, ok := p.pool.Get().(*ButtonRequestGeoLocation)
	if !ok {
		return &ButtonRequestGeoLocation{}
	}
	return x
}

func (p *poolButtonRequestGeoLocation) Put(x *ButtonRequestGeoLocation) {
	x.Text = ""
	p.pool.Put(x)
}

var PoolButtonRequestGeoLocation = poolButtonRequestGeoLocation{}

const C_ButtonSwitchInline int64 = 3842667878

type poolButtonSwitchInline struct {
	pool sync.Pool
}

func (p *poolButtonSwitchInline) Get() *ButtonSwitchInline {
	x, ok := p.pool.Get().(*ButtonSwitchInline)
	if !ok {
		return &ButtonSwitchInline{}
	}
	return x
}

func (p *poolButtonSwitchInline) Put(x *ButtonSwitchInline) {
	x.Text = ""
	x.Query = ""
	x.SamePeer = false
	p.pool.Put(x)
}

var PoolButtonSwitchInline = poolButtonSwitchInline{}

const C_ButtonBuy int64 = 2992465437

type poolButtonBuy struct {
	pool sync.Pool
}

func (p *poolButtonBuy) Get() *ButtonBuy {
	x, ok := p.pool.Get().(*ButtonBuy)
	if !ok {
		return &ButtonBuy{}
	}
	return x
}

func (p *poolButtonBuy) Put(x *ButtonBuy) {
	x.Text = ""
	p.pool.Put(x)
}

var PoolButtonBuy = poolButtonBuy{}

func init() {
	registry.RegisterConstructor(3207405102, "ReplyKeyboardMarkup")
	registry.RegisterConstructor(2436413989, "ReplyInlineMarkup")
	registry.RegisterConstructor(3134153306, "ReplyKeyboardHide")
	registry.RegisterConstructor(258469686, "ReplyKeyboardForceReply")
	registry.RegisterConstructor(2222403758, "KeyboardButtonRow")
	registry.RegisterConstructor(2639543624, "KeyboardButtonEnvelope")
	registry.RegisterConstructor(1034594571, "Button")
	registry.RegisterConstructor(2309530052, "ButtonUrl")
	registry.RegisterConstructor(4007711268, "ButtonCallback")
	registry.RegisterConstructor(630958494, "ButtonRequestPhone")
	registry.RegisterConstructor(323515934, "ButtonRequestGeoLocation")
	registry.RegisterConstructor(3842667878, "ButtonSwitchInline")
	registry.RegisterConstructor(2992465437, "ButtonBuy")
}

func (x *ReplyKeyboardMarkup) DeepCopy(z *ReplyKeyboardMarkup) {
	z.SingleUse = x.SingleUse
	z.Selective = x.Selective
	z.Resize = x.Resize
	for idx := range x.Rows {
		if x.Rows[idx] != nil {
			xx := PoolKeyboardButtonRow.Get()
			x.Rows[idx].DeepCopy(xx)
			z.Rows = append(z.Rows, xx)
		}
	}
}

func (x *ReplyInlineMarkup) DeepCopy(z *ReplyInlineMarkup) {
	for idx := range x.Rows {
		if x.Rows[idx] != nil {
			xx := PoolKeyboardButtonRow.Get()
			x.Rows[idx].DeepCopy(xx)
			z.Rows = append(z.Rows, xx)
		}
	}
}

func (x *ReplyKeyboardHide) DeepCopy(z *ReplyKeyboardHide) {
	z.Selective = x.Selective
}

func (x *ReplyKeyboardForceReply) DeepCopy(z *ReplyKeyboardForceReply) {
	z.SingleUse = x.SingleUse
	z.Selective = x.Selective
}

func (x *KeyboardButtonRow) DeepCopy(z *KeyboardButtonRow) {
	for idx := range x.Buttons {
		if x.Buttons[idx] != nil {
			xx := PoolKeyboardButtonEnvelope.Get()
			x.Buttons[idx].DeepCopy(xx)
			z.Buttons = append(z.Buttons, xx)
		}
	}
}

func (x *KeyboardButtonEnvelope) DeepCopy(z *KeyboardButtonEnvelope) {
	z.Constructor = x.Constructor
	z.Data = append(z.Data[:0], x.Data...)
}

func (x *Button) DeepCopy(z *Button) {
	z.Text = x.Text
}

func (x *ButtonUrl) DeepCopy(z *ButtonUrl) {
	z.Text = x.Text
	z.Url = x.Url
}

func (x *ButtonCallback) DeepCopy(z *ButtonCallback) {
	z.Text = x.Text
	z.Data = append(z.Data[:0], x.Data...)
}

func (x *ButtonRequestPhone) DeepCopy(z *ButtonRequestPhone) {
	z.Text = x.Text
}

func (x *ButtonRequestGeoLocation) DeepCopy(z *ButtonRequestGeoLocation) {
	z.Text = x.Text
}

func (x *ButtonSwitchInline) DeepCopy(z *ButtonSwitchInline) {
	z.Text = x.Text
	z.Query = x.Query
	z.SamePeer = x.SamePeer
}

func (x *ButtonBuy) DeepCopy(z *ButtonBuy) {
	z.Text = x.Text
}
