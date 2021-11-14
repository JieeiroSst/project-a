package http

import (
	"errors"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/upload/internal/usecase"
	"github.com/aws/aws-sdk-go/aws/session"
)

type uploadHttp struct {
	usecase usecase.UploadUsecase
}

type UploadHttp interface {
	AddFileS3(s *session.Session, fileDir string, image model.Image, option string) error
	ReadFile(data string) (model.Image, error)
}

func NewUploadHttp(usecase usecase.UploadUsecase) UploadHttp {
	return &uploadHttp{usecase:usecase}
}

func(u *uploadHttp) AddFileS3(s *session.Session, fileDir string, image model.Image, option string) error {
	switch option {
	case "stream":
		if err := u.usecase.AddFileToS3Stream(s,fileDir,image); err != nil {
			return err
		}
	case "no_stream":
		if err := u.usecase.AddFileToS3(s,fileDir,image); err != nil {
			return err
		}
	default:
		return errors.New("select option is not suitable")
	}
	return nil
}

func(u *uploadHttp) ReadFile(data string) (model.Image, error) {
	email,image, err := u.usecase.ReadFile(data)
	if err != nil {
		return model.Image{}, err
	}

	if len(email) ==0 {
		return model.Image{}, errors.New("image name could not be found")
	}
	return image, nil
}