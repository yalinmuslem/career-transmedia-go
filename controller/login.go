package controller

import "github.com/gin-gonic/gin"

type LoginController struct{}

func (ctrl *LoginController) ShowLoginForm(c *gin.Context) {
	// Menampilkan halaman form login
	c.HTML(200, "login-form.html", gin.H{})
}

func (ctrl *LoginController) HandleLogin(c *gin.Context) {
	// Logika penanganan login
	// ...
	// Contoh: Verifikasi kredensial, buat sesi, dll.

	// Mengarahkan pengguna setelah login berhasil
	c.Redirect(302, "/dashboard")
}
