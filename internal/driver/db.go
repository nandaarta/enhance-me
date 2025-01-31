package driver

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() {
	dbURL := "postgres://admin:admin@localhost/test_repo?sslmode=disable"

	db, errDb := sql.Open("postgres", dbURL)
	if errDb != nil {
		log.Fatal(errDb)
	}
	defer db.Close()

	driver, errDrv := postgres.WithInstance(db, &postgres.Config{})
	if errDrv != nil {
		log.Fatal(errDrv)
	}

	mig, errMig := migrate.NewWithDatabaseInstance("file://internal/migrations", "postgres", driver)
	if errMig != nil {
		log.Fatal(errMig)
	}

	err := mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migration completed.")
}
