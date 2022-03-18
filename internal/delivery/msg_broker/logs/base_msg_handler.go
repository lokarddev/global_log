package logs

import (
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/internal/delivery/msg_broker"
)

// supported commands list
const (
	infoLog = "infoLog"
)

type BaseLogsHandler struct {
	InfoLogsHandler *InfoLogsHandler
}

func (h *BaseLogsHandler) MatchLogsHandlers(d broker.DispatcherInterface) {
	d.Attach(infoLog, h.InfoLogsHandler)
}

func NewBaseLogsHandler(service delivery.LogsServiceInterface) *BaseLogsHandler {
	return &BaseLogsHandler{
		InfoLogsHandler: &InfoLogsHandler{service: service},
	}
}
