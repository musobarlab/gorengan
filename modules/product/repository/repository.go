package repository

import (
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// ProductRepository interface
type ProductRepository interface {
	Save(*domain.Product) shared.Output
	Delete(*domain.Product) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}

// CategoryRepository interface
type CategoryRepository interface {
	Save(*domain.Category) shared.Output
	Delete(*domain.Category) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
