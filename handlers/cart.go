package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/usecase"
)

func AddToCart(c *gin.Context) {
	pid := c.Query("id")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Cart, err := usecase.AddToCart(pid, Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added to cart successfully", "Cart": Cart})
}

func ViewCart(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserCart, err := usecase.ViewCart(Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart details", "Cart": UserCart})

}

func RemoveProductsFromCart(c *gin.Context) {
	id := c.Query("id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Cart,err := usecase.RemoveProductsFromCart(id, Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product removed from cart successfully","Cart":Cart})

}

func IncreaseQuantityUpdate(c *gin.Context) {
	pid := c.Query("product")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err=usecase.UpdateQuantityIncrease(Token,pid)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to add quantity"})
		return
	}

	err=usecase.UpdatePriceAdd(Token,pid)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to add quantity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "quantity added successfully"})

}


func DecreaseQuantityUpdate(c *gin.Context) {
	pid := c.Query("product")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err=usecase.UpdateQuantityDecrease(Token,pid)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to decrease quantity"})
		return
	}

	err=usecase.UpdatePriceDecrease(Token,pid)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to decrease quantity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "quantity decreased by 1 successfully"})
}

func EraseCartAfterOrder(c *gin.Context){
	Token,err:=c.Cookie("Authorisation")
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	cart,err:=usecase.EraseCart(Token)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"cart emptied successfully","cart":cart})
}