package http

import (
	"github.com/JieeiroSst/itjob/casbin/internal/http"
	"github.com/JieeiroSst/itjob/model"
)

type deliveryHttp struct {
	http http.Http
}

type DeliveryHttp interface {
	CasbinRuleAll() (model.Pagination, error)
	CasbinRuleById(id int) (model.CasbinRule,error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,name string, option string) error
}

func NewDeliveryHttp(http http.Http) DeliveryHttp {
	return &deliveryHttp{http:http}
}

func (h *deliveryHttp) CasbinRuleAll() (model.Pagination, error) {
	casbins, err := h.http.CasbinRuleAll()
	if err != nil {
		return model.Pagination{}, err
	}
	return casbins, nil
}

func (h *deliveryHttp) CasbinRuleById(id int) (model.CasbinRule,error) {
	casbin, err := h.http.CasbinRuleById(id)
	if err != nil {
		return model.CasbinRule{}, err
	}
	return casbin, nil
}

func (h *deliveryHttp) CreateCasbinRule(casbin model.CasbinRule) error {
	if err := h.http.CreateCasbinRule(casbin); err != nil {
		return err
	}
	return nil
}

func (h *deliveryHttp) DeleteCasbinRule(id int) error {
	if err := h.http.DeleteCasbinRule(id); err != nil {
		return err
	}
	return nil
}

func (h *deliveryHttp) UpdateCasbinRulePtype(id int,name string, option string) error {
	if err := h.http.UpdateCasbinRulePtype(id, name, option); err != nil {
		return err
	}
	return nil
}