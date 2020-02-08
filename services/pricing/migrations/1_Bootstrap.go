package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("Migrating 1_Bootstrap...")

		_, err := db.Exec(`
			CREATE TABLE city (
				id VARCHAR(32),
				base_fees SMALLINT NOT NULL,
				distance_fees SMALLINT NOT NULL,
				dynamic_fees SMALLINT NOT NULL,
				minute_fees SMALLINT NOT NULL,
				service_fees SMALLINT NOT NULL,
				PRIMARY KEY(id)
			);

			INSERT INTO city (id, base_fees, distance_fees, dynamic_fees, minute_fees, service_fees) VALUES
				('CURITIBA', 200, 50, 100, 80, 100),
				('RIO_DE_JANEIRO', 300, 60, 100, 95, 140),
				('SALVADOR', 150, 20, 100, 75, 120),
				('SAO_PAULO', 350, 50, 100, 100, 75);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping 1_Bootstrap...")

		_, err := db.Exec(("DROP TABLE city;"))
		return err
	})
}
