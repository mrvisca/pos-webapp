package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FillResponsePaymentMethod(paymet models.PaymentMethod) models.ResponsePaymentMethod {
	return models.ResponsePaymentMethod{
		ID:          paymet.ID,
		BusinessId:  paymet.BusinessId,
		WarehouseId: paymet.WarehouseId,
		Tipe:        paymet.Tipe,
		Name:        paymet.Name,
		Norek:       paymet.Norek,
		Admin:       paymet.Admin,
	}
}

func ListPaymentMethod(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan data model payment method kedalam variabel
	datas := []models.PaymentMethod{}

	// Panggil data model payment method sesuai dengan id bisnis, cabang yang terdapat pada token
	settings.DB.Where("business_id = ? AND warehouse_id = ?", businessid, warehouseid).Find(&datas)

	// Deklarasikan model struct response untuk dimasukan data iterasi
	list := []models.ResponsePaymentMethod{}
	for _, datas := range datas {
		list = append(list, FillResponsePaymentMethod(datas))
	}

	// Gunakan helper untuk mengirim data list yang telah di iterasi sebelumnya
	helpers.DataResponse(c, list)
}

func CreatePaymentMethod(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Masukan data request body kedalam masing-masing variabel untuk memudahkan proses input data
	tipe := c.PostForm("tipe")
	nameAkun := c.PostForm("name")
	norek := c.PostForm("norek")
	admin := c.PostForm("admin")

	// Konversi tipe data norek dan admin menjadi integer angka
	valueNorek, _ := strconv.Atoi(norek)
	valueAdmin, _ := strconv.Atoi(admin)

	// Definisikan model payment method kedalam sebuah variabel
	var payMethod models.PaymentMethod

	// Cek data duplikat
	if !settings.DB.First(&payMethod, "business_id = ? AND warehouse_id = ? AND tipe = ? AND name = ?", businessid, warehouseid, tipe, nameAkun).RecordNotFound() {
		helpers.ElorResponse(c, "Data metode bayar sudah ada dalam database!")
		c.Abort()
		return
	}

	// Sesuaikan variabel data sesuai dengan struct metode bayar
	simpan := models.PaymentMethod{
		BusinessId:  businessid,
		WarehouseId: warehouseid,
		Tipe:        tipe,
		Name:        nameAkun,
		Norek:       int64(valueNorek),
		Admin:       int64(valueAdmin),
	}

	// Simpan data kedalam database metode bayar
	settings.DB.Create(&simpan)

	// Masukan struct response yang rapi kedalam variabel untuk ditampilkan
	response := FillResponsePaymentMethod(simpan)

	// Panggil helper response untuk menampilkan hasil response
	helpers.SuksesWithDataResponse(c, "Berhasil membuat data metode bayar baru!", response)
}

func UpdatePaymentMethod(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Definisikan param query id kedalam variabel
	id := c.Param("id")

	// Konversi variabel norek dan admin menjadi integer angka
	norek, _ := strconv.Atoi(c.PostForm("norek"))
	admin, _ := strconv.Atoi(c.PostForm("admin"))

	// Definsikan model payment method kedalam sebuah variabel
	var payMethod models.PaymentMethod

	// Kondisi saat data id tidak ditemukan dalam database
	if settings.DB.First(&payMethod, "business_id = ? AND warehouse_id = ? AND id = ?", businessid, warehouseid, id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id metode bayar tidak ditemukan dalam sistem kami!")
		c.Abort()
		return
	}

	// Update data metode bayar
	settings.DB.Model(&payMethod).Where("id = ?", id).Updates(models.PaymentMethod{
		Tipe:  c.PostForm("tipe"),
		Name:  c.PostForm("name"),
		Norek: int64(norek),
		Admin: int64(admin),
	})

	// Masukan nilai inputan kedalam struct response dan deklarasikan pada sebuah variabel response
	response := FillResponsePaymentMethod(payMethod)

	// Masukan input data update yang telah dimasukan data
	helpers.SuksesWithDataResponse(c, "Berhasil melakukan update data metode bayar!", response)
}

func HapusPaymentMethod(c *gin.Context) {
	// Klasifikasi jwt_business_id dan jwt_warehouse_id kedalam masing-masing variabel
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))

	// Deklarasikan nilai param id kedalam variabel id
	id := c.Param("id")

	// Definisikan model metode bayar dalam sebuah variabel
	var payMethod models.PaymentMethod

	// Kondisi bila data id metode pembayaran tidak ditemukan
	if settings.DB.First(&payMethod, "business_id = ? AND warehouse_id = ? AND id = ?", businessid, warehouseid, id).RecordNotFound() {
		helpers.ElorResponse(c, "Data id metode bayar tidak ditemukan dalam sistem kami!")
		c.Abort()
		return
	}

	// Hapus jika id ada dalam database
	settings.DB.Where("id = ?", id).Delete(&payMethod)

	// Panggil helper response sukses hapus bila data berhasil terhapus
	helpers.SuksesResponse(c, "Berhasil melakukan hapus data metode bayar!")
}

func SupportMasterPayment(c *gin.Context) {
	// Definisikan param query method kedalam sebuah variabel
	method := c.Param("method")

	// Definisikan model role master payment kedalam sebuah variabel
	sumapa := []models.MasterPayment{}

	// Panggil semua data master payment yang berada dalam database
	settings.DB.Where("tipe = ?", method).Find(&sumapa)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.ResMasterPayment{}
	for _, sumapa := range sumapa {
		list = append(list, FillMapaRes(sumapa))
	}

	// Gunakan helper untuk mengirim response data list yang telah di iterasi sebelumnya
	helpers.DataResponse(c, list)
}
