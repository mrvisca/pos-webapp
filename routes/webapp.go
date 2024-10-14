package routes

import (
	"log"
	"net/http"
	"pos-webapp/controllers"
	"pos-webapp/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebAppRoute() {
	// Memanggil fungsi route dari framework gin golang
	router := gin.Default()

	// Menambahkan cors pada settingan route gin golang
	// router.Use(cors.Default())

	// Menggunakan middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://example.com"}, // Ganti dengan origin yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Memuat template dengan ekstensi .tmpl dari direktori view
	router.LoadHTMLGlob("view/*.tmpl")

	// Menyajikan file statis dari direktori assets
	router.Static("/assets", "./assets")

	// Route Website
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Login.tmpl", nil) // Render template Login.tmpl
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Register.tmpl", nil) // Render template Register.tmpl
	})
	router.GET("/aplikasi/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Dashboard.tmpl", nil) // Render template Dashboard.tmpl
	})
	router.GET("/aplikasi/staff/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ListPegawai.tmpl", nil) // Render template ListPegawai.tmpl
	})

	v1 := router.Group("api/v1/")
	{
		// Route khusus development
		devs := v1.Group("/dev-only/")
		{
			roleDev := devs.Group("/role/")
			{
				roleDev.POST("/create", controllers.CreateDevRole)
				roleDev.GET("/list", controllers.ListDevRole)
				roleDev.PUT("/update/:id", controllers.UpdateDevRole)
				roleDev.DELETE("/hapus/:id", controllers.HapusDevRole)
			}

			subscriptionDev := devs.Group("/master-subscription/")
			{
				subscriptionDev.POST("/create", controllers.CreateDevSubscription)
				subscriptionDev.GET("/list", controllers.ListDevSubcription)
				subscriptionDev.PUT("/update/:id", controllers.UpdateDevSubscription)
				subscriptionDev.DELETE("/hapus/:id", controllers.HapusDevSubscription)
			}

			devs.POST("/kirim-email", controllers.TestKirimEmail)
		}

		// Route untuk aplikasi
		oth := v1.Group("/autentikasi/")
		{
			oth.POST("/pendaftaran", controllers.RegisterAcc)
			oth.POST("/login", controllers.LoginCheck)
			oth.GET("/verifikasi/:kode", controllers.Verifikasi)
		}

		dashboardpage := v1.Group("/dashboard/")
		{
			dashboardpage.GET("/profile-check", middleware.IsAuth(), controllers.ProfilePengguna)
			dashboardpage.GET("/support/cabang", middleware.IsAuth(), controllers.SupportListCabang)
			dashboardpage.PUT("/ubah-cabang/:id", middleware.IsAuth(), controllers.UbahPilihCabang)
		}

		staffpage := v1.Group("/staff/")
		{
			staffpage.GET("/list", middleware.IsAuth(), controllers.ListPegawai)
			staffpage.POST("/tambah-data", middleware.IsAuth(), controllers.TambahPegawai)
			staffpage.GET("/support/role", middleware.IsAuth(), controllers.SupportDataRole)
			staffpage.PUT("/update/:id", middleware.IsAuth(), controllers.UpdatePegawai)
			staffpage.DELETE("/hapus/:id", middleware.IsAuth(), controllers.HapusStaff)
		}

		// Route Logout
		v1.GET("logout", controllers.Logout)
	}

	// Menampilkan log server berjalan dengan port 8080
	log.Println("Server started on: http://127.0.0.1:8080")

	// Menjalankan server ke port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
