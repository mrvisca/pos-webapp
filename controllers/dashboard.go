package controllers

import (
	"pos-webapp/helpers"
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
	BusinessId  uint
	WarehouseId uint
	Name        interface{}
	Email       interface{}
}

func ProfilePengguna(c *gin.Context) {
	userid := uint(c.MustGet("jwt_user_id").(float64))
	roleid := uint(c.MustGet("jwt_role_id").(float64))
	businessid := uint(c.MustGet("jwt_business_id").(float64))
	warehouseid := uint(c.MustGet("jwt_warehouse_id").(float64))
	name := c.MustGet("jwt_name")
	email := c.MustGet("jwt_email")

	profile := ProfileRes{
		UserId:      userid,
		RoleId:      roleid,
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
