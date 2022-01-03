package migrate

import (
	"fmt"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/migrate"
	"github.com/JieeiroSst/itjob/pkg/mysql"
)

type autoMigrate struct {
	config config.Config
}

type Migrate interface {
	AutoMigrate() error
}

func NewAutoMigrate(config config.Config) Migrate {
	return &autoMigrate{
		config:config,
	}
}

func (m *autoMigrate) AutoMigrate() error {
	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.config.Mysql.MysqlUser,
		m.config.Mysql.MysqlPassword,
		m.config.Mysql.MysqlHost,
		m.config.Mysql.MysqlPort,
		m.config.Mysql.MysqlDbname,
	)
	mysqlOrm:= mysql.NewMysqlConn(dns)
	newAutoMigrate := migrate.NewAutoMigrate(mysqlOrm)

	if err := newAutoMigrate.AutoMigrate(&model.Users{},&model.Posts{},
								&model.Email{},&model.Image{},
								&model.Ip{}); err != nil {
		return err
	}

	return nil
}