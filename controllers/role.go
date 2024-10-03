package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FillResrole(roled models.Role) models.Resrole {
	return models.Resrole{
		ID:       roled.ID,
		Name:     roled.Name,
		Desc:     roled.Desc,
		IsActive: roled.IsActive,
	}
}

func CreateDevRole(c *gin.Context) {
	// validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Definisikan model role kedalam variabel data
	var daro models.Role

	// Deklarasi request body kedalam variabel
	name := c.PostForm("name")
	desc := c.PostForm("desc")

	// Cek data untuk menghindari duplikasi data
	if !settings.DB.First(&daro, "name = ?", name).RecordNotFound() {
		helpers.ElorResponse(c, "Gagal membuat data role baru, data sudah ada dalam database!")
		c.Abort()
		return
	}

	// Buat struktur untuk menyimpan data dalam database
	simpan := models.Role{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}

	// Simpan data kedalam database
	settings.DB.Create(&simpan)

	// Buat variabel response
	response := FillResrole(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data role baru!", response)
}

func ListDevRole(c *gin.Context) {
	// Definisikan data model role kedalam variabel
	datas := []models.Role{}

	// Panggil data model role sesuai dengan variabel model role
	settings.DB.Find(&datas)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.Resrole{}
	for _, datas := range datas {
		list = append(list, FillResrole(datas))
	}

	// Gunakan helper untuk mengirim response data list yang telah di inputkan sebelumnya
	helpers.DataResponse(c, list)
}

func UpdateDevRole(c *gin.Context) {
	// Validasi api dengan passcode dev
	passcode := c.PostForm("passcode")
	if passcode != "into_muros" {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Deklarasi nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model role sebagai veriabel mod
	var mod models.Role

	// Sisipkan fungsi model user untuk mengganti 0 bila id role di nonaktifkan
	var usr models.User

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&mod, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Update data role gagal, data role tidak ditemukan!")
		c.Abort()
		return
	}

	// Ubah nilai request string is_active sebagai type data boolean
	is_active := c.PostForm("is_active")
	nilai, _ := strconv.ParseBool(is_active)

	// Jika role di nonaktifkan maka relasi role_id akan berubah menjadi 0 (nilai default)
	if !nilai {
		settings.DB.Model(&usr).Where("role_id = ?", id).Update("role_id", 0)
	}

	// Update data role
	settings.DB.Model(&mod).Where("id = ?", id).Updates(map[string]interface{}{
		"name":      c.PostForm("name"),
		"desc":      c.PostForm("desc"),
		"is_active": nilai,
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
	response := FillResrole(mod)

	// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
	helpers.SuksesWithDataResponse(c, "Berhasil update data role baru!", response)
}

func HapusDevRole(c *gin.Context) {
	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model role sebagai variabel mod
	var mod models.Role

	// Sisipkan fungsi model user untuk mengganti bila role dihapus
	var usr models.User

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&mod, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Hapus data role gagal, data role tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data id yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&mod)

	// Ubah data role_id yang telah dihapus menjadi 0
	settings.DB.Model(&usr).Where("role_id = ?", id).Update("role_id", 0)

	// Panggil response sukses bila berhasil hapus data
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data role!")
}
