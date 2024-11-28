package dto

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}
