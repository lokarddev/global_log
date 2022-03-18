package analytics_services

import (
	"github.com/lokarddev/global_log/internal/service"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
)

type AnalyticsService struct {
	storage      service.AnalyticsPostgresInterface
	logger       logger.LoggerInterface
	cacheStorage cache.CachingInterface
}

func NewAnalyticsService(storage service.AnalyticsPostgresInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *AnalyticsService {
	return &AnalyticsService{storage: storage, logger: logger, cacheStorage: cacheStorage}
}
