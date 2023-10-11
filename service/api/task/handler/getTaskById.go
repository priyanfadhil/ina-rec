package handler

import (
	"context"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (s *taskHandler) GetTaskById(ctx context.Context, id int) (taskModel.TaskResponse, error) {
	task, err := s.taskRepository.GetTaskById(ctx, s.dbManager.GetMaster(), id)
	if err != nil {
		return taskModel.TaskResponse{}, err
	}

	return task, nil
}
