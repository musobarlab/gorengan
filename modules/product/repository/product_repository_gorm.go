package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// ProductRepositoryGorm struct
type ProductRepositoryGorm struct {
	db *gorm.DB
}

// NewProductRepositoryGorm function
func NewProductRepositoryGorm(db *gorm.DB) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{db: db}
}

// Save function
func (r *ProductRepositoryGorm) Save(product *domain.Product) shared.Output {
	err := r.db.Save(product).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: product}
}

// Delete function
func (r *ProductRepositoryGorm) Delete(product *domain.Product) shared.Output {
	err := r.db.Delete(product).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: product}
}

// FindByID function
func (r *ProductRepositoryGorm) FindByID(id string) shared.Output {
	var product domain.Product

	err := r.db.Where(&domain.Product{ID: id}).Take(&product).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: &product}
}

// FindAll function
func (r *ProductRepositoryGorm) FindAll(params *shared.Parameters) shared.Output {
	var products domain.Products

	db := r.db.Preload("Category").Offset(params.Offset).Limit(params.Limit)

	if len(params.OrderBy) > 0 {
		db = db.Order(params.OrderBy)
	}

	err := db.Find(&products).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: products}
}

// Count function
func (r *ProductRepositoryGorm) Count(params *shared.Parameters) shared.Output {
	var count int

	db := r.db.Model(&domain.Product{})

	if len(params.OrderBy) > 0 {
		db = db.Order(params.OrderBy)
	}

	err := db.Count(&count).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: count}
}
