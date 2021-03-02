package tmp

import (
	"github.com/go-redis/redis/v8"
	"time"
)

func NewClient(c *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		DB:           c.Db,
		Username:     c.Username,
		Password:     c.Password,
		PoolSize:     c.PoolSize,
		DialTimeout:  time.Duration(c.DialTimeout),
		ReadTimeout:  time.Duration(c.ReadTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
	})
	if c.OpenTrace {
		rdb.AddHook(TracingHook{})
	}

	return rdb
}
