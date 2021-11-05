package repository

import (
	"errors"
	"github.com/JieeiroSst/itjob/users/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserRepository_CheckIP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDoer := mocks.NewMockUserRepository(ctrl)

	testTable := [] struct {
		description string
		got         string
		want        error
	}{
		{
			description: "satisfy the condition",
			got: "127.0.0.1",
			want: nil,
		},
		{
			description: "condition is not satisfied ",
			got: "127.0.88.100",
			want: errors.New("IP does not satisfy the condition"),
		},
	}

	for _, tt := range testTable {
		mockDoer.EXPECT().CheckIP(tt.got).Return(tt.want)
	}


}

func TestUserRepository_CheckAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_CheckAccountExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_CheckEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_CheckPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_FindAllUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_LockAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_RequestIpComputer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func TestUserRepository_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}