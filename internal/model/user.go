package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // Ini akan membuat ID, CreatedAt, UpdatedAt, DeletedAt secara otomatis
	Name       string `json:"name"`
	Email      string `gorm:"unique" json:"email"`
	Password   string `json:"-"` // Jangan tampilkan password di JSON
	Role       string `json:"role"`
}

// Hook BeforeSave untuk hashing password otomatis sebelum disimpan
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return
}
