package controllers

import (
	"fmt"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func New(config *config.Config) *Server {
	router := gin.Default()

	server := &Server{
		Router: router,
		Config: config,
	}

	server.SetupRoutes()

	return server
}

func (s *Server) Run() {
	reader := credentials.NewConfigReader()
	var config config.Config
	if err := reader.Read("debug", &config); err != nil {
		panic(err)
	}
	fmt.Println("ServerAddress:", config.ServerAddress)

	fmt.Println("ServerAddress:", s.Config.ServerAddress)
	s.Router.Run(s.Config.ServerAddress)
}
