package delivery

import (
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/api"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/usecase"
)

type grpc struct {
	usecase usecase.GrpcUsecase
}

type Grpc interface {
	GetData() ([]*api.Post,error)
}

func NewGrpc(usecase usecase.GrpcUsecase) Grpc{
	return &grpc{
		usecase : usecase,
	}
}

func (grpc *grpc) GetData() ([]*api.Post,error) {
	posts,err:=grpc.usecase.GetData()
	if err!=nil{
		return nil, err
	}
	return posts,nil
}
