package utils

import (
	"sync"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type Cache struct {
	Logger *zap.Logger
	Pool   *redis.Pool
}

type CacheService interface {
	Ping() error
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int64) error
	Exists(key string) (bool, error)
	Delete(key string) error
}

var (
	cacheInstance CacheService
	once          sync.Once
)

func NewCacheService(pool *redis.Pool, logger *zap.Logger) CacheService {
	once.Do(func() {
		cacheInstance = &Cache{Pool: pool, Logger: logger}
	})
	return cacheInstance
}

func CreateRedisPool(addr, password string, maxIdle int, logger *zap.Logger) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle: maxIdle,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				logger.Error("failed to connect to redis", zap.Error(err))
				return nil, err
			}

			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					logger.Error("failed to authenticate to redis", zap.Error(err))
					return nil, err
				}
			}

			return conn, nil
		},
	}

	return pool, nil
}

func (c *Cache) Ping() error {
	conn := c.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		c.Logger.Error("failed to ping redis", zap.Error(err))
		return err
	}

	c.Logger.Info("redis ping successful")
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	conn := c.Pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		c.Logger.Error("failed to get data from redis", zap.Error(err))
		return nil, err
	}

	c.Logger.Info("data retrieved from redis", zap.String("key", key))
	return data, nil
}

func (c *Cache) Set(key string, value []byte, ttl int64) error {
	conn := c.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		c.Logger.Error("failed to set data to redis", zap.Error(err))
		return err
	}

	if ttl > 0 {
		_, err := conn.Do("EXPIRE", key, ttl)
		if err != nil {
			c.Logger.Error("failed to set ttl to redis", zap.Error(err))
			return err
		}
	}

	c.Logger.Info("data set to redis", zap.String("key", key))
	return nil
}

func (c *Cache) Exists(key string) (bool, error) {
	conn := c.Pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		c.Logger.Error("failed to check if key exists in redis", zap.Error(err))
		return false, err
	}

	c.Logger.Info("key exists in redis", zap.String("key", key))
	return exists, nil
}

func (c *Cache) Delete(key string) error {
	conn := c.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		c.Logger.Error("failed to delete key from redis", zap.Error(err))
		return err
	}

	c.Logger.Info("key deleted from redis", zap.String("key", key))
	return nil
}
