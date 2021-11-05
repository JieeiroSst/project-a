package db

import (
	"errors"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"gorm.io/gorm"
)

type userDB struct {
	db *gorm.DB
}

type UserDB interface {
	CheckAccount(user model.Users) (int, string, error)
	CheckAccountExists(user model.Users) error
	CreateAccount(user model.Users) error
	RequestIpComputer(ip model.Ip) error
	FindAllUser() (model.Users, error)
	LockAccount(id int) error
	UpdateProfile(id int, user model.Users) error
}

func NewUserDB(db *gorm.DB) UserDB {
	return &userDB{db:db}
}

func (d *userDB) UpdateProfile(id int,user model.Users) error {
	err:=d.db.Model(model.Users{}).Where("id = ? ", id).Updates(user).Error
	if err != nil {
		log.NewLog().Error("update profile user failed")
		return errors.New("update profile user failed")
	}
	log.NewLog().Info("update profile user success")
	return nil
}

func (d *userDB) LockAccount(id int) error {
	err := d.db.Model(&model.Users{}).Where("id = ?", id).Update("checked", false).Error
	if err != nil {
		log.NewLog().Error("lock account failed")
		return errors.New("lock account failed")
	}
	log.NewLog().Info("lock account success")
	return nil
}

func (d *userDB) FindAllUser() (model.Users, error) {
	var user model.Users
	err := d.db.Select("id, username").Find(&user).Error
	if err != nil {
		log.NewLog().Error("query error")
		return model.Users{}, errors.New("")
	}
	log.NewLog().Info("information user id user name")
	return user, nil
}

func (d *userDB) CheckAccount(user model.Users) (int, string, error) {
	var result model.Users
	r := d.db.Where("username = ?", user.Username).Limit(1).Find(&result)

	if r.Error != nil{
		log.NewLog().Error("query error")
		return -1, "", errors.New("Query error")
	}

	if result.Id == 0 {
		log.NewLog().Error("user does not exist")
		return -1, "", errors.New("user does not exist")
	}
	log.NewLog().Info("information user id password")
	return result.Id, result.Password, nil
}

func (d *userDB) CheckAccountExists(user model.Users) error {
	var result model.Users
	r := d.db.Where("username = ?", user.Username).Limit(1).Find(&result)
	if r.Error != nil{
		log.NewLog().Error("query error")
		return errors.New("query error")
	}

	if result.Id !=0 {
		log.NewLog().Error("user does exist")
		return errors.New("user does exist")
	}
	log.NewLog().Info("check account exists")
	return nil
}

func (d *userDB) CreateAccount(user model.Users) error {
	if err := d.db.Create(&user).Error; err!=nil {
		log.NewLog().Error("create user error")
		return err
	}
	log.NewLog().Info("create user success")
	return nil
}

func (d *userDB) RequestIpComputer(ip model.Ip) error {
	if err := d.db.Create(&ip).Error; err != nil {
		return errors.New("get request ip client failed")
	}
	log.NewLog().Info("get request ip client success")
	return nil
}
