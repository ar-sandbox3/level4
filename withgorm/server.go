package main

import (
	"log"
	"net/http"

	"github.com/ar-sandbox3/level4/withgorm/db"
	"github.com/ar-sandbox3/level4/withgorm/web"
)

func main() {
	// Open a database connection.
	conn, err := db.New(db.DefaultDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	http.ListenAndServe(":8080", web.Handler(conn))
}
