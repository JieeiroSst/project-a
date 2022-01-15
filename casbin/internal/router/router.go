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

// CasbinRuleAll godoc
// @Summary CasbinRuleAll Permission
// @Description CasbinRuleAll Permission
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin [get]
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

// CasbinRuleById godoc
// @Summary UpdateProfile Permission
// @Description UpdateProfile Permission
// @Accept  json
// @Produce  json
// @Param id path int true "Casbin ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin/:id [get]
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

// CreateCasbinRule godoc
// @Summary UpdateProfile Permission
// @Description UpdateProfile Permission
// @Accept  json
// @Produce  json
// @Param ptype query string false "ptype is p/g in json casbin"
// @Param V0 query string false "v0 is username in json casbin"
// @Param V1 query string false "v1 is URL in json casbin"
// @Param V2 query string false "v2 is method in json casbin"
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin [post]
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

// DeleteCasbinRule godoc
// @Summary DeleteCasbinRule Permission
// @Description DeleteCasbinRule Permission
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin/:id [delete]
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

// UpdateCasbinRulePtype godoc
// @Summary UpdateProfile Permission
// @Description UpdateProfile Permission
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param option query string false "option "
// @Param ptype query string false "ptype is p/g in json casbin"
// @Param V0 query string false "v0 is username in json casbin"
// @Param V1 query string false "v1 is URL in json casbin"
// @Param V2 query string false "v2 is method in json casbin"
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin [put]
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

// OptionList godoc
// @Summary OptionList Permission
// @Description OptionList Permission
// @Accept  json
// @Produce  json
// @Success 200 {array} map[string]interface{}
// @Router /v1/casbin/option [get]
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