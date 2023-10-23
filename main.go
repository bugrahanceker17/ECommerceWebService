package main

import (
	"github.com/bugrahanceker17/ECommerceWebService/controllers"
	"github.com/bugrahanceker17/ECommerceWebService/database"
	"github.com/bugrahanceker17/ECommerceWebService/middleware"
	"github.com/bugrahanceker17/ECommerceWebService/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := controllers.NewApp(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/add-to-cart", app.AddToCart())
	router.GET("remove-item", app.RemoveItem())
	router.GET("cart-checkout", app.BuyFromCart())
	router.GET("instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
