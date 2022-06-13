package mock

import (
	"github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/shared"
	"gorm.io/gorm"
)

// CategoryRepositoryMock struct
type CategoryRepositoryMock struct {
	db map[string]*domain.Category
}

// NewCategoryRepositoryMock function
func NewCategoryRepositoryMock() *CategoryRepositoryMock {
	db := make(map[string]*domain.Category)
	db["1"] = &domain.Category{
		ID:   "1",
		Name: "Smart Phone",
	}

	db["2"] = &domain.Category{
		ID:   "2",
		Name: "Music",
	}
	return &CategoryRepositoryMock{db: db}
}

// Save function
func (r *CategoryRepositoryMock) Save(category *domain.Category) shared.Output {
	r.db[category.ID] = category
	return shared.Output{Result: category}
}

// Delete function
func (r *CategoryRepositoryMock) Delete(category *domain.Category) shared.Output {
	category, ok := r.db[category.ID]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	delete(r.db, category.ID)

	return shared.Output{Result: category}
}

// FindByID function
func (r *CategoryRepositoryMock) FindByID(id string) shared.Output {
	category, ok := r.db[id]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	return shared.Output{Result: category}
}

// FindAll function
func (r *CategoryRepositoryMock) FindAll(params *shared.Parameters) shared.Output {
	var categories domain.Categories

	for _, category := range categories {
		categories = append(categories, category)
	}

	return shared.Output{Result: categories}
}

// Count function
func (r *CategoryRepositoryMock) Count(params *shared.Parameters) shared.Output {
	return shared.Output{Result: len(r.db)}
}
