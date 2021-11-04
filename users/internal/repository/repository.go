package repository

import (
	"errors"
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/pkg/log"
	"gorm.io/gorm"
	"regexp"
)

type userRepository struct {
	db *gorm.DB
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

func NewUserRepository(db *gorm.DB) UserRepository {
	if err := db.AutoMigrate(&model.Users{}); err!=nil {
		log.NewLog().Error(err)
		return nil
	}
	log.NewLog().Info("infomation connect db gorm")
	return &userRepository{db:db}
}

func (d *userRepository) FindAllUser() (model.Users, error) {
	var user model.Users
	err := d.db.Select("id, username").Find(&user).Error
	if err != nil {
		log.NewLog().Error("query error")
		return model.Users{}, errors.New("")
	}
	log.NewLog().Info("information user id user name")
	return user, nil
}

func (d *userRepository) CheckAccount(user model.Users) (int, string, error) {
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

func (d *userRepository) CheckAccountExists(user model.Users) error {
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

func (d *userRepository) CreateAccount(user model.Users) error {
	if err := d.db.Create(&user).Error; err!=nil {
		log.NewLog().Error("create user error")
		return err
	}
	log.NewLog().Info("create user success")
	return nil
}

func (d *userRepository) RequestIpComputer(ip model.Ip) error {
	if err := d.db.Create(&ip).Error; err != nil {
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
	err:=d.db.Model(model.Users{}).Where("id = ? ", id).Updates(user).Error
	if err != nil {
		log.NewLog().Error("update profile user failed")
		return errors.New("update profile user failed")
	}
	log.NewLog().Info("update profile user success")
	return nil
}

func (d *userRepository) LockAccount(id int) error {
	err := d.db.Model(&model.Users{}).Where("id = ?", id).Update("checked", false).Error
	if err != nil {
		log.NewLog().Error("")
		return errors.New("")
	}
	log.NewLog().Info("")
	return nil
}