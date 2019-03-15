package blog

import (
	"github.com/go-chi/chi"
)

type Blog struct {
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Descrption string `json:"description"`
	CreateDate string `json:"created"`
	UpdateDate string `json:"updated"`
	Author     string `json:"author"`
}

type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateDate string `json:"created"`
	UpdateDate string `json:"updated"`
	Author     string `json:"author"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	// blogs
	router.Get("/", GetAllBlogs)
	router.Post("/", CreateBlog)
	router.Get("/{blogID}", GetBlog)
	router.Put("/{blogID}", UpdateBlog)
	router.Delete("/{blogID}", DeleteBlog)
	// articles
	router.Get("/{blogID}/articles", GetAllBlogArticles)
	router.Post("/{blogID}/articles", CreateBlogArticle)
	router.Get("/{blogID}/articles/{articleID}", GetBlogArticle)
	router.Put("/{blogID}/articles/{articleID}", UpdateBlogArticle)
	router.Delete("/{blogID}/articles/{articleID}", DeleteBlogArticle)
	return router
}
