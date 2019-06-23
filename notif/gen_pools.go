package notif

import (
	"sync"
)

var (
	PoolCancelRequest = sync.Pool{
		New: func() interface{} {
			m := new(CancelRequest)
			return m
		},
	}
	PoolNotification = sync.Pool{
		New: func() interface{} {
			m := new(Notification)
			return m
		},
	}
	PoolNotificationData = sync.Pool{
		New: func() interface{} {
			m := new(NotificationData)
			return m
		},
	}
)
