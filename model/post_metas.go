package model

import "time"

type PostMetas struct {
	Id 			int 	`gorm:"primary_key" json:"id"`
	PostId 		int 	`json:"post_id"`
	TextKey 	string  `json:"text_key"`
	Content 	string  `json:"content"`
	CreatedAt   time.Time 	`json:"created_at"`
}

type RequestPostMetas struct {
	PostId 	int 	`json:"post_id"`
	TextKey string  `json:"text_key"`
	Content string  `json:"content"`
}