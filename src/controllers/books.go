package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "models"

)

func FindBooks(c *gin.Context) {

	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books })

}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
  }

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
  }


// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput

	err := c.ShouldBindJSON(&input)


	if  err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }



  // GET /books/:id
// Find a book
func FindBook(c *gin.Context) {  // Get model if exist
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
  
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }
  

  // PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&book).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }


  // DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
  }
  
