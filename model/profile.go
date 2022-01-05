package model

import "time"

type Profiles struct {
	Id 				int 			`gorm:"primary_key" json:"id"`
	UserId 			int 			`json:"user_id"`
	FirstName 		string 			`json:"first_name"`
	MiddleName 		string 			`json:"middle_name"`
	LastName 		string 			`json:"last_name"`
	Mobile 			string 			`json:"mobile"`
	Email 			string 			`json:"email"`
	RegisteredAt 	time.Time 		`json:"registered_at"`
	CreatedAt   	time.Time 	`json:"created_at"`
	Profile 		string 			`json:"profile"`
	PostComments    []PostComments  `gorm:"foreignKey:ParentId"`
}

type RequestProfile struct {
	UserId 			int 			`json:"user_id"`
	FirstName 		string 			`json:"first_name"`
	MiddleName 		string 			`json:"middle_name"`
	LastName 		string 			`json:"last_name"`
	Mobile 			string 			`json:"mobile"`
	Email 			string 			`json:"email"`
	RegisteredAt 	time.Time 		`json:"registered_at"`
	Profile 		string 			`json:"profile"`
}