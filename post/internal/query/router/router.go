package router

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/JieeiroSst/itjob/post/internal/query/delivery"
	"github.com/JieeiroSst/itjob/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type postRouter struct {
	http delivery.PostHttp
	pagination utils.PaginationPage
	snowflake snowflake.SnowflakeData
}

type PostRouter interface {
	CreatePosts(c *gin.Context)
	UpdatePosts(c *gin.Context)
	DeletePosts(c *gin.Context)
	PostsById(c *gin.Context)
	CreateProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	ProfileById(c *gin.Context)
	CreatePostMetas(c *gin.Context)
	UpdatePostMetas(c *gin.Context)
	DeletePostMetas(c *gin.Context)
	PostMetasById(c *gin.Context)
	CreateComment(c *gin.Context)
	CreateCategories(c *gin.Context)
	UpdateCategories(c *gin.Context)
	DeleteCategories(c *gin.Context)
	CategoriesById(c *gin.Context)
	PostsAll(c *gin.Context)
	ProfileAll(c *gin.Context)
	PostMetasAll(c *gin.Context)
	CategoriesAll(c *gin.Context)
	CommentAllPost(c *gin.Context)
	PublishPost(c *gin.Context)
	RemoveComment(c *gin.Context)
	ListPublishPost(c *gin.Context)
	ListNotPublishPost(c *gin.Context)
}

func NewPostRouter(http delivery.PostHttp, pagination utils.PaginationPage,snowflake snowflake.SnowflakeData) PostRouter {
	return &postRouter{
		http: http,
		pagination: pagination,
		snowflake:snowflake,
	}
}

func (p *postRouter) CreatePosts(c *gin.Context) {
	var request model.RequestPost
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	post := model.Posts{
		Id:           p.snowflake.GearedID(),
		AuthorId:     request.AuthorId,
		Title:        request.Title,
		MetaTitle:    request.MetaTitle,
		Slug:         request.Slug,
		Summary:      request.Summary,
		Published:    0,
		CreatedAt:    time.Now(),
		UpdatedAt:    nil,
		PublishedAt:  nil,
		Content:      request.Content,
	}
	if err := p.http.CreatePosts(&post); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "create data post success",
		"error": nil,
	})
}

func (p *postRouter) UpdatePosts(c *gin.Context) {
	var request model.RequestPost
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(c.Query("name"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	post := model.Posts{
		AuthorId:     request.AuthorId,
		Title:        request.Title,
		MetaTitle:    request.MetaTitle,
		Slug:         request.Slug,
		Summary:      request.Summary,
		Published:    0,
		UpdatedAt:    time.Now(),
		PublishedAt:  nil,
		Content:      request.Content,
	}

	if err := p.http.UpdatePosts(id, &post); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "update data post success",
		"error": nil,
	})
}

func (p *postRouter) DeletePosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("name"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.DeletePosts(id); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data post success",
		"error": nil,
	})
}

func (p *postRouter) PostsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("name"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	post, err := p.http.PostsById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "get data post success",
		"error": nil,
		"data": post,
	})
}

func (p *postRouter) CreateProfile(c *gin.Context) {
	var request model.RequestProfile
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	profile := model.Profiles{
		Id:           p.snowflake.GearedID(),
		UserId:       request.UserId,
		FirstName:    request.FirstName,
		MiddleName:   request.MiddleName,
		LastName:     request.LastName,
		Mobile:       request.Mobile,
		Email:        request.Email,
		RegisteredAt: time.Now(),
		CreatedAt:    time.Now(),
		Profile:      request.Profile,
	}

	if err := p.http.CreateProfile(&profile); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "create data profile success",
		"error": nil,
	})
}

func (p *postRouter) UpdateProfile(c *gin.Context) {
	var request model.RequestProfile
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	profile := model.Profiles{
		UserId:       request.UserId,
		FirstName:    request.FirstName,
		MiddleName:   request.MiddleName,
		LastName:     request.LastName,
		Mobile:       request.Mobile,
		Email:        request.Email,
		Profile:      request.Profile,
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.UpdateProfile(id, &profile); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "update data profile success",
		"error": nil,
	})
}

func (p *postRouter) ProfileById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	profile, err := p.http.ProfileById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "get data profile success",
		"error": nil,
		"data": profile,
	})
}

