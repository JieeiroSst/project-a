package cmd

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/db"
	"github.com/JieeiroSst/itjob/pkg/mysql"
)

type dbCMD struct {
	config *config.Config
}

type DbCMD interface {
	Run() error
}

func NewDbCMD(config *config.Config) DbCMD {
	return &dbCMD{config:config}
}

func (d *dbCMD) Run() error {
	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.config.Mysql.MysqlUser,
		d.config.Mysql.MysqlPassword,
		d.config.Mysql.MysqlHost,
		d.config.Mysql.MysqlPort,
		d.config.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)

	db := db.NewAutoMigrate(mysqlOrm)

	if err := db.RunAutoMigrate(); err != nil {
		return err
	}
	return nil
}

