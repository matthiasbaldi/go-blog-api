package models

type Blog struct {
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Descrption string `json:"description"`
	CreateDate string `json:"created"`
	UpdateDate string `json:"updated"`
	Author     string `json:"author"`
}
