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
		ID:            desta.ID,
		RoleId:        desta.RoleId,
		RoleName:      desta.Role.Name,
		BusinessId:    desta.BusinessId,
		BusinessName:  desta.Business.Name,
		WarehouseId:   desta.WarehouseId,
		WarehouseName: desta.Warehouse.Name,
		Name:          desta.Name,
		Email:         desta.Email,
		Phone:         desta.Phone,
		IsVerified:    desta.IsVerified,
		Kode:          desta.Kode,
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
	settings.DB.Where("business_id = ? AND warehouse_id = ? AND role_id != ?", businessid, warehouseid, 2).Find(&datas)

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
