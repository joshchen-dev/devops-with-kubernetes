package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	c := 0
	f, err := os.Create(filepath.Join(os.TempDir(), "pingpong"))
	check(err)

	defer f.Close()

	fmt.Fprintln(f, c)

	mux := http.NewServeMux()

	mux.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		c++
		err = f.Truncate(0)
		check(err)
		fmt.Fprintln(w, "pong "+strconv.Itoa(c))
		fmt.Fprintln(f, c)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
