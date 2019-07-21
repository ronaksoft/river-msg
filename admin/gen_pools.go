package admin

import (
	"sync"
)

var (
	PoolUser = sync.Pool{
		New: func() interface{} {
			m := new(User)
			return m
		},
	}
)
