package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/still-breath/go-pos-backend/internal/config" // Ganti dengan path modul Anda
	"github.com/still-breath/go-pos-backend/internal/model"  // Ganti dengan path modul Anda
	"golang.org/x/crypto/bcrypt"
)

// DTO (Data Transfer Object) untuk registrasi
type RegisterInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=3"`
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
		Password: input.Password, // Hashing akan dilakukan oleh hook BeforeSave
		Role:     role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// TODO: Implement session creation here
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
