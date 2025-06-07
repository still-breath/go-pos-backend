package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// Koneksi DB dan Model tidak diperlukan karena kita menggunakan data hardcode
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
			"id":           1,
			"name":         "Apple",
			"total":        "6 kg", // Field ini dibutuhkan oleh ProductListItem.tsx
			"imageUrl":     "https://images.unsplash.com/photo-1560806887-1e4cd0b69665?w=50&h=50&fit=crop",
			"inStock":      true, // Field ini dibutuhkan oleh ProductListItem.tsx
			"sellingPrice": 100000,
			// Field di bawah ini ada di 'types/index.ts' dan baik untuk disertakan agar konsisten
			"stockQuantity": 6,
			"status":        "Active",
			"costPrice":     80000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
		},
		{
			"id":            2,
			"name":          "Watermelon",
			"total":         "5 kg",
			"imageUrl":      "https://images.unsplash.com/photo-1589984662646-e7b2e4962f18?w=50&h=50&fit=crop",
			"inStock":       true,
			"sellingPrice":  100000,
			"stockQuantity": 5,
			"status":        "Active",
			"costPrice":     75000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
		},
		{
			"id":            3,
			"name":          "Mango",
			"total":         "2 kg",
			"imageUrl":      "https://images.unsplash.com/photo-1591073113125-e46713c829ed?w=50&h=50&fit=crop",
			"inStock":       false, // Contoh produk habis
			"sellingPrice":  100000,
			"stockQuantity": 0,
			"status":        "Inactive",
			"costPrice":     85000,
			"category":      gin.H{"id": 1, "name": "Popular Fruits"},
		},
	}

	// 3. Data Hardcode untuk Popular Foods
	popularFoods := []gin.H{
		{
			"id":            5,
			"name":          "Beras Porang",
			"total":         "8 kg",
			"imageUrl":      "https://images.unsplash.com/photo-1586201375765-c1265b014357?w=50&h=50&fit=crop",
			"inStock":       true,
			"sellingPrice":  100000,
			"stockQuantity": 8,
			"status":        "Active",
			"costPrice":     70000,
			"category":      gin.H{"id": 2, "name": "Popular Foods"},
		},
		{
			"id":            6,
			"name":          "Beras Shirataki",
			"total":         "6 kg",
			"imageUrl":      "https://images.unsplash.com/photo-1536304993881-ff6e9eefa2a7?w=50&h=50&fit=crop",
			"inStock":       true,
			"sellingPrice":  100000,
			"stockQuantity": 6,
			"status":        "Active",
			"costPrice":     95000,
			"category":      gin.H{"id": 2, "name": "Popular Foods"},
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
