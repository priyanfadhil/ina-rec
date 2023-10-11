package handler

import (
	"context"

	bcrypt "github.com/priyanfadhil/ina-rec/helper"
	t "github.com/priyanfadhil/ina-rec/helper/time"
	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (s *userHandler) RegisterUser(ctx context.Context, usr *userModel.CreateUserRequest) error {
	tx := s.dbManager.GetMaster()
	var (
		insert userModel.CreateUserRequest
	)

	currentTime := t.Now().Format("2006-01-02 15:04:05")
	pass, _ := bcrypt.HashPassword(usr.Password)
	insert = userModel.CreateUserRequest{
		Name:      usr.Name,
		Email:     usr.Email,
		Password:  pass,
		CreatedAt: currentTime,
		CreatedBy: usr.Name,
	}

	err := s.userRepository.RegisterUser(ctx, tx, &insert)
	if err != nil {
		return err
	}

	return nil
}
