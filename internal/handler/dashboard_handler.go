package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/still-breath/go-pos-backend.git/internal/config" // Ganti path
	"github.com/still-breath/go-pos-backend.git/internal/model"  // Ganti dengan path modul Anda
)

func GetDashboard(c *gin.Context) {
	// Untuk saat ini, kita gunakan data dummy untuk statistik
	dashboardStats := gin.H{
		"dailySales":        500000,
		"monthlyRevenue":    50000000,
		"totalTransactions": 25,
	}

	// Ambil produk dan kategori
	var categories []model.Category
	config.DB.Preload("Products").Order("random()").Limit(2).Find(&categories)

	// Format data produk
	var popularProducts []gin.H
	for _, category := range categories {
		var products []gin.H
		productLimit := 5
		for i, product := range category.Products {
			if i >= productLimit {
				break
			}
			products = append(products, gin.H{
				"id":            product.ID,
				"name":          product.Name,
				"imageUrl":      product.ImageURL,
				"stockQuantity": product.StockQuantity,
				"status":        product.Status,
				"sellingPrice":  product.SellingPrice,
				"costPrice":     product.HPP,
			})
		}
		popularProducts = append(popularProducts, gin.H{
			"categoryName": category.Name,
			"products":     products,
		})
	}

	// Data dummy untuk chart
	overviewData := []gin.H{
		{"name": "JAN", "sales": 4000, "revenue": 2400},
		{"name": "FEB", "sales": 3000, "revenue": 1398},
		// ... tambahkan data bulan lain jika perlu
	}

	c.JSON(http.StatusOK, gin.H{
		"dashboardStats":  dashboardStats,
		"popularProducts": popularProducts,
		"overviewData":    overviewData,
	})
}
