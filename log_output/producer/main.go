package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func main() {
	rs := uuid.NewString()
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	f, err := os.Create(filepath.Join(os.TempDir(), "random_string_output"))
	check(err)

	fmt.Fprintln(f, time.Now().Format(time.RFC3339Nano)+" "+rs)

	for range ticker.C {
		err = f.Truncate(0)
		check(err)
		fmt.Fprintln(f, time.Now().Format(time.RFC3339Nano)+" "+rs)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
