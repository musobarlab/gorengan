package repository

import (
	"github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// CategoryRepository interface
type CategoryRepository interface {
	Save(*domain.Category) shared.Output
	Delete(*domain.Category) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
