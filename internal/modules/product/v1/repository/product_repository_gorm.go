package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/product/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
	"gorm.io/gorm"
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
func (r *ProductRepositoryGorm) Save(product *domain.Product) shared.Output[*domain.Product] {
	err := r.db.Save(product).Error
	if err != nil {
		return shared.Output[*domain.Product]{Err: err}
	}

	return shared.Output[*domain.Product]{Result: product}
}

// Delete function
func (r *ProductRepositoryGorm) Delete(product *domain.Product) shared.Output[*domain.Product] {
	err := r.db.Delete(product).Error
	if err != nil {
		return shared.Output[*domain.Product]{Err: err}
	}

	return shared.Output[*domain.Product]{Result: product}
}

// FindByID function
func (r *ProductRepositoryGorm) FindByID(id string) shared.Output[*domain.Product] {
	var product domain.Product

	err := r.db.Preload("Category").Where(&domain.Product{ID: id}).Take(&product).Error
	if err != nil {
		return shared.Output[*domain.Product]{Err: err}
	}

	return shared.Output[*domain.Product]{Result: &product}
}

// FindAll function
func (r *ProductRepositoryGorm) FindAll(params *shared.Parameters) shared.Output[domain.Products] {
	var products domain.Products

	db := r.db.Preload("Category").Offset(params.Offset).Limit(params.Limit)

	if len(params.OrderBy) > 0 {
		db = db.Order(params.OrderBy)
	}

	err := db.Find(&products).Error
	if err != nil {
		return shared.Output[domain.Products]{Err: err}
	}

	return shared.Output[domain.Products]{Result: products}
}

// Count function
func (r *ProductRepositoryGorm) Count(params *shared.Parameters) shared.Output[int64] {
	var count int64

	db := r.db.Model(&domain.Product{})

	if len(params.OrderBy) > 0 {
		db = db.Order(params.OrderBy)
	}

	err := db.Count(&count).Error
	if err != nil {
		return shared.Output[int64]{Err: err}
	}

	return shared.Output[int64]{Result: count}
}
