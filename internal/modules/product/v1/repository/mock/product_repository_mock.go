package mock

import (
	cd "github.com/musobarlab/gorengan/internal/modules/category/v1/domain"
	"github.com/musobarlab/gorengan/internal/modules/product/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
	"gorm.io/gorm"
)

// ProductRepositoryMock struct
type ProductRepositoryMock struct {
	db map[string]*domain.Product
}

// NewProductRepositoryMock function
func NewProductRepositoryMock() *ProductRepositoryMock {
	db := make(map[string]*domain.Product)
	db["1"] = &domain.Product{
		ID:         "1",
		Name:       "Nokia 6",
		Quantity:   6,
		CategoryID: "1",
		Category: cd.Category{
			ID:   "1",
			Name: "Smart Phone",
		},
	}

	db["2"] = &domain.Product{
		ID:         "2",
		Name:       "Guitar Electric",
		Quantity:   6,
		CategoryID: "2",
		Category: cd.Category{
			ID:   "2",
			Name: "Music",
		},
	}
	return &ProductRepositoryMock{db: db}
}

// Save function
func (r *ProductRepositoryMock) Save(product *domain.Product) shared.Output[*domain.Product] {
	r.db[product.ID] = product
	return shared.Output[*domain.Product]{Result: product}
}

// Delete function
func (r *ProductRepositoryMock) Delete(product *domain.Product) shared.Output[*domain.Product] {
	product, ok := r.db[product.ID]
	if !ok {
		return shared.Output[*domain.Product]{Err: gorm.ErrRecordNotFound}
	}

	delete(r.db, product.ID)
	return shared.Output[*domain.Product]{Result: product}
}

// FindByID function
func (r *ProductRepositoryMock) FindByID(id string) shared.Output[*domain.Product] {
	product, ok := r.db[id]
	if !ok {
		return shared.Output[*domain.Product]{Err: gorm.ErrRecordNotFound}
	}

	return shared.Output[*domain.Product]{Result: product}
}

// FindAll function
func (r *ProductRepositoryMock) FindAll(params *shared.Parameters) shared.Output[domain.Products] {
	var products domain.Products
	for _, product := range r.db {
		products = append(products, product)
	}
	return shared.Output[domain.Products]{Result: products}
}

// Count function
func (r *ProductRepositoryMock) Count(params *shared.Parameters) shared.Output[int64] {
	return shared.Output[int64]{Result: int64(len(r.db))}
}
