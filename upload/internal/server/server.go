package server

import (
	"github.com/JieeiroSst/itjob/upload/internal/router"
	"github.com/gin-gonic/gin"
)

type uploadServer struct {
	router router.UploadRouter
	server *gin.Engine
}

type UploadServer interface {
	RunServer() error
}

func NewUploadServer(router router.UploadRouter, server *gin.Engine) UploadServer {
	return &uploadServer{
		router:router,
		server:server,
	}
}

func (u *uploadServer) RunServer() error {
	group := u.server.Group("/v1")
	group.POST("/upload",u.router.AddFileS3).
		GET("/upload",u.router.ReadFile).GET("/upload/option",u.router.Option)

	return nil
}
