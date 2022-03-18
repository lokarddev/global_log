package web

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lokarddev/global_log/pkg/env"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Run() {
	if err := s.Router.Run(fmt.Sprintf(":%s", env.Port)); err != nil {
		log.Fatalf("error occured while running server: %s\n", err.Error())
	}
}

func NewServer() *Server {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), cors.Default())
	router.Use(static.Serve("/", static.LocalFile("../templates/dist", false)))
	router.NoRoute(func(c *gin.Context) { c.File("../templates/dist/index.html") })
	return &Server{Router: router}
}
