package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FillResSub(dsub models.Subscription) models.ResSub {
	return models.ResSub{
		ID:    dsub.ID,
		Name:  dsub.Name,
		Days:  dsub.Days,
		Price: dsub.Price,
		Tipe:  dsub.Tipe,
		Desc:  dsub.Desc,
	}
}

func CreateDevSubscription(c *gin.Context) {
	// Validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Definisikan model role kedalam variabel data
	var dasub models.Subscription

	// Deklarasi request body kedalam variabel
	name := c.PostForm("name")
	days := c.PostForm("days")
	price := c.PostForm("price")
	tipe := c.PostForm("tipe")
	desc := c.PostForm("desc")

	// Konversikan variabel days dan price dengan tipe data string kedalam bentuk angka integer
	numdays, _ := strconv.Atoi(days)
	numprice, _ := strconv.Atoi(price)

	// Cek data untuk menghindari duplikasi data
	if !settings.DB.First(&dasub, "name = ?", name).RecordNotFound() {
		helpers.ElorResponse(c, "Gagal membuat data subscription baru, data sudah ada dalam database!")
		c.Abort()
		return
	}

	// Buat struktur untuk menyimpan data dalam database
	simpan := models.Subscription{
		Name:  name,
		Days:  int64(numdays),
		Price: int64(numprice),
		Tipe:  tipe,
		Desc:  desc,
	}

	// Simpan data kedalam database
	settings.DB.Create(&simpan)

	// Buat variabel response
	response := FillResSub(simpan)

	// Panggil helper response untuk menampilkan hasil response (list data)
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data subcription master baru!", response)
}

func ListDevSubcription(c *gin.Context) {
	// Definisikan data model role kedalam variabel
	datas := []models.Subscription{}

	// Panggil data model role sesuai dengan variabel model role
	settings.DB.Find(&datas)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.ResSub{}
	for _, datas := range datas {
		list = append(list, FillResSub(datas))
	}

	// Gunakan helper untuk mengirim response data list yang telah diinputkan sebelumnya
	helpers.DataResponse(c, list)
}

func UpdateDevSubscription(c *gin.Context) {
	// Validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Deklarasi nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model role sebagai variabel mod
	var mod models.Subscription

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&mod, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Update data subscription gagal, data subscription tidak ditemukan!")
		c.Abort()
		return
	}

	// Deklarasikan request body kedalam variabel dan sebagian di konversi sesuai tipe data yang terdapat pada database
	name := c.PostForm("name")
	days := c.PostForm("days")
	price := c.PostForm("price")
	tipe := c.PostForm("tipe")
	desc := c.PostForm("desc")

	// Konversikan variabel days dan price dengan tipe data string kedalam bentuk angka integer
	numdays, _ := strconv.Atoi(days)
	numprice, _ := strconv.Atoi(price)

	// Update data role
	settings.DB.Model(&mod).Where("id = ?", id).Updates(map[string]interface{}{
		"name":  name,
		"days":  numdays,
		"price": numprice,
		"tipe":  tipe,
		"desc":  desc,
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
	response := FillResSub(mod)

	// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
	helpers.SuksesWithDataResponse(c, "Berhasil update data subscription baru!", response)
}

func HapusDevSubscription(c *gin.Context) {
	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model subscription sebagai variabel mod
	var mod models.Subscription

	// Sisipkan fungsi model warehouse untuk mengganti bila data subscription dihapus
	var gudang models.Warehouse

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&mod, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Hapus data subscription gagal, data subscription tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data id yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&mod)

	// Ubah data subscription_id yang telah di hapus menjadi 0
	settings.DB.Model(&gudang).Where("subscription_id = ?", id).Update("subscription_id", 0)

	// Panggil response sukses bila berhasil melakukan hapus data
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data subscription master!")
}
