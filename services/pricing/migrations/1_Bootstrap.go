package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("Migrating 1_Bootstrap...")

		_, err := db.Exec(`
			CREATE TABLE fees (
				id VARCHAR(28),
				base INT NOT NULL,
				distance INT NOT NULL,
				dynamic INT NOT NULL,
				minute INT NOT NULL,
				service INT NOT NULL,
				PRIMARY KEY(id)
			);

			INSERT INTO fees (id, base, distance, dynamic, minute, service) VALUES
				('CURITIBA', 200, 50, 100, 80, 100),
				('RIO_DE_JANEIRO', 300, 60, 100, 95, 140),
				('SALVADOR', 150, 20, 100, 75, 120),
				('SAO_PAULO', 350, 50, 100, 100, 75);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping 1_Bootstrap...")

		_, err := db.Exec(("DROP TABLE fees;"))
		return err
	})
}
