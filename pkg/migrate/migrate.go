package migrate

import (
	"gorm.io/gorm"
)

type migrate struct {
	db *gorm.DB
}

type Migrate interface {
	AutoMigrate(table ...interface{}) error
}

func NewAutoMigrate(db *gorm.DB) Migrate {
	return &migrate{
		db:db,
	}
}

func (m *migrate) AutoMigrate(table ...interface{}) error {
	if err := m.db.AutoMigrate(&table); err != nil {
		return err
	}
	return nil
}