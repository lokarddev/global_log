package analytics_http

import (
	"github.com/gin-gonic/gin"
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
)

type AnalyticsHTTPHandler struct {
	logger       logger.LoggerInterface
	service      delivery.AnalyticsServiceInterface
	cacheStorage cache.CachingInterface
}

func NewAnalyticsHTTPHandler(service delivery.AnalyticsServiceInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *AnalyticsHTTPHandler {
	return &AnalyticsHTTPHandler{service: service, logger: logger, cacheStorage: cacheStorage}
}

func (h *AnalyticsHTTPHandler) RegisterRoutes(api *gin.RouterGroup) {
}
