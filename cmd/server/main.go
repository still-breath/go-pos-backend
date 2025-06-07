package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/still-breath/go-pos-backend.git/internal/config"  // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/handler" // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/model"   // Ganti path
)

func main() {
	// Muat konfigurasi dari .env
	config.LoadEnv()

	// Hubungkan ke database
	config.ConnectDB()

	// Jalankan AutoMigrate (seperti php artisan migrate)
	// Ini akan membuat tabel jika belum ada
	err := config.DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Inisialisasi Gin router
	r := gin.Default()

	// Grup rute untuk API
	api := r.Group("/api")
	{
		// Rute publik
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		// TODO: Grup rute yang dilindungi middleware
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
