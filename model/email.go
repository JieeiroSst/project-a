package model

import "time"

type Email struct {
	Id               int       `json:"id"`
	NameEmailSend    string    `json:"name_email_send"`
	NameEmailReceive string    `json:"name_email_receive"`
	SubjectEmail     string    `json:"subject_email"`
	Content          string    `json:"content"`
	CreatedAt        time.Time `json:"created_at"`
}

type RequestEmail struct {
	NameEmailSend    string    `json:"name_email_send"`
	NameEmailReceive string    `json:"name_email_receive"`
	SubjectEmail     string    `json:"subject_email"`
	Content          string    `json:"content"`
}