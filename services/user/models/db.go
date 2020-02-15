package models

import (
	"github.com/go-pg/pg/v9"
	"github.com/weslenng/petssenger/services/user/config"
)

var db *pg.DB

// InitDB initialize a global postgres connection in a models context
func InitDB() *pg.DB {
	db = pg.Connect(&config.Default.PgConnOpts)
	return db
}
