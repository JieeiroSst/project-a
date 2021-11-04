package http

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/users/internal/usecase"
)

type userHttp struct {
	usecase usecase.UserUsecase
}

type UserHttp interface {
	Login(user model.Users) (int, string , error)
	SignUp(user model.Users) error
	RequestIpComputer(ip model.Ip) error
	UpdateProfile(id int, user model.Users) error
	LockAccount(id int) error
	FindAllUser() (model.Users, error)
}

func NewUserHttp(usecase usecase.UserUsecase) UserHttp {
	return &userHttp{
		usecase:usecase,
	}
}

func (h *userHttp) FindAllUser() (model.Users, error) {
	users, err := h.usecase.FindAllUser()
	if err != nil {
		log.NewLog().Error("not found all users")
		return model.Users{}, err
	}
	log.NewLog().Info("found all users")
	return users, nil
}

func (h *userHttp) Login(user model.Users) (int, string , error) {
	id, token, err := h.usecase.Login(user)
	if err != nil {
		log.NewLog().Error("error create token, login failed")
		return 0, "", err
	}
	log.NewLog().Info("login success")
	return id, token, nil
}

func (h *userHttp) SignUp(user model.Users) error {
	if err := h.usecase.SignUp(user); err != nil {
		log.NewLog().Error("create account user failed")
		return err
	}
	log.NewLog().Info("create account user success")
	return nil
}

func (h *userHttp) RequestIpComputer(ip model.Ip) error {
	if err := h.usecase.RequestIpComputer(ip); err != nil {
		log.NewLog().Error("request ip client access failed")
		return err
	}
	log.NewLog().Info("request ip client access success")
	return nil
}

func (h *userHttp) UpdateProfile(id int, user model.Users) error {
	if err := h.usecase.UpdateProfile(id, user); err != nil {
		log.NewLog().Error("update profile user failed")
		return err
	}
	log.NewLog().Info("update profile user success")
	return nil
}

func (h *userHttp) LockAccount(id int) error {
	if err := h.usecase.LockAccount(id); err != nil {
		log.NewLog().Error("lock account failed")
		return err
	}
	log.NewLog().Info("lock account success")
	return nil
}