package router

import (
	"context"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/api"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/delivery"
)

type grpcRouter struct{
	http delivery.Grpc
}

type GrpcRouter interface {
	UpdatePost(ctx context.Context,req *api.RequestPost) (*api.ResponsePost,error)
}

func NewGRPCServer(http delivery.Grpc) GrpcRouter {
	return &grpcRouter{
		http:http,
	}
}

func (s *grpcRouter) UpdatePost(ctx context.Context,req *api.RequestPost) (*api.ResponsePost,error) {
	posts,err := s.http.GetData()
	if err!=nil{
		return &api.ResponsePost{}, err
	}
	return &api.ResponsePost{Posts:posts}, nil
}