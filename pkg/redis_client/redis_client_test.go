package redis_client

import (
	"context"
	"encoding/json"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/pkg/env"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
	"time"
)

func newTestRedisClient() (*miniredis.Miniredis, error) {
	mr, err := miniredis.Run()
	err = os.Setenv("REDIS_HOST", mr.Host())
	err = os.Setenv("REDIS_PORT", mr.Port())
	err = os.Setenv("REDIS_USER", "default")
	err = os.Setenv("REDIS_PROTOCOL", "redis_client")
	err = os.Setenv("DEBUG", "true")
	return mr, err
}

func TestNewRedisClient(t *testing.T) {
	mr, err := newTestRedisClient()
	assert.NoError(t, err)
	defer mr.Close()
	defer os.Clearenv()

	t.Run("Valid", func(t *testing.T) {
		err = env.InitEnvVariables()
		assert.NoError(t, err)

		client, err := NewRedisClient()

		assert.Nil(t, err)
		assert.NotNil(t, client)
		assert.IsType(t, &redis.Client{}, client)

		err = client.Close()
		assert.NoError(t, err)
	})
	t.Run("Invalid redis_client config", func(t *testing.T) {
		err = os.Setenv("REDIS_PORT", "invalidPort")
		err = env.InitEnvVariables()
		assert.NoError(t, err)

		client, err := NewRedisClient()

		assert.Error(t, err)
		assert.IsType(t, &redis.Client{}, client)
	})
	t.Run("No redis_client instance running", func(t *testing.T) {
		mr.Close()

		client, err := NewRedisClient()

		assert.Error(t, err)
		assert.IsType(t, &redis.Client{}, client)
	})
}

func TestMsgBusRedis_Publish(t *testing.T) {
	mr, err := newTestRedisClient()
	assert.NoError(t, err)
	err = env.InitEnvVariables()
	defer mr.Close()
	defer os.Clearenv()
	r, err := NewRedisClient()
	msgBus := NewMsgBusRedis(r)
	assert.NoError(t, err)

	t.Run("Valid", func(t *testing.T) {
		busChan := "test_chan"
		result := entity.ResultMessage{
			MsgId:   "test-id",
			Success: false,
			Error:   "some-error",
		}

		err := msgBus.Publish(busChan, result)

		assert.NoError(t, err)
	})
	t.Run("Invalid", func(t *testing.T) {
		_ = os.Setenv("REDIS_PORT", "1")
		_ = env.InitEnvVariables()

		r, err = NewRedisClient()
		bus := NewMsgBusRedis(r)
		busChan := "test_chan"
		result := entity.ResultMessage{
			MsgId:   "test-id",
			Success: false,
			Error:   "some-error",
		}

		err = bus.Publish(busChan, result)
		assert.Error(t, err)
	})
}

func TestMsgBusRedis_Subscribe(t *testing.T) {
	mr, err := newTestRedisClient()
	assert.NoError(t, err)
	err = env.InitEnvVariables()
	defer mr.Close()
	defer os.Clearenv()
	r, err := NewRedisClient()
	msgBus := NewMsgBusRedis(r)
	assert.NoError(t, err)

	t.Run("Valid with payload", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
		busChan := "test_chan"
		payloadChan := make(chan []byte, 5)
		defer cancel()
		defer close(payloadChan)
		msg, err := json.Marshal(struct{ Payload string }{Payload: "some_value"})
		assert.NoError(t, err)
		var expectedErr error

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			expectedErr = msgBus.Subscribe(ctx, busChan, payloadChan)
			wg.Done()
		}(&wg)
		_ = msgBus.client.Publish(busChan, msg).Err()
		wg.Wait()

		assert.NoError(t, expectedErr)
	})
	t.Run("Valid without message (context cancellation)", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
		busChan := "test_chan"
		payloadChan := make(chan []byte, 5)
		defer cancel()
		defer close(payloadChan)

		err = msgBus.Subscribe(ctx, busChan, payloadChan)

		assert.NoError(t, err)
	})
}

func TestMsgBusRedis_Stop(t *testing.T) {
	mr, err := newTestRedisClient()
	assert.NoError(t, err)
	err = env.InitEnvVariables()
	defer mr.Close()
	defer os.Clearenv()

	t.Run("Valid", func(t *testing.T) {
		r, err := NewRedisClient()
		msgBus := NewMsgBusRedis(r)
		assert.NoError(t, err)

		err = msgBus.Stop()

		assert.NoError(t, err)
	})
}

func TestNewMsgBusRedis(t *testing.T) {
	mr, err := newTestRedisClient()
	assert.NoError(t, err)
	defer mr.Close()
	defer os.Clearenv()

	t.Run("Valid", func(t *testing.T) {
		err = env.InitEnvVariables()
		r, err := NewRedisClient()
		msgBus := NewMsgBusRedis(r)

		assert.NoError(t, err)
		assert.IsType(t, &MsgBusRedis{}, msgBus)
	})
	t.Run("With redis_client client error", func(t *testing.T) {
		err = os.Setenv("REDIS_PORT", "invalidPort")
		err = env.InitEnvVariables()

		r, err := NewRedisClient()
		msgBus := NewMsgBusRedis(r)

		assert.Error(t, err)
		assert.IsType(t, &MsgBusRedis{}, msgBus)
	})
}
