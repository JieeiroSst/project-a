package usecase

import (
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/api"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/repository"
)

type GrpcUsecase interface {
	GetData() ([]*api.Post,error)
}

type grpcUsecase struct {
	repo repository.GrpcRepository
}

func NewGrpcUsecase(repo repository.GrpcRepository) GrpcUsecase{
	return &grpcUsecase{
		repo:repo,
	}
}

func (grpc *grpcUsecase) GetData() ([]*api.Post,error){
	posts,err:=grpc.repo.GetData()
	if err!=nil{
		return nil, err
	}
	return posts,nil
}