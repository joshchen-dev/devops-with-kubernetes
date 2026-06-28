package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	PORT := "8080"
	mux := http.NewServeMux()

	if !imageExist(imageName) || imageStale(imageName, cacheMinutes) {
		updateImage(imageName)
	}

	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(os.TempDir(), imageName)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Add("Cache-Control", "no-cache")
		http.ServeFile(w, r, path)

		if imageStale(imageName, cacheMinutes) {
			updateImage(imageName)
		}
	})

	fmt.Printf("Server started in port %s\n", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
