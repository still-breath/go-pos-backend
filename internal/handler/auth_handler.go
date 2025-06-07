package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/still-breath/go-pos-backend.git/internal/config" // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Store adalah variabel global untuk session store, akan diinisialisasi di main.go
var Store *sessions.CookieStore

// DTO (Data Transfer Object) untuk registrasi
type RegisterInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"` // Diubah ke min=8 untuk keamanan
	PasswordConfirm string `json:"password_confirmation" binding:"required"`
	Role            string `json:"role"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	role := "admin" // Default role
	if input.Role != "" {
		role = input.Role
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password, // Hashing akan dilakukan oleh hook BeforeSave di model
		Role:     role,
	}

	// Mencoba membuat user baru di database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user, email may already exist"})
		return
	}

	// Kembalikan pesan sukses agar pengguna login secara manual
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully. Please log in."})
}

// DTO untuk login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// PERBAIKAN KRITIS: Membandingkan hash dari DB dengan password dari input
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Membuat sesi setelah login berhasil
	session, _ := Store.Get(c.Request, "pos-session")
	session.Values["user_id"] = user.ID
	session.Options.MaxAge = 60 * 60 * 24 // Sesi berlaku selama 1 hari

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, user) // Kembalikan data user ke frontend
}

func Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "pos-session")

	// Hapus data dari sesi
	session.Values["user_id"] = nil
	session.Options.MaxAge = -1 // Memberitahu browser untuk segera menghapus cookie

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func GetUser(c *gin.Context) {
	// Ambil user ID dari context yang di-set oleh middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in database"})
		return
	}

	c.JSON(http.StatusOK, user)
}
