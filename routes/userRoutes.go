package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func UserRoutes(r *gin.Engine) {

	r.POST("/signup", handlers.UserSignup)
	r.POST("/verify", handlers.VerifySignupOtp)
	r.POST("/login", handlers.UserLogin)
	r.POST("/login/verify", handlers.VerifyLoginOtp)
	r.POST("/logout", handlers.UserLogout)

	// r.POST("/address",middlewares.UserAuthMiddleware,handlers.UserAddress)

}