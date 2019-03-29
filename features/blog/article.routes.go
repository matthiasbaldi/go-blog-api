package blog

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/matthiasbaldi/go-blog-api/models"
)

func GetAllBlogArticles(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{
		{
			Title:   "Hello world",
			Content: "Hello world from planet earth",
		},
	}
	render.JSON(w, r, articles)
}

func GetBlogArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	article := models.Article{
		ID:      id,
		Title:   "Hello world",
		Content: "Hello world from planet earth",
	}
	render.JSON(w, r, article)
}

func CreateBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Blog successfully"
	render.JSON(w, r, response)
}

func UpdateBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Updated Blog successfully"
	render.JSON(w, r, response)
}

func DeleteBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Blog successfully"
	render.JSON(w, r, response)
}
