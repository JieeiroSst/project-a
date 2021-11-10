package repository

import (
	"github.com/JieeiroSst/itjob/casbin/internal/db"
	"github.com/JieeiroSst/itjob/model"
)

type casbinRuleRepository struct {
	db db.CasbinDB
}

type CasbinRuleRepository interface {
	CasbinRuleAll() (model.Pagination, error)
	CasbinRuleById(id int) (model.CasbinRule, error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,ptype string) error
	UpdateCasbinRuleName(id int,name string) error
	UpdateCasbinRuleEndpoint(id int,endpoint string) error
	UpdateCasbinMethod(id int,method string) error
}

func NewCasbinRuleRepository(db db.CasbinDB) CasbinRuleRepository{
	return &casbinRuleRepository{
		db:db,
	}
}

func (repo *casbinRuleRepository) CasbinRuleAll() (model.Pagination, error){
	pagination, err := repo.db.CasbinRuleAll()
	if err != nil {
		return model.Pagination{}, err
	}
	return pagination, nil
}

func (repo *casbinRuleRepository) CasbinRuleById(id int) (model.CasbinRule, error){
	casbin, err := repo.db.CasbinRuleById(id)
	if err != nil {
		return model.CasbinRule{}, nil
	}
	return casbin, err
}

func (repo *casbinRuleRepository)  UpdateCasbinRulePtype(id int,ptype string) error {
	if err := repo.db.UpdateCasbinRulePtype(id, ptype); err != nil {
		return err
	}
	return nil
}

func (repo *casbinRuleRepository)  UpdateCasbinRuleName(id int,name string) error {
	if err := repo.db.UpdateCasbinRuleName(id, name); err != nil {
		return err
	}
	return nil
}
func (repo *casbinRuleRepository)  UpdateCasbinRuleEndpoint(id int,endpoint string) error {
	if err := repo.db.UpdateCasbinRuleEndpoint(id, endpoint); err != nil {
		return err
	}
	return nil
}
func (repo *casbinRuleRepository)  UpdateCasbinMethod(id int,method string) error {
	if err := repo.db.UpdateCasbinMethod(id, method); err != nil {
		return err
	}
	return nil
}

func (repo *casbinRuleRepository) CreateCasbinRule(casbin model.CasbinRule) error {
	if err := repo.db.CreateCasbinRule(casbin); err != nil {
		return err
	}
	return nil
}

func (repo *casbinRuleRepository) DeleteCasbinRule(id int) error {
	if err := repo.db.DeleteCasbinRule(id); err != nil {
		return err
	}
	return nil
}