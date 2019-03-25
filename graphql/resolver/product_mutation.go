package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"golang.org/x/net/context"
)

// ProductInputArgs input
type ProductInputArgs struct {
	Product ProductInput
}

// DeleteProductArgs input
type DeleteProductArgs struct {
	ID graphql.ID
}

// CreateProduct mutation
func (r *Resolver) CreateProduct(ctx context.Context, args *ProductInputArgs) (*ProductResolver, error) {
	var product domain.Product
	product.ID = args.Product.ID
	product.Name = args.Product.Name
	product.Quantity = uint(args.Product.Quantity)
	product.CategoryID = args.Product.Category
	product.Created = time.Now()
	product.LastModified = time.Now()

	output := r.ProductUsecase.CreateProduct(&product)
	if output.Err != nil {
		return nil, output.Err
	}

	productSaved := output.Result.(*domain.Product)

	return &ProductResolver{productSaved}, nil
}

// DeleteProduct mutation
func (r *Resolver) DeleteProduct(ctx context.Context, args *DeleteProductArgs) (*ProductResolver, error) {
	return nil, nil
}