func (p *postRouter) CreatePostMetas(c *gin.Context) {
	var request model.RequestPostMetas
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	postMetas := model.PostMetas{
		Id:      p.snowflake.GearedID(),
		PostId:  request.PostId,
		TextKey: request.TextKey,
		Content: request.Content,
		CreatedAt:time.Now(),
	}
	if err := p.http.CreatePostMetas(&postMetas); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "200",
		"message": "create data PostMetas success",
		"error": nil,
	})
}

func (p *postRouter) UpdatePostMetas(c *gin.Context) {
	var request model.RequestPostMetas
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	postMetas := model.PostMetas{
		PostId:  request.PostId,
		TextKey: request.TextKey,
		Content: request.Content,
	}

	id ,err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	if err := p.http.UpdatePostMetas(id, &postMetas); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "update data PostMetas success",
		"error": nil,
	})
}

func (p *postRouter) DeletePostMetas(c *gin.Context) {
	id ,err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.DeletePostMetas(id); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data PostMetas success",
		"error": nil,
	})
}

func (p *postRouter) PostMetasById(c *gin.Context) {
	id ,err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	postMetas, err := p.http.PostMetasById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data PostMetas success",
		"error": nil,
		"data": postMetas,
	})
}

func (p *postRouter) CreateComment(c *gin.Context) {
	var request model.RequestPostComments
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	comment := model.PostComments{
		Id:          p.snowflake.GearedID(),
		PostId:      request.PostId,
		ParentId:    request.ParentId,
		Title:       request.Title,
		Published:   0,
		CreatedAt:   time.Now(),
		PublishedAt: time.Now(),
		Content:     request.Content,
	}
	if err := p.http.CreateComment(&comment); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "create data comment success",
		"error": nil,
	})
}

func (p *postRouter) CreateCategories(c *gin.Context) {
	var request model.RequestCategory
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	category := model.Categories{
		Id:        p.snowflake.GearedID(),
		ParentId:  request.ParentId,
		Title:     request.Title,
		MetaTitle: request.MetaTitle,
		Slug:      request.Slug,
		Content:   request.Content,
		CreatedAt: time.Now(),
	}
	if err := p.http.CreateCategories(&category); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "create data category success",
		"error": nil,
	})
}

func (p *postRouter) UpdateCategories(c *gin.Context) {
	var request model.RequestCategory
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	category := model.Categories{
		ParentId:  request.ParentId,
		Title:     request.Title,
		MetaTitle: request.MetaTitle,
		Slug:      request.Slug,
		Content:   request.Content,
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.UpdateCategories(id, &category); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "update data category success",
		"error": nil,
	})
}

func (p *postRouter) DeleteCategories(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	if err := p.http.DeleteCategories(id); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data category success",
		"error": nil,
	})
}

func (p *postRouter) CategoriesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}

	category, err := p.http.CategoriesById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data category success",
		"error": nil,
		"data": category,
	})
}

func (p *postRouter) PostsAll(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var post model.Posts
	posts, err := p.http.PostsAll(&post, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data post success",
		"error": nil,
		"data": posts,
	})
}

func (p *postRouter) ProfileAll(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var profile model.Profiles
	profiles, err := p.http.ProfileAll(&profile, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data profile success",
		"error": nil,
		"data": profiles,
	})
}

func (p *postRouter) PostMetasAll(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var postmeta model.PostMetas
	postmetas, err := p.http.PostMetasAll(&postmeta, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data PostMetas success",
		"error": nil,
		"data": postmetas,
	})
}

func (p *postRouter) CategoriesAll(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var category model.Categories
	categories, err := p.http.CategoriesAll(&category, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data categories success",
		"error": nil,
		"data": categories,
	})
}

func (p *postRouter) CommentAllPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var comment model.PostComments
	comments, err := p.http.CommentAllPost(id,&comment, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data comments success",
		"error": nil,
		"data": comments,
	})
}

func (p *postRouter) PublishPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.PublishPost(id); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "publish data post success",
		"error": nil,
	})
}

func (p *postRouter) RemoveComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	if err := p.http.RemoveComment(id); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "remove data comment success",
		"error": nil,
	})
}

func (p *postRouter) ListPublishPost(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var post model.Posts
	posts, err := p.http.ListPublishPost(&post, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data post success",
		"error": nil,
		"data": posts,
	})
}

func (p *postRouter) ListNotPublishPost(c *gin.Context) {
	pagination := p.pagination.GeneratePaginationFromRequest(c)
	var post model.Posts
	posts, err := p.http.ListNotPublishPost(&post, &pagination)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "delete data post success",
		"error": nil,
		"data": posts,
	})
}