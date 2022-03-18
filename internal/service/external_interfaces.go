package service

import "github.com/lokarddev/global_log/internal/entity"

type AnalyticsPostgresInterface interface {
}

type LogsPostgresInterface interface {
	GetAllLogs() ([]entity.CleanLog, error)
	CreateLog(msg entity.LogMsg) (entity.PostgresLog, error)
}
