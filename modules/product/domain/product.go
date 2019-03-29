package domain

import (
	"errors"

	categoryDomain "github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// Product domain
type Product struct {
	ID         string `gorm:"column:ID; primary_key:yes"`
	Name       string `gorm:"column:NAME"`
	Quantity   uint   `gorm:"column:QUANTITY"`
	CategoryID string `gorm:"column:CATEGORY_ID"`
	Category   categoryDomain.Category
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

// Products type list of Product
type Products []*Product

// Validate function
func (p *Product) Validate() error {
	if len(p.ID) <= 0 {
		return errors.New("product id is required")
	}

	if len(p.Name) <= 0 {
		return errors.New("product name is required")
	}

	if len(p.CategoryID) <= 0 {
		return errors.New("category id is required")
	}

	return nil
}
