package repository_test

import (
	mock_repository "github.com/JieeiroSst/itjob/users/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

var (
	a *mock_repository.MockUserRepository
)

func TestUserRepository_CheckIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
