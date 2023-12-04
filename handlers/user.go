package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

func UserSignup(c *gin.Context) {

	var User models.UserSignUpDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}
	data, err := utils.Validation(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": data})
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
	Error, err := utils.Validation(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
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

	c.JSON(http.StatusOK, gin.H{"message": "Enter otp and new password"})

}

func ResetForgottenPassword(c *gin.Context) {
	var Newpassword models.ForgotPassword

	if c.ShouldBindJSON(&Newpassword) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details in correct format"})
		return
	}

	err := usecase.ResetForgottenPassword(Newpassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}

func ViewUserAddress(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Address, err := usecase.ViewUserAddress(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Address", "Address": Address})

}

func AddNewAddressDetails(c *gin.Context) {
	var Address models.Address

	if c.ShouldBindJSON(&Address) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details correctly"})
	}

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	AddressRes, err := usecase.AddAddress(Address, Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address added successfully", "Address": AddressRes})
}

func EditUserAddress(c *gin.Context) {
	var UpdateAddress models.Address

	if c.ShouldBindJSON(&UpdateAddress) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter constraints correctly"})
	}

	Aid := c.Query("id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	UpdatedAddress, err := usecase.UpdateAddress(Token, Aid, UpdateAddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully", "Address": UpdatedAddress})

}

func RemoveUserAddress(c *gin.Context) {
	Aid := c.Query("id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	err = usecase.DeleteAddress(Token, Aid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address removed successfully"})
}

func UserProfile(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserDetails, err := usecase.UserProfile(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Profile", "profile": UserDetails})

}

func UpdateUserProfile(c *gin.Context) {
	var UserDetails models.UserProfile

	if c.ShouldBindJSON(&UserDetails) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details correctly"})
	}

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUserDetails, err := usecase.UpdateUserProfile(Token, UserDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated User Profile", "profile": updatedUserDetails})

}
