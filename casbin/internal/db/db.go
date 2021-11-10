package db

import (
	"fmt"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/pkg/pagination"
	"gorm.io/gorm"
)

type casbinDB struct {
	db *gorm.DB
	pagination pagination.PaginationPage

}

type CasbinDB interface {
	CasbinRuleAll() (model.Pagination, error)
	CasbinRuleById(id int) (model.CasbinRule, error)
	CreateCasbinRule(casbin model.CasbinRule) error
	DeleteCasbinRule(id int) error
	UpdateCasbinRulePtype(id int,ptype string) error
	UpdateCasbinRuleName(id int,name string) error
	UpdateCasbinRuleEndpoint(id int,endpoint string) error
	UpdateCasbinMethod(id int,method string) error
}

func NewCasbinDB(db *gorm.DB, pagination pagination.PaginationPage) CasbinDB {
	return &casbinDB{db: db, pagination: pagination}
}

func (c *casbinDB) CasbinRuleAll() (model.Pagination, error) {
	var casbinRules []model.CasbinRule
	scope, pagination := c.pagination.Paginate(casbinRules, c.db)
	err := c.db.Scopes(scope).Table("casbin_rule").Scan(&casbinRules)
	pagination.Rows = casbinRules
	if err != nil {
		log.NewLog().Error("not found data")
		return model.Pagination{}, err.Error
	}
	log.NewLog().Info("found data")
	return pagination, nil
}

func (c *casbinDB) CasbinRuleById(id int) (model.CasbinRule, error) {
	var casbinRule model.CasbinRule
	err := c.db.Table("casbin_rule").Where("id = ?",id).Scan(&casbinRule)
	if err != nil {
		log.NewLog().Error("cant't casbin rule get by id data")
		return model.CasbinRule{}, err.Error
	}
	log.NewLog().Info("can casbin rule get by id data")
	return casbinRule, nil
}

func (c *casbinDB) CreateCasbinRule(casbin model.CasbinRule) error {
	stmtString := fmt.Sprintf("INSERT INTO `casbin_rule` (ptype,v0,v1,v2) VALUES ('%s','%s','%s','%s');", casbin.Ptype, casbin.V0, casbin.V1, casbin.V2)
	err := c.db.Exec(stmtString)
	if err != nil {
		log.NewLog().Error("can't create casbin")
		return err.Error
	}
	log.NewLog().Info("create casbin success")
	return nil
}

func (c *casbinDB) DeleteCasbinRule(id int) error {
	err := c.db.Exec("DELETE FROM `casbin_rule` where id = ? ", id)
	if err != nil {
		log.NewLog().Error("delete casbin failed")
		return err.Error
	}
	log.NewLog().Info("delete casbin success")
	return nil
}

func (c *casbinDB) UpdateCasbinRulePtype(id int, ptype string) error {
	err := c.db.Exec("UPDATE `casbin_rule` SET ptype = ?  WHERE id = ?", ptype, id)
	if err != nil {
		log.NewLog().Error("update casbin failed")
		return err.Error
	}
	log.NewLog().Info("update casbin success")
	return nil
}

func (c *casbinDB) UpdateCasbinRuleName(id int, name string) error {
	err := c.db.Exec("UPDATE `casbin_rule` SET v0 = ?  WHERE id = ?", name, id)
	if err != nil {
		log.NewLog().Error("update casbin failed")
		return err.Error
	}
	log.NewLog().Info("update casbin success")
	return nil
}

func (c *casbinDB) UpdateCasbinRuleEndpoint(id int, endpoint string) error {
	err := c.db.Exec("UPDATE `casbin_rule` SET v1 = ? WHERE id = ? ", endpoint, id)
	if err != nil {
		log.NewLog().Error("")
		return err.Error
	}
	log.NewLog().Info("update casbin success")
	return nil
}

func (c *casbinDB) UpdateCasbinMethod(id int, method string) error {
	err := c.db.Exec("UPDATE `casbin_rule` SET v2 = ?  WHERE id = ?", method, id)
	if err != nil {
		log.NewLog().Error("update casbin failed")
		return err.Error
	}
	log.NewLog().Info("update casbin success")
	return nil
}