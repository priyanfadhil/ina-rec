package handler

import (
	// "context"
	"crypto/rand"
	"encoding/base64"

	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *userHandler) GenerateStateOauthCookie(ctx *gin.Context) (string, int) {
	// c, ok := ctx
	// if !ok {
	// 	// Handle the case where the context is not a Gin context
	// 	fmt.Println("masuk error nih")
	// 	return "", http.StatusInternalServerError
	// }

	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	cookie := &http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,                    // You may want to set this for security.
		SameSite: http.SameSiteStrictMode, // You can adjust this as needed.
	}

	http.SetCookie(ctx.Writer, cookie)

	return state, http.StatusOK
}
