package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/usecase"
)

func ViewUserWishlist(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in Access Token"})
		return
	}
	WishedProducts, err := usecase.ViewUserWishlist(Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Products in wishlist", "wishlist": WishedProducts})

}

func AddProductToWishlist(c *gin.Context) {
	ProductID := c.Query("product_id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in Access Token"})
		return
	}
	err = usecase.AddProductToWishlist(ProductID, Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added to wishlist successfully"})
}

func RemoveProductFromWishlist(c *gin.Context) {
	ProductID := c.Query("product_id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in Access Token"})
		return
	}
	err = usecase.RemoveProductFromWishlist(ProductID, Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed from wishlist successfully"})
}

