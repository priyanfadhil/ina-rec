package handler

import (
	"context"
	"net/http"

	"github.com/priyanfadhil/ina-rec/common/config"
	bcrypt "github.com/priyanfadhil/ina-rec/helper"
	"github.com/priyanfadhil/ina-rec/service/middleware"
)

func (s *userHandler) LoginUser(ctx context.Context, name, password string) (string, int) {
	user, _ := s.userRepository.GetOneUserByName(ctx, s.dbManager.GetMaster(), name)

	match := bcrypt.CheckPasswordHash(password, user.Password)
	if !match {
		return "", http.StatusUnauthorized
	}

	token, err := middleware.CreateToken(user.ID, user.Name, user.Email, config.GetConfig().JWTPublicKey)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
