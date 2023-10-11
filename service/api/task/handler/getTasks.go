package handler

import (
	"context"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (s *taskHandler) GetTasks(ctx context.Context) ([]taskModel.TaskResponse, error) {
	task, err := s.taskRepository.GetTasks(ctx, s.dbManager.GetMaster())
	if err != nil {
		return []taskModel.TaskResponse{}, err
	}

	return task, nil
}
