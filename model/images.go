package model

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Option string

const (
	Avatar	Option = "AVATAR"
	News 	Option = "NEWS"
)

type Image struct {
	Id     		int			`gorm:"primaryKey" json:"id"`
	Name   		string 		`json:"name"`
	UserRefer 	int			`json:"user_refer"`
	Option      Option		`json:"option" orm:"type:enum('AVATAR', 'NEWS');default:'NEWS'"`
	CreatedTime time.Time   `json:"created_time"`
	UpdatedTime time.Time   `json:"updated_time"`
}

func (e *Option) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	*e = Option(string(asBytes))
	return nil
}

func (e Option) Value() (driver.Value, error) {
	return string(e), nil
}