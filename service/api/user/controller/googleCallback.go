package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) OauthGoogleCallback(c *gin.Context) {
	// Read oauthState from Cookie
	oauthState, err := c.Cookie("oauthstate")
	if err != nil {
		log.Println("error reading oauthstate cookie:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	if c.DefaultPostForm("state", "") != oauthState {
		log.Println("invalid oauth google state")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	data, err := ctrl.UserService.GetUserDataFromGoogle(c.DefaultPostForm("code", ""))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "internal",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": "success",
		"data":     data,
	})
}
