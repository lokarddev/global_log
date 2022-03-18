package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type PostgresLog struct {
	Id         pgtype.Int8        `json:"id" db:"id"`
	CreatedAt  pgtype.Timestamptz `json:"created_at" db:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at" db:"updated_at"`
	LogLevelId pgtype.Int4        `json:"log_level_id" db:"log_level_id"`
	Payload    pgtype.Varchar     `json:"payload" db:"payload"`
	EventId    pgtype.Int8        `json:"event_id" db:"event_id"`
}

func (e *PostgresLog) ToCLeanLog() CleanLog {
	return CleanLog{
		Id:         e.Id.Int,
		CreatedAt:  e.CreatedAt.Time,
		UpdatedAt:  e.UpdatedAt.Time,
		LogLevelId: e.LogLevelId.Int,
		Payload:    e.Payload.String,
		EventId:    e.EventId.Int,
	}
}

type CleanLog struct {
	Id         int64     `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	LogLevelId int32     `json:"log_level_id" db:"log_level_id"`
	Payload    string    `json:"payload" db:"payload"`
	EventId    int64     `json:"event_id" db:"event_id"`
}

type LogLevel struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Code        pgtype.Varchar `json:"code" db:"code"`
	Value       pgtype.Varchar `json:"value" db:"value"`
	Description pgtype.Varchar `json:"description" db:"description"`
}
