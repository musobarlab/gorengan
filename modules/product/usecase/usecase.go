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
	RemoveProduct(string) shared.Output
}

// CategoryUsecase interface
type CategoryUsecase interface {
	CreateCategory(*domain.Category) shared.Output
	GetCategory(string) shared.Output
}
