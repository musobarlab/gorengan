package mock

import (
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// CategoryUsecaseMock struct
type CategoryUsecaseMock struct {
}

// NewCategoryUsecaseMock function
func NewCategoryUsecaseMock() *CategoryUsecaseMock {
	return &CategoryUsecaseMock{}
}

// CreateCategory function
func (u *CategoryUsecaseMock) CreateCategory(category *domain.Category) shared.Output[*domain.Category] {
	categorySaved := &domain.Category{
		ID:   "1",
		Name: "Smart Phone",
	}
	return shared.Output[*domain.Category]{Result: categorySaved}
}

// GetCategory function
func (u *CategoryUsecaseMock) GetCategory(id string) shared.Output[*domain.Category] {
	categoryExist := &domain.Category{
		ID:   "1",
		Name: "Smart Phone",
	}

	return shared.Output[*domain.Category]{Result: categoryExist}
}
