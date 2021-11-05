package repository

import (
	"errors"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"github.com/JieeiroSst/itjob/users/internal/db"
	"regexp"
)

type userRepository struct {
	db db.UserDB
}

type UserRepository interface {
	CheckAccount(user model.Users) (int, string, error)
	CheckAccountExists(user model.Users) error
	CreateAccount(user model.Users) error
	RequestIpComputer(ip model.Ip) error
	CheckPassword(password string) error
	CheckEmail(email string) error
	CheckIP(ip string) error
	UpdateProfile(id int, user model.Users) error
	LockAccount(id int) error
	FindAllUser() (model.Users, error)
}

func NewUserRepository(db db.UserDB) UserRepository {
	log.NewLog().Info("infomation connect db gorm")
	return &userRepository{db:db}
}

func (d *userRepository) FindAllUser() (model.Users, error) {
	users, err := d.db.FindAllUser()
	if err != nil {
		log.NewLog().Error("query error")
		return model.Users{}, errors.New("")
	}
	log.NewLog().Info("information user id user name")
	return users, nil
}

func (d *userRepository) CheckAccount(user model.Users) (int, string, error) {
	id, password, err := d.db.CheckAccount(user)
	if err != nil{
		log.NewLog().Error("query error")
		return -1, "", errors.New("query error")
	}
	log.NewLog().Info("information user id password")
	return id, password, nil
}

func (d *userRepository) CheckAccountExists(user model.Users) error {
	err := d.db.CheckAccountExists(user)
	if err != nil {
		log.NewLog().Error(err.Error())
		return err
	}
	log.NewLog().Info("check account exists")
	return nil
}

func (d *userRepository) CreateAccount(user model.Users) error {
	if err := d.db.CreateAccount(user); err != nil {
		log.NewLog().Error("create user error")
		return err
	}
	log.NewLog().Info("create user success")
	return nil
}

func (d *userRepository) RequestIpComputer(ip model.Ip) error {
	if err := d.db.RequestIpComputer(ip); err != nil {
		return errors.New("get request ip client failed")
	}
	log.NewLog().Info("get request ip client success")
	return nil
}

//String contains 2 uppercase letters or not?
//Contain special characters or not?
//Does it contain 2 numeric and lowercase characters?
// Finally, is the string enough 8 characters or not?
//true: AA@99sds
func (d *userRepository) CheckPassword(password string) error {
	regex := `^(?=.*[A-Z].*[A-Z])(?=.*[!@#$&*])(?=.*[0-9].*[0-9])(?=.*[a-z].*[a-z].*[a-z]).{8}$`
	matched, err := regexp.MatchString(regex, password)
	if !matched {
		log.NewLog().Error("password does not satisfy the condition")
		return errors.New("password does not satisfy the condition")
	}
	if err != nil {
		log.NewLog().Error("password does not satisfy the condition")
		return errors.New("password does not satisfy the condition")
	}
	log.NewLog().Info("password does satisfy the condition")
	return nil
}

func (d *userRepository) CheckEmail(email string) error {
	regex := `^[a-z][a-z0-9_\.]{5,32}@[a-z0-9]{2,}(\.[a-z0-9]{2,4}){1,2}$`
	matched, err := regexp.MatchString(regex, email)
	if !matched {
		log.NewLog().Error("email does not satisfy the condition")
		return errors.New("email does not satisfy the condition")
	}
	if err != nil {
		log.NewLog().Error("email does not satisfy the condition")
		return errors.New("email does not satisfy the condition")
	}
	log.NewLog().Info("email does satisfy the condition")
	return nil
}

func (d *userRepository) CheckIP(ip string) error {
	regex := `/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/`
	matched, err := regexp.MatchString(regex, ip)
	if !matched {
		log.NewLog().Error("IP does not satisfy the condition")
		return errors.New("IP does not satisfy the condition")
	}
	if err != nil {
		log.NewLog().Error("IP does not satisfy the condition")
		return errors.New("IP does not satisfy the condition")
	}
	log.NewLog().Info("IP does satisfy the condition")
	return nil
}

func (d *userRepository) UpdateProfile(id int,user model.Users) error {
	if err:=d.db.UpdateProfile(id, user); err != nil {
		log.NewLog().Error("update profile user failed")
		return errors.New("update profile user failed")
	}
	log.NewLog().Info("update profile user success")
	return nil
}

func (d *userRepository) LockAccount(id int) error {
	if err := d.db.LockAccount(id); err != nil {
		log.NewLog().Error("lock account failed")
		return errors.New("lock account failed")
	}
	log.NewLog().Info("lock account success")
	return nil
}