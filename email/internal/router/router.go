package router

import (
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/email/internal/delivery"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
	"strconv"
	"time"
)

type emailRouter struct {
	http      delivery.EmailHttp
	config    config.Config
	snowflake snowflake.SnowflakeData
}

type EmailRouter interface {
	UserSendEmail(c *gin.Context)
	AdminSendEmail(c *gin.Context)
}

func NewEmailRouter(http delivery.EmailHttp,config config.Config,snowflake snowflake.SnowflakeData) EmailRouter {
	return &emailRouter{
		http:http,
		config:config,
		snowflake:snowflake,
	}
}

func (e *emailRouter) UserSendEmail(c *gin.Context) {
	var requestEmail model.RequestEmail
	if err :=c.ShouldBind(requestEmail); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"message": err.Error(),

		})
		return
	}
	email := model.Email{
		Id:               e.snowflake.GearedID(),
		NameEmailSend:    requestEmail.NameEmailSend,
		NameEmailReceive: requestEmail.NameEmailReceive,
		SubjectEmail:     requestEmail.SubjectEmail,
		Content:          requestEmail.Content,
		CreatedAt:        time.Now(),
	}
	if err := e.http.CreateSendEmail(email); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	gogmail := gomail.NewMessage()
	gogmail.SetHeader("From", requestEmail.NameEmailSend)
	gogmail.SetHeader("To", requestEmail.NameEmailReceive)
	gogmail.SetHeader("Subject", requestEmail.SubjectEmail)
	gogmail.SetBody("text/plain", requestEmail.Content)

	port, err := strconv.Atoi(e.config.Email.Port)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}
	a := gomail.NewDialer(e.config.Email.Host, port, e.config.Email.NameEmail, e.config.Email.PasswordEmail)
	if err := a.DialAndSend(gogmail); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "send email success",
		"from":    requestEmail.NameEmailSend,
		"to":      requestEmail.NameEmailReceive,
	})
}

func (e *emailRouter) AdminSendEmail(c *gin.Context)  {
	var requestEmail model.RequestEmail
	if err :=c.ShouldBind(requestEmail); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"message": err.Error(),

		})
		return
	}
	email := model.Email{
		Id:               e.snowflake.GearedID(),
		NameEmailSend:    e.config.Email.NameEmail,
		NameEmailReceive: requestEmail.NameEmailReceive,
		SubjectEmail:     requestEmail.SubjectEmail,
		Content:          requestEmail.Content,
		CreatedAt:        time.Now(),
	}
	if err := e.http.CreateSendEmail(email); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),

		})
		return
	}

	gogmail := gomail.NewMessage()
	gogmail.SetHeader("From",e.config.Email.NameEmail)
	gogmail.SetHeader("To", requestEmail.NameEmailReceive)
	gogmail.SetHeader("Subject", requestEmail.SubjectEmail)
	gogmail.SetBody("text/plain", requestEmail.Content)

	port, err := strconv.Atoi(e.config.Email.Port)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"message":err.Error(),
		})
		return
	}
	a := gomail.NewDialer(e.config.Email.Host, port, e.config.Email.NameEmail, e.config.Email.PasswordEmail)
	if err := a.DialAndSend(gogmail); err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 	200,
		"message": 	"send email success",
		"from": 	e.config.Email.NameEmail,
		"to": 		requestEmail.NameEmailReceive,
	})
}