package usecase

import (
	"context"
	"github.com/JieeiroSst/itjob/post/internal/search/proto"
	"github.com/JieeiroSst/itjob/post/internal/search/repository"
)

type elasticsearchUsecase struct {
	repo repository.ElasticisesRepository
}


type ElasticsearchUsecase interface {
	Insert(ctx context.Context,data proto.Post) error
	Query(ctx context.Context,name string) (interface{},error)
}

func NewElasticsearchUsecase(repo repository.ElasticisesRepository) ElasticsearchUsecase {
	return &elasticsearchUsecase{
		repo:repo,
	}
}

func (e *elasticsearchUsecase) Insert(ctx context.Context,data proto.Post) error{
	return e.repo.Insert(ctx,data)
}

func (e *elasticsearchUsecase) Query(ctx context.Context,name string) (interface{},error){
	return e.repo.Query(ctx,name)
}