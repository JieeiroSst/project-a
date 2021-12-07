package server

import (
	"github.com/JieeiroSst/itjob/access_control"
	"github.com/JieeiroSst/itjob/casbin/internal/router"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
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

	resource := s.server.Group("/api")

	resource.Use(s.auth.Authenticate())
	{
		resourceCasbin:=resource.Group("/casbin")

		resourceCasbin.POST("",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.CreateCasbinRule)
		resourceCasbin.PUT("/:id",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.UpdateCasbinRulePtype)
		resourceCasbin.GET("",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.CasbinRuleAll)
		resourceCasbin.GET("/:id",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.CasbinRuleById)
		resourceCasbin.GET("/option",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.OptionList)
		resourceCasbin.DELETE("/:id",s.auth.Authorize("/api/casbin/*","POST",adapter),s.router.DeleteCasbinRule)
	}

	return nil
}