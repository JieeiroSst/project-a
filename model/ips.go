package model

import "time"

type Ip struct {
	Id     int          `gorm:"primary_key" json:"id"`
	Ip     string       `json:"ip"`
	Method string       `json:"method"`
	RequestAt time.Time `json:"created_at"`
}