package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func AdminAuthMiddleware(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		fmt.Println("he")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	role, err := utils.GetRoleFromToken(Token)
	if err != nil {
		fmt.Println("ji")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if role == "admin" {
		c.Next()
	}else{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}

