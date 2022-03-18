package delivery

import "github.com/lokarddev/global_log/internal/entity"

type LogsServiceInterface interface {
	GetAllLogs() ([]entity.CleanLog, error)
	CreateLog(msg entity.LogMsg) error
}

type AnalyticsServiceInterface interface {
}
