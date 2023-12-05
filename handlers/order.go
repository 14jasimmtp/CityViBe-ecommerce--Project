package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

func OrderFromCart(c *gin.Context) {
	var OrderInput models.CheckOut

	if c.ShouldBindJSON(&OrderInput) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Enter constraints correctly"})
		return
	}
	Error,err:=utils.Validation(OrderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":Error})
		return
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if OrderInput.PaymentID == 1 || OrderInput.PaymentID == 2 {
		OrderDetails, err := usecase.ExecutePurchase(Token, OrderInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}

	} else if OrderInput.PaymentID == 3 {
		OrderDetails, err := usecase.ExecutePurchaseWallet(Token, OrderInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter a valid payment method"})
		return
	}
}

func ViewCheckOut(c *gin.Context) {
	var coupon models.CheckoutCoupon
	if c.ShouldBindJSON(&coupon) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter coupon correctly"})
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.CheckOut(Token, coupon.Coupon)
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

	orderId := c.Query("order_id")
	product_id := c.Query("product_id")

	err = usecase.CancelOrder(Token, orderId, product_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})

}

func CancelOrderByAdmin(c *gin.Context) {
	orderID := c.Query("user_id")
	productID:=c.Query("product_id")
	userID:=c.Query("user_id")
	err := usecase.CancelOrderByAdmin(userID,orderID,productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't cancel the order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func ShipOrderByAdmin(c *gin.Context) {
	orderId := c.Query("OrderId")
	pid := c.Query("productId")
	userID:=c.Query("userID")
	err := usecase.ShipOrders(userID,orderId, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Shipped successfully"})
}

func DeliverOrderByAdmin(c *gin.Context) {
	id := c.Query("orderId")
	pid := c.Query("productId")
	userID:=c.Query("userID")

	err := usecase.DeliverOrder(userID,id, pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order delivered successfully"})

}

func ReturnOrder(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Query("order_id")
	productId := c.Query("product_id")

	err = usecase.ReturnOrder(Token, orderId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order returned successfully.Amount will be credited to wallet."})
}
