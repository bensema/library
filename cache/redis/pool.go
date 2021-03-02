package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Pool struct {
	rp *redis.Pool
	hooks
}

func NewPool(c *Config) *Pool {
	_p := &Pool{}
	_p.rp = &redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: time.Duration(c.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", c.Addr,
				redis.DialPassword(c.Password),
				redis.DialDatabase(c.Db),
				redis.DialConnectTimeout(time.Duration(c.DialTimeout)),
				redis.DialReadTimeout(time.Duration(c.ReadTimeout)),
				redis.DialWriteTimeout(time.Duration(c.WriteTimeout)))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	if c.OpenTrace {
		_p.AddHook(OpenTelemetryHook{})
	}
	return _p
}

func (p *Pool) Get() TraceConn {
	conn := p.rp.Get()
	return TraceConn{conn: conn, hs: p.hooks, ctx: context.Background()}
}
func (p *Pool) GetContext(ctx context.Context) (redis.Conn, error) {
	conn, err := p.rp.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	return &TraceConn{ctx: ctx, conn: conn, hs: p.hooks}, err
}
