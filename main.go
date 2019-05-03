package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/matthiasbaldi/go-blog-api/features/blog"
	"github.com/matthiasbaldi/go-blog-api/features/database"
)

// const (
//     dbName = "go-mysql-crud"
//     dbPass = "12345"
//     dbHost = "localhost"
//     dbPort = "33066"
// )

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/blogs", blog.Routes())
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	fs := http.StripPrefix("/", http.FileServer(http.Dir(filesDir)))
	router.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

	return router
}

func main() {
	database.Initialize()
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":4001", router))
}
