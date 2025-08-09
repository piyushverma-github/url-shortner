package main

import (
	"crypto/md5"      // Used to create a hash for generating the short URL
	"encoding/hex"    // Converts the MD5 hash bytes into a readable hex string
	"encoding/json"   // Handles JSON encoding/decoding for API requests/responses
	"fmt"             // For printing to console
	"net/http"        // For building the HTTP server
	"sync"            // Provides a mutex to handle concurrent map access
)

// URL struct to store original and shortened URLs
// It is also used to parse incoming JSON requests and send JSON responses
type URL struct {
	Original string `json:"original"` // Original long URL
	Short    string `json:"short"`    // Shortened URL
}

var (
	urlMap = make(map[string]string) // In-memory map to store shortCode -> originalURL
	mutex  sync.Mutex                 // Mutex to avoid race conditions when multiple requests modify the map
)

// shortenURLHandler handles POST requests to create shortened URLs
func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure only POST requests are accepted
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON request into the 'req' variable
	var req URL
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a 6-character short code from the MD5 hash of the original URL
	hash := md5.Sum([]byte(req.Original))
	short := hex.EncodeToString(hash[:])[:6]

	// Lock the map before writing (thread-safety)
	mutex.Lock()
	urlMap[short] = req.Original
	mutex.Unlock()

	// Create the response with original and short URL
	resp := URL{
		Original: req.Original,
		Short:    fmt.Sprintf("http://localhost:3000/%s", short),
	}

	// Send the response as JSON
	json.NewEncoder(w).Encode(resp)
}

// redirectHandler handles redirection when a short URL is accessed
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short code from the URL path (everything after '/')
	code := r.URL.Path[1:]

	// Lock map before reading (thread-safety)
	mutex.Lock()
	original, exists := urlMap[code]
	mutex.Unlock()

	// If the short code exists, redirect to the original URL
	if exists {
		http.Redirect(w, r, original, http.StatusFound)
	} else {
		// If not found, return 404
		http.NotFound(w, r)
	}
}

func main() {
	// Serve static files (like index.html) from the "static" folder at the root path
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API endpoint to shorten a URL
	http.HandleFunc("/shorten", shortenURLHandler)

	// API endpoint for redirecting short URLs
	// Note: This is currently bound to "/redirect", but real short URLs like /abc123
	// will also be served by this handler if routing is adjusted
	http.HandleFunc("/redirect", redirectHandler)

	// Start the server on port 3000
	fmt.Println("🚀 GO URL Shortener running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
