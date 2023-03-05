package main

import (
	"github.com/jesperoja/PasswordBank/PasswordBank/initializers"
	"github.com/jesperoja/PasswordBank/PasswordBank/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.PasswordUnit{})
}
