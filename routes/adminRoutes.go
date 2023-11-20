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
	r.DELETE("/product/:id/remove", middlewares.AdminAuthMiddleware, handlers.DeleteProduct)

	//category
	r.GET("admin/category", middlewares.AdminAuthMiddleware, handlers.GetCategory)
	r.POST("admin/category", middlewares.AdminAuthMiddleware, handlers.AddCategory)
	r.PUT("admin/category", middlewares.AdminAuthMiddleware, handlers.UpdateCategory)
	r.DELETE("admin/category", middlewares.AdminAuthMiddleware, handlers.DeleteCategory)

	r.POST("/password/forgot",handlers.ForgotPassword)

	//cart
	r.POST("/cart",middlewares.UserAuthMiddleware,handlers.AddToCart)
	r.GET("/cart",middlewares.UserAuthMiddleware,handlers.ViewCart)
	r.DELETE("/cart",middlewares.UserAuthMiddleware,handlers.RemoveProductsFromCart)
	r.PUT("/cart",middlewares.UserAuthMiddleware,handlers.ProductQuantity)

	//Address
	r.GET("/address",middlewares.UserAuthMiddleware,handlers.ViewUserAddress)
	r.PUT("/address",middlewares.UserAuthMiddleware,handlers.EditUserAddress)
	r.DELETE("/address",middlewares.UserAuthMiddleware,handlers.RemoveUserAddress)
	r.POST("/address",middlewares.UserAuthMiddleware,handlers.AddNewAddressDetails)

	//wishlist
	r.GET("/wishlist",middlewares.UserAuthMiddleware,handlers.ViewUserWishlist)
	r.POST("/wishlist",middlewares.UserAuthMiddleware,handlers.AddProductToWishlist)
	r.DELETE("/wishlist",middlewares.UserAuthMiddleware,handlers.RemoveProductFromWishlist)

	//orders
	r.GET("/orders",middlewares.UserAuthMiddleware,handlers.ViewOrderDetails)



}
