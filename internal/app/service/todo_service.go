package service

import (
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/dto"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/entity"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/app/repository"
	"github.com/gdgc-ub/web-todolist-backend-informal/internal/pkg/response"
	"log"
	"net/http"
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
		return response.NewErrorResponse(http.StatusInternalServerError, "something went wrong")
	}

	return nil
}

func (s *TodoService) ReadAll() ([]*entity.Todo, error) {
	todos, err := s.r.ReadAll()
	if err != nil {
		log.Println("Error reading todos: ", err)
		return nil, response.NewErrorResponse(http.StatusInternalServerError, "something went wrong")
	}

	return todos, nil
}

func (s *TodoService) ReadByID(req dto.ReadTodoByIDRequest) (*entity.Todo, error) {
	todo, err := s.r.ReadByID(req.ID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, response.NewErrorResponse(http.StatusNotFound, "todo not found")
		}

		log.Println("Error reading todo: ", err)
		return nil, response.NewErrorResponse(http.StatusInternalServerError, "something went wrong")
	}

	return todo, nil
}
