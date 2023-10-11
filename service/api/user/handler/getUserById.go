package handler

import (
	"context"

	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (s *userHandler) GetUserById(ctx context.Context, id int) (userModel.GetUserResponse, error) {
	user, err := s.userRepository.GetUserById(ctx, s.dbManager.GetMaster(), id)
	if err != nil {
		return userModel.GetUserResponse{}, err
	}

	return user, nil
}
