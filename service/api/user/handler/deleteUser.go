package handler

import (
	"context"
)

func (s *userHandler) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepository.DeleteUser(ctx, s.dbManager.GetMaster(), id)
	if err != nil {
		return err
	}
	return nil
}
