package main

import (
	casbinCMD "github.com/JieeiroSst/itjob/casbin/cmd"
	"github.com/JieeiroSst/itjob/config"
	dbCMD "github.com/JieeiroSst/itjob/db/cmd"
	"github.com/JieeiroSst/itjob/pkg/log"
	uploadCMD "github.com/JieeiroSst/itjob/upload/cmd"
	userCMD "github.com/JieeiroSst/itjob/users/cmd"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){
	router := gin.Default()

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	if err := dbCMD.NewDbCMD(conf); err != nil {
		log.NewLog().Error(err)
	}

	userMain := userCMD.NewUserMain(router)
	casbinCMD := casbinCMD.NewCasbinCMD(router)
	uploadCMD := uploadCMD.NewUploadCMD(router)

	if err := userMain.Run(); err != nil {
		log.NewLog().Error("run server user failed")
	}

	if err := casbinCMD.Run(); err != nil {
		log.NewLog().Error("run server casbin failed")
	}

	if err := uploadCMD.Run(); err != nil {
		log.NewLog().Error("run server upload failed")
	}

	go func() {
		if err := router.Run(conf.Server.PortServer); err != nil {
			log.NewLog().Errorf("run port server failed port %s: ",conf.Server.PortServer)
		}
	}()
}