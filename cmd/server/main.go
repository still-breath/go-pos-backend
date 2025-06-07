package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions" // <-- 1. Impor paket Sesi

	// PERBAIKAN PATH: Hapus .git dari path modul Anda
	"github.com/still-breath/go-pos-backend.git/internal/config"
	"github.com/still-breath/go-pos-backend.git/internal/handler"
	"github.com/still-breath/go-pos-backend.git/internal/middleware" // <-- 2. Impor Middleware
	"github.com/still-breath/go-pos-backend.git/internal/model"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// --- 3. INISIALISASI SESSION STORE ---
	// Mengambil kunci sesi dari file .env. Ini WAJIB.
	sessionKey := config.GetEnv("SESSION_KEY", "yuhj4suTdZvRPOuRnk9meVYGAIOWDxFE1")
	// Membuat session store dan menyimpannya di variabel global handler.Store
	handler.Store = sessions.NewCookieStore([]byte(sessionKey))

	// Jalankan AutoMigrate untuk membuat tabel jika belum ada
	err := config.DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()

	// Konfigurasi CORS (kode Anda sudah benar)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-XSRF-TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- 4. PERBAIKI STRUKTUR RUTE ---
	api := r.Group("/api")
	{
		// Rute publik (tidak perlu login)
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		// Rute yang dilindungi (harus login)
		// Semua rute di dalam grup ini akan melewati AuthMiddleware terlebih dahulu
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/user", handler.GetUser)
			protected.POST("/logout", handler.Logout)
			protected.GET("/owner/dashboard", handler.GetDashboard)
			// Tambahkan rute lain yang memerlukan login di sini
		}
	}

	// Jalankan server
	port := config.GetEnv("APP_PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
