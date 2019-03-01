package blog

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Blog struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateDate string `json:"created"`
	UpdateDate string `json:"updated"`
	Author     string `json:"author"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{articleID}", GetBlogArticle)
	router.Delete("/{articleID}", DeleteBlogArticle)
	router.Post("/", CreateBlogArticle)
	router.Get("/", GetAllBlogArticle)
	return router
}

func GetBlogArticle(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	blog := Blog{
		Slug:    articleID,
		Title:   "Hello world",
		Content: "Hello world from planet earth",
	}
	render.JSON(w, r, blog)
}

func DeleteBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Article successfully"
	render.JSON(w, r, response)
}

func CreateBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Article successfully"
	render.JSON(w, r, response)
}

func GetAllBlogArticle(w http.ResponseWriter, r *http.Request) {
	blog := []Blog{
		{
			Slug:    "slug",
			Title:   "Hello world",
			Content: "Hello world from planet earth",
		},
	}
	render.JSON(w, r, blog)
}
