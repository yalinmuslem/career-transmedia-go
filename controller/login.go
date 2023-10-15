package login

import (
	"time"
	"transmediacareer/data"

	"github.com/gin-gonic/gin"
)

func ShowLoginForm(c *gin.Context, transmedia data.JSONData) {

	// Menampilkan halaman form login

	tahunSaatIni := time.Now().Year()
	c.HTML(200, "login-form.html", gin.H{
		"DeskripsiPerusahaan": transmedia.TRANSMEDIA.DeskripsiPerusahaan,
		"AlamatPerusahaan":    transmedia.TRANSMEDIA.AlamatPerusahaan,
		"Tahun":               tahunSaatIni,
		"Phone":               transmedia.TRANSMEDIA.KONTAK.PHONE,
		"Fax":                 transmedia.TRANSMEDIA.KONTAK.FAX,
		"Twitter":             transmedia.TRANSMEDIA.SosialMedia.Twitter,
		"Facebook":            transmedia.TRANSMEDIA.SosialMedia.Facebook,
		"Youtube":             transmedia.TRANSMEDIA.SosialMedia.Youtube,
		"Instagram":           transmedia.TRANSMEDIA.SosialMedia.Instagram,
	})
}

func HandleLogin(c *gin.Context) {
	// Logika penanganan login
	// ...
	// Contoh: Verifikasi kredensial, buat sesi, dll.

	// Mengarahkan pengguna setelah login berhasil
	c.Redirect(302, "/dashboard")
}
