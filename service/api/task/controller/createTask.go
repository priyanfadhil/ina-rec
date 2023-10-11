package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
	modelTask "github.com/priyanfadhil/ina-rec/service/model/task"
)

func (ctrl *TaskController) CreateTask(c *gin.Context) {
	var u modelTask.ModelTask
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

	ctx := c.Request.Context()
	user, _ := ctx.Value("user").(auth.JWTUser)

	err := ctrl.TaskHandler.CreateTask(ctx, &u, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			api_response.BuildErrorResponse("Task Gagal Dibuat", 400, err, ""))
		return
	}

	c.JSON(http.StatusCreated,
		api_response.BuildSuccessResponse("Task Berhasil Dibuat", 201, ""))
}
