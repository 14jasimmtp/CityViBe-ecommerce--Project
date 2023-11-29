package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func OrderFromCart(c *gin.Context) {
	var OrderInput models.CheckOut

	if c.ShouldBindJSON(&OrderInput) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Enter constraints correctly"})
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.OrderFromCart(Token, OrderInput.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})

}

func ViewCheckOut(c *gin.Context) {
	var coupon models.CheckoutCoupon
	if c.ShouldBindJSON(&coupon) != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Enter coupon correctly"})
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.CheckOut(Token,coupon.Coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CheckOut Page loaded successfully", "order Details": OrderDetails})
}

func ViewOrders(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.ViewUserOrders(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "orders", "order Details": OrderDetails})

}

func CancelOrder(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Query("id")

	err = usecase.CancelOrder(Token, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})

}

func CancelOrderByAdmin(c *gin.Context) {
	orderID := c.Query("id")
	err := usecase.CancelOrderByAdmin(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't cancel the order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func ShipOrderByAdmin(c *gin.Context) {
	orderId := c.Query("id")
	err := usecase.ShipOrders(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Shipped successfully"})
}

func DeliverOrderByAdmin(c *gin.Context) {
	id := c.Query("orderId")

	err := usecase.DeliverOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order delivered successfully"})

}

func CancelSingleProduct(c *gin.Context) {
	id := c.Query("pid")
	orderID := c.Query("oid")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderDetails, err := usecase.CancelSingleProduct(id, Token, orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cancelled product from orders.If you have paid online money will be refunded within 2 days", "Order Details": orderDetails})
}
