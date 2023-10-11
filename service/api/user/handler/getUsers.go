package handler

import (
	"context"

	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (s *userHandler) GetUsers(ctx context.Context) ([]userModel.GetUserResponse, error) {
	user, err := s.userRepository.GetUsers(ctx, s.dbManager.GetMaster())
	if err != nil {
		return []userModel.GetUserResponse{}, err
	}

	return user, nil
}
