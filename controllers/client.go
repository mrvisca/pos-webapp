package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FillClientList(clientData models.Client) models.Pelanggan {
	return models.Pelanggan{
		ID:          clientData.ID,
		BusinessId:  clientData.BusinessId,
		WarehouseId: clientData.WarehouseId,
		Name:        clientData.Name,
		Email:       clientData.Email,
		Phone:       clientData.Phone,
		Address:     clientData.Address,
		Join:        clientData.Join,
	}
}

func ListClient(c *gin.Context) {
	// Definisikan model client kedalam variabel cli
	cli := []models.Client{}

	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan field column sortby dalam sebuah variabel
	columns := []string{
		"id",
		"name",
		"email",
		"phone",
		"address",
		"join",
	}

	// Definisikan request dalam bentuk variabel untuk memudahkan sort data
	start, _ := strconv.Atoi(c.PostForm("start")) // menggunakan PostForm untuk mendapatkan nilai dari request body
	limit, _ := strconv.Atoi(c.PostForm("length"))
	draw, _ := strconv.Atoi(c.PostForm("draw"))

	// Periksa apakah query order dan column ada
	orderColumn := "id" // nilai default
	if orderQuery := c.PostForm("order[0][column]"); orderQuery != "" {
		if columnIndex, err := strconv.Atoi(orderQuery); err == nil && columnIndex < len(columns) {
			orderColumn = columns[columnIndex]
		}
	}
	dir := c.PostForm("order[0][dir]")
	if dir == "" {
		dir = "asc" // nilai default
	}
	search := c.PostForm("search[value]") // Menggunakan PostForm untuk search value

	// Hitung total keseluruhan
	var totalRecord int64
	settings.DB.Model(&cli).Where("business_id = ? AND warehouse_id = ?", businessid, warehouseid).Count(&totalRecord)

	// Ambil data client
	settings.DB.Where("business_id = ? AND warehouse_id = ?", businessid, warehouseid).
		Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%").
		Order(orderColumn + " " + dir).
		Offset(start).Limit(limit).Find(&cli)

	// Mengubah data client menjadi format response
	clientList := []models.Pelanggan{}
	for _, cli := range cli {
		clientList = append(clientList, FillClientList(cli))
	}

	// Panggil helper response untuk memunculkan response yang di harapkan
	helpers.ListPaginate(c, clientList, totalRecord, draw)
}

func TambahPelanggan(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan request post body kedalam masing-masing variabel
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	Address := c.PostForm("address")

	// Set waktu untuk mengisi kolom join
	sekarang := time.Now()
	join := sekarang.Format("2006-01-02 15:04:05")

	// Definisikan model client kedalam variabel client
	var client models.Client

	if !settings.DB.First(&client, "email LIKE ?", "%"+email+"%").RecordNotFound() {
		helpers.ElorResponse(c, "Email sudah digunakan oleh pelanggan lain!")
		c.Abort()
		return
	}

	// Sesuaikan variabel data sesuai dengan struct client
	simpan := models.Client{
		BusinessId:  businessid,
		WarehouseId: warehouseid,
		Name:        name,
		Email:       email,
		Phone:       phone,
		Address:     Address,
		Join:        join,
	}

	// Simpan data kedalam database pelanggan
	settings.DB.Create(&simpan)

	// Masukan struct response yang rapi kedalam variabel untuk ditampilkan
	response := FillClientList(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data pelanggan baru!", response)
}

func UpdatePelanggan(c *gin.Context) {
	// Klasifikasi param id kedalam variabel id
	id := c.Param("id")

	// Klasifikasi body request put kedalam masing-masing variabel
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	// Definisikan model client kedalam sebuah variabel
	var pel models.Client

	// Buat kondisi bila data id pelanggan tidak ditemukan
	if settings.DB.First(&pel, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id pelanggan tidak ditemukan!")
		c.Abort()
		return
	}

	// Kondisi email tidak boleh sama dengan akun lainnya
	if !settings.DB.First(&pel, "email LIKE ? AND id != ?", "%"+email+"%", id).RecordNotFound() {
		helpers.ElorResponse(c, "Email sudah terdaftar di data pelanggan lainnya!")
		c.Abort()
		return
	}

	// Update data pelanggan jika id pelanggan ditemukan
	settings.DB.Model(&pel).Where("id = ?", id).Updates(models.Client{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada variabel untuk ditampilkan
	response := FillClientList(pel)

	// Masukan input data yang telah dilakukan akan ditampilkan dalam struct response
	helpers.SuksesWithDataResponse(c, "Sukses update data pelanggan!", response)
}

func HapusPelanggan(c *gin.Context) {
	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Sisipkan fungsi model pelanggan sebagai variabel pelanggan
	var pelanggan models.Client

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&pelanggan, "id = ?", id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id staff tidak ditemukan!")
		c.Abort()
		return
	}

	// Hapus data pelanggan dari id yang terdapat pada database
	settings.DB.Where("id = ?", id).Delete(&pelanggan)

	// Panggil response sukses helper bila data berhasil dihapus
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data pelanggan!")
}
