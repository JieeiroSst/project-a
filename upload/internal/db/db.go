package db

import (
	"github.com/JieeiroSst/itjob/model"
	"gorm.io/gorm"
)

type uploadDB struct {
	db *gorm.DB
}

type UploadDB interface {
	SaveImageForUser(image model.Image) error
	FindByIdImage(name string) (model.Image, error)
}

func NewUploadDB(db *gorm.DB) UploadDB {
	return &uploadDB{db:db}
}

func (u *uploadDB) SaveImageForUser(image model.Image) error {
	if err := u.db.Create(&image).Error; err != nil {
		return err
	}
	return nil
}

func (u *uploadDB) FindByIdImage(name string) (model.Image, error) {
	var image model.Image
	if err := u.db.Where("name = ?", name).Find(&image).Error; err != nil {
		return model.Image{}, err
	}
	return image, nil
}