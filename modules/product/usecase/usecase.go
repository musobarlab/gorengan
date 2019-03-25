package usecase

import (
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// ProductUsecase interface
type ProductUsecase interface {
	CreateProduct(*domain.Product) shared.Output
	GetProduct(string) shared.Output
	GetAllProduct(*shared.Parameters) shared.Output
	GetTotalProduct(*shared.Parameters) shared.Output
}
