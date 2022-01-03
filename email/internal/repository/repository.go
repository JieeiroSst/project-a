package repository

import (
	"github.com/JieeiroSst/itjob/model"
	"gorm.io/gorm"
)

type emailRepository struct {
	db *gorm.DB
}

type EmailRepository interface {
	CreateSendEmail(email model.Email) error
}

func NewEmailRepository(db *gorm.DB) EmailRepository {
	return &emailRepository{
		db:db,
	}
}

func (e *emailRepository) CreateSendEmail(email model.Email) error {
	if err := e.db.Create(&email).Error; err != nil {
		return err
	}
	return nil
}
