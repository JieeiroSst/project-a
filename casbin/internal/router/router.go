package router

import (
	deliveryHttp "github.com/JieeiroSst/itjob/casbin/internal/delivery/http"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type casbinRouter struct {
	deliveryHttp deliveryHttp.DeliveryHttp
	snowflakeData snowflake.SnowflakeData
}

type CasbinRouter interface {
	CasbinRuleAll(c *gin.Context)
	CasbinRuleById(c *gin.Context)
	CreateCasbinRule(c *gin.Context)
	DeleteCasbinRule(c *gin.Context)
	UpdateCasbinRulePtype(c *gin.Context)
	OptionList(c *gin.Context)
}

func NewCasbinRouter(deliveryHttp deliveryHttp.DeliveryHttp, snowflakeData snowflake.SnowflakeData) CasbinRouter {
	return &casbinRouter{
		deliveryHttp: deliveryHttp,
		snowflakeData: snowflakeData,
	}
}

func (r *casbinRouter) CasbinRuleAll(c *gin.Context) {
	casbin, err := r.deliveryHttp.CasbinRuleAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": casbin,
	})
}

func (r *casbinRouter) CasbinRuleById(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	casbin, err := r.deliveryHttp.CasbinRuleById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": casbin,
	})
}


func (r *casbinRouter) CreateCasbinRule(c *gin.Context) {
	var casbin model.CasbinRule
	if err := c.ShouldBind(&casbin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	casbinRule := model.CasbinRule{
		ID:    r.snowflakeData.GearedID(),
		Ptype: casbin.Ptype,
		V0:    casbin.V0,
		V1:    casbin.V1,
		V2:    casbin.V2,
	}


	if err := r.deliveryHttp.CreateCasbinRule(casbinRule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create casbin rule success",
	})
}

func (r *casbinRouter) DeleteCasbinRule(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.deliveryHttp.DeleteCasbinRule(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete casbin rule success",
		"id": id,
	})
}

func (r *casbinRouter) UpdateCasbinRulePtype(c *gin.Context) {
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := c.Query("name")
	option := c.Query("option")

	if err  := r.deliveryHttp.UpdateCasbinRulePtype(id, name, option); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update casbin rule success",
		"id": id,
	})
}

func (r *casbinRouter) OptionList(c *gin.Context) {
	option := make( map[int]string)
	option[1] = "Ptype"
	option[2] = "Name"
	option[3] = "EndPoint"
	option[4] = "Method"

	c.JSON(http.StatusOK, gin.H{
		"data": option,
	})
}