package controllers

import (
	"fmt"
	"gin-db/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// to export a function to another file it has to be in capital

// FindBooks find all book
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook create a new book
func CreateBook(c *gin.Context) {
	var input models.CreateBooksInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": err})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author, DateCreated: time.Now() }
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

// GetOneBook get one book using the id
func GetOneBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook update a book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		fmt.Println(c.Param("id"))
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found, data might have been deleted or moved"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := models.Book{Title: input.Title, Author: input.Author, DateUpdated: time.Now() }


	models.DB.Model(&book).Updates(&payload)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook delete a book
func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
