package usecase

import (
	"github.com/JieeiroSst/itjob/casbin/internal/repository"
	"github.com/JieeiroSst/itjob/model"
)

type casbinRuleUseCase struct {
	repo repository.CasbinRuleRepository
}

type CasbinRuleUseCase interface {
	CasbinRuleAll() (model.Pagination, error)
	CasbinRuleById(id int) (model.CasbinRule,error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,ptype string) error
	UpdateCasbinRuleName(id int,name string) error
	UpdateCasbinRuleEndpoint(id int,endpoint string) error
	UpdateCasbinMethod(id int,method string) error
}

func NewCasbinRuleUseCase(repo repository.CasbinRuleRepository) CasbinRuleUseCase {
	return &casbinRuleUseCase{
		repo:repo,
	}
}

func(repo *casbinRuleUseCase) CasbinRuleAll() (model.Pagination, error) {
	casbin,err := repo.repo.CasbinRuleAll()
	if err != nil {
		return model.Pagination{}, err
	}
	return casbin, nil
}

func(repo *casbinRuleUseCase) CasbinRuleById(id int) (model.CasbinRule,error) {
	casbin,err:=repo.repo.CasbinRuleById(id)
	if err != nil {
		return model.CasbinRule{}, err
	}
	return casbin, nil
}

func(repo *casbinRuleUseCase) CreateCasbinRule(casbin model.CasbinRule) error {
	if err := repo.repo.CreateCasbinRule(casbin); err != nil {
		return err
	}
	return nil
}

func(repo *casbinRuleUseCase) DeleteCasbinRule(id int) error {
	if err:= repo.repo.DeleteCasbinRule(id); err != nil {
		return err
	}
	return nil
}

func(repo *casbinRuleUseCase) UpdateCasbinRulePtype(id int,ptype string) error {
	if err := repo.repo.UpdateCasbinRulePtype(id,ptype); err != nil {
		return err
	}
	return nil
}

func(repo *casbinRuleUseCase) UpdateCasbinRuleName(id int,name string) error {
	if err := repo.repo.UpdateCasbinRuleName(id,name); err != nil {
		return err
	}
	return nil
}

func(repo *casbinRuleUseCase) UpdateCasbinRuleEndpoint(id int,endpoint string) error {
	if err := repo.repo.UpdateCasbinRuleEndpoint(id,endpoint); err != nil {
		return err
	}
	return nil
}

func(repo *casbinRuleUseCase) UpdateCasbinMethod(id int,method string) error {
	if err := repo.repo.UpdateCasbinMethod(id,method); err != nil {
		return err
	}
	return nil
}