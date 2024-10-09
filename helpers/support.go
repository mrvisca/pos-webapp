package helpers

import (
	"fmt"
	"math/rand"
	"os"
	"pos-webapp/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func CreateToken(user *models.User) string {
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

// Fungsi untuk membuat string acak 6 karakter
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Buat random generator baru dengan seed berbasis waktu
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, n)
	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}
