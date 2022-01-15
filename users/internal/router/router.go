package router

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	deliveryHttp "github.com/JieeiroSst/itjob/users/internal/delivery/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type userRouter struct {
	deliveryHttp deliveryHttp.DeliveryHttp
	snowflakeData snowflake.SnowflakeData
}

type UserRouter interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	UpdateProfile(c *gin.Context)
	LockAccount(c *gin.Context)
}

func NewRouter(deliveryHttp deliveryHttp.DeliveryHttp,snowflakeData snowflake.SnowflakeData) UserRouter {
	return &userRouter{
		deliveryHttp: deliveryHttp,
		snowflakeData: snowflakeData,
	}
}

// Login godoc
// @Summary Login Account
// @Description login account
// @Accept  json
// @Produce  json
// @Param username query string false "username in json login"
// @Param password query string false "password in json login"
// @Success 200 {array} map[string]interface{}
// @Router /v1/login [post]
func (r *userRouter) Login(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.Users{
		Username: login.Username,
		Password: login.Password,
	}
	id, token, err := r.deliveryHttp.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"token": "couldn't find the token just created "})
		return
	}
	ip := model.Ip{
		Id:        r.snowflakeData.GearedID(),
		Ip:        c.ClientIP(),
		Method:    "user login :" + login.Username + time.Now().String(),
		RequestAt: time.Now(),
	}
	if err := r.deliveryHttp.RequestIpComputer(ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"token": token,
		"message": "login success",
	})
}

// SignUp godoc
// @Summary SignUp Account
// @Description SignUp account
// @Accept  json
// @Produce  json
// @Param username query string false "username in json login"
// @Param password query string false "password in json login"
// @Success 200 {array} map[string]interface{}
// @Router /v1/register [post]
func (r *userRouter) SignUp(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataUser:=model.Users{
		Id:       	r.snowflakeData.GearedID(),
		Username: 	user.Username,
		Password: 	user.Password,
		Checked:  	true,
		CreateTime: time.Now(),
	}
	if err := r.deliveryHttp.SignUp(dataUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ip := model.Ip{
		Id:        r.snowflakeData.GearedID(),
		Ip:        c.ClientIP(),
		Method:    "user signup :" + user.Username + time.Now().String(),
		RequestAt: time.Now(),
	}
	if err := r.deliveryHttp.RequestIpComputer(ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"mesage": "signup success",
	})
}

// UpdateProfile godoc
// @Summary UpdateProfile Account
// @Description UpdateProfile account
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param username query string false "username in json login"
// @Param password query string false "password in json login"
// @Success 200 {array} map[string]interface{}
// @Router /v1//update/profile [post]
func (r *userRouter) UpdateProfile(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.Users
	user.UpdateTime = time.Now()
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.deliveryHttp.UpdateProfile(id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ip := model.Ip{
		Id:        r.snowflakeData.GearedID(),
		Ip:        c.ClientIP(),
		Method:    "user update profile :" + user.Username + time.Now().String(),
		RequestAt: time.Now(),
	}
	if err := r.deliveryHttp.RequestIpComputer(ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"mesage": "update profile user success",
	})
}

// LockAccount godoc
// @Summary LockAccount Account
// @Description LockAccount account
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1//update/profile [post]
func (r *userRouter) LockAccount(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.deliveryHttp.LockAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ip := model.Ip{
		Id:        r.snowflakeData.GearedID(),
		Ip:        c.ClientIP(),
		Method:    "user lock time is:" + time.Now().String(),
		RequestAt: time.Now(),
	}
	if err := r.deliveryHttp.RequestIpComputer(ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"mesage": "lock user success",
	})
}