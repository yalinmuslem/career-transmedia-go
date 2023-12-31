package main

import (
	"encoding/json"
	login "my-app/controllers/login"
	"my-app/controllers/registrasi"
	userboard "my-app/controllers/userboard"
	"my-app/data"
	forgotpassword "my-app/data/forgot-password"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
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

	// CSRF
	const CSRFTokenName = "csrf_token_career"
	const CSRFTokenMaxAge = 86400
	const CSRFTokenPath = "/"
	const CSRFTokenDomain = ""
	const CSRFTokenSecure = true
	const CSRFTokenHTTPOnly = true
	const CSRFTokenSameSite = http.SameSiteStrictMode
	const CSRFTokenKey = "aa6328f9-758d-4188-a80f-b842eb650f65"
	const CSRFHeaderName = "csrf_token_career"

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:" + CSRFHeaderName,
		ContextKey:     CSRFTokenKey,
		CookieName:     CSRFTokenName,
		CookieMaxAge:   CSRFTokenMaxAge,
		CookiePath:     CSRFTokenPath,
		CookieDomain:   CSRFTokenDomain,
		CookieSecure:   CSRFTokenSecure,
		CookieHTTPOnly: CSRFTokenHTTPOnly,
		CookieSameSite: CSRFTokenSameSite,
	}))

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
		}
		registrasi.ShowRegistrasiPage(c, jsonData, data, nil)
		return nil
	})

	e.POST("do-registrasi", func(c echo.Context) error {
		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
		}
		registrasi.HandleRegistrasi(c, jsonData, data)
		return nil
	})

	e.GET("/login", func(c echo.Context) error {

		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
			// "corps": map[string]interface{}{"corporate": corporate.SelectData(c)},
		}
		// fmt.Println(data)
		login.ShowLoginForm(c, jsonData, data, nil, nil)
		return nil
	})

	e.GET("/captcha/:captchaID", func(c echo.Context) error {
		captchaID := c.Param("captchaID")
		login.ShowCaptcha(c, captchaID)
		return nil
	})

	e.POST("/recaptcha", func(c echo.Context) error {
		login.ReloadCaptcha(c)
		return nil
	})

	e.POST("/login", func(c echo.Context) error {
		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
		}

		login.HandleLogin(c, jsonData, data)
		return nil
	})

	e.GET("forgot-password", func(c echo.Context) error {
		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
		}
		login.ForgotPassword(c, jsonData, data, nil, nil)
		return nil
	})

	e.GET("test-forgot-password", func(c echo.Context) error {

		c.Request().Method = http.MethodPost
		c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Request().PostForm = map[string][]string{"email": {"yalin.muslem@gmail.com"}, "perusahaan": {"TRANSTV"}}

		res := map[string]interface{}{"result": forgotpassword.SelectData(c)}
		return c.JSON(200, res)
	})

	e.POST("/forgot-password", func(c echo.Context) error {
		data := map[string]interface{}{
			"csrf": map[string]interface{}{
				"token": c.Get(CSRFTokenKey).(string),
				"name":  CSRFTokenName,
			},
		}
		login.HandleForgotPassword(c, jsonData, data)
		return nil
	})

	u := e.Group("/user")
	u.Use(session.MiddlewareWithConfig(session.Config{
		Store: sessions.NewCookieStore([]byte("ee6779d4-decc-40a8-849b-ae01e02b15c5")),
	}))

	u.GET("/userboard", func(c echo.Context) error {
		checkSession(c)

		userboard.UserBoard(c)
		return nil
	})

	e.Logger.Fatal(e.Start(":1322"))
}

func checkSession(c echo.Context) {
	session, _ := session.Get("SESSION_ID", c)
	if session.Values["email"] == nil {
		c.Redirect(http.StatusSeeOther, "/")
	}
}
