package cache_test

import (
	"testing"

	"github.com/avila-r/social/cache"
)

func Test_RedisConn(t *testing.T) {
	c := cache.Client

	if err := c.Verify(); err != nil {
		t.Error(err.Error())
	}
}
