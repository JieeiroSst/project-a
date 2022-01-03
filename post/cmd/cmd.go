package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/pkg/elasticsearch"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/mysql"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/api"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/delivery"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/repository"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/router"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/usecase"
	searchDelivery "github.com/JieeiroSst/itjob/post/internal/search/delivery"
	searchHttp "github.com/JieeiroSst/itjob/post/internal/search/http"
	searchRepository "github.com/JieeiroSst/itjob/post/internal/search/repository"
	searchRouter "github.com/JieeiroSst/itjob/post/internal/search/router"
	searchUsecase "github.com/JieeiroSst/itjob/post/internal/search/usecase"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
)

type serverGrpcPost struct {
	engine *gin.Engine
}

type ServerGrpcPost interface {
	RunServerGrpc() error
	RunClientGRPC() error
}

func NewServerGrpcPost(engine *gin.Engine) ServerGrpcPost {
	return &serverGrpcPost{
		engine:engine,
	}
}

func (s *serverGrpcPost) RunServerGrpc() error {
	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}

	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql.MysqlUser,
		conf.Mysql.MysqlPassword,
		conf.Mysql.MysqlHost,
		conf.Mysql.MysqlPort,
		conf.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	grpcRepository := repository.NewGrpcRepository(mysqlOrm)
	grpcUsecase := usecase.NewGrpcUsecase(grpcRepository)
	grpcDelivery := delivery.NewGrpc(grpcUsecase)
	grpcRouter := router.NewGRPCServer(grpcDelivery)

	server :=grpc.NewServer()
	api.RegisterHandleServiceServer(server, grpcRouter)
	listen, err := net.Listen("tcp", conf.Server.PprofPort)
	if err != nil {
		log.NewLog().Error(err)
		return err
	}
	if err := server.Serve(listen); err != nil {
		log.NewLog().Error(err)
		return err
	}
	return nil
}

func (s *serverGrpcPost) RunClientGRPC() error {
	conf, err := config.ReadConf("config/conf-docker.yml")
	if err != nil {
		log.NewLog().Error(err.Error())
	}
	elasticsearchConn :=elasticsearch.NewGetElasticsearchConn(conf.Elasticsearch.Dns)
	searchRepository := searchRepository.NewElasticsearchRepository(elasticsearchConn)
	searcgUsecase := searchUsecase.NewElasticsearchUsecase(searchRepository)
	searchHttp := searchHttp.NewHttp(searcgUsecase)
	searchDelivery := searchDelivery.NewElasticsearcDelivery(searchHttp)
	searchRouter :=searchRouter.NewElasticsearcRouter(searchDelivery,conf)
	group := s.engine.Group("/")
	group.GET("/", searchRouter.Query)
	group.POST("/search",searchRouter.InsertPost)

	return nil
}