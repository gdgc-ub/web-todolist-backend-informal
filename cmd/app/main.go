package main

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/handler"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/repository"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/service"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	v1 := r.Group("/v1")

	v1.POST("/todos", todoHandler.Create())
	v1.GET("/todos", todoHandler.ReadAll())
	v1.GET("/todos/:id", todoHandler.ReadByID())
	v1.PUT("/todos/:id", todoHandler.Update())
	v1.DELETE("/todos/:id", todoHandler.Delete())

	r.Run(":" + os.Getenv("APP_PORT"))
}
