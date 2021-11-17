package db

import (
	"gorm.io/gorm"
	"github.com/JieeiroSst/itjob/db/migration"
)

type autoMigrate struct {
	db *gorm.DB
}

type AutoMigrate interface {
	RunAutoMigrate() error
}

func NewAutoMigrate(db *gorm.DB) AutoMigrate {
	return &autoMigrate{
			db:db,
		}
}

func (a *autoMigrate) RunAutoMigrate() error {
	t := migration.Migrate1636995600add_table(a.db)

	if err := t.Migrate(); err != nil {
		return err
	}
	return nil
}