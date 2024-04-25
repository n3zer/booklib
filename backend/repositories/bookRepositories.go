package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/n3zer/booklib/database"
	"github.com/n3zer/booklib/models"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Db.Find(&books)
	c.JSON(http.StatusOK, &books)
}

func AddBook(c *gin.Context) {
	book := models.Book{}
	c.ShouldBindHeader(&book)
	database.DB.Db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func DeleteBook(c *gin.Context) {
	book := models.Book{}
	c.BindHeader(&book)

	result := database.DB.Db.Delete(&models.Book{}, book)
	if result.RowsAffected != 0 {
		c.Status(http.StatusOK)
	}
	c.Status(http.StatusOK)
}

func UpdateBook(c *gin.Context) {
	oldBook := models.Book{}
	book := models.Book{}
	c.BindHeader(&book)

	result := database.DB.Db.First(&oldBook, book.ID).Updates(book)
	if result.RowsAffected != 0 {
		c.Status(http.StatusOK)
	}
	c.Status(http.StatusOK)
}
