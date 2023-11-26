package registrasi

import (
	"fmt"
	"html/template"
	"my-app/data"
	"time"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
)

func ShowRegistrasiPage(e echo.Context, transmedia data.JSONData, data map[string]interface{}, validasi any) error {
	var captchaID = captcha.New()
	tahunSaatIni := time.Now().Year()
	var tmpl = template.Must(template.ParseFiles(
		"templates/user/registration/view-registrasi.html",
		"templates/user/registration/content.html",
		"templates/user/others/menu.html",
	))

	if validasi != nil {
		for key, e := range validasi.(map[string]string) {
			switch key {
			case "FullName":
				data["fullnameError"] = map[string]interface{}{"Info": e}
			case "Email":
				data["emailError"] = map[string]interface{}{"Info": e}
			case "Password":
				data["passwordError"] = map[string]interface{}{"Info": e}
			case "RetypePassword":
				data["retypeError"] = map[string]interface{}{"Info": e}
			case "DateOfBirth":
				data["dobError"] = map[string]interface{}{"Info": e}
			case "Captcha":
				data["captchaError"] = map[string]interface{}{"Info": e}
			}
		}
	}

	fmt.Println(data)

	tmpl.ExecuteTemplate(e.Response().Writer, "view-registrasi.html", map[string]interface{}{
		"title":      "Registrasi",
		"transmedia": transmedia,
		"Tahun":      tahunSaatIni,
		"captchaID":  captchaID,
		"data":       data,
	})

	fmt.Println(tahunSaatIni)
	return nil
}
