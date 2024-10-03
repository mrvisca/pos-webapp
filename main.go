package main

import (
	"pos-webapp/routes"
	"pos-webapp/settings"
)

func main() {
	// Panggil fungsi koneksi ke database
	settings.InitDB()
	defer settings.DB.Close()

	// Panggil fungsi route webapp
	routes.WebAppRoute()
}
