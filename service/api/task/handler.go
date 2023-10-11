package task

import (
	"context"

	"github.com/priyanfadhil/ina-rec/service/model/auth"
	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
)

type Handler interface {
	GetTaskById(ctx context.Context, id int) (taskModel.TaskResponse, error)
	CreateTask(ctx context.Context, task *taskModel.ModelTask, user *auth.JWTUser) error
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task *taskModel.CreateTaskRequest, user *auth.JWTUser) error
	GetTasks(ctx context.Context) ([]taskModel.TaskResponse, error)
}
