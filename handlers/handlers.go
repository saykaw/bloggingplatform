package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/saykaw/bloggingplatform/models"
	"gorm.io/gorm"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the home page. Welcome!")
}

func CreatePostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		blogPost := models.BlogPost{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&blogPost)

		result := db.Create(&blogPost)
		if result.Error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("something went wrong: %v \n", result.Error)
			rw.Write([]byte("could not create the blog post"))
			return
		}

		if result.RowsAffected == 0 {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("No rows were inserted.\n")
			rw.Write([]byte("could not create the post blog post"))
			return
		}

		rw.WriteHeader(http.StatusCreated)
		fmt.Fprintf(rw, "Your blog is created successfully!")
	}
}

func ListPostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		blogPost := []models.BlogPost{}
		result := db.Find(&blogPost)
		if result.Error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("something went wrong: %v \n", result.Error)
			rw.Write([]byte("could not display the blog posts"))
			return
		}

		if result.RowsAffected == 0 {
			fmt.Printf("No rows were found.\n")
			rw.Write([]byte("no blogs are found"))
			return
		} else {
			fmt.Printf("number of rows found: %v\n", result.RowsAffected)
		}
		rw.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(rw)
		encoder.Encode(&blogPost)
	}
}

func GetPostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Printf("theres an error in converting id into uint")
			rw.Write([]byte("id should be numerical"))
			return
		}
		ctx := context.Background()
		post, err := gorm.G[models.BlogPost](db).Where("id = ?", id).First(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Printf("something went wrong: %v \n", err)
			rw.Write([]byte("could not display the blog post with id, post with id not found"))
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(rw)
		encoder.Encode(&post)
	}
}

func UpdatePostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Printf("theres an error in converting id into uint")
			return
		}
		blogPost := models.BlogPost{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&blogPost)
		ctx := context.Background()
		rows, err := gorm.G[models.BlogPost](db).Where("id = ?", id).Updates(ctx, models.BlogPost{Title: blogPost.Title, Content: blogPost.Content, Category: blogPost.Category})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not update"))
			fmt.Printf("theres an error :%v\n", err)
			return
		}
		if rows == 0 {
			fmt.Printf("no blog updated, you might have to create it")
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("blog not found, you might have to create it!"))
			return
		}
	}
}

func DeletePostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Printf("theres an error in converting id into uint")
			rw.Write([]byte("id should be numerical"))
			return
		}
		res := db.Delete(&models.BlogPost{}, id)
		if res.Error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Delete Error: %v", res.Error)
			return
		}
		if res.RowsAffected == 0 {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("Post not found for deletion"))
			return
		}
		rw.WriteHeader(http.StatusNoContent)
	}
}

//TODO:the tags field in db
