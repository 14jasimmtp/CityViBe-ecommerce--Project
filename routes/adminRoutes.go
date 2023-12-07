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
	r.DELETE("admin/products/remove/:id", middlewares.AdminAuthMiddleware, handlers.DeleteProduct)

	//category
	r.GET("admin/category", middlewares.AdminAuthMiddleware, handlers.GetCategory)
	r.POST("admin/category", middlewares.AdminAuthMiddleware, handlers.AddCategory)
	r.PUT("admin/category", middlewares.AdminAuthMiddleware, handlers.UpdateCategory)
	r.DELETE("admin/category", middlewares.AdminAuthMiddleware, handlers.DeleteCategory)

	//order
	r.GET("admin/orders", middlewares.AdminAuthMiddleware, handlers.OrderDetailsForAdmin)
	r.POST("admin/orders/ship", middlewares.AdminAuthMiddleware, handlers.ShipOrderByAdmin)
	r.POST("admin/orders/cancel", middlewares.AdminAuthMiddleware, handlers.CancelOrderByAdmin)
	r.GET("admin/orders/details",middlewares.AdminAuthMiddleware,handlers.OrderDetailsforAdminWithID)
	r.POST("admin/orders/deliver",middlewares.AdminAuthMiddleware,handlers.DeliverOrderByAdmin)

	//coupons
	r.POST("admin/coupon",middlewares.AdminAuthMiddleware,handlers.MakeCoupon)
	r.PUT("admin/coupon/disable",middlewares.AdminAuthMiddleware,handlers.DisableCoupon)
	r.PUT("admin/coupon/enable",middlewares.AdminAuthMiddleware,handlers.EnableCoupon)
	r.GET("admin/coupon",middlewares.AdminAuthMiddleware,handlers.ViewCouponsAdmin)
	r.PUT("admin/coupon/update",middlewares.AdminAuthMiddleware,handlers.UpdateCoupon)

	//salesreport
	r.POST("admin/salesreport",middlewares.AdminAuthMiddleware,handlers.Salesreport)
	// r.GET("admin/Dashboard",middlewares.AdminAuthMiddleware,handlers.AdminDashboard)

}
