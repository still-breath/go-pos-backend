package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database yang bisa diakses dari paket lain
var DB *gorm.DB

// ConnectDB membuat koneksi ke database PostgreSQL menggunakan variabel dari .env
func ConnectDB() {
	var err error

	// Mengambil variabel lingkungan menggunakan fungsi GetEnv yang ada di config.go
	// Jika variabel tidak ada, akan digunakan nilai fallback (parameter kedua)
	dbUser := GetEnv("DB_USER", "postgres")
	dbPassword := GetEnv("DB_PASSWORD", "postgres")
	dbHost := GetEnv("DB_HOST", "127.0.0.1")
	dbPort := GetEnv("DB_PORT", "5432")
	dbName := GetEnv("DB_NAME", "pos_laravel_db")

	// Membuat DSN (Data Source Name) untuk koneksi PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	// Membuka koneksi ke database menggunakan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	log.Println("Koneksi database berhasil dibuka")
}
