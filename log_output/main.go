package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	rs := uuid.NewString()
	log.Println(rs)
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	go reportRandomString(rs, ticker)

	mux := http.NewServeMux()

	mux.HandleFunc("/now", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().Format(time.RFC3339Nano)+" "+rs)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func reportRandomString(rs string, t *time.Ticker) {
	for range t.C {
		log.Println(rs)
	}
}
