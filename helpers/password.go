package helpers

import "golang.org/x/crypto/bcrypt"

// Fungsi untuk melakukan hash pada password
func HashPassword(password string) (string, error) {
	// Gunakan bcrypt dengan cost default (biasanya 10)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Fungsi untuk memeriksa apakah password cocok dengan hash
func CheckPassword(hashedPassword, password string) bool {
	// Cek password dengan hash yang disimpan
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // Jika tidak ada error, berarti cocok
}
