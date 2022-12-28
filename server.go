package main

import (
	"fmt"
	"go-self-payroll-system/config"
	"go-self-payroll-system/delivery"
	"go-self-payroll-system/repository"
	"go-self-payroll-system/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

type server struct {
	httpServer *echo.Echo
	cfg			config.Config
}

type Server interface {
	Running()
}

func InitServer(cfg config.Config) Server {
	e := echo.New()

	return &server{
		httpServer: e,
		cfg: 		cfg,
	}
}

func (s *server) Running() {
	positionRepo := repository.NewPositionRepo(s.cfg)
	positionUsecase := usecase.NewPositionUsecase(positionRepo)
	positionDelivery := delivery.NewPositionDelivery(positionUsecase)
	posGroup := s.httpServer.Group("/positions")
	positionDelivery.Mounting(posGroup)


	err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort()))
	if err != nil {
		log.Panic(err)
	}
	
}