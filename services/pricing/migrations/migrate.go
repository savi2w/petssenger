package main

import (
	"flag"
	"log"

	"github.com/go-pg/migrations/v7"
	"github.com/weslenng/petssenger/services/pricing/postgres"
)

func main() {
	flag.Parse()
	conn := postgres.Connect()
	_, _, err := migrations.Run(conn, flag.Args()...)
	if err != nil {
		log.Fatalf("Error when migrating: %v", err)
	}
}
