package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors" // <-- 1. Impor paket CORS
	"github.com/gin-gonic/gin"

	"github.com/still-breath/go-pos-backend.git/internal/config"  // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/handler" // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/model"   // Ganti dengan path modul Anda
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Jalankan AutoMigrate untuk membuat tabel jika belum ada
	err := config.DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()

	// --- KONFIGURASI CORS DIMULAI DI SINI ---
	r.Use(cors.New(cors.Config{
		// 2. AllowOrigins berisi daftar alamat frontend yang diizinkan.
		//    Ganti dengan URL frontend Anda jika berbeda.
		AllowOrigins: []string{"http://localhost:5173"},

		// 3. AllowMethods menentukan metode HTTP apa saja yang diizinkan.
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		// 4. AllowHeaders menentukan header apa saja yang boleh dikirim oleh frontend.
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-XSRF-TOKEN"},

		// 5. ExposeHeaders memungkinkan frontend membaca header tertentu dari respons backend.
		ExposeHeaders: []string{"Content-Length"},

		// 6. AllowCredentials WAJIB true agar frontend bisa mengirim dan menerima cookie (penting untuk sesi login).
		AllowCredentials: true,

		// 7. MaxAge menentukan berapa lama hasil preflight request bisa di-cache oleh browser.
		MaxAge: 12 * time.Hour,
	}))
	// --- KONFIGURASI CORS SELESAI ---

	// Grup rute untuk API
	api := r.Group("/api")
	{
		// Rute publik
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		// Rute yang dilindungi (Contoh, perlu implementasi middleware autentikasi)
		// protected := api.Group("/")
		// protected.Use(middleware.AuthMiddleware())
		// {
		//     protected.GET("/dashboard", handler.GetDashboard)
		// }
	}

	// Jalankan server
	port := config.GetEnv("APP_PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
