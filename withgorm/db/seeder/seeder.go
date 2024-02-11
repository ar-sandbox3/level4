package seeder

import (
	"errors"
	"fmt"

	"github.com/ar-sandbox3/level4/withgorm/db/models"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {
	// Migrate the schema. This is an example for departments.
	departments := &models.Department{}
	if db.Migrator().HasTable(departments) {
		db.Migrator().DropTable(departments)
	}
	if err := db.AutoMigrate(departments); err != nil {
		return fmt.Errorf("failed to migrate departments: %w", err)
	}

	// Validate the table creation.
	if !db.Migrator().HasTable(departments) {
		return errors.New("failed to create departments table")
	}

	// When we have the table, we can create some records (seed data).
	if db.Migrator().HasTable(departments) {
		if err := db.Model(departments).CreateInBatches([]models.Department{
			{
				Name: "Engineering",
			},
			{
				Name: "IT",
			},
			{
				Name: "Customer Service",
			},
		}, 10).Error; err != nil {
			return fmt.Errorf("failed to seed departments: %w", err)
		}
	}
	return nil
}
