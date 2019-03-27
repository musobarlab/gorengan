package mock

import (
	"github.com/jinzhu/gorm"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
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
		Category: domain.Category{
			ID:   "1",
			Name: "Smart Phone",
		},
	}

	db["2"] = &domain.Product{
		ID:         "2",
		Name:       "Guitar Electric",
		Quantity:   6,
		CategoryID: "2",
		Category: domain.Category{
			ID:   "2",
			Name: "Music",
		},
	}
	return &ProductRepositoryMock{db: db}
}

// Save function
func (r *ProductRepositoryMock) Save(product *domain.Product) shared.Output {
	r.db[product.ID] = product
	return shared.Output{Result: product}
}

// Delete function
func (r *ProductRepositoryMock) Delete(product *domain.Product) shared.Output {
	product, ok := r.db[product.ID]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	delete(r.db, product.ID)
	return shared.Output{Result: product}
}

// FindByID function
func (r *ProductRepositoryMock) FindByID(id string) shared.Output {
	product, ok := r.db[id]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	return shared.Output{Result: product}
}

// FindAll function
func (r *ProductRepositoryMock) FindAll(params *shared.Parameters) shared.Output {
	var products domain.Products
	for _, product := range r.db {
		products = append(products, product)
	}
	return shared.Output{Result: products}
}

// Count function
func (r *ProductRepositoryMock) Count(params *shared.Parameters) shared.Output {
	return shared.Output{Result: len(r.db)}
}
