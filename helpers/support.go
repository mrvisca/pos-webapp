package helpers

import (
	"math/rand"
	"time"
)

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
