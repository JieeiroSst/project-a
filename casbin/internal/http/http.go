package http

import (
	"errors"
	"github.com/JieeiroSst/itjob/casbin/internal/usecase"
	"github.com/JieeiroSst/itjob/model"
)

type http struct {
	usecase usecase.CasbinRuleUseCase
}

type Http interface {
	CasbinRuleAll() (model.Pagination, error)
	CasbinRuleById(id int) (model.CasbinRule,error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,name string, option string) error
}

func NewHttp(usecase usecase.CasbinRuleUseCase) Http{
	return &http{
		usecase: usecase,
	}
}

func(http *http) CasbinRuleAll() (model.Pagination, error) {
	casbin,err:=http.usecase.CasbinRuleAll()
	return casbin, err
}

func(http *http) CasbinRuleById(id int) (model.CasbinRule, error) {
	casbin,err:=http.usecase.CasbinRuleById(id)
	return casbin,err
}

func(http *http) CreateCasbinRule(casbin model.CasbinRule) error {
	return http.usecase.CreateCasbinRule(casbin)
}

func(http *http) DeleteCasbinRule(id int) error {
	return http.usecase.DeleteCasbinRule(id)
}

func (http *http) UpdateCasbinRulePtype(id int,name string, option string) error {
	switch option  {
	case "Ptype":
		if err := http.usecase.UpdateCasbinRulePtype(id, name); err != nil {
			return err
		}
	case "Name":
		if err := http.usecase.UpdateCasbinRuleName(id, name); err != nil {
			return err
		}
	case "EndPoint":
		if err := http.usecase.UpdateCasbinRuleEndpoint(id, name); err != nil {
			return err
		}
	case "Method":
		if err := http.usecase.UpdateCasbinMethod(id, name); err != nil {
			return err
		}
	default:
		return errors.New("condition is not satisfied option")
	}
	return nil
}