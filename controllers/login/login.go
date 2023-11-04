package login

import (
	"html/template"
	"my-app/data"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ShowLoginForm(e echo.Context, transmedia data.JSONData, data map[string]interface{}, errorValidation any, errorCaptcha any) error {
	captchaID := captcha.New()
	tahunSaatIni := time.Now().Year()
	var tmpl = template.Must(template.ParseFiles(
		"templates/user/others/templates.html",
		"templates/user/login/content.html",
		"templates/user/login/js.html",
		"templates/user/others/menu.html",
	))

	if errorCaptcha != nil || errorValidation != nil {

		if errorValidation != nil {
			for _, err := range errorValidation.(validator.ValidationErrors) {
				switch err.Field() {
				case "Email":
					data["emailError"] = map[string]interface{}{"Info": err.Field() + " tidak boleh kosong atau format salah", "Value": err.Value()}
				case "Password":
					data["passwordError"] = map[string]interface{}{"Info": err.Field() + " tidak boleh kosong", "Value": err.Value()}
				case "Captcha":
					data["captchaError"] = map[string]interface{}{"Info": err.Field() + " tidak boleh kosong atau format salah", "Value": err.Value()}
				}
			}
		}

		if errorCaptcha != nil {
			data["captchaError"] = errorCaptcha
		}

		return tmpl.ExecuteTemplate(e.Response().Writer, "templates.html", map[string]interface{}{
			"title":      "Login",
			"Tahun":      tahunSaatIni,
			"transmedia": transmedia,
			"captchaID":  captchaID,
			"data":       data,
		})

	}

	return tmpl.ExecuteTemplate(e.Response().Writer, "templates.html", map[string]interface{}{
		"title":      "Login",
		"Tahun":      tahunSaatIni,
		"transmedia": transmedia,
		"captchaID":  captchaID,
		"data":       data,
	})
}
