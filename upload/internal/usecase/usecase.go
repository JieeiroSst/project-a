package usecase

import (
	"bytes"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/upload/internal/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type uploadUsecase struct {
	repo repository.UploadRepository
	s3 *session.Session
	config *config.Config
}

type UploadUsecase interface {
	AddFileToS3(s *session.Session, fileDir string, image model.Image) error
	AddFileToS3Stream(s *session.Session, data string, image model.Image) error
	ReadFile(data string) (string, model.Image, error)
}

func NewUploadUsecase(repo repository.UploadRepository, s3 *session.Session, config *config.Config) UploadUsecase {
	return &uploadUsecase{
		repo:repo,
		s3:s3,
		config:config,
	}
}

func(u *uploadUsecase) AddFileToS3(s *session.Session, fileDir string, image model.Image) error {
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	if _ , err =s3.New(s).PutObject(&s3.PutObjectInput{
		ACL:                       aws.String(u.config.AmazonS3.S3ACL),
		Body:                      bytes.NewReader(buffer),
		Bucket:                    aws.String(u.config.AmazonS3.S3Bucket),
		ContentDisposition:        aws.String("attachment"),
		ContentLength:              aws.Int64(size),
		ContentType:               aws.String(http.DetectContentType(buffer)),
		Key:                       aws.String(fileDir),
		ServerSideEncryption:      aws.String("AES256"),

	});err != nil {
		return err
	}
	if err := u.repo.SaveImageForUser(image); err != nil {
		return err
	}
	return nil
}
func(u *uploadUsecase) AddFileToS3Stream(s *session.Session, data string, image model.Image) error {
	reader := strings.NewReader("Geeks")

	buffer := make([]byte, 4)

	n, err := io.ReadFull(reader, buffer)
	if err != nil {
		return err
	}
	fileDir:=data[0:4]+time.Now().String()

	if _ , err =s3.New(s).PutObject(&s3.PutObjectInput{
		ACL:                       aws.String(u.config.AmazonS3.S3ACL),
		Body:                      bytes.NewReader(buffer),
		Bucket:                    aws.String(u.config.AmazonS3.S3Bucket),
		ContentDisposition:        aws.String("attachment"),
		ContentLength:              aws.Int64(int64(n)),
		ContentType:               aws.String(http.DetectContentType(buffer)),
		Key:                       aws.String(fileDir),
		ServerSideEncryption:      aws.String("AES256"),

	}); err != nil {
		return err
	}

	if err := u.repo.SaveImageForUser(image); err != nil {
		return err
	}
	return err

}
func(u *uploadUsecase) ReadFile(data string) (string, model.Image, error) {
	results, err := s3.New(u.s3).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(u.config.AmazonS3.S3Bucket),
		Key:    aws.String(data),
	})
	if err != nil {
		return "", model.Image{}, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return "", model.Image{}, err
	}

	image, err := u.repo.FindByIdImage(string(buf.Bytes()))
	if err != nil {
		return "", model.Image{}, err
	}

	return string(buf.Bytes()), image, nil
}

