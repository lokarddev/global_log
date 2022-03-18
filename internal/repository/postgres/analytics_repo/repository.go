package analytics_repo

import (
	"context"
	"github.com/lokarddev/global_log/internal/repository/postgres"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
)

type AnalyticsRepository struct {
	db           postgres.PgxPoolInterface
	ctx          context.Context
	logger       logger.LoggerInterface
	cacheStorage cache.CachingInterface
}

func NewAnalyticsRepository(db postgres.PgxPoolInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *AnalyticsRepository {
	return &AnalyticsRepository{db: db, ctx: context.Background(), logger: logger, cacheStorage: cacheStorage}
}
