package controller

import "github.com/priyanfadhil/ina-rec/service/api/user"

type UserController struct {
	UserService user.Handler
}

func NewController(userService user.Handler) *UserController {
	return &UserController{
		UserService: userService,
	}
}
