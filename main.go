package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saykaw/bloggingplatform/db"
	"github.com/saykaw/bloggingplatform/handlers"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = db.ConnectToDb()
	if DB == nil {
		fmt.Printf("something went wrong")
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", handlers.Home)
	r.Get("/posts", handlers.ListPostHandler(DB))
	r.Get("/posts/{id}", handlers.GetPostHandler(DB))
	r.Post("/post", handlers.CreatePostHandler(DB))
	r.Put("/posts/{id}", handlers.UpdatePostHandler(DB))
	r.Delete("/posts/{id}", handlers.DeletePostHandler(DB))
	// http.Handle("/", http.HandlerFunc(handlers.Home))
	// http.Handle("/post", http.HandlerFunc(handlers.CreatePostHandler(DB))) //dependency injection
	// http.Handle("/posts", http.HandlerFunc(handlers.ListPostHandler(DB)))
	http.ListenAndServe(":8080", r)
}
