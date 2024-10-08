package controllers

import (
	"net/http"
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"

	"github.com/gin-gonic/gin"
)

func Verifikasi(c *gin.Context) {
	// Deklarasikan query param kode kedalam variabel kode
	kode := c.Param("kode")

	// Buat variabel models role
	var usr models.User

	// Kondisi bila data kode tidak ditemukan
	if settings.DB.First(&usr, "kode = ?", kode).RecordNotFound() {
		helpers.ElorResponse(c, "Verifikasi akun gagal, data tidak ditemukan!")
		c.Abort()
		return
	}

	if !usr.IsVerified {
		helpers.ElorResponse(c, "Verifikasi akun gagal, akun telah terverifikasi")
		c.Abort()
		return
	}

	settings.DB.Model(&usr).Where("kode = ?", kode).Update("is_verified", true)

	c.Redirect(http.StatusFound, "http://localhost:8080/")
}
