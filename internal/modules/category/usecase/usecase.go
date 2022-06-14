package usecase

import (
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// CategoryUsecase interface
type CategoryUsecase interface {
	CreateCategory(*domain.Category) shared.Output
	GetCategory(string) shared.Output
}
