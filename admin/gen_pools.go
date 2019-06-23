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
	PoolFileLocation = sync.Pool{
		New: func() interface{} {
			m := new(FileLocation)
			return m
		},
	}
)
