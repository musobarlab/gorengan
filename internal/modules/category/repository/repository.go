package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// CategoryRepository interface
type CategoryRepository interface {
	Save(*domain.Category) shared.Output
	Delete(*domain.Category) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
