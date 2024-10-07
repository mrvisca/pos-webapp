package routes

import (
	"log"
	"net/http"
	"pos-webapp/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebAppRoute() {
	// Memanggil fungsi route dari framework gin golang
	router := gin.Default()

	// Menambahkan cors pada settingan route gin golang
	router.Use(cors.Default())

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
		}
	}

	// Menampilkan log server berjalan dengan port 8080
	log.Println("Server started on: http://127.0.0.1:8080")

	// Menjalankan server ke port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
