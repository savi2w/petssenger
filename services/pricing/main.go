package main

import (
	"log"

	"github.com/weslenng/petssenger/services/pricing/config"
)

func main() {
	_, err := config.ServerListen()
	if err != nil {
		log.Fatalf("Error when listening: %v", err)
	}
}
