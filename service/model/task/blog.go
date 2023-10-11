package task

type ModelTask struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}

func (ModelTask) ModelName() string {
	return "tasks"
}

func (ModelTask) TableName() string {
	return "public.tasks"
}

type GetTaskReq struct {
	Id int `uri:"id"`
}

type CreateTaskRequest struct {
	ID          int    `json:"-" uri:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}

func (CreateTaskRequest) ModelName() string {
	return "tasks"
}

func (CreateTaskRequest) TableName() string {
	return "public.tasks"
}

type TaskResponse struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
