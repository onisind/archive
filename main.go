package main

import (
	"archive-system/databaseProvaider"
	"archive-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseProvaider.ConnectPostgres()
	databaseProvaider.ConnectMongo()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
