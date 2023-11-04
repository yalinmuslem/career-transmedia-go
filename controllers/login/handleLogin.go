package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"my-app/data"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type LoginStruct struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Captcha  string `validate:"required,numeric"`
}

func generateRandomToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func HandleLogin(e echo.Context, transmedia data.JSONData, data map[string]interface{}) {

	captchaError := captcha.VerifyString(e.Request().FormValue("captchaId"), e.Request().FormValue("captcha"))

	if !captchaError {
		ShowLoginForm(e, transmedia, data, nil, map[string]interface{}{"Info": "Captcha tidak sesuai", "Value": e.Request().FormValue("captcha")})
	}

	validate := validator.New()

	loginForm := LoginStruct{
		Email:    e.Request().FormValue("email"),
		Password: e.Request().FormValue("password"),
		Captcha:  e.Request().FormValue("captcha"),
	}

	err := validate.Struct(loginForm)
	if err != nil {
		// fmt.Println(err)
		ShowLoginForm(e, transmedia, data, err, nil)
	}

	randomToken := generateRandomToken()

	store := sessions.NewCookieStore([]byte("ee6779d4-decc-40a8-849b-ae01e02b15c5"))

	sessions, _ := store.Get(e.Request(), "SESSION_ID")
	sessions.Values["email"] = loginForm.Email
	sessions.Values["token"] = randomToken
	sessions.Save(e.Request(), e.Response().Writer)

	fmt.Println(randomToken)

	e.Redirect(http.StatusSeeOther, "/user/userboard")

}
