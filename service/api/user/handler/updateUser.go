package handler

import (
	"context"

	t "github.com/priyanfadhil/ina-rec/helper/time"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (s *userHandler) UpdateUser(ctx context.Context, req *userModel.UpdateUserRequest, user *auth.JWTUser) error {
	tx := s.dbManager.GetMaster()
	var (
		insert userModel.UpdateUserRequest
	)

	currentTime := t.Now().Format("2006-01-02 15:04:05")
	insert = userModel.UpdateUserRequest{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		UpdatedAt: currentTime,
		UpdatedBy: user.Name,
	}

	err := s.userRepository.UpdateUser(ctx, tx, &insert)
	if err != nil {
		return err
	}

	return nil
}
