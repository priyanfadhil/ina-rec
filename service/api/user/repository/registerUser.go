package repository

import (
	"context"

	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

func (r *UserRepository) RegisterUser(ctx context.Context, db *gorm.DB, insert *userModel.CreateUserRequest) error {

	err := db.
		Create(&insert).Error

	if err != nil {
		return err
	}

	return nil
}
