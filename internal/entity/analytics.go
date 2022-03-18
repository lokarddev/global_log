package entity

import "github.com/jackc/pgtype"

type Analytics struct {
	Id            pgtype.Int8        `json:"id" db:"id"`
	CreatedAt     pgtype.Timestamptz `json:"created_at" db:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at" db:"updated_at"`
	MetricsTypeId pgtype.Int4        `json:"metrics_type_id" db:"metrics_type_id"`
	Payload       pgtype.JSONB       `json:"payload" db:"payload"`
	EventId       pgtype.Int8        `json:"event_id" db:"event_id"`
}

type MetricsType struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Code        pgtype.Varchar `json:"code" db:"code"`
	Value       pgtype.Varchar `json:"value" db:"value"`
	Description pgtype.Varchar `json:"description" db:"description"`
}
