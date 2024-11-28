package handler

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/dto"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	s *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{s}
}

func (h *TodoHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := h.s.Create(req.Title); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Status(201)
	}
}
