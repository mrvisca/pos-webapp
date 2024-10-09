package controllers

import (
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	blacklist   = make(map[string]struct{}) // Blacklist untuk menyimpan token yang di-logout
	blacklistMu sync.Mutex                  // Mutex untuk mengamankan akses ke blacklist
)

type ProfileRes struct {
	UserId      uint
	RoleId      uint
	RoleName    string
	BusinessId  uint
	WarehouseId uint
	Name        interface{}
	Email       interface{}
}

func FillSupportCabang(cabs models.Warehouse) models.SupportCabang {
	return models.SupportCabang{
		ID:   cabs.ID,
		Name: cabs.Name,
	}
}

func ProfilePengguna(c *gin.Context) {
	userid := uint(c.MustGet("jwt_user_id").(float64))
	roleid := uint(c.MustGet("jwt_role_id").(float64))
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))
	name := c.MustGet("jwt_name")
	email := c.MustGet("jwt_email")

	// Buat variabel untuk dimasukan nilai role model
	var romod models.Role

	// Panggil data role untuk mengambil nama role
	settings.DB.First(&romod, "id = ?", roleid)

	profile := ProfileRes{
		UserId:      userid,
		RoleId:      roleid,
		RoleName:    romod.Name,
		BusinessId:  businessid,
		WarehouseId: warehouseid,
		Name:        name,
		Email:       email,
	}

	helpers.SuksesWithDataResponse(c, "Berhasil memanggil data profile", profile)
}

// Fungsi Logout
func Logout(c *gin.Context) {
	// Mengambil token dari header Authorization
	authHeader := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) == 2 {
		// Simpan token di blacklist
		blacklistMu.Lock()
		blacklist[bearerToken[1]] = struct{}{}
		blacklistMu.Unlock()

		helpers.SuksesResponse(c, "Berhasil logout akun!")
	} else {
		helpers.ElorResponse(c, "Format token tidak valid!")
		c.Abort()
		return
	}
}

func SupportListCabang(c *gin.Context) {
	// Ambil data jwt_business_id dari parsing bearer token
	businessid := uint(c.MustGet("jwt_business_id").(float64))

	// Buat variabel untuk digunakan oleh model cabang
	cabang := []models.Warehouse{}

	// Panggil data cabang untuk digunakan sebagai data support
	settings.DB.Where("business_id = ?", businessid).Find(&cabang)

	// Deklarasikan model struct response untuk dimasukan data
	list := []models.SupportCabang{}
	for _, cabang := range cabang {
		list = append(list, FillSupportCabang(cabang))
	}

	// Gunakan helper untuk mengirim response data list yang telah diinputkan sebelumnya
	helpers.DataResponse(c, list)
}

func UbahPilihCabang(c *gin.Context) {
	// Masukan jwt_user_id dan jwt_role_id kedalam variabel untuk pencarian data
	roleid := uint(c.MustGet("jwt_role_id").(float64))
	userid := uint(c.MustGet("jwt_user_id").(float64))

	// Masukan query param ke dalam variabel
	cabangid := c.Param("id")

	// Definisikan model user ke dalam variabel
	var user models.User

	// Hanya owner yang bisa mengganti cabang
	if roleid != 2 {
		helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
		c.Abort()
		return
	}

	// Kondisi bila data id tidak ditemukan
	if settings.DB.First(&user, "id = ?", userid).RecordNotFound() {
		helpers.ElorResponse(c, "Ubah cabang gagal!, data pengguna tidak ditemukan!")
		c.Abort()
		return
	}

	// Update data cabang_id pada user
	settings.DB.Model(&user).Where("id = ?", userid).Update("warehouse_id", cabangid)

	// Buat token bila data sudah terupdate dalam database
	token := helpers.CreateToken(&user)

	// Tambahkan response sukses untuk menampilkan response dan juga token autentikasi
	helpers.SuksesLogin(c, "Berhasil melakukan perubahan data cabang owner", token, int64(user.RoleId))
}
