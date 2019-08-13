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
)
