package main

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/database"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/handler"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/repository"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	db := database.NewPostgresDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	r.POST("/todos", todoHandler.Create())
	r.GET("/todos", todoHandler.ReadAll())

	r.Run(":8080")
}
