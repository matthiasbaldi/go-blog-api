package blog

import (
	"github.com/go-chi/chi"
)

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
