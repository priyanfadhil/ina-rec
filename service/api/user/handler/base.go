package handler

import (
	"github.com/priyanfadhil/ina-rec/common/database"
	"github.com/priyanfadhil/ina-rec/service/api/user"
)

type userHandler struct {
	dbManager      database.DatabaseManager
	userRepository user.Repository
}

func NewHandler(
	dbManager database.DatabaseManager,
	userRepository user.Repository,
) user.Handler {
	return &userHandler{
		dbManager:      dbManager,
		userRepository: userRepository,
	}
}
