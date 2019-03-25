package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

// CategoryInputArgs input
type CategoryInputArgs struct {
	Category CategoryInput
}

// ProductInputArgs input
type ProductInputArgs struct {
	Product ProductInput
}

// DeleteProductArgs input
type DeleteProductArgs struct {
	ID graphql.ID
}

// CreateCategory mutation
func (r *Resolver) CreateCategory(ctx context.Context, args *CategoryInputArgs) (*CategoryResolver, error) {
	return nil, nil
}

// CreateProduct mutation
func (r *Resolver) CreateProduct(ctx context.Context, args *ProductInputArgs) (*ProductResolver, error) {
	return nil, nil
}

// DeleteProduct mutation
func (r *Resolver) DeleteProduct(ctx context.Context, args *DeleteProductArgs) (*ProductResolver, error) {
	return nil, nil
}
