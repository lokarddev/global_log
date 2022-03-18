package analytics

import (
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/internal/delivery/msg_broker"
)

// supported commands list
const (
	landingMetrics = "landing_metrics"
)

type BaseAnalyticsHandler struct {
	LandingMetricsHandler *LandingMetricsHandler
}

func (h *BaseAnalyticsHandler) MatchAnalyticsHandlers(d broker.DispatcherInterface) {
	d.Attach(landingMetrics, h.LandingMetricsHandler)
}

func NewBaseAnalyticsHandler(service delivery.AnalyticsServiceInterface) *BaseAnalyticsHandler {
	return &BaseAnalyticsHandler{
		LandingMetricsHandler: &LandingMetricsHandler{service: service},
	}
}
