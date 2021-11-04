package server

import (
	"github.com/JieeiroSst/itjob/users/internal/router"
	"github.com/gin-gonic/gin"
)

type userServer struct {
	router router.UserRouter
	server *gin.Engine
}

type UserServer interface {
	RunServer() error
}

func NewUserServer(router router.UserRouter, server *gin.Engine) UserServer {
	return &userServer{
		router: router,
		server: server,
	}
}

func (r *userServer) RunServer() error {
	group := r.server.Group("/v1")
	group.POST("/login", r.router.Login).
		  POST("/register", r.router.SignUp).
		  POST("/update/profile", r.router.UpdateProfile).
		  POST("/lock_user", r.router.LockAccount)
	return nil
}