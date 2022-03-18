package logger

import "github.com/stretchr/testify/mock"

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) Info(_ string) {

}

func (l *LoggerMock) Error(_ string) {

}

func (l *LoggerMock) Warning(_ string) {

}

func (l *LoggerMock) Critical(_ string) {

}

func (l *LoggerMock) Debug(_ string) {

}
