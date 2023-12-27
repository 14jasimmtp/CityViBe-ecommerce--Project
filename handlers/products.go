package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/usecase"
	"main.go/utils"
)

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product with details and an image.
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product formData models.AddProduct true "Product details"
// @Param image formData file true "Product image"
// @Success 200 {object} string "message": "product added successfully", "product": models.AddProduct
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/products [post]
func AddProduct(c *gin.Context) {
	var product models.AddProduct

	if c.ShouldBind(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Product details correctly"})
		return
	}

	image, _ := c.FormFile("image")

	Error, err := utils.Validation(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	NewProduct, err := usecase.AddProduct(product,image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully", "product": NewProduct})
}

// EditProductDetails godoc
// @Summary Edit product details
// @Description Edit the details of an existing product by providing the product ID.
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query string true "Product ID to be edited"
// @Param product body models.AddProduct true "Updated product details"
// @Success 200 {object} string "message": "product updated successfully", "product": models.AddProduct
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/products [put]
func EditProductDetails(c *gin.Context) {
	var product models.AddProduct
	id := c.Query("product_id")
	if c.ShouldBindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details correctly"})
		return
	}

	Error, err := utils.Validation(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
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

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by providing the product ID.
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID to be deleted"
// @Success 200 {object} string "message": "product removed successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/products/remove/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := usecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed successfully"})
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "message": "products list", "products": [object]
// @Failure 404 {object} string "error": "Not Found"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products, err := usecase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "products list", "products": products})

}

// AllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products.
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "message": "successfully retrieved products", "products": [object]
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /admin/products [get]
func AllProducts(c *gin.Context) {
	products, err := usecase.SeeAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully retrieved products", "products": products})

}

// ShowSingleProduct godoc
// @Summary Get details of a single product
// @Description Retrieve details of a specific product by providing its ID.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} string "message": "product details", "product": object
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /products/{id} [get]
func ShowSingleProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := usecase.GetSingleProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product details", "product": product})
}

// FilterProducts godoc
// @Summary Filter products based on specified criteria
// @Description Filter products based on category, size, and price range.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category query string false "Product category"
// @Param size query string false "Product size"
// @Param minPrice query number false "Minimum price"
// @Param maxPrice query number false "Maximum price"
// @Success 200 {object} string "message": "filtered products", "products": object
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /products/filter [get]
func FilterProducts(c *gin.Context) {
	category := c.Query("category")
	size := c.Query("size")
	minPrice := c.Query("minPrice")
	maxPrice := c.Query("maxPrice")

	min, _ := strconv.ParseFloat(minPrice, 64)
	max, _ := strconv.ParseFloat(maxPrice, 64)
	Products, err := usecase.FilterProducts(category, size, min, max)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "filtered products", "products": Products})
}

// SearchProducts godoc
// @Summary Search for products based on a keyword
// @Description Search for products using a specified keyword.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param search body models.Search true "Search keyword"
// @Success 200 {object} string "Products": object, "message": "Searched Products"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /products/search [Get]
func SearchProducts(c *gin.Context) {
	var Search models.Search

	if c.ShouldBindJSON(&Search) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Constraints correctly"})
		return
	}
	products, err := usecase.SearchProduct(Search.Search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Products": products, "message": "Searched Products"})
}