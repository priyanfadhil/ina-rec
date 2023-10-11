package repository

import task "github.com/priyanfadhil/ina-rec/service/api/task"

type TaskRepository struct {
}

func NewTaskRepository() task.Repository {
	return &TaskRepository{}
}
