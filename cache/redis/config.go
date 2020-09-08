package redis

import (
	xtime "github.com/bensema/library/time"
)

type Config struct {
	Name string // redis name, for trace

	MaxIdle         int
	MaxActive       int
	IdleTimeout     xtime.Duration
	MaxConnLifetime xtime.Duration

	Addr         string
	Db           int
	Password     string
	DialTimeout  xtime.Duration
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
}
