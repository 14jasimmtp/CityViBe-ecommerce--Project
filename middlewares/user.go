package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func UserAuthMiddleware(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login to view page"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	role, err := utils.GetRoleFromToken(Token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if role == "user" {
		c.Next()

	}
}
