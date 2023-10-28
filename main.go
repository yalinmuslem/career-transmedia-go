package main

import (
	"encoding/json"
	"os"

	"path/filepath"

	login "transmediacareer/controller"
	"transmediacareer/data"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Mendefinisikan path ke file JSON dalam folder config
	configFilePath := "data/config.json"
	absPath, _ := filepath.Abs(configFilePath)

	// Baca data JSON dari file
	jsonBytes, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	// Unmarshal data JSON ke dalam struktur data yang sesuai
	var jsonData data.JSONData
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		panic(err)
	}

	r.HTMLRender = loadTemplates("./templates")
	r.Static("/assets", "assets")

	transmedia := jsonData

	r.GET("/login", func(ctx *gin.Context) {

		login.ShowLoginForm(ctx, transmedia)
	})
	r.POST("/login", func(ctx *gin.Context) {
		// login.HandleLogin(ctx)
	})
	r.GET("/captcha/:captcha", func(ctx *gin.Context) {
		captchaID := ctx.Param("captcha")
		login.ShowCaptcha(ctx, captchaID)
	})
	r.Run()
}

func loadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles("login-form",
		templatesDir+"/user/login/login-form.html",
		templatesDir+"/user/login/content.html",
		templatesDir+"/user/menu.html")
	return r
}
