package web

import (
	"log"
	"net/http"
	"path/filepath"
)

// ListenAndServe starts a static server on the given port
// root should be the path to the web directory
func ListenAndServe(addr string, root string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(root, "index.html")
		http.ServeFile(w, r, path)
	})
	http.HandleFunc("/7s.js", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(root, "7s.js")
		http.ServeFile(w, r, path)
	})
	log.Printf("Listening on %s", addr)
	return http.ListenAndServe(addr, nil)
}
