package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"

	"github.com/gin-gonic/gin"
)

func FillMapaRes(mapa models.MasterPayment) models.ResMasterPayment {
	return models.ResMasterPayment{
		ID:   mapa.ID,
		Tipe: mapa.Tipe,
		Name: mapa.Name,
	}
}

func ListMasterPayment(c *gin.Context) {
	// Definisikan data model master payment kedalam sebuah variabel
	datas := []models.MasterPayment{}

	// Panggil data model master payment sesuai dengan struct model
	settings.DB.Find(&datas)

	// Deklarasikan model struct response dan iterasi data sesuai dengan struct data response untuk ditampilkan
	list := []models.ResMasterPayment{}
	for _, datas := range datas {
		list = append(list, FillMapaRes(datas))
	}

	// Gunakan helper untuk mengirim response data list yang telah diinputkan sebelumnya
	helpers.DataResponse(c, list)
}

func CreateMasterPayment(c *gin.Context) {
	// Validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Deklarasikan request body kedalam masing masing variabel
	tipe := c.PostForm("tipe")
	name := c.PostForm("name")

	// Buat struktur untuk menyimpan data master payment kedalam database
	simpan := models.MasterPayment{
		Tipe: tipe,
		Name: name,
	}

	// Simpan data kedalam database
	settings.DB.Create(&simpan)

	// Buat variabel response
	response := FillMapaRes(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data master payment baru!", response)
}

func UpdateMasterPayment(c *gin.Context) {
	// Validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Masukan model master payment kedalam variabel untuk memudahkan filter data
	var modelMapa models.MasterPayment

	// Buat kondisi bila data id tidak ditemukan dalam database
	if settings.DB.First(&modelMapa, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data master payment tidak ditemukan!")
		c.Abort()
		return
	}

	// Update data master payment dengan data request body baru
	settings.DB.Model(&modelMapa).Where("id= ?", id).Updates(models.MasterPayment{
		Tipe: c.PostForm("tipe"),
		Name: c.PostForm("name"),
	})

	// Masukan nilai input update kedalam variabel response untuk ditampilkan
	response := FillMapaRes(modelMapa)

	// Panggil helper untuk menampilkan data response
	helpers.SuksesWithDataResponse(c, "Sukses update data master payment!", response)
}

func HapusMasterPayment(c *gin.Context) {
	// Definisikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Masukan fungsi model master payment kedalam sebuah variabel
	var modelMapa models.MasterPayment

	// Buat kondisi bila data id master payment tidak ditemukan dalam database
	if settings.DB.First(&modelMapa, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id master payment tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data master payment dari id param yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&modelMapa)

	// Panggil response helper untuk memunculkan response sukses
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data master payment")
}
