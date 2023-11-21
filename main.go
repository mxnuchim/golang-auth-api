package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mxnuchim/golang-auth-api/controllers"
	"github.com/mxnuchim/golang-auth-api/initializers"
	"github.com/mxnuchim/golang-auth-api/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.MigrateDB()
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/signup", controllers.SignUp)
		api.POST("/login", controllers.Login)
		api.GET("/validate", middleware.RequireAuth,  controllers.Validate)
	}

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "Server's running",})
	})

	router.Run()
}