package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type redisClient interface {
	Subscribe(channels ...string) *redis.PubSub
	Publish(channel string, message interface{}) *redis.IntCmd
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	ZAdd(key string, members ...redis.Z) *redis.IntCmd
	ZRem(key string, members ...interface{}) *redis.IntCmd
	ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd
	HMGet(key string, fields ...string) *redis.SliceCmd
	HMSet(key string, fields map[string]interface{}) *redis.StatusCmd
	Exists(keys ...string) *redis.IntCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	Del(keys ...string) *redis.IntCmd
	ZCount(key, min, max string) *redis.IntCmd
	HGet(key, field string) *redis.StringCmd
	HSet(key string, field string, values interface{}) *redis.BoolCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
}

// Pool is struct to handle all rooms and clients in the server
type Pool struct {
	// Reference to main Redis connection
	rdb redisClient

	// DB pool
	db *pgxpool.Pool
}

var singletonPool *Pool
var once sync.Once

// GetPool will return single instance of pool that manage every clients
func GetPool() *Pool {
	once.Do(func() {
		var rdb redisClient
		redisMode := os.Getenv("REDIS_MODE")
		redisEndpoint := os.Getenv("REDIS_ENDPOINT")

		if redisMode == "single" {
			rdb = redis.NewClient(&redis.Options{
				Addr: redisEndpoint,
			})
		} else {
			rdb = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs: []string{redisEndpoint},
			})
		}

		db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}

		singletonPool = &Pool{
			rdb: rdb,
			db:  db,
		}
	})

	return singletonPool
}

func (pool *Pool) Start() {
	log.Println("Starting pool...")
}