package main

import (
	"flag"
	"log"

	"github.com/go-pg/migrations/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
)

func main() {
	flag.Parse()
	conn := config.PostgresConnect()
	_, _, err := migrations.Run(conn, flag.Args()...)
	if err != nil {
		log.Fatalf("Error when migrating: %v", err)
	}
}
