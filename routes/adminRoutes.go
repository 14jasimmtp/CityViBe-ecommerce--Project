package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
	"main.go/middlewares"
)

func AdminRoutes(r *gin.Engine) {

	//USER
	r.POST("/admin/login", handlers.AdminLogin)
	r.GET("/admin/users", middlewares.AdminAuthMiddleware, handlers.GetAllUsers)
	r.POST("/admin/users/block", middlewares.AdminAuthMiddleware, handlers.BlockUser)
	r.POST("/admin/users/unblock", middlewares.AdminAuthMiddleware, handlers.UnBlockUser)

	//product
	r.GET("/admin/products", middlewares.AdminAuthMiddleware, handlers.AllProducts)
	r.POST("/admin/product/add", middlewares.AdminAuthMiddleware, handlers.AddProduct)
	r.PUT("/admin/product/update", middlewares.AdminAuthMiddleware, handlers.EditProductDetails)
	r.DELETE("/product/:id/remove", middlewares.AdminAuthMiddleware, handlers.DeleteProduct)

	//category
	r.GET("admin/category", middlewares.AdminAuthMiddleware, handlers.GetCategory)
	r.POST("admin/category", middlewares.AdminAuthMiddleware, handlers.AddCategory)
	r.PUT("admin/category", middlewares.AdminAuthMiddleware, handlers.UpdateCategory)
	r.DELETE("admin/category", middlewares.AdminAuthMiddleware, handlers.DeleteCategory)
}
