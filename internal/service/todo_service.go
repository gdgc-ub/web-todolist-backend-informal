package service

import (
	"errors"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/repository"
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
