package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var u modelUser.CreateUserRequest
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Validation Error.", 102, err, &u))
		return
	}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Validation Error.", 102, err, &u))
		return
	}

	ctx := context.Background()
	err := ctrl.UserService.RegisterUser(ctx, &u)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("User Gagal Dibuat", 400, err, ""))
		return
	}

	c.JSON(http.StatusCreated,
		api_response.BuildSuccessResponse("User Berhasil Dibuat", 201, ""))
}
