package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/logger"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	logger       logger.LoggerInterface
	cacheStorage cache.CachingInterface
}

var upgrader = websocket.Upgrader{
	CheckOrigin: originChecker,
}

func originChecker(_ *http.Request) bool {
	return true
}

func NewWebSocketHandler(l logger.LoggerInterface, cacheStorage cache.CachingInterface) *WebSocketHandler {
	return &WebSocketHandler{
		logger:       l,
		cacheStorage: cacheStorage,
	}
}

func (h *WebSocketHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/logger-ws-conn", h.HandleWS)
}

func (h *WebSocketHandler) HandleWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	connId, _ := uuid.NewV4()
	h.cacheStorage.NewWsConnection(connId.String(), conn)
	defer func(conn *websocket.Conn) {
		h.cacheStorage.DropWsConnection(connId.String())
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
	}(conn)
	if err != nil {
		log.Println(err)
	}
	if err = h.listenToWS(conn); err != nil {
		log.Println(err)
	}
}

func (h *WebSocketHandler) listenToWS(conn *websocket.Conn) error {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return err
		}
	}
}
