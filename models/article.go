package models

type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateDate string `json:"created"`
	UpdateDate string `json:"updated"`
	Author     string `json:"author"`
}
