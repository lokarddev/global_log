package entity

import "github.com/jackc/pgtype"

type Event struct {
	Id          pgtype.Int8        `json:"id" db:"id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at" db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at" db:"updated_at"`
	EventTypeId pgtype.Int4        `json:"event_type_id" db:"event_type_id"`
	SourceId    pgtype.Int4        `json:"source_id" db:"source_id"`
}

type EventType struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Code        pgtype.Varchar `json:"code" db:"code"`
	Value       pgtype.Varchar `json:"value" db:"value"`
	Description pgtype.Varchar `json:"description" db:"description"`
}

type Source struct {
	Id          pgtype.Int4    `json:"id" db:"id"`
	Code        pgtype.Varchar `json:"code" db:"code"`
	Value       pgtype.Varchar `json:"value" db:"value"`
	Description pgtype.Varchar `json:"description" db:"description"`
}
