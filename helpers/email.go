package helpers

import "net/smtp"

// Fungsi untuk mengirim email
func SendRegisEmail(to string, subject string, body string) error {
	from := "your_email@example.com" // Ganti dengan email Anda
	password := "your_password"      // Ganti dengan password email Anda

	// Set up the SMTP server configuration
	smtpHost := "smtp.example.com" // Ganti dengan SMTP host Anda
	smtpPort := "587"              // Port SMTP (umumnya 587 untuk TLS)

	// Buat message
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Kirim email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
