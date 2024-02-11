package main

import (
	"log"

	"github.com/ar-sandbox3/level4/withgorm/db"
	"github.com/ar-sandbox3/level4/withgorm/db/seeder"
)

func main() {
	conn, err := db.New(db.DefaultDSN)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := seeder.MigrateAndSeed(conn); err != nil {
		log.Fatalf("failed to migrate and seed the database: %v", err)
	}
}
