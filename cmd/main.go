package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	initialisers "main.go/Initialisers"
	"main.go/routes"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.DBInitialise()

}

func main() {
	fmt.Println("running at locatioon:http://localhost:3000")

	r := gin.Default()
	r.LoadHTMLGlob("/home/jasim/CityViBe-Project-Ecommerce/template/*")
	routes.AdminRoutes(r)
	routes.UserRoutes(r)
	r.Run(":3000")
}
