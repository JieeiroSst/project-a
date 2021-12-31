package repository

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/post/internal/grpc/pkg/api"
	"github.com/golang/protobuf/ptypes"
	"gorm.io/gorm"
)

type GrpcRepository interface {
	GetData() ([]*api.Post,error)
}

type grpcRepository struct {
	db *gorm.DB
}

func NewGrpcRepository(db *gorm.DB) GrpcRepository{
	return &grpcRepository{
		db:db,
	}
}

func (grpc *grpcRepository) GetData() ([]*api.Post,error){
	var postAll []model.Posts
	grpc.db.Find(&postAll)
	var posts []*api.Post
	for _,post := range postAll{
		createdAt, err := ptypes.TimestampProto(post.CreatedAt)
		if err != nil {

		}
		updatedAt, err := ptypes.TimestampProto(post.UpdatedAt)
		if err != nil {

		}
		publishedAt, err := ptypes.TimestampProto(post.PublishedAt)
		if err != nil {

		}
		data := api.Post{
			Id:          int32(post.Id),
			AuthorId:    int32(post.AuthorId),
			Title:       post.Title,
			MetaTitle:   post.MetaTitle,
			Slug:        post.Slug,
			Summary:     post.Summary,
			Published:   int32(post.Published),
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			PublishedAt: publishedAt,
			Content:     post.Content,
		}

		posts=append(posts,&data)

	}
	return posts,nil
}