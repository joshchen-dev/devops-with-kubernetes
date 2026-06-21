package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dat, err := os.ReadFile(filepath.Join(os.TempDir(), "random_string_output"))
		check(err)
		fmt.Fprintln(w, string(dat))
	})

	http.ListenAndServe(":3000", mux)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
