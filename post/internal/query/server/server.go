package server

import (
	"github.com/JieeiroSst/itjob/access_control"
	_ "github.com/JieeiroSst/itjob/post/internal/query/docs"
	"github.com/JieeiroSst/itjob/post/internal/query/router"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type postServer struct {
	accessControl access_control.Authorization
	router 		  router.PostRouter
	engine        *gin.Engine
	db 			  *gorm.DB
}


type PostServer interface {
	RunServer() error
}

func NewPostServer(accessControl access_control.Authorization, router router.PostRouter, engine *gin.Engine,db *gorm.DB) PostServer {
	return &postServer{
		accessControl: accessControl,
		router:        router,
		engine:        engine,
		db:			   db,
	}
}

func (p *postServer) RunServer() error {
	adapter, err:=gormadapter.NewAdapterByDB(p.db)
	if err != nil {
		return err
	}

	resource := p.engine.Group("/v1")
	url := ginSwagger.URL("http://localhost:3000/swagger/post/doc.json") // The url pointing to API definition
	resource.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	resourceGuest := resource.Group("/guest")

	resourceGuest.GET("/post", p.router.PostsAll)


	resource.Use(p.accessControl.Authenticate())
	{
		resourceWriter := resource.Group("/writer")

		resourceWriter.POST("/post", p.accessControl.Authorize("/v1/writer/*","POST", adapter), p.router.CreatePosts)
		resourceWriter.PUT("/post", p.accessControl.Authorize("/v1/writer/*","PUT", adapter), p.router.UpdatePosts)
		resourceWriter.DELETE("/post", p.accessControl.Authorize("/v1/writer/*","DELETE", adapter), p.router.DeletePosts)
		resourceWriter.GET("/post", p.accessControl.Authorize("/v1/writer/*","GET", adapter), p.router.PostsById)


		resourceWriter.POST("/profile", p.accessControl.Authorize("/v1/writer/*","POST", adapter), p.router.CreateProfile)
		resourceWriter.PUT("/profile", p.accessControl.Authorize("/v1/writer/*","PUT", adapter), p.router.UpdateProfile)
		resourceWriter.GET("/profile", p.accessControl.Authorize("/v1/writer/*","GET", adapter), p.router.ProfileById)


		resourceAdmin := resource.Group("/admin")

		resourceAdmin.POST("/post-metas", p.accessControl.Authorize("/v1/admin/*","POST", adapter), p.router.CreatePostMetas)
		resourceAdmin.PUT("/post-metas", p.accessControl.Authorize("/v1/admin/*","PUT", adapter), p.router.UpdatePostMetas)
		resourceAdmin.DELETE("/post-metas", p.accessControl.Authorize("/v1/admin/*","DELETE", adapter), p.router.DeletePostMetas)
		resourceAdmin.GET("/post-metas/:id", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.PostMetasById)
		resourceAdmin.GET("/post-metas", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.PostMetasAll)


		resourceAdmin.POST("/category", p.accessControl.Authorize("/v1/admin/*","POST", adapter), p.router.CreateCategories)
		resourceAdmin.PUT("/category", p.accessControl.Authorize("/v1/admin/*","PUT", adapter), p.router.UpdateCategories)
		resourceAdmin.DELETE("/category", p.accessControl.Authorize("/v1/admin/*","DELETE", adapter), p.router.DeleteCategories)
		resourceAdmin.GET("/category", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.CategoriesById)

		resourceAdmin.GET("/profile", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.ProfileAll)


		resourceAdmin.POST("/post/publish", p.accessControl.Authorize("/v1/admin/*","POST", adapter), p.router.PublishPost)

		resourceAdmin.GET("/post/list/publish", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.ListPublishPost)
		resourceAdmin.GET("/post/list/not/publish", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.ListNotPublishPost)

		resourceAdmin.GET("/category", p.accessControl.Authorize("/v1/admin/*","GET", adapter), p.router.CategoriesAll)

		resourceClient := resource.Group("/client")

		resourceClient.GET("/post", p.accessControl.Authorize("/v1/client/*","GET", adapter), p.router.PostsAll)

		resourceClient.POST("/comment", p.accessControl.Authorize("/v1/client/*","POST", adapter), p.router.CreateComment)

		resourceClient.POST("/profile", p.accessControl.Authorize("/v1/client/*","POST", adapter), p.router.CreateProfile)
		resourceClient.PUT("/profile", p.accessControl.Authorize("/v1/client/*","PUT", adapter), p.router.UpdateProfile)
		resourceClient.GET("/profile", p.accessControl.Authorize("/v1/client/*","GET", adapter), p.router.ProfileById)

		resourceClient.GET("/comment", p.accessControl.Authorize("/v1/client/*","GET", adapter), p.router.CommentAllPost)
		resourceClient.DELETE("/comment", p.accessControl.Authorize("/v1/client/*","GET", adapter), p.router.RemoveComment)

	}
	

	return nil
}