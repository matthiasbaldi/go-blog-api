package database

import (
	"fmt"
	"log"

	"github.com/matthiasbaldi/go-blog-api/models"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

const (
	url  = "localhost"
	port = 28015
)

var (
	session *r.Session
)

// https://github.com/rethinkdb/rethinkdb-go/wiki/Simple-Example

func Initialize() {
	s, err := r.Connect(r.ConnectOpts{
		Address: url,
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		session = s
	}
	res, err := r.Expr(fmt.Sprintf("Database running %s:%d", url, port)).Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
	CreateTable("blogs")
	CreateTable("articles")
}

func CreateTable(tableName string) {
	result, err := r.TableCreate(tableName).RunWrite(session)
	if err != nil {
		log.Printf("check table %s - %s\n", tableName, err.Error())
	} else {
		log.Printf("check table %s - created %d\n", tableName, result.Created)
	}
}

func InsertBlog(blog models.Blog) int {
	// var data = map[string]interface{}{
	// 	"Name":  "David Davidson",
	// 	"Place": "Somewhere",
	// }

	// var resBlog models.Blog
	result, err := r.Table("blogs").Insert(blog).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return result.Created
	}
	return result.Created
	// resBlog := res.

	// return result.GeneratedKeys[0]
}

func FetchAllBlogs() []models.Blog {
	rows, err := r.Table("blogs").Run(session)
	var blogs []models.Blog
	if err != nil {
		fmt.Println(err)
		return blogs
	}

	err2 := rows.All(&blogs)
	if err2 != nil {
		fmt.Println(err2)
		return blogs
	}
	println(len(blogs))
	return blogs
}
