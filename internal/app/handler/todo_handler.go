package handler

import (
	"errors"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/dto"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/service"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.s.Create(req.Title); err != nil {
			var errResp *response.ErrorResponse
			errors.As(err, &errResp)
			c.JSON(errResp.Code, errResp)
			return
		}

		c.Status(http.StatusCreated)
	}
}

func (h *TodoHandler) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := h.s.ReadAll()
		if err != nil {
			var errResp *response.ErrorResponse
			errors.As(err, &errResp)
			c.JSON(errResp.Code, errResp)
			return
		}

		c.JSON(http.StatusOK, todos)
	}
}

func (h *TodoHandler) ReadByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.TodoByIDRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		todo, err := h.s.ReadByID(req)
		if err != nil {
			var errResp *response.ErrorResponse
			errors.As(err, &errResp)
			c.JSON(errResp.Code, errResp)
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func (h *TodoHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.UpdateTodoRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := h.s.Update(req); err != nil {
			var errResp *response.ErrorResponse
			errors.As(err, &errResp)
			c.JSON(errResp.Code, errResp)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (h *TodoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.TodoByIDRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.s.Delete(req); err != nil {
			var errResp *response.ErrorResponse
			errors.As(err, &errResp)
			c.JSON(errResp.Code, errResp)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
