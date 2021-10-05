package main

import (
	"flag"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	"github.com/nglabo/petssenger/services/pricing/config"
)

func main() {
	flag.Parse()
	db := pg.Connect(&config.Default.PgConnOpts)
	_, _, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		panic(err)
	}
}
