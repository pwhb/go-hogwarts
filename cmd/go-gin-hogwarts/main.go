package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pwhb/go-gin-hogwarts/pkg/database"
	"github.com/pwhb/go-gin-hogwarts/pkg/routes"
)

func main() {
	database.SetupDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
