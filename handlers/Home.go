package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"welcome to homepage.",
	})		
}