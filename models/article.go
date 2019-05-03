package models

import "time"

type Article struct {
	ID         string    `json:"id"`
	BlogID     string    `json:"blogId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"created"`
	UpdateDate time.Time `json:"updated"`
	Author     string    `json:"author"`
}
