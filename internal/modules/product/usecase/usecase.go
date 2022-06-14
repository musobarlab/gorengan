package usecase

import (
	"github.com/musobarlab/gorengan/internal/modules/product/domain"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// ProductUsecase interface
type ProductUsecase interface {
	CreateProduct(*domain.Product) shared.Output
	GetProduct(string) shared.Output
	GetAllProduct(*shared.Parameters) shared.Output
	GetTotalProduct(*shared.Parameters) shared.Output
	RemoveProduct(string) shared.Output
}
