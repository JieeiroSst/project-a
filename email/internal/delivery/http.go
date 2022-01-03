package delivery

import (
	"github.com/JieeiroSst/itjob/email/internal/usecase"
	"github.com/JieeiroSst/itjob/model"
)

type emailHttp struct {
	usecase usecase.EmailUsecase
}

type EmailHttp interface {
	CreateSendEmail(email model.Email) error
}

func NewEmailHttp(usecase usecase.EmailUsecase) EmailHttp {
	return &emailHttp{
		usecase:usecase,
	}
}

func (e *emailHttp) CreateSendEmail(email model.Email) error {
	if err := e.usecase.CreateSendEmail(email); err != nil {
		return err
	}
	return nil
}