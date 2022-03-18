package logger

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/lokarddev/global_log/internal/entity"
	"log"
	"os"
)

// BASE log event types
const (
	LogEvent       = "log"
	AnalyticsEvent = "analytics"
)

// BASE log sources
const (
	Calculate = "go_calculate"
	GoLog     = "go_log"
)

// BASE log levels
const (
	InfoLvl     = "info"
	DebugLvl    = "debug"
	WarningLvl  = "warning"
	CriticalLvl = "critical"
	ErrorLvl    = "error"
)

// BASE command topic levels
const (
	InfoTopic     = "infoLog"
	DebugTopic    = "debugLog"
	ErrorTopic    = "errorLog"
	WarningTopic  = "warningLog"
	CriticalTopic = "criticalLog"
)

var (
	logInfo     = infoLogger()
	logErr      = errorLogger()
	logDebug    = debugLogger()
	logWarning  = warningLogger()
	logCritical = criticalLogger()
)

func infoLogger() *log.Logger {
	return log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Llongfile)
}

func errorLogger() *log.Logger {
	return log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Llongfile)
}

func debugLogger() *log.Logger {
	return log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Llongfile)
}

func warningLogger() *log.Logger {
	return log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Llongfile)
}

func criticalLogger() *log.Logger {
	return log.New(os.Stdout, "CRITICAL ", log.Ldate|log.Ltime|log.Llongfile)
}

type LoggerInterface interface {
	Info(payload string)
	Error(payload string)
	Warning(payload string)
	Critical(payload string)
	Debug(payload string)
}

type Logger struct {
	source           string
	publisher        *redis.Client
	msgBrokerChannel string
}

type LoggerConfig struct {
	Source    string
	MsgClient *redis.Client
}

func NewLogger(cfg LoggerConfig) *Logger {
	if ch := os.Getenv("LOGGER_CHAN"); ch == "" {
		log.Fatalf("logger channel for msgBroker is not provided! Create `LOGGER_CHAN` OS env!")
	}

	return &Logger{
		source:           cfg.Source,
		publisher:        cfg.MsgClient,
		msgBrokerChannel: os.Getenv("LOGGER_CHAN")}
}

func (l *Logger) Info(payload string) {
	_ = logInfo.Output(2, payload)
	l.pushLog(InfoLvl, InfoTopic, payload)
}

func (l *Logger) Error(payload string) {
	_ = logErr.Output(2, payload)
	l.pushLog(ErrorLvl, ErrorTopic, payload)
}

func (l *Logger) Warning(payload string) {
	_ = logWarning.Output(2, payload)
	l.pushLog(WarningLvl, WarningTopic, payload)
}

func (l *Logger) Critical(payload string) {
	_ = logCritical.Output(2, payload)
	l.pushLog(CriticalLvl, CriticalTopic, payload)
}

func (l *Logger) Debug(payload string) {
	_ = logDebug.Output(2, payload)
	l.pushLog(DebugLvl, DebugTopic, payload)
}

func (l *Logger) pushLog(lvl, topic, payload string) {
	msgId, err := uuid.NewV4()
	loggerMessage := entity.LogMsg{
		Event:   LogEvent,
		Source:  l.source,
		Level:   lvl,
		Payload: payload,
	}
	bytes, err := json.Marshal(loggerMessage)
	msg := entity.IncomeMessage{
		MsgId:   msgId.String(),
		TopicId: topic,
		Payload: bytes,
	}
	res, _ := json.Marshal(msg)
	err = l.publisher.Publish(l.msgBrokerChannel, res).Err()
	if err != nil {
		fmt.Println(err)
	}
}
