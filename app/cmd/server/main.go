package main

import (
	"log"

	"job-hunting-service-management-backend/app/infrastructure/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB接続
	database, err := db.New()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	log.Println("Database connection successful!")

	// Ginルーター設定
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"status":  "Database connected successfully",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
