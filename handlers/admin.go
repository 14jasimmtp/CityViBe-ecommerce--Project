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

// AdminLogin godoc
// @Summary Admin login
// @Description Authenticate and log in as an admin.
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_details body models.Admin true "Admin credentials for login"
// @Success 200 {object} string "message": "Admin logged in successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/login [post]
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

// GetAllUsers godoc
// @Summary view users
// @Description Retrieve a list of all users.
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Users are", "users": Users
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/users [get]
func GetAllUsers(c *gin.Context) {
	Users, err := usecase.GetAllUsers()
	if err != nil {
		fmt.Println("usecase error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users are", "users": Users})
}

// BlockUser godoc
// @Summary Block user
// @Description Block a user by their ID.
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Param id query string true "User ID to be blocked"
// @Success 200 {object} string "message": "User successfully blocked"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/users/block [post]
func BlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.BlockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully blocked"})
}

// UnBlockUser godoc
// @Summary Unblock user
// @Description Unblock a user by their ID.
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Param id query string true "User ID to be unblocked"
// @Success 200 {object} string "message": "User successfully unblocked"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/users/unblock [post]
func UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.UnBlockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully unblocked"})
}

// OrderDetailsForAdmin godoc
// @Summary Get all order details for admin
// @Description Retrieve all order details for administrative purposes.
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Order details retrieved successfully", "All orders": allOrderDetails
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/orders [get]
func OrderDetailsForAdmin(c *gin.Context) {
	allOrderDetails, err := usecase.GetAllOrderDetailsForAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't retrieve order details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order details retrieved successfully", "All orders": allOrderDetails})
}

// OrderDetailsforAdminWithID godoc
// @Summary view single orders
// @Description Retrieve order details for administrative purposes based on the given order ID.
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Param orderID query string true "Order ID to retrieve details for"
// @Success 200 {object} string"Order Products": OrderDetails
// @Failure 500 {object} string"error": "Internal Server Error"
// @Router /admin/orders/details [get]
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

// AddProductOffer godoc
// @Summary Add offer to a product
// @Description Add an offer to a product based on the provided product ID and offer ID.
// @Tags Admin Offer management
// @Accept json
// @Produce json
// @Param productid formData integer true "Product ID to add offer for"
// @Param offer formData integer true "Offer ID to be associated with the product"
// @Success 200 {object} string "offer added": prod
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/product/offer [post]
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

// AddCategoryOffer godoc
// @Summary Add offer to a category
// @Description Add an offer to a category based on the provided category ID and offer ID.
// @Tags Admin Offer AManagement
// @Accept json
// @Produce json
// @Param categoryid formData integer true "Category ID to add offer for"
// @Param offer formData integer true "Offer ID to be associated with the category"
// @Success 200 {object} string "offer added": productlist
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/category/offer [post]
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

// DashBoard godoc
// @Summary Get admin dashboard information
// @Description Retrieve information for the admin dashboard.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Admin dashboard", "dashboard": adminDashboard
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/dashboard [get]
func DashBoard(c *gin.Context) {
	adminDashboard, err := usecase.DashBoard()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "admin dashboard ", "dashboard": adminDashboard})
}
