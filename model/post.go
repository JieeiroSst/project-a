package model

import "time"

type Posts struct {
	Id int `gorm:"primary_key" json:"id"`
	AuthorId int `json:"author_id"`
	Title string `json:"title"`
	MetaTitle string `json:"meta_title"`
	Slug string `json:"slug"`
	Summary string `json:"summary"`
	Published int `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:null"`
	PublishedAt time.Time `json:"published_at" gorm:"default:null"`
	Content string `json:"content"`
}
