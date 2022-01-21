package server

import (
	"github.com/JieeiroSst/itjob/email/internal/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	url := ginSwagger.URL("http://localhost:5000/swagger/email/doc.json") // The url pointing to API definition
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	group.POST("/client_send", e.router.AdminSendEmail)
	group.POST("/admin_send", e.router.UserSendEmail)

	return nil
}
