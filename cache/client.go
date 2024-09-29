package cache

import (
	"github.com/avila-r/social/config"
	"github.com/avila-r/xgo/pkg/cache"
	"github.com/redis/go-redis/v9"
)

var (
	Client = func() *cache.Client {
		uri := config.RedisProperties.URi

		return cache.NewClient(&redis.Options{
			Addr: uri,
		})
	}()
)
