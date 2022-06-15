package repository

import (
	"github.com/musobarlab/gorengan/internal/modules/product/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// ProductRepository interface
type ProductRepository interface {
	Save(*domain.Product) shared.Output[*domain.Product]
	Delete(*domain.Product) shared.Output[*domain.Product]
	FindByID(string) shared.Output[*domain.Product]
	FindAll(*shared.Parameters) shared.Output[domain.Products]
	Count(*shared.Parameters) shared.Output[int64]
}
