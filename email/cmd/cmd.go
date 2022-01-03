package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/email/internal/delivery"
	"github.com/JieeiroSst/itjob/email/internal/repository"
	"github.com/JieeiroSst/itjob/email/internal/router"
	"github.com/JieeiroSst/itjob/email/internal/server"
	"github.com/JieeiroSst/itjob/email/internal/usecase"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/mysql"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/gin-gonic/gin"
)

type emailCMD struct {
	engine *gin.Engine
}

type EmailCMD interface {
	RunEmailCMD() error
}

func NewEmailCMD(engine *gin.Engine) EmailCMD {
	return &emailCMD{
		engine:engine,
	}
}

func (e *emailCMD) RunEmailCMD() error {
	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	var (
		snowflakeData snowflake.SnowflakeData
	)

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.MysqlUser,
		conf.Mysql.MysqlPassword,
		conf.Mysql.MysqlHost,
		conf.Mysql.MysqlPort,
		conf.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)
	emailRepository := repository.NewEmailRepository(mysqlOrm)
	emailUsecase := usecase.NewEmailUsecase(emailRepository)
	http := delivery.NewEmailHttp(emailUsecase)
	emailRouter := router.NewEmailRouter(http, *conf, snowflakeData)
	emailServer := server.NewEmailServer(emailRouter)

	if err := emailServer.Run(); err != nil {
		log.NewLog().Error(err.Error())
	}

	return nil
}
