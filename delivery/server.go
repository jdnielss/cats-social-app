package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"jdnielss.dev/cats-social-app/config"
	"jdnielss.dev/cats-social-app/delivery/controller"
	"jdnielss.dev/cats-social-app/manager"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewCatController(s.ucManager.CatUseCase(), rg).Route()
	controller.NewAuthController(s.ucManager.AuthUseCase(), rg).Route()
	controller.NewUserController(s.ucManager.UserUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	s.engine.Use()
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
