package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/usecase"
)

// ViewUserWishlist godoc
// @Summary View products in user's wishlist
// @Description Retrieve and display the products currently present in the user's wishlist.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "message": "Products in wishlist", "wishlist": []models.Product
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /wishlist [get]
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

// AddProductToWishlist godoc
// @Summary Add a product to the user's wishlist
// @Description Add a specific product to the wishlist of the authenticated user.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query string true "Product ID to add to wishlist"
// @Success 200 {object} string "message": "product added to wishlist successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /wishlist [post]
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

// RemoveProductFromWishlist godoc
// @Summary Remove a product from the user's wishlist
// @Description Remove a specific product from the wishlist of the authenticated user.
// @Tags Wishlist
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query string true "Product ID to remove from wishlist"
// @Success 200 {object} string "message": "product removed from wishlist successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /wishlist [delete]
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

