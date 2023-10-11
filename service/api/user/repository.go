package user

import (
	"context"

	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

type Repository interface {
	GetOneUserByName(ctx context.Context, db *gorm.DB, name string) (user modelUser.User, err error)
	GetUserById(ctx context.Context, db *gorm.DB, id int) (modelUser.GetUserResponse, error)
	RegisterUser(ctx context.Context, db *gorm.DB, insert *modelUser.CreateUserRequest) error
	UpdateUser(ctx context.Context, db *gorm.DB, insert *modelUser.UpdateUserRequest) error
	DeleteUser(ctx context.Context, db *gorm.DB, id int) error
	GetUsers(ctx context.Context, db *gorm.DB) ([]modelUser.GetUserResponse, error)
}
