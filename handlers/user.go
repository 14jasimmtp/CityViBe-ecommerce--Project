package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func UserSignup(c *gin.Context) {

	var User models.UserSignUpDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}

	err := usecase.SignUp(User)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enter otp to signup"})

}

func UserLogin(c *gin.Context) {
	var User models.UserLoginDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}

	err := usecase.UserLogin(User)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enter otp to login"})

}

func UserLogout(c *gin.Context) {
	c.SetCookie("Authorisation", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user logged out successfully"})
	fmt.Println("cookie deleted")
}
