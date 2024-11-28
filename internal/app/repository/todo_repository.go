package repository

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/entity"
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

func (r *TodoRepository) ReadAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) ReadByID(id uint) (*entity.Todo, error) {
	var todo entity.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Update(newTodo *entity.Todo) error {
	return r.db.Select("*").Updates(newTodo).Error
}
