package entity

import "encoding/json"

type ResultMessage struct {
	MsgId   string `json:"msg_id,omitempty"`
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}

type IncomeMessage struct {
	MsgId   string          `json:"msg_id,omitempty"`
	TopicId string          `json:"topic_id,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type LogMsg struct {
	Event   string `json:"event,omitempty"`
	Source  string `json:"source,omitempty"`
	Level   string `json:"level,omitempty"`
	Payload string `json:"payload,omitempty"`
}
