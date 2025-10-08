package models

import "time"

type BlogPost struct {
	Id       uint16 `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	//Tags      []string  `gorm:"type:text[]" json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//we currently dont want anything to be null
