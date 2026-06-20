package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	c := 0

	mux := http.NewServeMux()

	mux.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong "+strconv.Itoa(c))
		c++
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
