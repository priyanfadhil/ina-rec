package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var req modelUser.UpdateUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Validation Error.", 102, err, nil))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Validation Error.", 102, err, &req))
		return
	}

	ctx := c.Request.Context()
	user, _ := ctx.Value("user").(auth.JWTUser)

	err := ctrl.UserService.UpdateUser(ctx, &req, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("User Gagal Diperbarui", 400, err, ""))
		return
	}

	c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("User Berhasil Diperbarui", 201, ""))
}
