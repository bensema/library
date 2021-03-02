package tmp

import (
	xtime "github.com/bensema/library/time"
)

type Config struct {
	Name      string // redis name, for trace
	OpenTrace bool

	Addr         string
	Db           int
	Username     string
	Password     string
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration

	PoolSize int
}
