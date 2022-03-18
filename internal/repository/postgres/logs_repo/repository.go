package logs_repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/internal/repository/postgres"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
)

type LogsRepository struct {
	db           postgres.PgxPoolInterface
	ctx          context.Context
	logger       logger.LoggerInterface
	cacheStorage cache.CachingInterface
}

func NewLogsRepository(db postgres.PgxPoolInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *LogsRepository {
	return &LogsRepository{db: db, ctx: context.Background(), logger: logger, cacheStorage: cacheStorage}
}

func (r *LogsRepository) GetAllLogs() ([]entity.CleanLog, error) {
	var logs []entity.CleanLog
	query := fmt.Sprintf("SELECT id, created_at, updated_at, log_level_id, payload, event_id FROM %s", postgres.LogsTable)
	rows, err := r.db.Query(r.ctx, query)
	defer rows.Close()
	for rows.Next() {
		var log entity.PostgresLog
		err = rows.Scan(&log.Id, &log.CreatedAt, &log.UpdatedAt, &log.LogLevelId, &log.Payload, &log.EventId)
		logs = append(logs, log.ToCLeanLog())
	}
	return logs, err
}

func (r *LogsRepository) CreateLog(msg entity.LogMsg) (entity.PostgresLog, error) {
	var singleLog entity.PostgresLog
	eventId, err := r.createEvent(msg)
	if err != nil {
		return singleLog, err
	}
	query := fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, log_level_id, payload, event_id) VALUES 
		(now(), now(), 
		(select id FROM %s WHERE code=$1 LIMIT 1), 
		$2, $3) RETURNING id, created_at, log_level_id, payload, event_id`, postgres.LogsTable, postgres.LogLevelTable)
	err = r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.db.Query(r.ctx, query, msg.Level, msg.Payload, eventId)
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&singleLog.Id, &singleLog.CreatedAt, &singleLog.LogLevelId, &singleLog.Payload, &singleLog.EventId)
			return err
		}
		return err
	})
	return singleLog, err
}

func (r *LogsRepository) createEvent(msg entity.LogMsg) (int, error) {
	var id int
	query := fmt.Sprintf(`
		INSERT INTO %s (created_at, updated_at, event_type_id, source_id) 
		VALUES (now(), now(), 
		(SELECT id FROM %s WHERE code=$1 LIMIT 1), 
		(SELECT id FROm %s WHERE code=$2 LIMIT 1)) RETURNING id`,
		postgres.EventTable, postgres.EventTypeTable, postgres.SourceTable)
	err := r.db.BeginTxFunc(r.ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(r.ctx, query, msg.Event, msg.Source)
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&id)
		}
		return err
	})
	return id, err
}
