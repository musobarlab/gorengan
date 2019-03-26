package usecase

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/product/repository"
	"github.com/musobarlab/gorengan/modules/shared"
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
func (u *CategoryUsecaseImpl) CreateCategory(category *domain.Category) shared.Output {
	categoryOutput := u.categoryRepositoryRead.FindByID(category.ID)
	if categoryOutput.Err != nil && categoryOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: categoryOutput.Err}
	}

	if categoryOutput.Result != nil {
		categoryExist := categoryOutput.Result.(*domain.Category)

		if categoryExist != nil {
			return shared.Output{Err: fmt.Errorf("category with id %s already exist", category.ID)}
		}
	}

	err := category.Validate()

	if err != nil {
		return shared.Output{Err: err}
	}

	categorySaveOutput := u.categoryRepositoryWrite.Save(category)
	if categorySaveOutput.Err != nil {
		return shared.Output{Err: categorySaveOutput.Err}
	}

	categorySaved := categorySaveOutput.Result.(*domain.Category)

	return shared.Output{Result: categorySaved}
}

// GetCategory function
func (u *CategoryUsecaseImpl) GetCategory(id string) shared.Output {
	categoryOutput := u.categoryRepositoryRead.FindByID(id)
	if categoryOutput.Err != nil && categoryOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: categoryOutput.Err}
	}

	categoryExist := categoryOutput.Result.(*domain.Category)

	return shared.Output{Result: categoryExist}
}
