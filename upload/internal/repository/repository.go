package repository

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/upload/internal/db"
)

type uploadRepository struct {
	db db.UploadDB
}

type UploadRepository interface {
	SaveImageForUser(image model.Image) error
	FindByIdImage(name string) (model.Image, error)
}

func NewUploadRepository(db db.UploadDB) UploadRepository {
	return &uploadRepository{db:db}
}

func (u *uploadRepository) SaveImageForUser(image model.Image) error {
	if err := u.db.SaveImageForUser(image); err != nil {
		return err
	}
	return nil
}

func (u *uploadRepository) FindByIdImage(name string) (model.Image, error) {
	image, err := u.db.FindByIdImage(name)
	if err != nil {
		return model.Image{}, nil
	}
	return image, nil
}