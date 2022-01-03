package migrate

import (
	"github.com/JieeiroSst/itjob/model"
	"gorm.io/gorm"
)

type migrate struct {
	db *gorm.DB
}

type Migrate interface {
	AutoMigrate() error
}

func NewAutoMigrate(db *gorm.DB) Migrate {
	return &migrate{
		db:db,
	}
}

func (m *migrate) AutoMigrate() error {
	if err := m.db.AutoMigrate(&model.Users{},&model.Posts{},
								&model.Email{},&model.Image{},
								&model.Ip{}); err != nil {
		return err
	}
	return nil
}