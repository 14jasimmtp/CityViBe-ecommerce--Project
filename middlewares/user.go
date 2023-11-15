package middlewares

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func UserAuthMiddleware(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.Abort()
	}
	role, err := utils.GetRoleFromToken(Token)
	if err != nil {
		c.Abort()
	}
	if role == "user" {
		c.Next()

	}
}