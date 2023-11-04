package login

import (
	"my-app/data"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ForgotPasswordStruct struct {
	Email     string `validate:"required,email"`
	Captcha   string `validate:"required,numeric"`
	Birthdate string `validate:"required"`
}

func validateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	return err
}
func HandleForgotPassword(c echo.Context, transmedia data.JSONData, data map[string]interface{}) error {

	captchaID := captcha.VerifyString(c.Request().FormValue("captchaId"), c.Request().FormValue("captcha"))

	if !captchaID {
		ForgotPassword(c, transmedia, data, nil, map[string]interface{}{"captcha": map[string]interface{}{"Info": "Captcha tidak sesuai", "Value": c.Request().FormValue("captcha")}})
	}

	birthdate := c.Request().FormValue("birthdate")
	errBirthdate := validateDate(birthdate)
	if errBirthdate != nil {
		ForgotPassword(c, transmedia, data, nil, map[string]interface{}{"birthdate": map[string]interface{}{"Info": "Format tanggal salah", "Value": c.Request().FormValue("birthdate")}})
	}

	validate := validator.New()

	forgotPasswordForm := ForgotPasswordStruct{
		Email:     c.Request().FormValue("email"),
		Captcha:   c.Request().FormValue("captcha"),
		Birthdate: c.Request().FormValue("birthdate"),
	}

	err := validate.Struct(forgotPasswordForm)
	if err != nil {
		// fmt.Println(err)
		ForgotPassword(c, transmedia, data, err, nil)
	}

	return nil
}
