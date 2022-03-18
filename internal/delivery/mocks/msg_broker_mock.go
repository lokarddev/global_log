package mocks

import (
	"context"
	"errors"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MsgBrokerMock struct {
	mock.Mock
	InvalidConfigTrigger bool
}

func (b *MsgBrokerMock) Subscribe(_ context.Context, publisherChan string, _ chan []byte) error {
	switch {
	case publisherChan == "":
		return errors.New("test error reached")
	default:
		return nil
	}
}

func (b *MsgBrokerMock) Publish(_ string, result entity.ResultMessage) error {
	switch {
	case result.MsgId == "":
		return errors.New("test error reached")
	default:
		return nil
	}
}

func (b *MsgBrokerMock) Stop() error {
	switch {
	case b.InvalidConfigTrigger:
		return errors.New("test error reached")
	default:
		return nil
	}
}
