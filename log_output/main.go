package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	rs := uuid.NewString()
	log.Println(rs)
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		log.Println(rs)
	}
}
