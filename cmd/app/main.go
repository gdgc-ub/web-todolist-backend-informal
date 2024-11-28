package main

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/handler"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/repository"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/service"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/pkg/database"
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
	r.GET("/todos/:id", todoHandler.ReadByID())
	r.PUT("/todos/:id", todoHandler.Update())
	r.DELETE("/todos/:id", todoHandler.Delete())

	r.Run(":8080")
}
