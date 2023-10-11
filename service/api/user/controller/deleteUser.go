package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
)

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	ctx := context.Background()
	var req modelUser.GetUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data User Kosong", 400, err, nil))
		return
	}

	err := ctrl.UserService.DeleteUser(ctx, req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data User Kosong", 400, err, nil))
		return
	}

	c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Sukses", 200, "Data Berhasil Dihapus"))
}
