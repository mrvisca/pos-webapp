package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Palares struct {
	ID         uint
	IsTax      bool
	Taxval     int64
	IsService  bool
	Serviceval int64
}

func GetDataPala(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan model pajaklayanan kedalam variabel pala
	var pala models.PajakLayanan

	// Panggil data master pajak layanan
	settings.DB.First(&pala, "business_id = ? AND warehouse_id = ?", businessid, warehouseid)

	// Sesuaikan response data dalam variabel data untuk ditampilkan kedalam helper response
	dataPala := Palares{
		ID:         pala.ID,
		IsTax:      pala.IsTax,
		Taxval:     pala.Taxval,
		IsService:  pala.IsService,
		Serviceval: pala.Serviceval,
	}

	// Kirim data ke helper response untuk mengembalikan response data
	helpers.SuksesWithDataResponse(c, "Berhasil memanggil data master pajak layanan!", dataPala)
}

func UpdatePala(c *gin.Context) {
	// Klasifikasi param id kedalam variabel id
	id := c.Param("id")

	// Klasifikasi body request put kedalam masing-masing variabel
	istax := c.PostForm("is_tax")
	taxval := c.PostForm("taxval")
	isservice := c.PostForm("is_service")
	serviceval := c.PostForm("serviceval")

	// Konversi taxval & serviceval string menjadi integer
	nilaitax, _ := strconv.Atoi(taxval)
	nilaiservice, _ := strconv.Atoi(serviceval)

	// Konversi request istax dan isservice string menjadi boolean
	booltax, _ := strconv.ParseBool(istax)
	boolservice, _ := strconv.ParseBool(isservice)

	// Definisikan model pajak layanan kedalam sebuah variabel
	var pala models.PajakLayanan

	// Buat kondisi bila data master pajak layanan id tidak ditemukan
	if settings.DB.First(&pala, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data master pajak layanan tidak ditemukan!")
		c.Abort()
		return
	}

	// Update data master pajak layanan
	settings.DB.Model(&pala).Where("id = ?", id).Updates(map[string]interface{}{
		"is_tax":     booltax,
		"taxval":     int64(nilaitax),
		"is_service": boolservice,
		"serviceval": int64(nilaiservice),
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
	response := Palares{
		ID:         pala.ID,
		IsTax:      booltax,
		Taxval:     int64(nilaitax),
		IsService:  boolservice,
		Serviceval: int64(nilaiservice),
	}

	// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
	helpers.SuksesWithDataResponse(c, "Berhasil melakukan update data master pajak!", response)
}
