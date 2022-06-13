package mock

import (
	categoryDomain "github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// ProductUsecaseMock struct
type ProductUsecaseMock struct {
}

// NewProductUsecaseMock function
func NewProductUsecaseMock() *ProductUsecaseMock {
	return &ProductUsecaseMock{}
}

// CreateProduct function
func (u *ProductUsecaseMock) CreateProduct(product *domain.Product) shared.Output {
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
	return shared.Output{Result: productSaved}
}

// RemoveProduct function
func (u *ProductUsecaseMock) RemoveProduct(id string) shared.Output {
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
	return shared.Output{Result: productSaved}
}

// GetProduct function
func (u *ProductUsecaseMock) GetProduct(id string) shared.Output {
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

	return shared.Output{Result: product}
}

// GetAllProduct function
func (u *ProductUsecaseMock) GetAllProduct(params *shared.Parameters) shared.Output {

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
	return shared.Output{Result: products}
}

// GetTotalProduct function
func (u *ProductUsecaseMock) GetTotalProduct(params *shared.Parameters) shared.Output {
	return shared.Output{Result: int64(2)}
}
