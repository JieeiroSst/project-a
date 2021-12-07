package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/access_control"
	"github.com/JieeiroSst/itjob/casbin/internal/db"
	deliveryHttp "github.com/JieeiroSst/itjob/casbin/internal/delivery/http"
	"github.com/JieeiroSst/itjob/casbin/internal/http"
	"github.com/JieeiroSst/itjob/casbin/internal/repository"
	"github.com/JieeiroSst/itjob/casbin/internal/router"
	casbinServer "github.com/JieeiroSst/itjob/casbin/internal/server"
	"github.com/JieeiroSst/itjob/casbin/internal/usecase"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/bigcache"
	"github.com/JieeiroSst/itjob/pkg/jwt"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/mysql"
	"github.com/JieeiroSst/itjob/pkg/pagination"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/gin-gonic/gin"
)

type casbinCMD struct {
	server *gin.Engine
}

type CasbinCMD interface {
	Run() error
}

func NewCasbinCMD(server *gin.Engine) CasbinCMD {
	return &casbinCMD{
			server:server,
	}
}

func (c *casbinCMD) Run() error {
	log.NewLog().Info("Starting user")

	conf, err := config.ReadConf("conf/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	var paginationPage = pagination.NewPaginationPage(model.Pagination{})
	var snowflakeData = snowflake.NewSnowflake()
	var tokenUser = jwt.NewTokenUser(conf)
	var cache = bigcache.NewBigCache()


	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.MysqlUser,
		conf.Mysql.MysqlPassword,
		conf.Mysql.MysqlHost,
		conf.Mysql.MysqlPort,
		conf.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	accessControl := access_control.NewAuthorization(cache,tokenUser)

	casbinDB := db.NewCasbinDB(mysqlOrm, paginationPage)
	casbinRuleRepository := repository.NewCasbinRuleRepository(casbinDB)
	casbinRuleUseCase := usecase.NewCasbinRuleUseCase(casbinRuleRepository)
	newHttp := http.NewHttp(casbinRuleUseCase)
	newDeliveryHttp := deliveryHttp.NewDeliveryHttp(newHttp)
	casbinRouter := router.NewCasbinRouter(newDeliveryHttp, snowflakeData)
	newCasbinServer := casbinServer.NewCasbinServer(c.server,mysqlOrm, casbinRouter, accessControl)

	if err := newCasbinServer.Run(); err != nil {
		log.NewLog().Error(err.Error())
	}

	return nil
}