package logs_http

import (
	"github.com/gin-gonic/gin"
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
	"net/http"
)

type LogsHTTPHandler struct {
	logger       logger.LoggerInterface
	service      delivery.LogsServiceInterface
	cacheStorage cache.CachingInterface
}

func NewLogsHTTPHandler(service delivery.LogsServiceInterface, logger logger.LoggerInterface, cacheStorage cache.CachingInterface) *LogsHTTPHandler {
	return &LogsHTTPHandler{service: service, logger: logger, cacheStorage: cacheStorage}
}

func (h *LogsHTTPHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/get-logs", h.GetLogs)
}

func (h *LogsHTTPHandler) GetLogs(c *gin.Context) {
	logs, err := h.service.GetAllLogs()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}
