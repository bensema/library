package redis

import (
	"time"
)

type Config struct {
	Name string // redis name, for trace

	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration

	Addr         string
	Db           int
	Password     string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func e() {
	p := NewPool(&Config{})
	conn := p.rp.Get()
	defer conn.Close()
	conn.Send()
}
