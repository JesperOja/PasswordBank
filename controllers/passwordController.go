package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jesperoja/PasswordBank/PasswordBank/initializers"
	"github.com/jesperoja/PasswordBank/PasswordBank/models"
)

func AddPassword(ctx *gin.Context) {

	var body struct {
		Site     string
		Password string
	}

	ctx.Bind(&body)

	pass := models.PasswordUnit{Site: body.Site, Password: body.Password}

	result := initializers.DB.Create(&pass)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"password": pass,
	})
}

func GetPassword(ctx *gin.Context) {
	site := ctx.Query("site")
	var pass models.PasswordUnit

	initializers.DB.Where("site = ?", site).First(&pass)

	ctx.JSON(200, gin.H{
		"password": pass,
	})
}

func UpdatePassword(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Site     string
		Password string
	}

	ctx.Bind(&body)

	var pass models.PasswordUnit

	initializers.DB.First(&pass, id)

	initializers.DB.Model(&pass).Updates(models.PasswordUnit{
		Site:     body.Site,
		Password: body.Password,
	})

	ctx.JSON(200, gin.H{
		"password": pass,
	})
}
