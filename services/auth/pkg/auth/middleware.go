package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const TokenKey string = "token"

func BearerTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		c.Set(TokenKey, token)
		c.Next()
	}
}
