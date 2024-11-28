package dto

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

type ReadTodoByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdateTodoRequest struct {
	ID    uint   `uri:"id" binding:"required"`
	Title string `json:"title" binding:"omitempty,min=1,max=255"`
	Done  bool   `json:"done" binding:"omitempty"`
}
