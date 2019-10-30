package dao

import (
	redis "github.com/go-redis/redis/v7"
)

type Dao struct {
	redis *redis.Client
}

type DaoConfig struct {
	Redis *redis.Options
}

func New(conf *DaoConfig) *Dao {
	return &Dao{
		redis: redis.NewClient(conf.Redis),
	}
}
