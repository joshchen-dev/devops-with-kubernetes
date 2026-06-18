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
	fmt.Printf("Server started in port %s\n", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
