package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
)

func AddProduct(c *gin.Context) {
	var product models.AddProduct

	if c.ShouldBindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Product details correctly"})
		return
	}
	NewProduct, err := usecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully", "product": NewProduct})
}

func EditProductDetails(c *gin.Context) {
	var product models.AddProduct
	id := c.Query("id")
	if c.ShouldBindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details correctly"})
		return
	}
	UpdatedProduct, err := usecase.EditProductDetails(id, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully", "product": UpdatedProduct})
	// err:=usecase.EditProductDetails(product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := usecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed successfully"})
}

func GetAllProducts(c *gin.Context) {
	products, err := usecase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "products list", "products": products})

}

func AllProducts(c *gin.Context) {
	products, err := usecase.SeeAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully retrieved products", "products": products})

}

func ShowSingleProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := usecase.GetSingleProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product details", "product": product})
}

func ViewProductCategoryWise(c *gin.Context){
	category:=c.Query("category")
	Products,err:=usecase.ViewProductCategoryWise(category)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusOK,gin.H{"message":"Categoried products","products":Products})
}
