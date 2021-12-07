package usecase

import (
	"errors"
	"github.com/JieeiroSst/itjob/config"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/jwt"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/users/internal/repository"
	"github.com/JieeiroSst/itjob/utils"
)

type userUsecase struct {
	userRepo repository.UserRepository
	conf config.Config
	hash utils.Hash
	jwt jwt.TokenUser
}

//go:generate mockgen --destination=./mocks/userUsecase.go usecase UserUsecase
type UserUsecase interface {
	Login(user model.Users) (int, string , error)
	SignUp(user model.Users) error
	RequestIpComputer(ip model.Ip) error
	UpdateProfile(id int, user model.Users) error
	LockAccount(id int) error
	FindAllUser() (model.Users, error)
}

func NewUserCase(userRepo repository.UserRepository, hash utils.Hash, jwt jwt.TokenUser, conf config.Config) UserUsecase {
	return &userUsecase{
		userRepo:userRepo,
		hash:hash,
		jwt:jwt,
		conf:conf,
	}
}

func (u *userUsecase) FindAllUser() (model.Users, error) {
	users, err := u.userRepo.FindAllUser()
	if err != nil {
		log.NewLog().Error("not found all users")
		return model.Users{}, err
	}
	log.NewLog().Info("found all users")
	return users, nil
}

func (u *userUsecase) Login(user model.Users) (int, string , error) {
	id, hashPassword,err  := u.userRepo.CheckAccount(user)
	if err != nil {
		log.NewLog().Error("user does not exist")
		return 0, "", errors.New("user does not exist")
	}
	if checkPass := u.hash.CheckPassowrd(user.Password, hashPassword); checkPass != nil {
		log.NewLog().Error("password entered incorrectly")
		return 0, "" ,errors.New("password entered incorrectly")
	}
	log.NewLog().Info("login success")
	token, _ := u.jwt.GenerateToken(user.Username)
	return id, token, nil
}

func (u *userUsecase) SignUp(user model.Users) error {
	if err := u.userRepo.CheckEmail(user.Email); err != nil {
		log.NewLog().Error("invalid email")
		return err
	}
	if err := u.userRepo.CheckPassword(user.Password); err != nil {
		log.NewLog().Error("wrong password")
		return err
	}
	check := u.userRepo.CheckAccountExists(user)
	if check != nil {
		log.NewLog().Error("user already exists")
		return errors.New("user already exists")
	}
	hashPassword, err := u.hash.HashPassword(user.Password)
	if err != nil {
		log.NewLog().Error("password failed")
		return errors.New("password failed")
	}
	account:= model.Users{
		Id:           user.Id,
		Username:     user.Username,
		Password:     hashPassword,
	}
	err = u.userRepo.CreateAccount(account)
	if err!=nil{
		log.NewLog().Error("create failed")
		return errors.New("create failed")
	}
	log.NewLog().Info("create success")
	return errors.New("create success")
}

func (u *userUsecase) RequestIpComputer(ip model.Ip) error {
	if err := u.userRepo.CheckIP(ip.Ip); err != nil {
		log.NewLog().Error("wrong ip")
		return err
	}

	if err := u.userRepo.RequestIpComputer(ip); err != nil {
		log.NewLog().Error("request failed")
		return errors.New("request failed")
	}
	log.NewLog().Info("request success")
	return nil
}

func (u *userUsecase) UpdateProfile(id int, user model.Users) error {
	if err := u.userRepo.UpdateProfile(id, user); err != nil {
		log.NewLog().Error("update failed")
		return err
	}
	log.NewLog().Info("update success")
	return nil
}
func (u *userUsecase) LockAccount(id int) error {
	if err := u.userRepo.LockAccount(id); err != nil {
		log.NewLog().Error("lock account failed")
		return err
	}
	log.NewLog().Info("lock account success")
	return nil
}