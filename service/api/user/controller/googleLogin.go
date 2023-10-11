package controller

import (
	// "context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/common/config"
	// "github.com/golang-jwt/jwt"
	// "github.com/priyanfadhil/ina-rec/service/model"
)

func (ctrl *UserController) OauthGoogleLogin(c *gin.Context) {
	oauthState, statusCode := ctrl.UserService.GenerateStateOauthCookie(c)
	switch statusCode {
	case http.StatusInternalServerError:
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "internal",
		})
		return
	}

	fmt.Println(oauthState)
	u := config.GoogleOauthConfig.AuthCodeURL(oauthState)
	fmt.Println(u)
	c.Redirect(http.StatusTemporaryRedirect, u)
}
