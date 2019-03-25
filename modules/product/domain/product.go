package domain

import (
	"errors"

	"github.com/musobarlab/gorengan/modules/shared"
)

// Product domain
type Product struct {
	ID         string `gorm:"column:ID; primary_key:yes"`
	Name       string `gorm:"column:NAME"`
	Quantity   uint   `gorm:"column:QUANTITY"`
	CategoryID string `gorm:"column:CATEGORY_ID"`
	Category   Category
	shared.BaseDomain
}

// TableName function
func (p Product) TableName() string {
	return "PRODUCTS"
}

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

// Products type list of Product
type Products []Product

// Categories type list of Category
type Categories []Category

// Validate function
func (p *Product) Validate() error {
	if len(p.ID) <= 0 {
		return errors.New("product id is required")
	}

	if len(p.Name) <= 0 {
		return errors.New("product name is required")
	}

	return nil
}
