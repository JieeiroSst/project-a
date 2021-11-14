package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/pkg/jwt"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/mysql"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/JieeiroSst/itjob/users/internal/db"
	deliveryHttp "github.com/JieeiroSst/itjob/users/internal/delivery/http"
	"github.com/JieeiroSst/itjob/users/internal/http"
	"github.com/JieeiroSst/itjob/users/internal/repository"
	userRouter "github.com/JieeiroSst/itjob/users/internal/router"
	userServer "github.com/JieeiroSst/itjob/users/internal/server"
	"github.com/JieeiroSst/itjob/users/internal/usecase"
	"github.com/JieeiroSst/itjob/utils"
	"github.com/gin-gonic/gin"
)

type userCMD struct {
	server *gin.Engine
}

type UserCMD interface {
	Run() error
}

func NewUserMain(server *gin.Engine) UserCMD {
	return &userCMD{
		server: server,
	}
}

func (m *userCMD) Run() error {
	log.NewLog().Info("Starting user")

	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	var (
		hash = utils.NewHash()
		tokenUser = jwt.NewTokenUser(conf)
		newSnowflake = snowflake.NewSnowflake()
	)

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.MysqlUser,
		conf.Mysql.MysqlPassword,
		conf.Mysql.MysqlHost,
		conf.Mysql.MysqlPort,
		conf.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	userDB := db.NewUserDB(mysqlOrm)
	userRepository := repository.NewUserRepository(userDB)
	userCase := usecase.NewUserCase(userRepository, *hash, tokenUser, *conf)
	userHttp := http.NewUserHttp(userCase)
	newDeliveryHttp := deliveryHttp.NewDeliveryHttp(userHttp)
	newRouter := userRouter.NewRouter(newDeliveryHttp, newSnowflake)
	newUserServer := userServer.NewUserServer(newRouter, m.server)

	if err := newUserServer.RunServer(); err != nil {
		log.NewLog().Error(err.Error())
	}

	return nil
}
