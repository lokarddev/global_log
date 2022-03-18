package logs_services

import (
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/internal/service"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
)

type LogsService struct {
	storage      service.LogsPostgresInterface
	logger       logger.LoggerInterface
	cacheStorage cache.CachingInterface
}

func NewLogsService(storage service.LogsPostgresInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *LogsService {
	return &LogsService{storage: storage, logger: logger, cacheStorage: cacheStorage}
}

func (s *LogsService) GetAllLogs() ([]entity.CleanLog, error) {
	return s.storage.GetAllLogs()
}

func (s *LogsService) CreateLog(msg entity.LogMsg) error {
	log, err := s.storage.CreateLog(msg)
	err = s.cacheStorage.PublishToAllWsCons(log.ToCLeanLog())
	return err
}
