package router

import (
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/post/internal/search/internal/delivery"
	"github.com/JieeiroSst/itjob/post/internal/search/internal/http"
	"github.com/JieeiroSst/itjob/post/internal/search/internal/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

type elasticsearcRouter struct {
	delivery delivery.ElasticsearcDelivery
	config   *config.Config
}

type ElasticsearcRouter interface {
	InsertPost(c *gin.Context)
	Query(c *gin.Context)
}

func NewElasticsearcRouter(delivery delivery.ElasticsearcDelivery) ElasticsearcRouter {
	return &elasticsearcRouter{
		delivery:delivery,
	}
}

func(e *elasticsearcRouter) InsertPost(c *gin.Context) {
	conn,err:=grpc.Dial("localhost"+e.config.Server.PprofPort,grpc.WithInsecure())
	if err != nil {
		log.Println("CLIENT IS NO DIAL",err)
		return
	}
	log.Printf("CLIENT IS DIAL AT %s...", e.config.Server.PprofPort)

	client:=proto.NewHandleServiceClient(conn)
	req :=&proto.RequestPost{}
	data,err:=client.UpdatePost(c,req)
	if err!=nil{
		c.JSON(500,"get data failed")
		return
	}
	for _,post := range data.Posts {
		err := e.delivery.InsertPost(c,*post)
		if err!=nil {
			c.JSON(500,"insert failed")
			return
		}
	}
	c.JSON(200,"insert success")

}


func(e *elasticsearcRouter) Query(c *gin.Context) {
	name := c.Query("name")
	result,err:=e.delivery.Query(c, name)
	if err!=nil{
		c.JSON(500,map[string]interface{}{"data":"no data"})
		return
	}
	c.JSON(200,map[string]interface{}{"result":result})
}