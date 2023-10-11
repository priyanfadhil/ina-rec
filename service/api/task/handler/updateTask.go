package handler

import (
	"context"

	t "github.com/priyanfadhil/ina-rec/helper/time"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (s *taskHandler) UpdateTask(ctx context.Context, task *taskModel.CreateTaskRequest, user *auth.JWTUser) error {
	tx := s.dbManager.GetMaster()
	var (
		insert taskModel.CreateTaskRequest
	)

	currentTime := t.Now().Format("2006-01-02 15:04:05")

	insert = taskModel.CreateTaskRequest{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UpdatedBy:   user.Name,
		UpdatedAt:   currentTime,
	}

	err := s.taskRepository.UpdateTask(ctx, tx, &insert)
	if err != nil {
		return err
	}

	return nil
}
