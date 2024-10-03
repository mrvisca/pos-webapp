package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisData struct {
	User   models.User
	Bisnis models.Business
	Gudang models.Warehouse
}

func RegisterAcc(c *gin.Context) {
	// Masukan request body kedalam variabel
	owner := c.PostForm("owner")
	bisnis := c.PostForm("bisnis")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	tipe := c.PostForm("tipe")

	// Enkripsi password untuk dimasukan dalam database
	inp_pass, _ := helpers.HashPassword(password)

	// Definisikan model user kedalam variabel data
	var usr models.User

	// Definisikan model bisnis kedalam variabel data
	var bis models.Business

	// Cek data untuk menghindari duplikasi data owner
	if !settings.DB.First(&usr, "name = ?", owner).RecordNotFound() {
		helpers.ElorResponse(c, "Gagal melakukan pendaftaran akun, nama owner sudah terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Cek data untuk menghindari duplikasi data bisnis
	if !settings.DB.First(&bis, "name = ?", bisnis).RecordNotFound() {
		helpers.ElorResponse(c, "Gagal melakukan pendaftaran akun, nama bisnis sudah terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Set waktu untuk tanggal join dan end date paket langganan
	sekarang := time.Now()
	join := sekarang.Format("2006-01-02 15:04:05")
	end_date := sekarang.AddDate(0, 0, 14) // 0 tahun, 0 bulan, 10 hari
	maktif := end_date.Format("2006-01-02 15:04:05")

	// Buat struktur untuk menyimpan data request kedalam database
	simbis := models.Business{
		Name:        bisnis,
		Branchlimit: 1,
		Tipe:        tipe,
	}
	settings.DB.Create(&bisnis)

	// Agar tidak elor
	_ = simbis.Name
	_ = simbis.Branchlimit
	_ = simbis.Tipe

	gudang := models.Warehouse{
		BusinessId:     simbis.ID,
		SubscriptionId: 1,
		Name:           "Cabang Utama",
		Address:        "Jakarta, Indonesia",
		Phone:          phone,
		Join:           join,
		EndDate:        maktif,
		IsDefault:      true,
	}
	settings.DB.Create(&gudang)

	user := models.User{
		RoleId:      2,
		BusinessId:  simbis.ID,
		WarehouseId: gudang.ID,
		Name:        owner,
		Email:       email,
		Password:    inp_pass,
		Phone:       phone,
		IsVerified:  false,
	}
	settings.DB.Create(&user)

	// Masukan hasil inputan simbis, gudang dan user dalam 1 variabel untuk ditampilkan pada response pendaftaran
	response := RegisData{
		User:   user,
		Bisnis: simbis,
		Gudang: gudang,
	}

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil melakukan pendaftaran akun baru!", response)
}
