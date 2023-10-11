package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	"github.com/priyanfadhil/ina-rec/service/model/user"
)

func (ctrl *UserController) GetUserById(c *gin.Context) {
	ctx := context.Background()
	var req user.GetUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Validation Error.", 102, err, nil))
		return
	}

	data, err := ctrl.UserService.GetUserById(ctx, req.Id)
	if err != nil {
		if err.Error() == "data user kosong" {
			c.JSON(http.StatusOK,
				api_response.BuildErrorResponse("Validation Error.", 102, err, data))
		} else {
			c.JSON(http.StatusInternalServerError,
				api_response.BuildErrorResponse("Internal Server Error.", 500, err, data))
		}
		return
	}

	c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Sukses", 200, data))
}
