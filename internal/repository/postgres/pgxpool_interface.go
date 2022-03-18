package postgres

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

const (
	AnalyticsTable = "analytics"
	LogsTable      = "logs"
	EventTable     = "event"
	EventTypeTable = "event_type"
	LogLevelTable  = "log_level"
	MetricsTable   = "metrics"
	SourceTable    = "source"
)

// PgxPoolInterface minimal interface for mocking and testing
type PgxPoolInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}
