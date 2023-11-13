package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
	"main.go/middlewares"
)

func AdminRoutes(r *gin.Engine) {

	//
	r.POST("/admin/login", handlers.AdminLogin)
	r.GET("/admin/users",middlewares.AdminAuthMiddleware, handlers.GetAllUsers)
	r.POST("/admin/users/block",middlewares.AdminAuthMiddleware,handlers.BlockUser)
	r.POST("/admin/users/unblock",middlewares.AdminAuthMiddleware,handlers.UnBlockUser)
	
	//product
	r.POST("/admin/product/add",middlewares.AdminAuthMiddleware,handlers.AddProduct)
	r.PUT("/admin/product/edit",middlewares.AdminAuthMiddleware,handlers.EditProductDetails)
	r.DELETE("/product/remove",middlewares.AdminAuthMiddleware,handlers.DeleteProduct)

}
