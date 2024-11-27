package main

import (
	"fmt"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	db := database.NewPostgresDB()

	fmt.Println(db) //TODO: Remove this line

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
