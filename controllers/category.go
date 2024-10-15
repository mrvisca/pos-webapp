package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"

	"github.com/gin-gonic/gin"
)

func FillCategoryData(dcat models.Category) models.DataCategory {
	return models.DataCategory{
		ID:          dcat.ID,
		BusinessId:  dcat.BusinessId,
		WarehouseId: dcat.WarehouseId,
		Name:        dcat.Name,
		Desc:        dcat.Desc,
		Item:        dcat.Item,
	}
}

func ListCategory(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan data model user kedalam variabel
	datas := []models.Category{}

	// Panggil data model list kategori produk sesuai dengan id bisnis dan id cabang yang terdaftar
	settings.DB.Where("business_id = ? AND warehouse_id = ?", businessid, warehouseid).Find(&datas)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.DataCategory{}
	for _, datas := range datas {
		list = append(list, FillCategoryData(datas))
	}

	// Gunakan helper untuk mengirim response data list yang telah diinputkan sebelumnya
	helpers.DataResponse(c, list)
}

func TambahKategori(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Masukan data request body kedalam masing-masing variabel untuk memudahkan proses pembuatan data
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	// Definisikan model kategori produk model kedalam sebuah variabel
	var kategori models.Category

	// Kondisi untuk mengatasi duplikasi data
	if !settings.DB.First(&kategori, "business_id = ? AND warehouse_id = ? AND name LIKE ?", businessid, warehouseid, "%"+name+"%").RecordNotFound() {
		helpers.ElorResponse(c, "Data kategori produk sudah tersedia dalam database!")
		c.Abort()
		return
	}

	// Masukan variabel data sesuai dengan struct kategori produk
	simpan := models.Category{
		BusinessId:  businessid,
		WarehouseId: warehouseid,
		Name:        name,
		Desc:        desc,
		Item:        0,
	}

	// Simpan data kedalam database tabel kategori produk
	settings.DB.Create(&simpan)

	// Masukan struct response yang rapi kedalam variabel untuk ditampilkan
	response := FillCategoryData(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data kategori produk baru!", response)
}

func UpdateCategory(c *gin.Context) {
	// Definisikan jwt_business_id dan jwt_warehouse_id dari token kedalam variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan request param id kedalam variabel id
	id := c.Param("id")

	// Definisikan request body ke dalam masing-masing variabel
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	// Definisikan model category kedalam variabel kategori
	var kategori models.Category

	// Buat kondisi bila data id kategori tidak ditemukan
	if settings.DB.First(&kategori, "id = ? AND business_id = ? AND warehouse_id = ?", id, businessid, warehouseid).RecordNotFound() {
		helpers.ElorResponse(c, "Data id kategori produk tidak ditemukan!")
		c.Abort()
		return
	}

	// Update data id kategori produk jika data ditemukan
	settings.DB.Model(&kategori).Where("id = ?", id).Updates(models.Category{
		Name: name,
		Desc: desc,
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
	response := FillCategoryData(kategori)

	// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
	helpers.SuksesWithDataResponse(c, "Berhasil melakukan update data kategori produk!", response)
}

func HapusKategori(c *gin.Context) {
	// Definisikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model kategori produk sebagai variabel kategori
	var kategori models.Category

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&kategori, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id kategori produk tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data kategori produk dari id yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&kategori)

	// Panggil response sukses helper bila data berhasil terhapus
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data kategori produk!")
}
