package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saykaw/bloggingplatform/models"
	"gorm.io/gorm"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(rw, "This is the home page. Welcome!")
}

func CreatePostHandler(db *gorm.DB) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
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

		if r.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
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
			fmt.Printf("number of rows found: %v", result.RowsAffected)
		}
	}
}
