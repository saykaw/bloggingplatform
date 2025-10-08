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
		decoder := json.NewDecoder(r.Body) //kisko decode karna hai
		decoder.Decode(&blogPost)          //kisme decode karna hai

		result := db.Create(&blogPost) //gorm create a record
		if result.Error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("something went wrong: %v \n", result.Error)
			return
		}

		if result.RowsAffected == 0 {
			fmt.Printf("No rows were inserted.\n")
			return
		}

		rw.WriteHeader(http.StatusCreated)
		fmt.Fprintf(rw, "Your blog is created successfully!")
	}
}

func ListPostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		blogPost := []models.BlogPost{}
		encoder := json.NewEncoder(rw)
		result := db.Find(&blogPost)
		encoder.Encode(&blogPost)
		if result.Error != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("something went wrong: %v \n", result.Error)
			return
		}

		if result.RowsAffected == 0 {
			fmt.Printf("No rows were found.\n")
			return
		} else {
			fmt.Printf("number of rows found: %v\n", result.RowsAffected)
		}
	}
}

func GetPostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Printf("theres an error in converting id into uint")
		}
		ctx := context.Background()
		post, err := gorm.G[models.BlogPost](db).Where("id = ?", id).First(ctx)
		if err != nil {
			fmt.Printf("theres an error in finding post with id %v\n", id)
		}
		encoder := json.NewEncoder(rw)
		encoder.Encode(&post)
	}
}

func UpdatePostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			fmt.Printf("theres an error in converting id into uint")
		}
		blogPost := models.BlogPost{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&blogPost)
		ctx := context.Background()
		rows, err := gorm.G[models.BlogPost](db).Where("id = ?", id).Updates(ctx, models.BlogPost{Title: blogPost.Title, Content: blogPost.Content, Category: blogPost.Category})
		if err != nil {
			fmt.Printf("theres an error :%v\n", err)
		}
		if rows == 0 {
			fmt.Printf("no blog updated, you might have to create it")
			rw.WriteHeader(http.StatusBadRequest)
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
			fmt.Printf("theres an error in converting id into uint")
		}
		db.Delete(&models.BlogPost{}, id)
	}
}

//TODO: when the blod is deleted , the user can still get the blod (although its empty fields), so if the id doesnt exist return bad request
//(might have to create the blog)
