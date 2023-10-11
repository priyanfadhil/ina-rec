package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/priyanfadhil/ina-rec/common/config"
	response "github.com/priyanfadhil/ina-rec/service/model/api_response"
	"github.com/priyanfadhil/ina-rec/service/model/auth"
)

type Middleware struct {
}

func NewMiddleware() Middle {
	return &Middleware{}
}

type Middle interface {
	JWTMiddleware() gin.HandlerFunc
}

func CreateToken(id int, name, email, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(168 * time.Hour).Unix()
	claims["iat"] = time.Now().Unix()
	claims["name"] = name
	claims["email"] = email
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

var bearer = "Bearer"

func (m *Middleware) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		publicKeyPem := config.GetConfig().JWTPublicKey
		publicKey := []byte(publicKeyPem)
		jwtToken := auth.NewJWT(publicKey)
		l := len(bearer)
		auth := c.GetHeader("Authorization")
		if len(auth) <= l+1 || !strings.EqualFold(auth[:l], bearer) {
			c.JSON(http.StatusBadRequest, response.AuthorizationError("unauthorized please login", 400, fmt.Errorf("invalid auth header"), ""))
			c.Abort()
			return
		}
		data, errGetToken := jwtToken.ValidateToken(auth[l+1:])
		if errGetToken != nil {
			c.JSON(http.StatusBadRequest, response.BuildErrorResponse("Error Get Token", 400, errGetToken, ""))
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user", data)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
