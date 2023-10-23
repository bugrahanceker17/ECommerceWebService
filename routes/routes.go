package routes

import (
	"github.com/bugrahanceker17/ECommerceWebService/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoute *gin.Engine) {
	incomingRoute.POST("/users/signup", controllers.SignUp())
	incomingRoute.POST("/users/login", controllers.Login())
	incomingRoute.POST("/admin/add-product", controllers.ProductViewerAdmin())
	incomingRoute.GET("/users/product-view", controllers.SearchProduct())
	incomingRoute.GET("/users/search", controllers.SearchProductByQuery())
}
