package main

import (
	"ProductService/internal/db"
	"ProductService/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("Starting the server")
	db.DB = db.Connect()
	defer db.DB.Close()
	log.Println("continuec")
	gin := gin.Default()
	routes.SetupRoutes(gin)
	gin.Run(":8080")
	log.Println("Server started on port 8080")
}
