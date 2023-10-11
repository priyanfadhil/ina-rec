package handler

import (
	"context"
	"strconv"

	t "github.com/priyanfadhil/ina-rec/helper/time"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (s *taskHandler) CreateTask(ctx context.Context, task *taskModel.ModelTask, user *auth.JWTUser) error {
	tx := s.dbManager.GetMaster()
	var (
		insert taskModel.ModelTask
	)

	currentTime := t.Now().Format("2006-01-02 15:04:05")
	usrId, _ := strconv.Atoi(user.Id)
	insert = taskModel.ModelTask{
		UserID:      usrId,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedBy:   user.Name,
		CreatedAt:   currentTime,
	}

	err := s.taskRepository.CreateTask(ctx, tx, &insert)
	if err != nil {
		return err
	}

	return nil
}
