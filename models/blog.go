package models

import "time"

type Blog struct {
	ID         string    `json:"id"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Descrption string    `json:"description"`
	CreateDate time.Time `json:"created"`
	UpdateDate time.Time `json:"updated"`
	Author     string    `json:"author"`
}
