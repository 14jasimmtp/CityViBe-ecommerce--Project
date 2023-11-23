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

	err = usecase.AddToCart(pid, Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added to cart successfully"})
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

	err = usecase.RemoveProductsFromCart(id, Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product removed from cart successfully"})

}

func UpdateQuantityFromCart(c *gin.Context){
	pid:=c.Query("id")
	quantity:=c.Query("quantity")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdatedCart,err:=usecase.UpdateQuantityFromCart(Token,pid,quantity)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"quantity updated successfully" ,"Cart":UpdatedCart})


}
