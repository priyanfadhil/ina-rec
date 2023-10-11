package handler

import (
	"context"
)

func (s *taskHandler) DeleteTask(ctx context.Context, id int) error {

	err := s.taskRepository.DeleteTask(ctx, s.dbManager.GetMaster(), id)
	if err != nil {
		return err
	}
	return nil
}
