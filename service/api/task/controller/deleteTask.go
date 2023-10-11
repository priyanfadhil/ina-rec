package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	modelTask "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	ctx := context.Background()
	var req modelTask.GetTaskReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data Task Kosong", 400, err, nil))
		return
	}

	err := ctrl.TaskHandler.DeleteTask(ctx, req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Data Task Kosong", 400, err, nil))
		return
	}

	c.JSON(http.StatusOK,
		api_response.BuildSuccessResponse("Sukses", 200, "Data Berhasil Dihapus"))
}
