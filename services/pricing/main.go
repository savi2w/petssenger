package main

import (
	"log"

	"github.com/weslenng/petssenger/services/pricing/server"
)

func main() {
	_, err := server.Listen()
	if err != nil {
		log.Fatalf("Error when listening: %v", err)
	}
}
