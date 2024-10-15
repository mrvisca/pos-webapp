package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FillStaffList(desta models.User) models.Staff {
	return models.Staff{
		ID:          desta.ID,
		RoleId:      desta.RoleId,
		RoleName:    desta.Role.Name,
		BusinessId:  desta.BusinessId,
		WarehouseId: desta.WarehouseId,
		Name:        desta.Name,
		Email:       desta.Email,
		Phone:       desta.Phone,
		IsVerified:  desta.IsVerified,
		Kode:        desta.Kode,
	}
}

func FillSupportRole(rol models.Role) models.SupportRole {
	return models.SupportRole{
		ID:   rol.ID,
		Name: rol.Name,
	}
}

func ListPegawai(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan data model user kedalam variabel
	datas := []models.User{}

	// Panggil data model user sesuai dengan variabel mode user
	settings.DB.Where("business_id = ? AND warehouse_id = ? AND role_id != ? AND role_id != ?", businessid, warehouseid, 2, 1).Preload("Role").Find(&datas)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.Staff{}
	for _, datas := range datas {
		list = append(list, FillStaffList(datas))
	}

	// Gunakan helper untuk mengirim response data list yang telah diinputkan sebelumnya
	helpers.DataResponse(c, list)
}

func TambahPegawai(c *gin.Context) {
	// Masukan data request body kedalam masing-masing variabel untuk memudahkan proses input data
	roleid := c.PostForm("role_id")
	businessid := c.PostForm("business_id")
	warehouseid := c.PostForm("warehouse_id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// Konversi request body string ke tipe data angka
	valroleid, _ := strconv.Atoi(roleid)
	valbussinessid, _ := strconv.Atoi(businessid)
	valwarehouseid, _ := strconv.Atoi(warehouseid)

	// Enkripsi password
	sandi, _ := helpers.HashPassword(password)

	// Kondisi filter gagal bila menambahkan data admin / owner pada pembuatan data pegawai baru
	if valroleid == 1 || valroleid == 2 {
		helpers.ElorResponse(c, "Gagal membuat data pegawai baru, akses ditangguhkan!")
		c.Abort()
		return
	}

	// Definisikan model user kedalam sebuah variabel
	var peg models.User

	// Cek duplikasi data email
	if !settings.DB.First(&peg, "email LIKE ?", email).RecordNotFound() {
		helpers.ElorResponse(c, "Gagal membuat data pegawai baru, email sudah terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Sesuaikan variabel data sesuai dengan struct user
	simpan := models.User{
		RoleId:      uint(valroleid),
		BusinessId:  uint(valbussinessid),
		WarehouseId: uint(valwarehouseid),
		Name:        name,
		Email:       email,
		Password:    sandi,
		Phone:       phone,
		IsVerified:  false,
		Kode:        helpers.RandomString(6),
	}

	// Simpan data kedalam database user pegawai
	settings.DB.Create(&simpan)

	// Masukan struct response yang rapi kedalam variabel untuk di tampilkan
	response := FillStaffList(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data pegawai baru!", response)
}

func UpdatePegawai(c *gin.Context) {
	// Klasifikasi param id kedalam variabel id
	id := c.Param("id")

	// Klasifikasi body request put kedalam masing-masing variabel
	roleid := c.PostForm("role_id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	verifikasi := c.PostForm("verifikasi")

	// Konversi request body string (integer) ke tipe data asli angka
	valroleid, _ := strconv.Atoi(roleid)

	// Konversi request verifikasi string menjadi boolean
	boolverifikasi, _ := strconv.ParseBool(verifikasi)

	if valroleid == 1 || valroleid == 2 {
		helpers.ElorResponse(c, "Gagal mengubah data pegawai, akses ditangguhkan")
		c.Abort()
		return
	}

	// Definisikan model user kedalam sebuah variabel
	var peg models.User

	// Buat kondisi ubah dengan password atau tanpa password
	if password != "" {
		// Enkripsi password
		sandi, _ := helpers.HashPassword(password)

		// Kondisi email tidak boleh sama dengan email akun lainnya
		if !settings.DB.First(&peg, "email LIKE ? AND id != ?", "%"+email+"%", id).RecordNotFound() {
			helpers.ElorResponse(c, "Gagal mengubah data pegawai, email sudah terdaftar dalam sistem kami!")
			c.Abort()
			return
		}

		// Update data user pegawai dengan password
		settings.DB.Model(&peg).Where("id = ?", id).Updates(map[string]interface{}{
			"role_id":     valroleid,
			"name":        name,
			"email":       email,
			"password":    sandi,
			"phone":       phone,
			"is_verified": boolverifikasi,
		})

		// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
		response := FillStaffList(peg)

		// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
		helpers.SuksesWithDataResponse(c, "Berhasil update data pegawai dengan password!", response)
	} else {
		// Kondisi email tidak boleh sama dengan email akun lainnya
		if !settings.DB.First(&peg, "email LIKE ? AND id != ?", email, id).RecordNotFound() {
			helpers.ElorResponse(c, "Gagal mengubah data pegawai, email sudah terdaftar dalam sistem kami!")
			c.Abort()
			return
		}

		// Update data user pegawai dengan password
		settings.DB.Model(&peg).Where("id = ?", id).Updates(map[string]interface{}{
			"role_id":     valroleid,
			"name":        name,
			"email":       email,
			"phone":       phone,
			"is_verified": boolverifikasi,
		})

		// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
		response := FillStaffList(peg)

		// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
		helpers.SuksesWithDataResponse(c, "Berhasil update data pegawai tanpa password!", response)
	}
}

func HapusStaff(c *gin.Context) {
	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model user sebagai variabel user
	var user models.User

	// kondisi bila data id tidak ditemukan
	if settings.DB.First(&user, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Hapus data staff gagal, data staff tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data user dari id yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&user)

	// Panggil response sukses bila berhasil hapus data
	helpers.SuksesResponse(c, "Berhasil menghapus data pegawai!")
}

func SupportDataRole(c *gin.Context) {
	// Definisikan model role kedalam sebuah variabel
	surole := []models.Role{}

	// Panggil semua data role yang berada pada database
	settings.DB.Where("id != ? AND id != ?", 1, 2).Find(&surole)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.SupportRole{}
	for _, surole := range surole {
		list = append(list, FillSupportRole(surole))
	}

	// Gunakan helper untuk mengirim response data list yang telah di inputkan sebelumnya
	helpers.DataResponse(c, list)
}
