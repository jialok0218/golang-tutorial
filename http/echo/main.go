package main

import (
	"net/http"
	"unicode"

	"github.com/labstack/echo/v4"
)

// isMandarin checks if the text contains only Mandarin characters
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
	// Create a new Echo instance
	e := echo.New()

	// Define the /check endpoint
	e.GET("/check", func(c echo.Context) error {

		text := c.QueryParam("text")

		// Check if the "text" parameter is missing
		if text == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Error: 'text' query parameter is required",
			})
		}

		// Check if the text contains only Mandarin characters
		if isMandarin(text) {
			return c.JSON(http.StatusOK, map[string]string{
				"message": "ok",
			})
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "only mandarin allowed",
			})
		}
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
