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
	Error, err := utils.Validation(OrderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
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
	var cancel models.AdminOrder
	if c.ShouldBindJSON(&cancel) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Constraints correctly"})
		return
	}

	Error, err := utils.Validation(cancel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	err = usecase.CancelOrderByAdmin(cancel.UserID, cancel.OrderID, cancel.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func ShipOrderByAdmin(c *gin.Context) {
	var Ship models.AdminOrder
	if c.ShouldBindJSON(&Ship) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter constraints correctly"})
		return
	}

	Error, err := utils.Validation(Ship)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
	}

	err = usecase.ShipOrders(Ship.UserID, Ship.OrderID, Ship.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Shipped successfully"})
}

func DeliverOrderByAdmin(c *gin.Context) {
	var Deliver models.AdminOrder

	if c.ShouldBindJSON(&Deliver) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter constraints correctly"})
		return
	}

	Error, err := utils.Validation(Deliver)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	err = usecase.DeliverOrder(Deliver.UserID, Deliver.OrderID, Deliver.ProductID)
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
