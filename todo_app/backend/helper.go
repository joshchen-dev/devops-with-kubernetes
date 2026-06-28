package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func updateImage(name string) {
	imagePath := filepath.Join(os.TempDir(), name)

	f, err := os.Create(imagePath)
	check(err)
	defer f.Close()

	r, err := http.Get("https://picsum.photos/1200")
	check(err)
	defer r.Body.Close()

	_, err = io.Copy(f, r.Body)
	check(err)
}

// Check if image has been stale for ${n} minutes. Assume image exists.
func imageStale(name string, n int) bool {
	imagePath := filepath.Join(os.TempDir(), name)

	fileinfo, err := os.Stat(imagePath)
	check(err)

	return fileinfo.ModTime().Before(time.Now().Add(-time.Duration(n) * time.Minute))
}

func imageExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		panic(err)
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
