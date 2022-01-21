package server

import (
	"github.com/JieeiroSst/itjob/access_control"
	_ "github.com/JieeiroSst/itjob/casbin/docs"
	"github.com/JieeiroSst/itjob/casbin/internal/router"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type casbinServer struct {
	server *gin.Engine
	db *gorm.DB
	auth access_control.Authorization
	router router.CasbinRouter
}

type CasbinServer interface {
	Run() error
}

func NewCasbinServer(server *gin.Engine,db *gorm.DB, router router.CasbinRouter,auth access_control.Authorization) CasbinServer{
	return &casbinServer{
		server:server,
		db:db,
		router:router,
		auth:auth,
	}
}

func (s *casbinServer) Run() error {
	adapter, err:=gormadapter.NewAdapterByDB(s.db)
	if err != nil {
		return err
	}

	resource := s.server.Group("/v1")

	url := ginSwagger.URL("http://localhost:5000/swagger/casbin/doc.json") // The url pointing to API definition
	resource.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	resource.Use(s.auth.Authenticate())
	{
		resourceCasbin:=resource.Group("/casbin")

		resourceCasbin.POST("",s.auth.Authorize("/v1/casbin/*","POST",adapter),s.router.CreateCasbinRule)
		resourceCasbin.PUT("/:id",s.auth.Authorize("/v1/casbin/*","PUT",adapter),s.router.UpdateCasbinRulePtype)
		resourceCasbin.GET("",s.auth.Authorize("/v1/casbin/*","GET",adapter),s.router.CasbinRuleAll)
		resourceCasbin.GET("/:id",s.auth.Authorize("/v1/casbin/*","GET",adapter),s.router.CasbinRuleById)
		resourceCasbin.GET("/option",s.auth.Authorize("/v1/casbin/*","GET",adapter),s.router.OptionList)
		resourceCasbin.DELETE("/:id",s.auth.Authorize("/v1/casbin/*","DELETE",adapter),s.router.DeleteCasbinRule)
	}

	return nil
}