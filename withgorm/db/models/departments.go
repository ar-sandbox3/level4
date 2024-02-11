package models

import (
	"gorm.io/gorm"
)

// Department represents a department in the company.
type Department struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

// DepartmentStore is a data store for departments.
type DepartmentStore struct {
	DB *gorm.DB
}

// Create creates a new department.
func (s *DepartmentStore) Create(name string) error {
	return s.DB.Create(&Department{Name: name}).Error
}

// Get retrieves a department by its ID.
func (s *DepartmentStore) Get(id int64) (Department, error) {
	var d Department
	err := s.DB.First(&d, id).Error
	return d, err
}

// GetAll retrieves all departments.
func (s *DepartmentStore) GetAll() ([]Department, error) {
	var d []Department
	err := s.DB.Find(&d).Error
	return d, err
}
