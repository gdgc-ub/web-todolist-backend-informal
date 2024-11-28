package service

import (
	"errors"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/entity"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/repository"
	"log"
)

type TodoService struct {
	r *repository.TodoRepository
}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{r}
}

func (s *TodoService) Create(title string) error {
	if err := s.r.Create(title); err != nil {
		log.Println("Error creating todo: ", err)
		return errors.New("something went wrong")
	}

	return nil
}

func (s *TodoService) ReadAll() ([]*entity.Todo, error) {
	todos, err := s.r.ReadAll()
	if err != nil {
		log.Println("Error reading todos: ", err)
		return nil, errors.New("something went wrong")
	}

	return todos, nil
}
