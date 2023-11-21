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
	r.GET("/")

	//password change
	r.POST("/password/forgot", handlers.ForgotPassword)
	r.POST("password/forgot/change", handlers.ResetForgottenPassword)
	r.POST("/")

	//Address
	r.GET("/address", middlewares.UserAuthMiddleware, handlers.ViewUserAddress)
	r.POST("/address", middlewares.UserAuthMiddleware, handlers.AddNewAddressDetails)
	r.PUT("/address", middlewares.UserAuthMiddleware, handlers.EditUserAddress)
	r.DELETE("/address", middlewares.UserAuthMiddleware, handlers.RemoveUserAddress)

	//profile
	// r.GET("/profile",middlewares.UserAuthMiddleware,handlers.ViewUserProfile)
	// r.PUT("/profile",middlewares.UserAuthMiddleware,handlers.UpdateUserprofile)

}
