// File: internal/model/product.go

package model

import (
	"gorm.io/gorm"
)

// Product mendefinisikan struktur untuk tabel 'products' di database.
// Ini adalah padanan dari class Product di Laravel.
type Product struct {
	// gorm.Model secara otomatis menambahkan field: ID, CreatedAt, UpdatedAt, DeletedAt.
	gorm.Model

	// Field-field ini akan dipetakan ke kolom database dengan nama snake_case secara default
	// (misal: Name -> name, ImageURL -> image_url, dll.)
	Name          string `json:"name"`
	ImageURL      string `json:"imageUrl"`      // Sesuai dengan 'imageUrl' di frontend
	StockQuantity int    `json:"stockQuantity"` // Sesuai dengan 'stockQuantity'
	Status        string `json:"status"`
	SellingPrice  int    `json:"sellingPrice"`                // Tipe integer sesuai migrasi sebelumnya
	HPP           int    `json:"costPrice" gorm:"column:hpp"` // Memetakan kolom 'hpp' DB ke 'costPrice' JSON
	Unit          string `json:"unit"`

	// Foreign key untuk relasi ke tabel categories.
	// Tag `json:"-"` menyembunyikan field ini dari output JSON untuk menghindari duplikasi data.
	CategoryID uint `json:"-"`

	// Category mendefinisikan relasi `belongsTo`.
	// GORM akan otomatis menggunakan CategoryID sebagai foreign key.
	// Saat dimuat (eager-loaded), data kategori akan dimasukkan ke dalam field ini.
	Category Category `json:"category"`
}
