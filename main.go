package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"models"
    "controllers"
)

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data":"Hello World"})

	})
	r.GET("/myname", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Name":"Lokesh Ajja"})
	})

	r.GET("/books", controllers.FindBooks) // finds all books
	r.POST("/books", controllers.CreateBook) // post books
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)

	r.DELETE("/books/:id", controllers.DeleteBook) // new


	models.ConnectDatabase() // new

	r.Run()

}


