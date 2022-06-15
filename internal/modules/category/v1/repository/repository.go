package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/category/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// CategoryRepository interface
type CategoryRepository interface {
	Save(*domain.Category) shared.Output[*domain.Category]
	Delete(*domain.Category) shared.Output[*domain.Category]
	FindByID(string) shared.Output[*domain.Category]
	FindAll(*shared.Parameters) shared.Output[domain.Categories]
	Count(*shared.Parameters) shared.Output[int64]
}
