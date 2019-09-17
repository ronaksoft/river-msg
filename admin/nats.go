package admin

import (
	"time"
)

const (
	UpsertUserSubj        = "upsert.user"
	UpsertUserGroup       = "upsert.user"
	UpsertUserDurName     = "upsert.user"
	UpsertUserAckWait     = 5 * time.Second
	UpsertUserMaxInflight = 50

	NewWelcMessageSubj = "new.welc.msg"

	GetWelcMessageSubj  = "get.welc.msg"
	GetWelcMessageGroup = "get.welc.msg"

	NewNotifTemplateSubj = "new.notif.tmpl"

	GetNotifTemplatesSubj  = "get.notif.tmpls"
	GetNotifTemplatesGroup = "get.notif.tmpls"

	GetPushProvSubj  = "get.push.prov"
	GetPushProvGroup = "get.push.prov"

	NewRiverVersionSubj = "new.river.vers"

	GetRiverVersionsSubj = "get.river.vers"
	GetRiverVersionsGroup = "get.river.vers"
)
