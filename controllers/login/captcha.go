package login

import (
	"bytes"
	"net/http"
	"path"
	"time"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
)

func ShowCaptcha(e echo.Context, captchaid string) {
	_, file := path.Split(captchaid)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if e.Request().FormValue("reload") != "" {
		captcha.Reload(id)
	}
	var content bytes.Buffer
	e.Response().Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	e.Response().Writer.Header().Set("Pragma", "no-cache")
	e.Response().Writer.Header().Set("Expires", "0")
	e.Response().Writer.Header().Set("Content-Type", "image/png")
	captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	http.ServeContent(e.Response().Writer, e.Request(), id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
}

func ReloadCaptcha(e echo.Context) {
	captchaID := captcha.New()
	e.JSON(http.StatusOK, map[string]any{
		"status":    true,
		"captchaID": captchaID,
	})
}
