package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// CustomValidatorMiddleware is a middleware that converts the app.Validator to a Gin Gonic middleware.
func CustomValidatorMiddleware(validator *CustomValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// You can perform any necessary validation here using your app.Validator
		// For example:
		// if err := validator.Validate(c); err != nil {
		//     // Handle validation error
		//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//     c.Abort()
		//     return
		// }

		// If validation passes, continue to the next handler
		c.Next()
	}
}
