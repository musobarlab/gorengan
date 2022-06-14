package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/product/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// ProductRepository interface
type ProductRepository interface {
	Save(*domain.Product) shared.Output
	Delete(*domain.Product) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
