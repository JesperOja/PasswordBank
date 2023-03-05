package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesperoja/PasswordBank/PasswordBank/controllers"
	"github.com/jesperoja/PasswordBank/PasswordBank/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/add", controllers.AddPassword)
	r.GET("/get", controllers.GetPassword)
	r.PUT("/update/:id", controllers.UpdatePassword)

	r.Run()
}
