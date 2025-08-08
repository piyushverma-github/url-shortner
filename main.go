package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type URL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

var (
	urlMap = make(map[string]string)
	mutex  sync.Mutex
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req URL
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash := md5.Sum([]byte(req.Original))
	short := hex.EncodeToString(hash[:])[:6]

	mutex.Lock()
	urlMap[short] = req.Original
	mutex.Unlock()

	resp := URL{
		Original: req.Original,
		Short:    fmt.Sprintf("http://localhost:3000/%s", short),
	}
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	mutex.Lock()
	original, exists := urlMap[code]
	mutex.Unlock()

	if exists {
		http.Redirect(w, r, original, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/redirect", redirectHandler)

	fmt.Println("🚀 GO URL Shortener running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
