package task

import (
	"context"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
	"gorm.io/gorm"
)

type Repository interface {
	GetTaskById(ctx context.Context, db *gorm.DB, id int) (taskModel.TaskResponse, error)
	CreateTask(ctx context.Context, db *gorm.DB, insert *taskModel.ModelTask) error
	DeleteTask(ctx context.Context, db *gorm.DB, id int) error
	UpdateTask(ctx context.Context, db *gorm.DB, insert *taskModel.CreateTaskRequest) error
	GetTasks(ctx context.Context, db *gorm.DB) ([]taskModel.TaskResponse, error)
}
