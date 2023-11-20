package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main.go/models"
	"main.go/usecase"
)

func UserSignup(c *gin.Context) {

	var User models.UserSignUpDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}
	err := validator.New().Struct(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "details not satisfied"})
		return
	}

	err = usecase.SignUp(User)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully signed up.Enter otp to login."})

}

func UserLogin(c *gin.Context) {
	var User models.UserLoginDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}

	err := validator.New().Struct(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "details not satisfied"})
		return
	}

	err = usecase.UserLogin(User)
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

func ForgotPassword(c *gin.Context) {
	var forgotPassword models.Phone
	if c.ShouldBindJSON(&forgotPassword) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter constraints correctly"})
	}
	err := usecase.ForgotPassword(forgotPassword.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})

}

func ResetForgottenPassword(c *gin.Context){
	var Newpassword models.ForgotPassword

	if c.ShouldBindJSON(&Newpassword) != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Enter details in correct format"})
		return
	}

	err:=usecase.ResetForgottenPassword(Newpassword)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{"message":"password changed successfully"})
}


func ViewUserAddress (c *gin .Context){

}

func EditUserAddress (c *gin.Context){

}

func RemoveUserAddress (c *gin.Context){

}

func AddNewAddressDetails (c *gin.Context){

}

