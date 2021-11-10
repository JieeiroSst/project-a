package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/casbin/internal/db"
	deliveryHttp "github.com/JieeiroSst/itjob/casbin/internal/delivery/http"
	"github.com/JieeiroSst/itjob/casbin/internal/http"
	"github.com/JieeiroSst/itjob/casbin/internal/repository"
	"github.com/JieeiroSst/itjob/casbin/internal/router"
	casbinServer "github.com/JieeiroSst/itjob/casbin/internal/server"
	"github.com/JieeiroSst/itjob/casbin/internal/usecase"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/model"
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

	var pagination = pagination.NewPaginationPage(model.Pagination{})
	var snowflake = snowflake.NewSnowflake()

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.CasbinMysql.MysqlUser,
		conf.CasbinMysql.MysqlPassword,
		conf.CasbinMysql.MysqlHost,
		conf.CasbinMysql.MysqlPort,
		conf.CasbinMysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	db := db.NewCasbinDB(mysqlOrm, pagination)
	repository := repository.NewCasbinRuleRepository(db)
	usecase := usecase.NewCasbinRuleUseCase(repository)
	http := http.NewHttp(usecase)
	deliveryHttp := deliveryHttp.NewDeliveryHttp(http)
	router := router.NewCasbinRouter(deliveryHttp, snowflake)
	casbinServer := casbinServer.NewCasbinServer(c.server,mysqlOrm,router)

	if err := casbinServer.Run(); err != nil {
		log.NewLog().Error(err.Error())
	}

	return nil
}