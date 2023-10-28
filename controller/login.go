package login

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"time"
	"transmediacareer/data"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func ShowCaptcha(c *gin.Context, captchaid string) {
	// fmt.Println(cap)
	_, file := path.Split(captchaid)
	// _, file := path.Split(c.Request.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	// fmt.Println("file : " + file)
	// fmt.Println("ext : " + ext)
	// fmt.Println("id : " + id)
	// fmt.Println("dir" + dir)
	if c.Request.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	var content bytes.Buffer
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
	captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	http.ServeContent(c.Writer, c.Request, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
}

func ShowLoginForm(c *gin.Context, transmedia data.JSONData) {
	// Menampilkan halaman form login
	captchaID := captcha.New()
	fmt.Printf("captchaID: %v\n", captchaID)
	tahunSaatIni := time.Now().Year()
	c.HTML(200, "login-form", gin.H{
		"captcha_id": captchaID,
		"transmedia": transmedia.TRANSMEDIA,
		"Tahun":      tahunSaatIni,
	})
}

func HandleLogin(c *gin.Context) {
	// Logika penanganan login
	// ...
	// Contoh: Verifikasi kredensial, buat sesi, dll.

	// Mengarahkan pengguna setelah login berhasil
	c.Redirect(302, "/dashboard")
}
