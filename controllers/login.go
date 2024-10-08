package controllers

import (
	"fmt"
	"os"
	"pos-webapp/helpers"
	"pos-webapp/models"
	"pos-webapp/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func createToken(user *models.User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      user.ID,
		"role_id":      user.RoleId,
		"business_id":  user.BusinessId,
		"warehouse_id": user.WarehouseId,
		"name":         user.Name,
		"email":        user.Email,
		"exp":          time.Now().AddDate(0, 0, 7).Unix(),
		"iat":          time.Now().Unix(),
	})

	// Sign and get the completed encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

func LoginCheck(c *gin.Context) {
	// Ambil body request form login dan deklarasikan pada masing-masing variabel
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Definisikan model role kedalam variabel
	var user models.User

	// Kondisi bila email user tidak ditemukan
	if settings.DB.First(&user, "email = ?", email).RecordNotFound() {
		helpers.ElorResponse(c, "Email tidak terdaftar dalam sistem kami!")
		c.Abort()
		return
	}

	// Kondisi bila akun belum terverifikasi
	if !user.IsVerified {
		helpers.ElorResponse(c, "Login gagal, akun belum terverifikasi!")
		c.Abort()
		return
	}

	// Kondisi pencocokan data password input form user dengan data password dalam database
	if !helpers.CheckPassword(user.Password, password) {
		helpers.ElorResponse(c, "Password tidak valid!")
		c.Abort()
		return
	}

	// Buat token bila data sudah sesuai dengan data dalam database
	token := createToken(&user)

	// Tambahkan response sukses untuk menampilkan response dan juga token autentikasi
	helpers.SuksesLogin(c, "Login aplikasi berhasil dilakukan!", token, int64(user.RoleId))
}
