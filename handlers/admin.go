package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

func AddOffer(c *gin.Context) {
	var offer models.Offer

	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := usecase.ExecuteAddOffer(&offer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "offer added sussefully"})
}

func AllOffer(c *gin.Context) {

	offerlist, err := usecase.ExecuteGetOffers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offers": offerlist})
}

func AddProductOffer(c *gin.Context) {
	strpro := c.PostForm("productid")
	stroffer := c.PostForm("offer")
	productid, err := strconv.Atoi(strpro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conv failed"})
		return
	}
	offer, err := strconv.Atoi(stroffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conv failed"})
		return
	}
	prod, err1 := usecase.ExecuteAddProductOffer(productid, offer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offer added ": prod})
}

func AddCategoryOffer(c *gin.Context) {
	strcat := c.PostForm("categoryid")
	stroffer := c.PostForm("offer")
	categoryid, err := strconv.Atoi(strcat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str1 conv failed"})
		return
	}
	offer, err := strconv.Atoi(stroffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conv failed"})
		return
	}
	productlist, err1 := usecase.ExecuteCategoryOffer(categoryid, offer)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offer addded": productlist})
}

func DashBoard(c *gin.Context) {
	adminDashboard, err := usecase.DashBoard()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "admin dashboard ", "dashboard": adminDashboard})
}
