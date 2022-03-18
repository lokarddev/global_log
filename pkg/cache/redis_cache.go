package cache

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/lokarddev/global_log/pkg/env"
	"time"
)

type CachingInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, obj []byte) error
	Del(key string) error
	Push(key string, val []byte) error
	Range(key string) ([]string, error)
	PublishToAllWsCons(val interface{}) error
	NewWsConnection(key string, conn *websocket.Conn)
	DropWsConnection(connId string)
}

type RedisCache struct {
	client  *redis.Client
	ttl     time.Duration
	wsConns map[string]*websocket.Conn
}

func NewRedisCache(c *redis.Client) *RedisCache {
	wsSet := make(map[string]*websocket.Conn)
	return &RedisCache{client: c, ttl: time.Minute * time.Duration(env.CacheTTL), wsConns: wsSet}
}

func (c *RedisCache) Get(key string) ([]byte, error) {
	r, err := c.client.Get(key).Bytes()
	return r, err
}

func (c *RedisCache) Set(key string, obj []byte) error {
	return c.client.Set(key, obj, c.ttl).Err()
}

func (c *RedisCache) Del(key string) error {
	return c.client.Del(key).Err()
}

func (c *RedisCache) Push(key string, val []byte) error {
	return c.client.RPush(key, val).Err()
}

func (c *RedisCache) Range(key string) ([]string, error) {
	val, err := c.client.LRange(key, 0, -1).Result()
	return val, err
}

func (c *RedisCache) NewWsConnection(key string, conn *websocket.Conn) {
	c.wsConns[key] = conn
}

func (c *RedisCache) PublishToAllWsCons(val interface{}) error {
	for _, conn := range c.wsConns {
		if err := conn.WriteJSON(val); err != nil {
			return err
		}
	}
	return nil
}

func (c *RedisCache) DropWsConnection(connId string) {
	delete(c.wsConns, connId)
}
