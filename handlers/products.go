package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func AddProduct(c *gin.Context) {
	var product models.Product

	if c.ShouldBindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Product details correctly"})
		return
	}
	productDetails, err := usecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully", "product": productDetails})
}

func EditProductDetails(c *gin.Context) {
	var product models.Product

	if c.ShouldBindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details correctly"})
		return
	}
	// err:=usecase.EditProductDetails(product)
}

func DeleteProduct(c *gin.Context) {

}
