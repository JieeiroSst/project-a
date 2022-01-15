package server

import (
	_ "github.com/JieeiroSst/itjob/users/docs"
	"github.com/JieeiroSst/itjob/users/internal/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	url := ginSwagger.URL("http://localhost:3000/swagger/users/doc.json") // The url pointing to API definition
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	group.POST("/login", r.router.Login).
		  POST("/register", r.router.SignUp).
		  POST("/update/profile", r.router.UpdateProfile).
		  POST("/lock_user", r.router.LockAccount)
	return nil
}