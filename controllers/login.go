package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"

	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) {
	// Ambil body request form login dan deklarasikan pada masing-masing variabel
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Definisikan model role kedalam variabel
	var user models.User

	// Kondisi bila email user tidak ditemukan
	if settings.DB.First(&user, "email = ?", email).RecordNotFound() {
		helpers.ElorResponse(c, "Email tidak terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Kondisi bila akun belum terverifikasi
	if !user.IsVerified {
		helpers.ElorResponse(c, "Login gagal, akun belum terverifikasi!")
		c.Abort()
		return
	}

	// Kondisi pencocokan data password input form user dengan data password dalam database
	if !helpers.CheckPassword(user.Password, password) {
		helpers.ElorResponse(c, "Password tidak valid!")
		c.Abort()
		return
	}

	// Buat token bila data sudah sesuai dengan data dalam database
	token := helpers.CreateToken(&user)

	// Tambahkan response sukses untuk menampilkan response dan juga token autentikasi
	helpers.SuksesLogin(c, "Login aplikasi berhasil dilakukan!", token, int64(user.RoleId))
}
