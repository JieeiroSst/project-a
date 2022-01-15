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

// CreatePosts godoc
// @Summary CreatePosts Account
// @Description CreatePosts account
// @Accept  json
// @Produce  json
// @Param AuthorId query string false "AuthorId in json post"
// @Param Title query string false "Title in json post"
// @Param MetaTitle query string false "MetaTitle in json post"
// @Param Slug query string false "Slug in json post"
// @Param Summary query string false "Summary in json post"
// @Param Content query string false "Content in json post"
// @Success 200 {array} map[string]interface{}
// @Router /v1/write/post [post]
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

// UpdatePosts godoc
// @Summary UpdatePosts Account
// @Description UpdatePosts account
// @Accept  json
// @Produce  json
// @Param id query int true "Post ID"
// @Param AuthorId query string false "AuthorId in json post"
// @Param Title query string false "Title in json post"
// @Param MetaTitle query string false "MetaTitle in json post"
// @Param Slug query string false "Slug in json post"
// @Param Summary query string false "Summary in json post"
// @Param Content query string false "Content in json post"
// @Success 200 {array} map[string]interface{}
// @Router /v1/writer/post [put]
func (p *postRouter) UpdatePosts(c *gin.Context) {
	var request model.RequestPost
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(c.Query("id"))
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

// DeletePosts godoc
// @Summary DeletePosts Account
// @Description DeletePosts account
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/write/post [delete]
func (p *postRouter) DeletePosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
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

// PostsById godoc
// @Summary PostsById Account
// @Description PostsById account
// @Accept  json
// @Produce  json
// @Param id query int true "Post ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/writer/post [post]
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

// CreateProfile godoc
// @Summary CreateProfile Account
// @Description CreateProfile account
// @Accept  json
// @Produce  json
// @Param UserId query string false "username in json profile"
// @Param FirstName query string false "FirstName in json profile"
// @Param MiddleName query string false "MiddleName in json profile"
// @Param LastName query string false "LastName in json profile"
// @Param Mobile query string false "Mobile in json profile"
// @Param Email query string false "Email in json profile"
// @Param Profile query string false "Profile in json profile"
// @Success 200 {array} map[string]interface{}
// @Router /v1/writer/profile [post]
// @Router /v1/client/profile [post]
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

// UpdateProfile godoc
// @Summary UpdateProfile Account
// @Description UpdateProfile account
// @Accept  json
// @Produce  json
// @Param id query int true "Profile ID"
// @Param UserId query string false "username in json profile"
// @Param FirstName query string false "FirstName in json profile"
// @Param MiddleName query string false "MiddleName in json profile"
// @Param LastName query string false "LastName in json profile"
// @Param Mobile query string false "Mobile in json profile"
// @Param Email query string false "Email in json profile"
// @Param Profile query string false "Profile in json profile"
// @Success 200 {array} map[string]interface{}
// @Router /v1/writer/profile [put]
// @Router /v1/client/profile [put]
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

// ProfileById godoc
// @Summary ProfileById Account
// @Description ProfileById account
// @Accept  json
// @Produce  json
// @Param id query int true "profile ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/write/profile [get]
// @Router /v1/client/profile [get]
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

// CreatePostMetas godoc
// @Summary CreatePostMetas Account
// @Description CreatePostMetas account
// @Accept  json
// @Produce  json
// @Param PostId query string false "PostId in json post-metas"
// @Param TextKey query string false "TextKey in json post-metas"
// @Param Content query string false "Content in json post-metas"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post-metas [post]
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

// UpdatePostMetas godoc
// @Summary UpdatePostMetas Account
// @Description UpdatePostMetas account
// @Accept  json
// @Produce  json
// @Param id query int true "Post-metas ID"
// @Param PostId query string false "PostId in json post-metas"
// @Param TextKey query string false "TextKey in json post-metas"
// @Param Content query string false "Content in json post-metas"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post-metas [put]
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

// DeletePostMetas godoc
// @Summary DeletePostMetas Account
// @Description DeletePostMetas account
// @Accept  json
// @Produce  json
// @Param id query int true "Post-metas ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post-metas [delete]
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

// PostMetasById godoc
// @Summary PostMetasById Account
// @Description PostMetasById account
// @Accept  json
// @Produce  json
// @Param id path int true "Post-metas ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post-metas [get]
func (p *postRouter) PostMetasById(c *gin.Context) {
	id ,err := strconv.Atoi(c.Param("id"))
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

// CreateComment godoc
// @Summary CreateComment Account
// @Description CreateComment account
// @Accept  json
// @Produce  json
// @Param PostId query string false "PostId in json Post"
// @Param ParentId query string false "ParentId in json Post"
// @Param Title query string false "Title in json Post"
// @Param Content query string false "Content in json Post"
// @Success 200 {array} map[string]interface{}
// @Router /v1/client/comment [post]
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

// CreateCategories godoc
// @Summary CreateCategories Account
// @Description CreateCategories account
// @Accept  json
// @Produce  json
// @Param ParentId query string false "ParentId in json category"
// @Param Title query string false "Title in json category"
// @Param MetaTitle query string false "MetaTitle in json category"
// @Param Slug query string false "Slug in json category"
// @Param Content query string false "Content in json category"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/category [post]
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

// UpdateCategories godoc
// @Summary UpdateCategories Account
// @Description UpdateCategories account
// @Accept  json
// @Produce  json
// @Param id query int true "Category ID"
// @Param ParentId query string false "ParentId in json category"
// @Param Title query string false "Title in json category"
// @Param MetaTitle query string false "MetaTitle in json category"
// @Param Slug query string false "Slug in json category"
// @Param Content query string false "Content in json category"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/category [put]
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

// DeleteCategories godoc
// @Summary DeleteCategories Account
// @Description DeleteCategories account
// @Accept  json
// @Produce  json
// @Param id query int true "category ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/category [delete]
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

// CategoriesById godoc
// @Summary CategoriesById Account
// @Description CategoriesById account
// @Accept  json
// @Produce  json
// @Param id query int true "category ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/category [get]
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

// PostsAll godoc
// @Summary PostsAll Account
// @Description PostsAll account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/guest/post [get]
// @Router /v1/client/post [get]
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

// ProfileAll godoc
// @Summary ProfileAll Account
// @Description ProfileAll account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/profile [get]
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

// PostMetasAll godoc
// @Summary PostMetasAll Account
// @Description PostMetasAll account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post-metas [get]
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

// CategoriesAll godoc
// @Summary CategoriesAll Account
// @Description CategoriesAll account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/category [get]
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

// CommentAllPost godoc
// @Summary CommentAllPost Account
// @Description CommentAllPost account
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/client/comment [get]
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

// PublishPost godoc
// @Summary PublishPost Account
// @Description PublishPost account
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin//post/publish [post]
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

// RemoveComment godoc
// @Summary RemoveComment Account
// @Description RemoveComment account
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/client/comment [delete]
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

// ListPublishPost godoc
// @Summary ListPublishPost Account
// @Description ListPublishPost account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post/list/publish [get]
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

// ListNotPublishPost godoc
// @Summary ListNotPublishPost Account
// @Description ListNotPublishPost account
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/admin/post/list/not/publish [get]
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