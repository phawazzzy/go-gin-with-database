package models

import "time"

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	DateCreated time.Time `json:"datetime"`
	DateUpdated time.Time `json:"updatedtime"`

}

type CreateBooksInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	DateCreated time.Time `json:"datetime"`
}
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	DateUpdated time.Time `json:"updatedTime"`
}
