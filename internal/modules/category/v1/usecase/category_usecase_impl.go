package usecase

import (
	"fmt"

	"github.com/musobarlab/gorengan/internal/modules/category/v1/domain"
	"github.com/musobarlab/gorengan/internal/modules/category/v1/repository"
	"github.com/musobarlab/gorengan/pkg/shared"
	"gorm.io/gorm"
)

// CategoryUsecaseImpl struct
type CategoryUsecaseImpl struct {
	categoryRepositoryRead  repository.CategoryRepository
	categoryRepositoryWrite repository.CategoryRepository
}

// NewCategoryUsecaseImpl function
func NewCategoryUsecaseImpl(categoryRepositoryRead, categoryRepositoryWrite repository.CategoryRepository) *CategoryUsecaseImpl {
	return &CategoryUsecaseImpl{categoryRepositoryRead: categoryRepositoryRead, categoryRepositoryWrite: categoryRepositoryWrite}
}

// CreateCategory function
func (u *CategoryUsecaseImpl) CreateCategory(category *domain.Category) shared.Output[*domain.Category] {
	categoryOutput := u.categoryRepositoryRead.FindByID(category.ID)
	if categoryOutput.Err != nil && categoryOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output[*domain.Category]{Err: categoryOutput.Err}
	}

	if categoryOutput.Result != nil {
		categoryExist := categoryOutput.Result

		if categoryExist != nil {
			return shared.Output[*domain.Category]{Err: fmt.Errorf("category with id %s already exist", category.ID)}
		}
	}

	err := category.Validate()

	if err != nil {
		return shared.Output[*domain.Category]{Err: err}
	}

	categorySaveOutput := u.categoryRepositoryWrite.Save(category)
	if categorySaveOutput.Err != nil {
		return shared.Output[*domain.Category]{Err: categorySaveOutput.Err}
	}

	categorySaved := categorySaveOutput.Result

	return shared.Output[*domain.Category]{Result: categorySaved}
}

// GetCategory function
func (u *CategoryUsecaseImpl) GetCategory(id string) shared.Output[*domain.Category] {
	categoryOutput := u.categoryRepositoryRead.FindByID(id)
	if categoryOutput.Err != nil && categoryOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output[*domain.Category]{Err: categoryOutput.Err}
	}

	categoryExist := categoryOutput.Result

	return shared.Output[*domain.Category]{Result: categoryExist}
}
