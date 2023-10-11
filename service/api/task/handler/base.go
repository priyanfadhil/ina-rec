package handler

import (
	"github.com/priyanfadhil/ina-rec/common/database"
	task "github.com/priyanfadhil/ina-rec/service/api/task"
)

type taskHandler struct {
	dbManager      database.DatabaseManager
	taskRepository task.Repository
}

func NewHandler(
	dbManager database.DatabaseManager,
	taskRepository task.Repository,
) task.Handler {
	return &taskHandler{
		dbManager:      dbManager,
		taskRepository: taskRepository,
	}
}
