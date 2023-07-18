package controller

import (
	"jul14/client"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Client client.Client
}

// Connect function is use for connect mysql database
// and create database if the database is not already created
func (s *Server) Connect() {
	s.Router = gin.Default()
	s.Client = client.StorjClient()
	s.InitializeRouter()
}

func (s *Server) InitializeRouter() {
	v1 := s.Router.Group("/api/v1")
	v1.POST("/upload", s.Upload)
}
