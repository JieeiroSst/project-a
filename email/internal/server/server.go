package server

import (
	"github.com/JieeiroSst/itjob/email/internal/router"
	"github.com/gin-gonic/gin"
)

type emailServer struct {
	router router.EmailRouter
	engine *gin.Engine
}

type EmailServer interface {
	Run() error
}

func NewEmailServer(router router.EmailRouter) EmailServer {
	return &emailServer{
		router:router,
	}
}

func (e *emailServer) Run() error {
	group := e.engine.Group("/v1/email")

	group.POST("/client_send", e.router.AdminSendEmail)
	group.POST("/admin_send", e.router.UserSendEmail)

	return nil
}
