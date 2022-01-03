package usecase

import (
	"github.com/JieeiroSst/itjob/email/internal/repository"
	"github.com/JieeiroSst/itjob/model"
)

type emailUsecase struct {
	repository repository.EmailRepository
}

type EmailUsecase interface {
	CreateSendEmail(email model.Email) error
}

func NewEmailUsecase(repository repository.EmailRepository) EmailUsecase {
	return &emailUsecase{
		repository:repository,
	}
}

func (e *emailUsecase) CreateSendEmail(email model.Email) error {
	if err := e.repository.CreateSendEmail(email); err != nil {
		return err
	}
	return nil
}
