package main

import (
	"transmedia/career/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Menggunakan controller login
	loginController := &controller.LoginController{}

	r.GET("/login", loginController.ShowLoginForm)
	r.POST("/login", loginController.HandleLogin)

	r.Run()
}
