package analytics

import (
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/pkg/logger"
)

type LandingMetricsHandler struct {
	service delivery.AnalyticsServiceInterface
	logger  logger.LoggerInterface
}

func (h *LandingMetricsHandler) ProcessTask(payload []byte) (bool, error) {
	return true, nil
}
