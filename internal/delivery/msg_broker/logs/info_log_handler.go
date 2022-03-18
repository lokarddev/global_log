package logs

import (
	"encoding/json"
	"github.com/lokarddev/global_log/internal/delivery"
	"github.com/lokarddev/global_log/internal/entity"
	"github.com/lokarddev/global_log/pkg/logger"
)

type InfoLogsHandler struct {
	service delivery.LogsServiceInterface
	logger  logger.LoggerInterface
}

func (h *InfoLogsHandler) ProcessTask(payload []byte) (bool, error) {
	var msg entity.LogMsg
	if err := json.Unmarshal(payload, &msg); err != nil {
		return false, err
	}
	if err := h.service.CreateLog(msg); err != nil {
		return false, err
	}
	return true, nil
}
