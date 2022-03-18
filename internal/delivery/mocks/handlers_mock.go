package mocks

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

type BrokerHandlerMock struct {
	mock.Mock
}

func (h *BrokerHandlerMock) ProcessTask(context []byte) (bool, error) {
	var empty []byte
	switch {
	case len(context) == len(empty):
		return false, errors.New("test error reached")
	default:
		return true, nil
	}
}
