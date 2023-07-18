package controller

import (
	"context"
	"jul14/client"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	Router *gin.Engine
	Client client.Client
}

// Connect function is use for connect mysql database
// and create database if the database is not already created
func (s *Server) Connect() {
	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ginSetMode(os.Getenv("R_MODE"))
	s.Router = gin.Default()
	s.Client = client.StorjClient(ctx)
	s.InitializeRouter()
}

func (s *Server) InitializeRouter() {
	v1 := s.Router.Group("/api/v1")
	v1.POST("/upload", s.Upload)
}

func ginSetMode(s string) {
	if s == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if s == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}
}
