package internalMsg

import (
	"sync"
)

var (
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
