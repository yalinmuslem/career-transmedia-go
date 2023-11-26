package registrasi

import (
	"fmt"
	"my-app/data"
	"net/http"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func translateAll(trans ut.Translator) (error map[string]string) {

	type RegistrasiStruct struct {
		FullName       string `validate:"required"`
		Email          string `validate:"required,email"`
		Password       string `validate:"required,min=8"`
		RetypePassword string `validate:"required,eqfield=Password"`
		DateOfBirth    string `validate:"required,datetime=2006-01-02"`
		Captcha        string `validate:"required,numeric"`
	}

	validasiRegistrasi := RegistrasiStruct{
		FullName:       "FullName",
		Email:          "Email",
		Password:       "Password",
		RetypePassword: "Retype Password",
		DateOfBirth:    "Date Of Birth",
		Captcha:        "Captcha",
	}

	err := validate.Struct(validasiRegistrasi)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		error = make(map[string]string)
		for _, err := range errs {
			translateErr := fmt.Errorf(err.Translate(trans))
			error[err.Field()] = translateErr.Error()
		}
		return error
	}

	return nil
}

func HandleRegistrasi(e echo.Context, transmedia data.JSONData, data map[string]interface{}) {
	var email = e.Request().FormValue("email")
	var password = e.Request().FormValue("password")
	var retypepassword = e.Request().FormValue("retypepassword")
	var dateofbirth = e.Request().FormValue("dob")
	var captcha = e.Request().FormValue("captcha")

	id := id.New()
	uni = ut.New(id, id)

	trans, _ := uni.GetTranslator("id")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	t := translateAll(trans)
	if t != nil {
		ShowRegistrasiPage(e, transmedia, data, t)
		return
	}
	// var tmpl = template.Must(template.ParseFiles(
	// 	"templates/user/registration/view-registrasi.html",
	// 	"templates/user/registration/content.html",
	// 	"templates/user/others/menu.html",
	// ))

	fmt.Println("Registrasi berhasil")
	e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Registrasi berhasil",
		"data": map[string]interface{}{
			"email":    email,
			"password": password,
			"retype":   retypepassword,
			"dob":      dateofbirth,
			"captcha":  captcha,
		},
	})
}
