package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ner-service-go/internal/config"
	"ner-service-go/internal/ner"
)

func main() {
	cfg := config.Load()

	nerService, err := ner.NewService(cfg.ModelPath)
	if err != nil {
		log.Fatalf("Failed to initialize NER service: %v", err)
	}
	defer nerService.Close()

	r := gin.Default()

	r.GET("/health", handleHealth)
	r.POST("/ner", handleNER(nerService))

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleNER(nerService *ner.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ner.ExtractRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		if req.Text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text field is required"})
			return
		}

		entities, err := nerService.ExtractEntities(req.Text)
		if err != nil {
			log.Printf("Error extracting entities: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract entities"})
			return
		}

		c.JSON(http.StatusOK, entities)
	}
}

func handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "ner-service-go",
	})
}