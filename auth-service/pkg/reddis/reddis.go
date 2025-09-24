package reddis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_valueTTL = 5 * time.Minute
)

type Reddis struct {
	client *redis.Client
}

func New(addr string, passwrod string, db int) *Reddis {
	rdb := &Reddis{}
	rdb.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwrod,
		DB:       db,
	})
	return rdb
}

func (r *Reddis) Set(ctx context.Context, key string, value interface{}) error {
	err := r.client.Set(ctx, key, value, _valueTTL).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Reddis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Close(rdb *Reddis) {
	if rdb.client != nil {
		rdb.client.Close()
	}
}
