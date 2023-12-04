package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
	"main.go/middlewares"
)

func UserRoutes(r *gin.Engine) {
	//login
	r.POST("/signup", handlers.UserSignup)
	r.POST("/login", handlers.UserLogin)

	r.POST("/verify", handlers.VerifyLoginOtp)
	r.POST("/logout", handlers.UserLogout)

	//products
	r.GET("/products", handlers.GetAllProducts)
	r.GET("/products/:id", handlers.ShowSingleProduct)
	r.GET("/products/search",handlers.SearchProducts)//search

	//filtering
	r.GET("/products/filter/",handlers.FilterProducts)

	//wishlist
	r.POST("/products/wishlist",middlewares.UserAuthMiddleware,handlers.AddProductToWishlist)
	r.GET("/products/wishlist",middlewares.UserAuthMiddleware,handlers.ViewUserWishlist)
	r.DELETE("/products/wishlist",middlewares.UserAuthMiddleware,handlers.RemoveProductFromWishlist)

	//profile
	r.GET("/profile",middlewares.UserAuthMiddleware,handlers.UserProfile)
	r.PUT("/profile",middlewares.UserAuthMiddleware,handlers.UpdateUserProfile)

	//password change
	r.POST("/password/forgot", handlers.ForgotPassword)
	r.POST("password/forgot/change", handlers.ResetForgottenPassword)

	//Address
	r.GET("/address", middlewares.UserAuthMiddleware, handlers.ViewUserAddress)
	r.POST("/address", middlewares.UserAuthMiddleware, handlers.AddNewAddressDetails)
	r.PUT("/address", middlewares.UserAuthMiddleware, handlers.EditUserAddress)
	r.DELETE("/address", middlewares.UserAuthMiddleware, handlers.RemoveUserAddress)

	//cart
	r.GET("/cart", middlewares.UserAuthMiddleware, handlers.ViewCart)
	r.POST("/cart", middlewares.UserAuthMiddleware, handlers.AddToCart)
	r.DELETE("/cart", middlewares.UserAuthMiddleware, handlers.RemoveProductsFromCart)
	r.PUT("/cart/add-quantity", middlewares.UserAuthMiddleware, handlers.IncreaseQuantityUpdate)
	r.PUT("/cart/reduce-quantity", middlewares.UserAuthMiddleware, handlers.DecreaseQuantityUpdate)

	//orders
	r.GET("/orders", middlewares.UserAuthMiddleware, handlers.ViewOrders)
	r.POST("/orders", middlewares.UserAuthMiddleware, handlers.OrderFromCart)
	r.GET("/checkout", middlewares.UserAuthMiddleware, handlers.ViewCheckOut)
	r.PUT("/orders/return", middlewares.UserAuthMiddleware, handlers.ReturnOrder)
	r.PUT("/orders/cancel",middlewares.UserAuthMiddleware,handlers.CancelOrder)

	//wishlist
	r.GET("/wishlist", middlewares.UserAuthMiddleware, handlers.ViewUserWishlist)
	r.POST("/wishlist", middlewares.UserAuthMiddleware, handlers.AddProductToWishlist)
	r.DELETE("/wishlist", middlewares.UserAuthMiddleware, handlers.RemoveProductFromWishlist)

	//payment
	r.GET("/payment/razorpay",handlers.ExecuteRazorPayPayment)	
	r.POST("/payment/verify",handlers.VerifyPayment)
}
