package mock

import (
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// CategoryUsecaseMock struct
type CategoryUsecaseMock struct {
}

// NewCategoryUsecaseMock function
func NewCategoryUsecaseMock() *CategoryUsecaseMock {
	return &CategoryUsecaseMock{}
}

// CreateCategory function
func (u *CategoryUsecaseMock) CreateCategory(category *domain.Category) shared.Output {
	categorySaved := &domain.Category{
		ID:   "1",
		Name: "Smart Phone",
	}
	return shared.Output{Result: categorySaved}
}

// GetCategory function
func (u *CategoryUsecaseMock) GetCategory(id string) shared.Output {
	categoryExist := &domain.Category{
		ID:   "1",
		Name: "Smart Phone",
	}

	return shared.Output{Result: categoryExist}
}
