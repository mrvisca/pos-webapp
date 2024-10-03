package helpers

import "net/smtp"

// Fungsi untuk mengirim email
func SendRegisEmail(to string, subject string, body string) error {
	from := "bimasaktiputra95@gmail.com" // Ganti dengan email Anda
	password := "lqsq dpej owqi qlts"    // Ganti dengan password email Anda

	// Set up the SMTP server configuration
	smtpHost := "smtp.gmail.com" // Ganti dengan SMTP host Anda
	smtpPort := "587"            // Port SMTP (umumnya 587 untuk TLS)

	// Buat message
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Kirim email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
