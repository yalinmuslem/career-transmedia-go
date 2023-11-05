package login

import (
	"html/template"
	"my-app/data"
	"time"

	"github.com/dchest/captcha"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ForgotPassword(c echo.Context, transmedia data.JSONData, data map[string]interface{}, errorValidation any, othersError any) error {

	var tmpl = template.Must(template.ParseFiles(
		"templates/user/others/templates.html",
		"templates/user/forgot-password/content.html",
		"templates/user/forgot-password/js.html",
		"templates/user/others/menu.html",
	))

	tahunSaatIni := time.Now().Year()
	captchaID := captcha.New()

	if othersError != nil || errorValidation != nil {
		if errorValidation != nil {
			for _, err := range errorValidation.(validator.ValidationErrors) {
				switch err.Field() {
				case "Email":
					data["emailError"] = map[string]interface{}{"Info": err.Field() + " tidak boleh kosong atau format salah", "Value": err.Value()}
				case "Captcha":
					data["captchaError"] = map[string]interface{}{"Info": err.Field() + " tidak boleh kosong atau format salah", "Value": err.Value()}
				}
			}
		}

		if othersError != nil {
			for key := range othersError.(map[string]interface{}) {
				switch key {
				case "birthdate":
					data["birthdateError"] = othersError.(map[string]interface{})["birthdate"]
				case "captcha":
					data["captchaError"] = othersError.(map[string]interface{})["captcha"]
				case "result":
					data["resultError"] = othersError.(map[string]interface{})["result"]
				}
			}
		}

		return tmpl.ExecuteTemplate(c.Response().Writer, "templates.html", map[string]interface{}{
			"title":      "Login",
			"Tahun":      tahunSaatIni,
			"transmedia": transmedia,
			"captchaID":  captchaID,
			"data":       data,
		})
	}

	return tmpl.ExecuteTemplate(c.Response().Writer, "templates.html", map[string]interface{}{
		"title":      "Login",
		"Tahun":      tahunSaatIni,
		"transmedia": transmedia,
		"captchaID":  captchaID,
		"data":       data,
	})
}
