package router

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	deliveryHttp "github.com/JieeiroSst/itjob/upload/internal/delivery/http"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type uploadRouter struct {
	deliveryHttp deliveryHttp.UploadDeliveryHttp
	snowflakeData snowflake.SnowflakeData
	session *session.Session
}

type UploadRouter interface {
	AddFileS3(c *gin.Context)
	ReadFile(c *gin.Context)
	Option(c *gin.Context)
}

func NewUploadRouter(deliveryHttp deliveryHttp.UploadDeliveryHttp, snowflakeData snowflake.SnowflakeData) UploadRouter {
	return &uploadRouter{
		deliveryHttp:  deliveryHttp,
		snowflakeData: snowflakeData,
	}
}

func (u *uploadRouter) AddFileS3(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	option := c.Query("option")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	image := model.Image{
		Id:          u.snowflakeData.GearedID(),
		Name:        file.Filename,
		UserRefer:   id,
		Option:      model.Option(option),
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}

	err = u.deliveryHttp.AddFileS3(u.session, file.Filename, image, option)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status":"200",
		"message":"upload file success",
	})
}

func (u *uploadRouter) ReadFile(c *gin.Context) {
	data := c.Query("data")
	image, err := u.deliveryHttp.ReadFile(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": image,
	})
}

func (u *uploadRouter) Option(c *gin.Context) {
	var option = map[int]string {
		1: "AVATAR",
		2: "NEWS",
	}

	c.JSON(http.StatusOK, gin.H{
		"option": option,
	})
}