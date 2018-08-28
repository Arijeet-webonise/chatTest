package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisServiceImpl implement Redis
type RedisServiceImpl struct {
	Client *redis.Client
}

// RedisService encapsulates redis functions
type RedisService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (*redis.StringCmd, error)
	Delete(key ...string) (int64, error)
}

// InitRedis inicilize Redis Server
func InitRedis(addr, password string, db int) (RedisService, string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	pong, err := client.Ping().Result()
	return &RedisServiceImpl{client}, pong, err
}

// Set save value to redis
func (rs *RedisServiceImpl) Set(key string, value interface{}, expiration time.Duration) error {
	_, err := rs.Client.Set(key, value, expiration).Result()
	return err
}

// Get return data from redis
func (rs *RedisServiceImpl) Get(key string) (*redis.StringCmd, error) {
	result := rs.Client.Get(key)
	if _, err := result.Result(); err != nil {
		return nil, err
	}
	return result, nil
}

// Delete deletes keys from redis
func (rs *RedisServiceImpl) Delete(key ...string) (int64, error) {
	return rs.Client.Del(key...).Result()
}
