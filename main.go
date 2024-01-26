package main

import (
	"github.com/anastasiya1155/rest-api/db"
	"github.com/anastasiya1155/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
