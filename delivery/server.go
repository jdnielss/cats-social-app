package delivery

import (
	"fmt"
	"log"

	"enigmacamp.com/be-lms-university/config"
	"enigmacamp.com/be-lms-university/delivery/controller"
	"enigmacamp.com/be-lms-university/manager"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewEnrollmentController(s.ucManager.EnrollmentUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		ucManager: ucManager,
		engine:    engine,
		host:      host,
	}
}
