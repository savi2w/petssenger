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
			CREATE TABLE users (
				id uuid NOT NULL DEFAULT uuid_generate_v4(),
				email VARCHAR NOT NULL UNIQUE,
				created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY(id)
			);

			INSERT INTO users (id, email) VALUES ('08842beb-a4fc-4cb2-9f87-d80f1a2d5045', 'test@petssenger.com');
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping user 1_Bootstrap...")
		_, err := db.Exec("DROP TABLE users;")
		return err
	})
}
