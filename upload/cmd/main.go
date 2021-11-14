package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/mysql"
	"github.com/JieeiroSst/itjob/pkg/s3"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/gin-gonic/gin"

	"github.com/JieeiroSst/itjob/upload/internal/db"
	deliverHttp "github.com/JieeiroSst/itjob/upload/internal/delivery/http"
	"github.com/JieeiroSst/itjob/upload/internal/http"
	"github.com/JieeiroSst/itjob/upload/internal/repository"
	"github.com/JieeiroSst/itjob/upload/internal/router"
	"github.com/JieeiroSst/itjob/upload/internal/server"
	"github.com/JieeiroSst/itjob/upload/internal/usecase"
)

type uploadCMD struct {
	server *gin.Engine
}

type UploadCMD interface {
	Run() error
}

func NewUploadCMD(server *gin.Engine) UploadCMD {
	return &uploadCMD{
		server: server,
	}
}

func (u *uploadCMD) Run() error {
	log.NewLog().Info("Starting user")

	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	var (
		newSnowflake = snowflake.NewSnowflake()
	)

	s3:=s3.NewS3(conf.AmazonS3.S3Region)

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.MysqlUser,
		conf.Mysql.MysqlPassword,
		conf.Mysql.MysqlHost,
		conf.Mysql.MysqlPort,
		conf.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	uploadDB := db.NewUploadDB(mysqlOrm)
	uploadRepository := repository.NewUploadRepository(uploadDB)
	uploadUsecase := usecase.NewUploadUsecase(uploadRepository,s3,conf)
	uploadHttp := http.NewUploadHttp(uploadUsecase)
	uploadDeliveryHttp := deliverHttp.NewUploadDeliveryHttp(uploadHttp)
	uploadRouter :=router.NewUploadRouter(uploadDeliveryHttp, newSnowflake)
	uploadServer := server.NewUploadServer(uploadRouter, u.server)

	if err := uploadServer.RunServer(); err != nil {
		log.NewLog().Error(err.Error())
	}
	return nil
}