package http

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/users/internal/http"
)

type deliveryHttp struct {
	userHttp http.UserHttp
}

type DeliveryHttp interface {
	Login(user model.Users) (int, string , error)
	SignUp(user model.Users) error
	RequestIpComputer(ip model.Ip) error
	UpdateProfile(id int, user model.Users) error
	LockAccount(id int) error
	FindAllUser() (model.Users, error)
}

func NewDeliveryHttp(userHttp http.UserHttp) DeliveryHttp {
	return &deliveryHttp{
		userHttp,
	}
}

func (d *deliveryHttp) Login(user model.Users) (int, string , error) {
	id, token, err := d.userHttp.Login(user)
	if err != nil {
		log.NewLog().Error("error create token, login failed")
		return 0, "", err
	}
	log.NewLog().Info("login success")
	return id, token, nil
}
func (d *deliveryHttp) SignUp(user model.Users) error {
	if err := d.userHttp.SignUp(user); err != nil {
		log.NewLog().Error("create account user failed")
		return err
	}
	log.NewLog().Info("create account user success")
	return nil
}
func (d *deliveryHttp) RequestIpComputer(ip model.Ip) error {
	if err := d.userHttp.RequestIpComputer(ip); err != nil {
		log.NewLog().Error("request ip client access failed")
		return err
	}
	log.NewLog().Info("request ip client access success")
	return nil
}
func (d *deliveryHttp) UpdateProfile(id int, user model.Users) error {
	if err := d.userHttp.UpdateProfile(id, user); err != nil {
		log.NewLog().Error("update profile user failed")
		return err
	}
	log.NewLog().Info("update profile user success")
	return nil
}
func (d *deliveryHttp) LockAccount(id int) error {
	if err := d.userHttp.LockAccount(id); err != nil {
		log.NewLog().Error("lock account failed")
		return err
	}
	log.NewLog().Info("lock account success")
	return nil
}
func (d *deliveryHttp) FindAllUser() (model.Users, error) {
	users, err := d.userHttp.FindAllUser()
	if err != nil {
		log.NewLog().Error("not found all users")
		return model.Users{}, err
	}
	log.NewLog().Info("found all users")
	return users, nil
}
