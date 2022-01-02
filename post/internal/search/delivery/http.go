package delivery

import (
	"context"
	"github.com/JieeiroSst/itjob/post/internal/search/http"
	"github.com/JieeiroSst/itjob/post/internal/search/proto"
)

type elasticsearcDelivery struct {
	http http.ElasticsearcHttp
}

type ElasticsearcDelivery interface {
	InsertPost(ctx context.Context,data proto.Post) error
	Query(ctx context.Context,name string) (interface{},error)
}

func NewElasticsearcDelivery(http http.ElasticsearcHttp) ElasticsearcDelivery {
	return &elasticsearcDelivery{
		http:http,
	}
}

func(e *elasticsearcDelivery) InsertPost(ctx context.Context,data proto.Post) error {
	return e.http.InsertPost(ctx, data)
}

func(e *elasticsearcDelivery) Query(ctx context.Context,name string) (interface{},error) {
	return e.http.Query(ctx, name)
}