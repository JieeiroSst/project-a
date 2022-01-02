package repository

import (
	"context"
	"github.com/JieeiroSst/itjob/post/internal/search/proto"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

type ElasticisesRepository interface {
	Insert(ctx context.Context, data proto.Post) error
	Query(ctx context.Context, name string) (interface{}, error)
}

type elasticisesRepository struct {
	elastic *elastic.Client
}


func NewElasticsearchRepository(elastic *elastic.Client) ElasticisesRepository {
	return &elasticisesRepository{elastic: elastic}
}

func (e *elasticisesRepository) Insert(ctx context.Context,data proto.Post) error{
	id, err := gonanoid.New()
	put1, err := e.elastic.Index().
		Index("posts").
		Type("post").
		Id(id).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return nil
}

func (e *elasticisesRepository) Query(ctx context.Context,name string) (interface{},error){
	query := elastic.NewMatchBoolPrefixQuery("title", name)
	result, err := e.elastic.Search().
		Index("posts").
		Type("post").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	var posts []proto.Post
	var post proto.Post
	for _, item := range result.Each(reflect.TypeOf(post)) {
		t := item.(proto.Post)
		post = proto.Post{
			Id:          t.Id,
			AuthorId:    t.AuthorId,
			Title:       t.Title,
			MetaTitle:   t.MetaTitle,
			Slug:        t.Slug,
			Summary:     t.Summary,
			Published:   t.Published,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
			PublishedAt: t.PublishedAt,
			Content:     t.Content,
		}
		posts= append(posts,post)
	}
	return posts,nil
}