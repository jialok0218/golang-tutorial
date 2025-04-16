package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode"
)

type Response struct {
	Message string `json:"message"`
}

func isMandarin(text string) bool {
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			continue
		}
		return false
	}
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	w.Header().Set("Content-Type", "application/json")

	// Check if the "text" parameter is missing
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Error: 'text' query parameter is required"})
		return
	}

	// Check if the text contains only Mandarin characters
	if isMandarin(text) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: "ok"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "only mandarin allowed"})
	}
}

func main() {
	fmt.Println("Starting server on http://localhost:8080")
	http.HandleFunc("/check", handler)
	http.ListenAndServe(":8080", nil)
}
