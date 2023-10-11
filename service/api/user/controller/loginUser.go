package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	// "github.com/priyanfadhil/ina-rec/service/model"
)

func (ctrl *UserController) LoginUserController(c *gin.Context) {
	userLogin := make(map[string]interface{})
	ctx := context.Background()

	c.Bind(&userLogin)

	token, statusCode := ctrl.UserService.LoginUser(ctx, userLogin["name"].(string), userLogin["password"].(string))
	switch statusCode {
	case http.StatusUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "name atau password salah",
		})
		return

	case http.StatusInternalServerError:
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": "success",
		"token":    token,
	})
}
