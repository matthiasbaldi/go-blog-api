package blog

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/matthiasbaldi/go-blog-api/features/database"
	"github.com/matthiasbaldi/go-blog-api/models"
	uuid "github.com/satori/go.uuid"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := database.FetchAllBlogs()
	render.JSON(w, r, blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "blogID")
	blog := database.FetchOneBlog(id)
	render.JSON(w, r, blog)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var blog models.Blog
	err := decoder.Decode(&blog)
	if err != nil {
		panic(err)
	}

	newBlog := models.Blog{
		ID:         uuid.Must(uuid.NewV4()).String(),
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		Author:     blog.Author,
		Slug:       blog.Slug,
		Title:      blog.Title,
		Descrption: blog.Descrption,
	}

	database.InsertBlog(newBlog)

	response := make(map[string]string)
	response["message"] = "Created Blog successfully"
	render.JSON(w, r, response)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "blogID")
	decoder := json.NewDecoder(r.Body)

	var blog models.Blog
	err := decoder.Decode(&blog)
	if err != nil {
		panic(err)
	}

	updatedBlog := models.Blog{
		ID:         id,
		Author:     blog.Author,
		Slug:       blog.Slug,
		Title:      blog.Title,
		Descrption: blog.Descrption,
		CreateDate: blog.CreateDate,
		UpdateDate: time.Now(),
	}

	database.UpdateBlog(id, updatedBlog)

	response := make(map[string]string)
	response["message"] = "Updated Blog successfully"
	render.JSON(w, r, response)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "blogID")
	database.DeleteBlog(id)

	response := make(map[string]string)
	response["message"] = "Deleted Blog successfully"
	render.JSON(w, r, response)
}
