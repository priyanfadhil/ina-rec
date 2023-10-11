package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	modelTask "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	var req modelTask.CreateTaskRequest
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

	err := ctrl.TaskHandler.UpdateTask(ctx, &req, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Task Gagal Diperbarui", 400, err, ""))
		return
	}

	c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Task Berhasil Diperbarui", 201, ""))
}
