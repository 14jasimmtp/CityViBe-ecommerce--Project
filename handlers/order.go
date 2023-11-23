package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/usecase"
)

func OrderFromCart(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CartId := c.Query("cartId")
	AddressId := c.Query("AddressId")

	OrderDetails, err := usecase.OrderFromCart(Token, CartId, AddressId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})

}

func ViewCheckOut(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := usecase.CheckOut(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
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
	orderId:=c.Query("id")
	err:=usecase.ShipOrders(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"couldn't ship order"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"Order Shipped successfully"})
}
