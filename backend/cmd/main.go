package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n3zer/booklib/database"
)

func main() {
	database.ConnectDb()
	app := gin.Default()
	SetupRoutes(app)
	app.Run(":3000")
}
