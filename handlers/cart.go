package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/usecase"
)

// AddToCart godoc
// @Summary Add product to user's cart
// @Description Add a product to the user's cart based on the provided product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID to add to the cart"
// @Success 200 {object} string "message": "Product added to cart successfully", "Cart": Cart
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /cart [post]
func AddToCart(c *gin.Context) {
	pid := c.Query("product_id")

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

// ViewCart godoc
// @Summary View user's cart
// @Description Retrieve details of the user's cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Cart details", "Cart": UserCart
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /cart [get]
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

// RemoveProductsFromCart godoc
// @Summary Remove product from user's cart
// @Description Remove a product from the user's cart based on the provided product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID to remove from the cart"
// @Success 200 {object} string "message": "Product removed from cart successfully", "Cart": Cart
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /cart [delete]
func RemoveProductsFromCart(c *gin.Context) {
	id := c.Query("product_id")
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

// IncreaseQuantityUpdate godoc
// @Summary Increase quantity of a product in the user's cart
// @Description Increase the quantity of a product in the user's cart based on the provided product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID to increase quantity"
// @Success 200 {object} string "message": "Quantity added successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /cart/add-quantity [put]
func IncreaseQuantityUpdate(c *gin.Context) {
	pid := c.Query("product_id")

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

// DecreaseQuantityUpdate godoc
// @Summary Decrease quantity of a product in the user's cart
// @Description Decrease the quantity of a product in the user's cart based on the provided product ID.
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID to decrease quantity"
// @Success 200 {object} string "message": "Quantity decreased by 1 successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /cart/reduce-quantity [put]
func DecreaseQuantityUpdate(c *gin.Context) {
	pid := c.Query("product_id")

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