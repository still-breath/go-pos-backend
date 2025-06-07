// File: internal/model/category.go

package model

import (
	"gorm.io/gorm"
)

// Category mendefinisikan struktur untuk tabel 'categories' di database.
// Ini adalah padanan dari class Category di Laravel.
type Category struct {
	// gorm.Model secara otomatis menambahkan field:
	// ID        uint           `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model

	// Name adalah nama dari kategori.
	// Tag `gorm:"unique"` memastikan setiap nama kategori unik di database.
	// Tag `json:"name"` menentukan bagaimana field ini akan dinamai saat diubah menjadi JSON.
	Name string `gorm:"unique" json:"name"`

	// Products mendefinisikan relasi one-to-many (HasMany).
	// Satu Category memiliki banyak Product.
	// GORM akan secara otomatis mencari foreign key 'CategoryID' di model Product.
	// Tag `json:"products,omitempty"` berarti field ini hanya akan muncul di JSON jika datanya dimuat (eager-loaded).
	Products []Product `json:"products,omitempty"`
}
