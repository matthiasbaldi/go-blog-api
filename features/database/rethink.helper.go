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

/* ############ BLOG HELPER ############ */

func CreateTable(tableName string) {
	result, err := r.TableCreate(tableName).RunWrite(session)
	if err != nil {
		log.Printf("check table %s - %s\n", tableName, err.Error())
	} else {
		log.Printf("check table %s - created %d\n", tableName, result.Created)
	}
}

func InsertBlog(blog models.Blog) int {
	result, err := r.Table("blogs").Insert(blog).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return result.Created
	}
	return result.Created
}

func UpdateBlog(id string, blog models.Blog) int {
	result, err := r.Table("blogs").Update(blog).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return result.Updated
	}
	return result.Updated
}

func DeleteBlog(id string) {
	result, err := r.Table("blogs").Delete().Run(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(result)
	return
}

func FetchAllBlogs() []models.Blog {
	rows, err := r.Table("blogs").Run(session)
	blogs := []models.Blog{}
	if err != nil {
		fmt.Println(err)
		return blogs
	}

	err2 := rows.All(&blogs)
	if err2 != nil {
		fmt.Println(err2)
		return blogs
	}
	return blogs
}

func FetchOneBlog(id string) models.Blog {
	cursor, err := r.Table("blogs").Get(id).Run(session)
	var blog models.Blog
	cursor.One(&blog)
	if err != nil {
		fmt.Println(err)
		return blog
	}

	return blog
}

/* ############ ARTICLE HELPER ############ */

func InsertArticle(article models.Article) int {
	result, err := r.Table("articles").Insert(article).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return result.Created
	}
	return result.Created
}

func FetchAllArticles(blogId string) []models.Article {
	rows, err := r.Table("articles").Run(session)
	articles := []models.Article{}
	if err != nil {
		fmt.Println(err)
		return articles
	}

	err2 := rows.All(&articles)
	if err2 != nil {
		fmt.Println(err2)
		return articles
	}
	return articles
}

func FetchOneArticle(blogId string, id string) models.Article {
	cursor, err := r.Table("articles").Filter(func(row r.Term) interface{} {
		return row.Field("id").Eq(id).And(row.Field("blogId").Eq(blogId))
	}).Run(session)

	var article models.Article
	cursor.One(&article)
	if err != nil {
		fmt.Println(err)
		return article
	}

	return article
}
