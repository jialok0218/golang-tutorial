package main

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

// type Response struct {
// 	Message string `json:"message"`
// }

func isMandarin(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			continue
		}
		return false
	}
	return true
}

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define the /check endpoint
	r.GET("/check", func(c *gin.Context) {
		text := c.Query("text")

		// Check if the "text" parameter is missing
		if text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error: 'text' query parameter is required"})
			return
		}

		// Check if the text contains only Mandarin characters
		if isMandarin(text) {
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "only mandarin allowed"})
		}
	})

	// Start the server on port 8080
	r.Run(":8080")
}
