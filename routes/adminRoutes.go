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
	r.POST("/admin/products", middlewares.AdminAuthMiddleware, handlers.AddProduct)
	r.PUT("/admin/products", middlewares.AdminAuthMiddleware, handlers.EditProductDetails)
	r.DELETE("admin/products/:id/remove", middlewares.AdminAuthMiddleware, handlers.DeleteProduct)

	//category
	r.GET("admin/category", middlewares.AdminAuthMiddleware, handlers.GetCategory)
	r.POST("admin/category", middlewares.AdminAuthMiddleware, handlers.AddCategory)
	r.PUT("admin/category", middlewares.AdminAuthMiddleware, handlers.UpdateCategory)
	r.DELETE("admin/category", middlewares.AdminAuthMiddleware, handlers.DeleteCategory)

	//order
	r.GET("admin/order-details", middlewares.AdminAuthMiddleware, handlers.OrderDetailsForAdmin)
	r.POST("admin/ship-order", middlewares.AdminAuthMiddleware, handlers.ShipOrderByAdmin)
	r.POST("admin/cancel-order", middlewares.AdminAuthMiddleware, handlers.CancelOrderByAdmin)
	r.GET("admin/order-single-details",middlewares.AdminAuthMiddleware,handlers.OrderDetailsforAdminWithID)
	r.POST("admin/deliver-order",middlewares.AdminAuthMiddleware,handlers.DeliverOrderByAdmin)

}
