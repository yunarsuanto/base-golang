package redis

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yunarsuanto/base-go/config"

	"github.com/go-redis/redis/v8"
)

type redisServer struct {
	cfg *config.RedisServer
}

// RedisServerInterface interface for redis server
type RedisServerInterface interface {
	Connect(ctx context.Context) (*redis.Client, error)
}

// NewRedisServer function to connect redisServer to RedisServerInterface
// Params:
// cfg: redis config
// Returns RedisServerInterface
func NewRedisServer(cfg *config.RedisServer) RedisServerInterface {
	return &redisServer{
		cfg: cfg,
	}
}

// Connect function to connect to redis server
func (r *redisServer) Connect(ctx context.Context) (*redis.Client, error) {
	timeout := time.Duration(r.cfg.Timeout) * time.Second
	rdb := redis.NewClient(&redis.Options{
		Addr:        r.cfg.Addr,
		Password:    r.cfg.Password,
		DialTimeout: timeout,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("cannot connect to redis")
		return nil, err
	}

	var additionalMessage string
	if os.Getenv("ENV") == "" {
		additionalMessage = r.cfg.Addr
	}
	log.Printf("success connect to redis %s", additionalMessage)
	return rdb, nil
}
