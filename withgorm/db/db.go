package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DefaultDSN is the default data source name for the database.
// For example you can run a local MySQL server with Docker:
//
//	docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --name mysql mysql:8
const DefaultDSN = "root:password@tcp(localhost:3306)/level3?multiStatements=true&parseTime=true"

// New returns a new database connection.
func New(dsn string) (*gorm.DB, error) {
	// Open a database connection.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
