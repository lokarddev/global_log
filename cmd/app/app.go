package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lokarddev/global_log/internal/delivery/msg_broker"
	"github.com/lokarddev/global_log/internal/delivery/msg_broker/analytics"
	"github.com/lokarddev/global_log/internal/delivery/msg_broker/logs"
	"github.com/lokarddev/global_log/internal/delivery/web_http"
	"github.com/lokarddev/global_log/internal/delivery/web_http/analytics_http"
	"github.com/lokarddev/global_log/internal/delivery/web_http/logs_http"
	"github.com/lokarddev/global_log/internal/delivery/web_http/ws"
	"github.com/lokarddev/global_log/internal/repository/postgres/analytics_repo"
	"github.com/lokarddev/global_log/internal/repository/postgres/logs_repo"
	"github.com/lokarddev/global_log/internal/service/analytics_services"
	"github.com/lokarddev/global_log/internal/service/logs_services"
	"github.com/lokarddev/global_log/pkg/cache"
	"github.com/lokarddev/global_log/pkg/database"
	"github.com/lokarddev/global_log/pkg/env"
	"github.com/lokarddev/global_log/pkg/logger"
	"github.com/lokarddev/global_log/pkg/redis_client"
	"log"
)

type App struct {
	server        *web.Server
	dbPool        *pgxpool.Pool
	msgDispatcher broker.DispatcherInterface
	logger        logger.LoggerInterface
	cacheStorage  cache.CachingInterface
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer a.dbPool.Close()
	defer a.msgDispatcher.StopBrokerListening()
	defer cancel()

	go a.msgDispatcher.RunBrokerListening(ctx)
	a.server.Run()
}

func (a *App) InitApp() {
	api := a.server.Router.Group("/api")

	a.initLogs(a.dbPool, api)
	a.initAnalytics(a.dbPool, api)
	a.initWsHandler(api)
}

func (a *App) initLogs(db *pgxpool.Pool, api *gin.RouterGroup) {
	repo := logs_repo.NewLogsRepository(db, a.logger, a.cacheStorage)
	service := logs_services.NewLogsService(repo, a.logger, a.cacheStorage)
	httpHandler := logs_http.NewLogsHTTPHandler(service, a.logger, a.cacheStorage)
	httpHandler.RegisterRoutes(api)

	msgBrokerHandler := logs.NewBaseLogsHandler(service)
	msgBrokerHandler.MatchLogsHandlers(a.msgDispatcher)
}

func (a *App) initAnalytics(db *pgxpool.Pool, api *gin.RouterGroup) {
	repo := analytics_repo.NewAnalyticsRepository(db, a.logger, a.cacheStorage)
	service := analytics_services.NewAnalyticsService(repo, a.logger, a.cacheStorage)
	httpHandler := analytics_http.NewAnalyticsHTTPHandler(service, a.logger, a.cacheStorage)
	httpHandler.RegisterRoutes(api)

	msgBrokerHandler := analytics.NewBaseAnalyticsHandler(service)
	msgBrokerHandler.MatchAnalyticsHandlers(a.msgDispatcher)
}

func (a *App) initWsHandler(api *gin.RouterGroup) {
	httpHandler := ws.NewWebSocketHandler(a.logger, a.cacheStorage)
	httpHandler.RegisterRoutes(api)
}

func (a *App) initLogger() *logger.Logger {
	r, err := redis_client.NewRedisClient()
	if err != nil {
		log.Fatalf("error initialising redis: %s", err.Error())
	}
	return logger.NewLogger(logger.LoggerConfig{
		Source:    logger.GoLog,
		MsgClient: r,
	})
}

func NewApplication() *App {
	if err := env.InitEnvVariables(); err != nil {
		log.Fatal(err)
	}
	redisClient, err := redis_client.NewRedisClient()
	msgBus := redis_client.NewMsgBusRedis(redisClient)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.InitDatabasePostgres()
	if err != nil {
		log.Fatalf("error initialising db: %s", err.Error())
	}
	app := &App{
		server:       web.NewServer(),
		dbPool:       db,
		cacheStorage: cache.NewRedisCache(redisClient),
	}
	app.logger = app.initLogger()
	app.msgDispatcher = broker.NewDispatcher(msgBus, app.logger)

	return app
}
