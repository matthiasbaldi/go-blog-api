package blog

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/matthiasbaldi/go-blog-api/features/database"
	"github.com/matthiasbaldi/go-blog-api/models"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := database.FetchAllBlogs()
	render.JSON(w, r, blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "blogID")
	blog := models.Blog{
		ID:         id,
		Slug:       "slug",
		Title:      "Hello world",
		Descrption: "Hello world from planet earth",
	}
	render.JSON(w, r, blog)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	var blog models.Blog
	database.InsertBlog(r.Body)
	response["message"] = "Created Blog successfully"
	// response["created"] = database.InsertBlog(r.Body())
	render.JSON(w, r, response)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Updated Blog successfully"
	render.JSON(w, r, response)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Blog successfully"
	render.JSON(w, r, response)
}
