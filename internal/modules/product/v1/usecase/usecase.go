package usecase

import (
	"github.com/musobarlab/gorengan/internal/modules/product/v1/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// ProductUsecase interface
type ProductUsecase interface {
	CreateProduct(*domain.Product) shared.Output[*domain.Product]
	GetProduct(string) shared.Output[*domain.Product]
	GetAllProduct(*shared.Parameters) shared.Output[domain.Products]
	GetTotalProduct(*shared.Parameters) shared.Output[int64]
	RemoveProduct(string) shared.Output[*domain.Product]
}
