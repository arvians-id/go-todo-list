package request

type CreateTodoRequest struct {
	Title           string `json:"title" validate:"required,max=100"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,number"`
	IsActive        bool   `json:"is_active" validate:"required,boolean"`
	Priority        string `json:"priority"`
}

type UpdateTodoRequest struct {
	Title    string `json:"title" validate:"required,max=100"`
	Priority string `json:"priority" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required,boolean"`
}
