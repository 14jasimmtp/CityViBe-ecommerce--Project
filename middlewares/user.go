package middlewares

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func UserAuthMiddleware(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.Abort()
	}
	role, err := utils.GetRoleFromToken(Token)
	if err != nil {
		c.Abort()
	}
	if role == "user" {
		c.Next()

	}
}

// func UserAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := helper.GetTokenFromHeader(authHeader)
// 		// Validate the token and extract the user ID
// 		if tokenString == "" {
// 			var err error
// 			tokenString, err = c.Cookie("Authorization")
// 			if err != nil {
// 				c.AbortWithStatus(http.StatusUnauthorized)
// 				return
// 			}
// 		}
// 		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
// 		if err != nil {
// 			fmt.Println("error is 👺 ", err)
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		c.Set("user_id", userID)
// 		c.Set("user_email", userEmail)
// 		c.Next()
// 	}
// }
//
