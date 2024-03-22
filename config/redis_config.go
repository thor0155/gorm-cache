package config

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

type RedisConfigMode int

const (
	RedisConfigModeOptions RedisConfigMode = 0
	RedisConfigModeRaw     RedisConfigMode = 1
)

type RedisConfig struct {
	Mode RedisConfigMode

	Options *redis.Options
	Client  redis.UniversalClient

	once sync.Once
}

func (c *RedisConfig) InitClient() redis.UniversalClient {
	c.once.Do(func() {
		if c.Mode == RedisConfigModeOptions {
			c.Client = redis.NewClient(c.Options)
		}
	})
	return c.Client
}
