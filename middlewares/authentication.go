package middlewares

import (
	"go-gorm/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unautheticate",
				"message": err.Error()})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}

}
