package usecase

import (
	"github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// CategoryUsecase interface
type CategoryUsecase interface {
	CreateCategory(*domain.Category) shared.Output
	GetCategory(string) shared.Output
}
