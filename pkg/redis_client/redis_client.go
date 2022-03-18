package redis_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/pkg/env"
	"log"
)

type CfgRedis struct {
	DB       string
	Host     string
	Port     string
	Username string
	Protocol string
	Password string
}

type MsgBusRedis struct {
	client *redis.Client
}

func NewMsgBusRedis(client *redis.Client) *MsgBusRedis {
	return &MsgBusRedis{client: client}
}

func NewRedisClient() (*redis.Client, error) {
	cfg := &CfgRedis{
		DB:       env.RedisDb,
		Host:     env.RedisHost,
		Port:     env.RedisPort,
		Username: env.RedisUser,
		Protocol: env.RedisProtocol,
		Password: env.RedisPass,
	}
	options, err := redis.ParseURL(fmt.Sprintf("%s://%s:%s@%s:%s/%+v", cfg.Protocol, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB))
	if err != nil {
		log.Println(fmt.Sprintf("ERROR INITIALIZING REDIS: %s", err.Error()))
		return &redis.Client{}, err
	}
	return redis.NewClient(options), nil
}

func (b *MsgBusRedis) Subscribe(ctx context.Context, publisherChan string, ch chan []byte) error {
	sub := b.client.Subscribe(publisherChan)
	msgChan := sub.Channel()
	log.Println(fmt.Sprintf("Started redis_client bus listening on channel %s", publisherChan))
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgChan:
			ch <- []byte(msg.Payload)
		}
	}
}

func (b *MsgBusRedis) Publish(subscriberChan string, result entity.ResultMessage) error {
	bytes, err := json.Marshal(result)
	if err = b.client.Publish(subscriberChan, bytes).Err(); err != nil {
		log.Println(err)
	}
	return err
}

func (b *MsgBusRedis) Stop() error {
	err := b.client.Close()
	return err
}
