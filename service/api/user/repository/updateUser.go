package repository

import (
	"context"

	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

func (r *UserRepository) UpdateUser(ctx context.Context, db *gorm.DB, insert *userModel.UpdateUserRequest) error {
	err := db.WithContext(ctx).Where("id = ?", insert.ID).Updates(insert).Error

	if err != nil {
		return err
	}

	return nil
}
