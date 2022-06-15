package mock

import (
	categoryDomain "github.com/musobarlab/gorengan/internal/modules/category/v1/domain"
	"github.com/musobarlab/gorengan/internal/modules/product/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// ProductUsecaseMock struct
type ProductUsecaseMock struct {
}

// NewProductUsecaseMock function
func NewProductUsecaseMock() *ProductUsecaseMock {
	return &ProductUsecaseMock{}
}

// CreateProduct function
func (u *ProductUsecaseMock) CreateProduct(product *domain.Product) shared.Output[*domain.Product] {
	productSaved := &domain.Product{
		ID:         "1",
		Name:       "Nokia 6",
		Quantity:   6,
		CategoryID: "1",
		Category: categoryDomain.Category{
			ID:   "1",
			Name: "Smart Phone",
		},
	}
	return shared.Output[*domain.Product]{Result: productSaved}
}

// RemoveProduct function
func (u *ProductUsecaseMock) RemoveProduct(id string) shared.Output[*domain.Product] {
	productSaved := &domain.Product{
		ID:         "1",
		Name:       "Nokia 6",
		Quantity:   6,
		CategoryID: "1",
		Category: categoryDomain.Category{
			ID:   "1",
			Name: "Smart Phone",
		},
	}
	return shared.Output[*domain.Product]{Result: productSaved}
}

// GetProduct function
func (u *ProductUsecaseMock) GetProduct(id string) shared.Output[*domain.Product] {
	product := &domain.Product{
		ID:         "1",
		Name:       "Nokia 6",
		Quantity:   6,
		CategoryID: "1",
		Category: categoryDomain.Category{
			ID:   "1",
			Name: "Smart Phone",
		},
	}

	return shared.Output[*domain.Product]{Result: product}
}

// GetAllProduct function
func (u *ProductUsecaseMock) GetAllProduct(params *shared.Parameters) shared.Output[domain.Products] {

	products := domain.Products{
		&domain.Product{
			ID:         "1",
			Name:       "Nokia 6",
			Quantity:   6,
			CategoryID: "1",
			Category: categoryDomain.Category{
				ID:   "1",
				Name: "Smart Phone",
			},
		},
		&domain.Product{
			ID:         "2",
			Name:       "Guitar Electric",
			Quantity:   6,
			CategoryID: "2",
			Category: categoryDomain.Category{
				ID:   "2",
				Name: "Music",
			},
		},
	}
	return shared.Output[domain.Products]{Result: products}
}

// GetTotalProduct function
func (u *ProductUsecaseMock) GetTotalProduct(params *shared.Parameters) shared.Output[int64] {
	return shared.Output[int64]{Result: int64(2)}
}
