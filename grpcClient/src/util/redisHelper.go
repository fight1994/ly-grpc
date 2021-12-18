package util

import (
	"github.com/monnand/goredis"
)

type RedisHelper struct {
	Client *goredis.Client
}

func (r *RedisHelper) ConnRedis() {
	if r.Client == nil {
		var client *goredis.Client
		client.Addr = "127.0.0.1:6379"
		r.Client = client
	}
}
