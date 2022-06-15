package domain

import (
	"errors"

	"github.com/musobarlab/gorengan/pkg/shared"
)

// Category struct
type Category struct {
	ID   string `gorm:"column:ID; primary_key:yes"`
	Name string `gorm:"column:NAME"`
	shared.BaseDomain
}

// TableName function
func (c Category) TableName() string {
	return "PRODUCT_CATEGORIES"
}

// Categories type list of Category
type Categories []*Category

// Validate function
func (c *Category) Validate() error {
	if len(c.ID) <= 0 {
		return errors.New("category id is required")
	}

	if len(c.Name) <= 0 {
		return errors.New("category name is required")
	}

	return nil
}
