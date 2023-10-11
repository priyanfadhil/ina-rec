package controller

import (
	task "github.com/priyanfadhil/ina-rec/service/api/task"
)

type TaskController struct {
	TaskHandler task.Handler
}

func NewController(taskHandler task.Handler) *TaskController {
	return &TaskController{
		TaskHandler: taskHandler,
	}
}
