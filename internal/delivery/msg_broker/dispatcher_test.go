package broker

import (
	"context"
	"encoding/json"
	"github.com/lokarddev/global_log/internal/delivery/mocks"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/pkg/env"
	"github.com/lokarddev/global_log/pkg/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestNewDispatcher(t *testing.T) {
	broker := &mocks.MsgBrokerMock{}

	t.Run("Valid", func(t *testing.T) {
		NewDispatcher(broker, &logger.LoggerMock{})
	})
}

func TestDispatcher_Attach(t *testing.T) {
	broker := &mocks.MsgBrokerMock{}

	t.Run("Valid", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		dispatcher.Attach("test-topic", &mocks.BrokerHandlerMock{})

		assert.Equal(t, &mocks.BrokerHandlerMock{}, dispatcher.handlers["test-topic"])
	})
}

func TestDispatcher_RunBrokerListening(t *testing.T) {
	broker := &mocks.MsgBrokerMock{}

	t.Run("Valid", func(t *testing.T) {
		err := os.Setenv("INCOME_CHAN", "test_chan")
		err = os.Setenv("WORKERS_COUNT", "1")
		env.LoggerChan = os.Getenv("INCOME_CHAN")
		env.WorkersCount, err = strconv.Atoi(os.Getenv("WORKERS_COUNT"))
		assert.NoError(t, err)
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2)
		defer cancel()
		dispatcher.RunBrokerListening(ctx)
		os.Clearenv()
	})
	t.Run("Invalid (subscriber error)", func(t *testing.T) {
		err := os.Setenv("WORKERS_COUNT", "1")
		env.LoggerChan = os.Getenv("INCOME_CHAN")
		env.WorkersCount, err = strconv.Atoi(os.Getenv("WORKERS_COUNT"))
		assert.NoError(t, err)
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2)
		defer cancel()
		dispatcher.RunBrokerListening(ctx)
		os.Clearenv()
	})
}

func TestDispatcher_StopBrokerListening(t *testing.T) {

	t.Run("Valid", func(t *testing.T) {
		broker := &mocks.MsgBrokerMock{}
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})

		err := dispatcher.StopBrokerListening()

		assert.NoError(t, err)
	})
	t.Run("Invalid (error closing broker)", func(t *testing.T) {
		broker := &mocks.MsgBrokerMock{InvalidConfigTrigger: true}
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})

		err := dispatcher.StopBrokerListening()

		assert.Error(t, err)
	})
}

func TestDispatcher_RunWorker(t *testing.T) {
	broker := &mocks.MsgBrokerMock{}

	t.Run("Valid (interruption error)", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
		defer cancel()
		payloadChan := make(chan []byte, 5)

		dispatcher.runWorker(ctx, payloadChan)
	})
	t.Run("Valid Payload handled", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
		defer cancel()
		payloadChan := make(chan []byte, 5)
		defer close(payloadChan)
		msg, err := json.Marshal(entity.IncomeMessage{
			MsgId:   "test-msg",
			TopicId: "test-topic",
			Payload: json.RawMessage{},
		})
		assert.NoError(t, err)
		payloadChan <- msg

		dispatcher.runWorker(ctx, payloadChan)
	})
	t.Run("Invalid publish error", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
		defer cancel()
		payloadChan := make(chan []byte, 5)
		defer close(payloadChan)
		msg, err := json.Marshal(entity.IncomeMessage{
			MsgId:   "",
			TopicId: "test-topic",
			Payload: json.RawMessage{},
		})
		assert.NoError(t, err)
		payloadChan <- msg

		dispatcher.runWorker(ctx, payloadChan)
	})
}

func TestDispatcher_CallHandler(t *testing.T) {
	broker := &mocks.MsgBrokerMock{}

	t.Run("Valid", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		dispatcher.Attach("test-topic", &mocks.BrokerHandlerMock{})
		payload := entity.IncomeMessage{
			MsgId:   "test-msg-id",
			TopicId: "test-topic",
			Payload: []byte(`"some": "data"`),
		}

		result := dispatcher.callHandler(payload)

		assert.NotNil(t, result)
		assert.IsType(t, entity.ResultMessage{}, result)
		assert.Equal(t, entity.ResultMessage{MsgId: "test-msg-id", Success: true, Error: ""}, result)
	})
	t.Run("Valid", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		payload := entity.IncomeMessage{
			MsgId:   "test-msg-id",
			TopicId: "test-topic",
			Payload: []byte(`"some": "data"`),
		}

		result := dispatcher.callHandler(payload)

		assert.NotNil(t, result)
		assert.IsType(t, entity.ResultMessage{}, result)
		assert.Equal(t, false, result.Success)
		assert.NotEqual(t, "", result.Error)
	})
	t.Run("Valid", func(t *testing.T) {
		dispatcher := NewDispatcher(broker, &logger.LoggerMock{})
		dispatcher.Attach("test-topic", &mocks.BrokerHandlerMock{})
		payload := entity.IncomeMessage{
			MsgId:   "test-msg-id",
			TopicId: "test-topic",
			Payload: json.RawMessage{},
		}

		result := dispatcher.callHandler(payload)

		assert.NotNil(t, result)
		assert.IsType(t, entity.ResultMessage{}, result)
		assert.Equal(t, false, result.Success)
		assert.NotEqual(t, "", result.Error)
	})
}
