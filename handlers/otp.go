package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"main.go/models"
	"main.go/repository"
	"main.go/utils"
)

func VerifySignupOtp(c *gin.Context) {
	var otp models.OTP

	if c.Bind(&otp) != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Enter otp to login",
		})
		return
	}

	err := utils.CheckOtp(otp.Phone, otp.Otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid otp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully  signed up"})

}

func VerifyLoginOtp(c *gin.Context) {
	var otp models.OTP

	if c.Bind(&otp) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter otp to login"})
		return
	}

	err := utils.CheckOtp(otp.Phone, otp.Otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid otp"})
		return
	}

	user, err := repository.FindUserByPhone(otp.Phone)
	if err != nil{
		return
	}
	var client models.ClientToken
	err = copier.Copy(&client, &user)
	if err != nil {
		return
	}
	Tokenstring, err := utils.TokenGenerate(&client, "user")
	if err != nil {
		return
	}
	c.SetCookie("Authorisation", Tokenstring, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user successfully logged in", "user": user})
}
