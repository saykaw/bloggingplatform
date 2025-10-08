package db

import (
	"fmt"

	"github.com/saykaw/bloggingplatform/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=sayali1605 dbname=blogsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("there was an error connecting to db: %v \n", err.Error())
		return nil
	}
	db.AutoMigrate(&models.BlogPost{}) //gorm automigrate: table created will be blog_posts
	return db
}

//ques:
//how do i see how many connections i have made to db
