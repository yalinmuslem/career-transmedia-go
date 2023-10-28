package login

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"transmediacareer/data"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func ShowCaptcha(c *gin.Context, cap string) {
	// fmt.Println(cap)
	http.Handle("/assets/captcha", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}

func ShowLoginForm(c *gin.Context, transmedia data.JSONData) {

	captchaDir := "assets/captcha"
	if err := os.MkdirAll(captchaDir, os.ModePerm); err != nil {
		panic(err)
	}

	captcha.SetCustomStore(&CustomStore{
		CaptchaDir: captchaDir,
	})

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

// CustomStore adalah penyimpanan CAPTCHA kustom
type CustomStore struct {
	CaptchaDir string
}

// Set mengatur CAPTCHA dengan ID ke dalam penyimpanan
func (s *CustomStore) Set(id string, digits []byte) {
	// Simpan gambar CAPTCHA ke dalam direktori CaptchaDir
	filePath := s.CaptchaDir + "/" + id + ".png"
	os.WriteFile(filePath, digits, 0644)
	// captcha.WriteImage()
}

// Get mengambil gambar CAPTCHA dengan ID dari penyimpanan
func (s *CustomStore) Get(id string, clear bool) (digits []byte) {
	filePath := s.CaptchaDir + "/" + id + ".png"
	digits, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if clear {
		_ = os.Remove(filePath)
	}
	return digits
}

// Clear menghapus CAPTCHA dengan ID dari penyimpanan
func (s *CustomStore) Clear(id string) {
	// Hapus gambar CAPTCHA dari direktori CaptchaDir
	filePath := s.CaptchaDir + "/" + id + ".png"
	_ = os.Remove(filePath)
}
