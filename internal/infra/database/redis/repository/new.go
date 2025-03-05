package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Iredis struct {
	Db *redis.Client
}

func NewRedisRepository(c *redis.Client) *Iredis {
	return &Iredis{
		Db: c,
	}
}
