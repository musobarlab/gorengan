package resolver

import (
	"github.com/musobarlab/gorengan/modules/product/usecase"
)

// Resolver for product module
type Resolver struct {
	ProductUsecase usecase.ProductUsecase
}
