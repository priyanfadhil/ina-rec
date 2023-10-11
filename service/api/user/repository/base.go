package repository

import "github.com/priyanfadhil/ina-rec/service/api/user"

type UserRepository struct {
}

func NewUserRepository() user.Repository {
	return &UserRepository{}
}
