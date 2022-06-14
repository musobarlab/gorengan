package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
	"gorm.io/gorm"
)

// CategoryRepositoryGorm struct
type CategoryRepositoryGorm struct {
	db *gorm.DB
}

// NewCategoryRepositoryGorm function
func NewCategoryRepositoryGorm(db *gorm.DB) *CategoryRepositoryGorm {
	return &CategoryRepositoryGorm{db: db}
}

// Save function
func (r *CategoryRepositoryGorm) Save(category *domain.Category) shared.Output[*domain.Category] {
	err := r.db.Save(category).Error
	if err != nil {
		return shared.Output[*domain.Category]{Err: err}
	}

	return shared.Output[*domain.Category]{Result: category}
}

// Delete function
func (r *CategoryRepositoryGorm) Delete(category *domain.Category) shared.Output[*domain.Category] {
	err := r.db.Delete(category).Error
	if err != nil {
		return shared.Output[*domain.Category]{Err: err}
	}

	return shared.Output[*domain.Category]{Result: category}
}

// FindByID function
func (r *CategoryRepositoryGorm) FindByID(id string) shared.Output[*domain.Category] {
	var category domain.Category

	err := r.db.Where(&domain.Category{ID: id}).Take(&category).Error
	if err != nil {
		return shared.Output[*domain.Category]{Err: err}
	}

	return shared.Output[*domain.Category]{Result: &category}
}

// FindAll function
func (r *CategoryRepositoryGorm) FindAll(params *shared.Parameters) shared.Output[domain.Categories] {
	var categories domain.Categories

	err := r.db.Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Find(&categories).Error
	if err != nil {
		return shared.Output[domain.Categories]{Err: err}
	}

	return shared.Output[domain.Categories]{Result: categories}
}

// Count function
func (r *CategoryRepositoryGorm) Count(params *shared.Parameters) shared.Output[int64] {
	var count int64
	err := r.db.Model(&domain.Category{}).Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Count(&count).Error
	if err != nil {
		return shared.Output[int64]{Err: err}
	}

	return shared.Output[int64]{Result: count}
}
