package blog

import (
	"net/http"

	"github.com/go-chi/render"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := []Blog{
		{
			Slug:       "slug",
			Title:      "Hello world",
			Descrption: "Hello world from planet earth",
		},
	}
	render.JSON(w, r, blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	blog := Blog{
		Slug:       "slug",
		Title:      "Hello world",
		Descrption: "Hello world from planet earth",
	}
	render.JSON(w, r, blog)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Blog successfully"
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
