package usecase

import (
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// CategoryUsecase interface
type CategoryUsecase interface {
	CreateCategory(*domain.Category) shared.Output[*domain.Category]
	GetCategory(string) shared.Output[*domain.Category]
}
