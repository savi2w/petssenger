package main

import (
	"flag"

	"github.com/go-pg/migrations/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
)

func main() {
	flag.Parse()
	conn := config.PricingPostgresConnect()
	_, _, err := migrations.Run(conn, flag.Args()...)
	if err != nil {
		panic("Error when migrating")
	}
}
