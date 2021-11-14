package http

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/JieeiroSst/itjob/upload/internal/http"
)

type uploadDeliveryHttp struct {
	http http.UploadHttp
}

type UploadDeliveryHttp interface {
	AddFileS3(s *session.Session, fileDir string, image model.Image, option string) error
	ReadFile(data string) (model.Image, error)
}

func NewUploadDeliveryHttp(http http.UploadHttp) UploadDeliveryHttp {
	return &uploadDeliveryHttp{http:http}
}

func (u *uploadDeliveryHttp) AddFileS3(s *session.Session, fileDir string, image model.Image, option string) error {
	if err := u.http.AddFileS3(s, fileDir, image, option); err != nil {
		return err
	}
	return nil
}

func (u *uploadDeliveryHttp) ReadFile(data string) (model.Image, error) {
	image, err:= u.http.ReadFile(data)
	if err != nil {
		return model.Image{}, nil
	}
	return image, nil
}
