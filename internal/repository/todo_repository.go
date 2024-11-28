package repository

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/entity"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) Create(title string) error {
	return r.db.Create(&entity.Todo{Title: title}).Error
}
