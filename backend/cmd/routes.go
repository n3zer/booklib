package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n3zer/booklib/repositories"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.GET("/getbooks", repositories.GetBooks)
	app.POST("/addbook", repositories.AddBook)
	app.POST("/deletebook", repositories.DeleteBook)
	app.POST("/updatebook", repositories.UpdateBook)
}
