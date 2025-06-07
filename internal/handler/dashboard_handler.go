package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "go-pos-wanti/internal/config" // Tidak perlu koneksi DB untuk saat ini
	// "go-pos-wanti/internal/model"  // Tidak perlu model untuk saat ini
)

func GetDashboard(c *gin.Context) {
	// 1. Data Hardcode untuk Statistik Cards
	dashboardStats := gin.H{
		"dailySales":        500000,
		"monthlyRevenue":    50000000,
		"totalTransactions": 25,
	}

	// 2. Data Hardcode untuk Popular Fruits
	popularFruits := []gin.H{
		{
			"id":            1,
			"name":          "Apple",
			"imageUrl":      "https://via.placeholder.com/150/FFC0CB/000000?Text=Apple",
			"stockQuantity": 6,
			"status":        "Active",
			"sellingPrice":  100000,
			"costPrice":     80000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
			"total":         "6 kg", // Menambahkan field 'total' sesuai UI awal
			"inStock":       true,   // Menambahkan field 'inStock' sesuai UI awal
		},
		{
			"id":            2,
			"name":          "Watermelon",
			"imageUrl":      "https://via.placeholder.com/150/90EE90/000000?Text=Watermelon",
			"stockQuantity": 5,
			"status":        "Active",
			"sellingPrice":  100000,
			"costPrice":     75000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
			"total":         "5 kg",
			"inStock":       true,
		},
		{
			"id":            3,
			"name":          "Mango",
			"imageUrl":      "https://via.placeholder.com/150/FFD700/000000?Text=Mango",
			"stockQuantity": 0,
			"status":        "Inactive",
			"sellingPrice":  100000,
			"costPrice":     85000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
			"total":         "2 kg",
			"inStock":       false,
		},
		{
			"id":            4,
			"name":          "Pear",
			"imageUrl":      "https://via.placeholder.com/150/B0E0E6/000000?Text=Pear",
			"stockQuantity": 1,
			"status":        "Active",
			"sellingPrice":  100000,
			"costPrice":     90000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
			"total":         "1 kg",
			"inStock":       true,
		},
	}

	// 3. Data Hardcode untuk Popular Foods
	popularFoods := []gin.H{
		{
			"id":            5,
			"name":          "Beras Porang",
			"imageUrl":      "https://via.placeholder.com/150/D2B48C/000000?Text=Beras",
			"stockQuantity": 8,
			"status":        "Active",
			"sellingPrice":  100000,
			"costPrice":     70000,
			"category":      gin.H{"id": 2, "name": "Popular Foods"},
			"total":         "8 kg",
			"inStock":       true,
		},
		{
			"id":            6,
			"name":          "Beras Shirataki",
			"imageUrl":      "https://via.placeholder.com/150/E6E6FA/000000?Text=Beras",
			"stockQuantity": 6,
			"status":        "Active",
			"sellingPrice":  100000,
			"costPrice":     95000,
			"category":      gin.H{"id": 2, "name": "Popular Foods"},
			"total":         "6 kg",
			"inStock":       true,
		},
	}

	// 4. Data Hardcode untuk Overview Chart
	overviewData := []gin.H{
		{"name": "JAN", "sales": 4000, "revenue": 2400},
		{"name": "FEB", "sales": 3000, "revenue": 1398},
		{"name": "MAR", "sales": 2000, "revenue": 9800},
		{"name": "APR", "sales": 2780, "revenue": 3908},
		{"name": "MAY", "sales": 1890, "revenue": 4800},
		{"name": "JUN", "sales": 2390, "revenue": 3800},
		{"name": "JUL", "sales": 3490, "revenue": 4300},
		{"name": "AUG", "sales": 4200, "revenue": 5100},
		{"name": "SEP", "sales": 3100, "revenue": 3500},
		{"name": "OCT", "sales": 2500, "revenue": 6000},
		{"name": "NOV", "sales": 3800, "revenue": 4500},
		{"name": "DEC", "sales": 4300, "revenue": 5200},
	}

	// 5. Kirim respons JSON dengan struktur yang DIHARAPKAN oleh frontend
	c.JSON(http.StatusOK, gin.H{
		"stats":         dashboardStats,
		"popularFruits": popularFruits,
		"popularFoods":  popularFoods,
		"overview":      overviewData,
	})
}
