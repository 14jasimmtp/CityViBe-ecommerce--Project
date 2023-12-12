package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

func AdminLogin(c *gin.Context) {
	var admin models.Admin

	if c.ShouldBindJSON(&admin) != nil {
		fmt.Println("binding error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details correctly"})
		return
	}

	Error, err := utils.Validation(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	admindetails, err := usecase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("Authorisation", admindetails.TokenString, 36000, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Admin logged in successfully"})

}

func GetAllUsers(c *gin.Context) {
	Users, err := usecase.GetAllUsers()
	if err != nil {
		fmt.Println("usecase error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users are", "users": Users})
}

func BlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.BlockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully blocked"})
}

func UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.UnBlockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully unblocked"})
}

func OrderDetailsForAdmin(c *gin.Context) {
	allOrderDetails, err := usecase.GetAllOrderDetailsForAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't retrieve order details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order details retrieved successfully", "All orders": allOrderDetails})
}

func OrderDetailsforAdminWithID(c *gin.Context) {
	orderID := c.Query("orderID")

	OrderDetails, err := usecase.GetOrderDetails(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Order Products": OrderDetails})
}

