package main

import (
	"fmt"
	"net/http"

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
	http.Handle("/", http.HandlerFunc(handlers.Home))
	http.Handle("/post", http.HandlerFunc(handlers.CreatePostHandler(DB))) //dependency injection
	http.Handle("/posts", http.HandlerFunc(handlers.ListPostHandler(DB)))
	http.ListenAndServe(":8080", nil)
}
