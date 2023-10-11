package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/priyanfadhil/ina-rec/service/model/auth"
	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
)

type Handler interface {
	LoginUser(ctx context.Context, name, password string) (string, int)
	RegisterUser(ctx context.Context, usr *userModel.CreateUserRequest) error
	UpdateUser(ctx context.Context, req *userModel.UpdateUserRequest, user *auth.JWTUser) error
	GetUserById(ctx context.Context, id int) (userModel.GetUserResponse, error)
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]userModel.GetUserResponse, error)
	GenerateStateOauthCookie(ctx *gin.Context) (string, int)
	GetUserDataFromGoogle(code string) ([]byte, error)
}
