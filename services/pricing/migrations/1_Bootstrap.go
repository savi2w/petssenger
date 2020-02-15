package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("Migrating pricing 1_Bootstrap...")

		_, err := db.Exec(`
			CREATE TABLE fees (
				id VARCHAR(28),
				base NUMERIC NOT NULL,
				distance NUMERIC NOT NULL,
				dynamic NUMERIC NOT NULL,
				minute NUMERIC NOT NULL,
				service NUMERIC NOT NULL,
				PRIMARY KEY(id)
			);

			INSERT INTO fees (id, base, distance, dynamic, minute, service) VALUES
				('CURITIBA', 2, 0.5, 1, 0.8, 1),
				('RIO_DE_JANEIRO', 3, 0.6, 1, 0.95, 1.4),
				('SALVADOR', 1.5, 0.2, 1, 0.75, 1.2),
				('SAO_PAULO', 3.5, 0.5, 1, 1, 0.75);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping pricing 1_Bootstrap...")

		_, err := db.Exec(("DROP TABLE fees;"))
		return err
	})
}
