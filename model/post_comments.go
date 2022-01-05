package model

import "time"

type PostComments struct {
	Id 			int 		`gorm:"primary_key" json:"id"`
	PostId 		int 		`json:"post_id"`
	ParentId 	int 		`json:"parent_id"`
	Title 		string 		`json:"title"`
	Published 	int 		`json:"published"`
	CreatedAt 	time.Time 	`json:"created_at"`
	PublishedAt time.Time 	`json:"published_at" gorm:"default:null"`
	Content 	string 		`json:"content"`
}

type RequestPostComments struct {
	PostId 		int 		`json:"post_id"`
	ParentId 	int 		`json:"parent_id"`
	Title 		string 		`json:"title"`
	Published 	int 		`json:"published"`
	Content 	string 		`json:"content"`
}