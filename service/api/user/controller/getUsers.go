package controller

import (
	"context"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
)

func (ctrl *UserController) GetUsers(c *gin.Context) {
	ctx := context.Background()

	data, err := ctrl.UserService.GetUsers(ctx)
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
