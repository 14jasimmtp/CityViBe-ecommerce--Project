package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	initialisers "main.go/Initialisers"
	_ "main.go/docs"
	"main.go/routes"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.DBInitialise()

}


func main() {
	fmt.Println("running at locatioon:http://localhost:3000")

	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLGlob("/home/jasim/CityViBe-Project-Ecommerce/template/*")
	routes.AdminRoutes(r)
	routes.UserRoutes(r)
	r.Run(":3000")

}