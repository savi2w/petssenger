package models

import (
	"github.com/go-pg/pg/v9"
	"github.com/savi2w/petssenger/services/pricing/config"
)

var db *pg.DB

func InitDB() *pg.DB {
	db = pg.Connect(&config.Default.PgConnOpts)
	return db
}
