package blog

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matthiasbaldi/go-blog-api/features/database"
	"github.com/matthiasbaldi/go-blog-api/models"
	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetAllBlogArticles(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "blogID")
	articles := database.FetchAllArticles(blogId)
	render.JSON(w, r, articles)
}

func GetBlogArticle(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "blogID")
	articleId := chi.URLParam(r, "articleID")
	article := database.FetchOneArticle(blogId, articleId)
	render.JSON(w, r, article)
}

func CreateBlogArticle(w http.ResponseWriter, r *http.Request) {
	blogId := chi.URLParam(r, "blogID")

	decoder := json.NewDecoder(r.Body)

	var article models.Article
	err := decoder.Decode(&article)
	if err != nil {
		panic(err)
	}

	newArticle := models.Article{
		ID:         uuid.Must(uuid.NewV4()).String(),
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		Author:     article.Author,
		BlogID:     blogId,
		Title:      article.Title,
		Content:    article.Content,
	}

	database.InsertArticle(newArticle)

	response := make(map[string]string)
	response["message"] = "Created Article successfully"
	render.JSON(w, r, response)
}

func UpdateBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Updated Article successfully --> DUMMY"
	render.JSON(w, r, response)
}

func DeleteBlogArticle(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Article successfully --> DUMMY"
	render.JSON(w, r, response)
}
