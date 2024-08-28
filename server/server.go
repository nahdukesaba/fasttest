package server

import (
	"github.com/gin-gonic/gin"

	"fasttest/service"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) RegisterRoute() {
	handler := service.NewService()
	s.router.GET("/products", handler.GetAllProduct)
	s.router.POST("/order", handler.CreateOrder)
	s.router.GET("/orders", handler.GetAllOrder)
}

func (s *Server) Run() {
	s.router.Run(":8081")
}
