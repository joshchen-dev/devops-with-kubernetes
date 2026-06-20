package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "static/index.html")
		}
	})

	fmt.Printf("Server started in port %s\n", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
