package broker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/pkg/env"
	"github.com/lokarddev/global_log/pkg/logger"
	"log"
)

var (
	ctxInterruptionErr = errors.New("context interruption reached")
)

type DispatcherInterface interface {
	RunBrokerListening(ctx context.Context)
	StopBrokerListening() error
	Attach(topic string, handler HandlerInterface)
}

type HandlerInterface interface {
	ProcessTask(context []byte) (bool, error)
}

type MsgBrokerInterface interface {
	Subscribe(ctx context.Context, publisherChan string, ch chan []byte) error
	Publish(subscriberChan string, result entity.ResultMessage) error
	Stop() error
}

type Dispatcher struct {
	msgClient MsgBrokerInterface
	handlers  map[string]HandlerInterface
	logger    logger.LoggerInterface
}

func NewDispatcher(broker MsgBrokerInterface, logger logger.LoggerInterface) *Dispatcher {
	return &Dispatcher{msgClient: broker, handlers: make(map[string]HandlerInterface), logger: logger}
}

func (d *Dispatcher) RunBrokerListening(ctx context.Context) {
	payloadChan := make(chan []byte, 5)
	defer close(payloadChan)
	for i := 0; i < env.WorkersCount; i++ {
		go d.runWorker(ctx, payloadChan)
	}
	if err := d.msgClient.Subscribe(ctx, env.LoggerChan, payloadChan); err != nil {
		log.Println(err)
	}
}

func (d *Dispatcher) runWorker(ctx context.Context, ch chan []byte) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctxInterruptionErr)
			return
		case payload := <-ch:
			var incomeMsg entity.IncomeMessage
			if err := json.Unmarshal(payload, &incomeMsg); err != nil {
				log.Println(err)
				continue
			}
			log.Println(fmt.Sprintf("Run: %s -- %s -- %s", incomeMsg.MsgId, incomeMsg.TopicId, incomeMsg.Payload))
			_ = d.callHandler(incomeMsg)
		}
	}
}

func (d *Dispatcher) callHandler(message entity.IncomeMessage) entity.ResultMessage {
	handler, ok := d.handlers[message.TopicId]
	switch {
	case !ok:
		err := errors.New(fmt.Sprintf("no handler provided! Topic -- %s. MsgId -- %s", message.TopicId, message.MsgId))
		log.Println(err)
		return entity.ResultMessage{MsgId: message.MsgId, Success: false, Error: err.Error()}
	default:
		success, err := handler.ProcessTask(message.Payload)
		if err != nil {
			return entity.ResultMessage{MsgId: message.MsgId, Success: false, Error: err.Error()}
		}
		return entity.ResultMessage{MsgId: message.MsgId, Success: success, Error: ""}
	}
}

func (d *Dispatcher) Attach(topic string, handler HandlerInterface) {
	d.handlers[topic] = handler
}

func (d *Dispatcher) StopBrokerListening() error {
	if err := d.msgClient.Stop(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
