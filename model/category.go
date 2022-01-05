package model

import "time"

type Categories struct {
	Id 			int 		`gorm:"primary_key" json:"id"`
	ParentId 	int 		`json:"parent_id"`
	Title 		string 		`json:"title"`
	MetaTitle 	string 		`json:"meta_title"`
	Slug 		string 		`json:"slug"`
	Content 	string 		`json:"content"`
	CreatedAt   time.Time 	`json:"created_at"`
}

type RequestCategory struct {
	ParentId 	int 	`json:"parent_id"`
	Title 		string 	`json:"title"`
	MetaTitle 	string 	`json:"meta_title"`
	Slug 		string 	`json:"slug"`
	Content 	string 	`json:"content"`
}