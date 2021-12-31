package http

import (
	"context"
	"github.com/JieeiroSst/itjob/post/internal/search/internal/proto"
	"github.com/JieeiroSst/itjob/post/internal/search/internal/usecase"
)

type ElasticsearcHttp interface {
	InsertPost(ctx context.Context,data proto.Post) error
	Query(ctx context.Context,name string) (interface{},error)
}

type elasticsearcHttp struct {
	usecase usecase.ElasticsearchUsecase
}

func NewHttp(usecase usecase.ElasticsearchUsecase) ElasticsearcHttp {
	return &elasticsearcHttp{usecase:usecase}
}

func (h *elasticsearcHttp) InsertPost(ctx context.Context,data proto.Post) error {
	return h.usecase.Insert(ctx,data)
}

func (h *elasticsearcHttp) Query(ctx context.Context,name string) (interface{},error) {
	return h.usecase.Query(ctx,name)
}