package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("Migrating user 1_Bootstrap...")

		_, err := db.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
			CREATE TABLE user (
				id uuid NOT NULL DEFAULT uuid_generate_v4(),
				name VARCHAR NOT NULL,
				created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping user 1_Bootstrap...")
		_, err := db.Exec("DROP TABLE user;")
		return err
	})
}
