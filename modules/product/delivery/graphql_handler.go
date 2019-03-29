package delivery

import (
	"math"
	"strconv"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/product/graphql/schema"
	"github.com/musobarlab/gorengan/modules/product/usecase"
	"github.com/musobarlab/gorengan/modules/shared"
	"golang.org/x/net/context"
)

// GraphQLProductHandler struct
// Handler means Resolver
type GraphQLProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

// ProductInputArgs input
type ProductInputArgs struct {
	Product schema.ProductSchemaInput
}

// DeleteProductArgs input
type DeleteProductArgs struct {
	ID graphql.ID
}

// ProductQueryArgs args
type ProductQueryArgs struct {
	ID graphql.ID
}

// ProductsArgs struct
type ProductsArgs struct {
	Query   *string
	Limit   *float64
	Page    *float64
	OrderBy *string
	Sort    *string
}

// ProductsQueryArgs args
type ProductsQueryArgs struct {
	ProductsArgs *ProductsArgs
}

// CreateProduct mutation
func (r *GraphQLProductHandler) CreateProduct(ctx context.Context, args *ProductInputArgs) (*schema.ProductSchema, error) {
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

	return &schema.ProductSchema{Product: productSaved}, nil
}

// DeleteProduct mutation
func (r *GraphQLProductHandler) DeleteProduct(ctx context.Context, args *DeleteProductArgs) (*schema.ProductSchema, error) {
	output := r.ProductUsecase.RemoveProduct(string(args.ID))

	if output.Err != nil {
		return nil, output.Err
	}

	product := output.Result.(*domain.Product)

	return &schema.ProductSchema{Product: product}, nil
}

// Product Query function
func (r *GraphQLProductHandler) Product(ctx context.Context, args *ProductQueryArgs) (*schema.ProductSchema, error) {
	output := r.ProductUsecase.GetProduct(string(args.ID))

	if output.Err != nil {
		return nil, output.Err
	}

	product := output.Result.(*domain.Product)

	return &schema.ProductSchema{Product: product}, nil

}

// Products Query function
func (r *GraphQLProductHandler) Products(ctx context.Context, args *ProductsQueryArgs) (*schema.ProductListResolver, error) {
	var (
		params           shared.Parameters
		productsResolver []*schema.ProductSchema
		meta             schema.MetaResolver
		result           schema.ProductListResolver
	)
	if args.ProductsArgs.Limit != nil {
		limitStr := strconv.Itoa(int(*args.ProductsArgs.Limit))
		params.StrLimit = limitStr
	}

	if args.ProductsArgs.Page != nil {
		pageStr := strconv.Itoa(int(*args.ProductsArgs.Page))
		params.StrPage = pageStr
	}

	if args.ProductsArgs.OrderBy != nil {
		params.OrderBy = *args.ProductsArgs.OrderBy
	}

	if args.ProductsArgs.Sort != nil {
		params.Sort = *args.ProductsArgs.Sort
	}

	productsOutput := r.ProductUsecase.GetAllProduct(&params)

	if productsOutput.Err != nil {
		return nil, productsOutput.Err
	}

	products := productsOutput.Result.(domain.Products)

	if len(products) <= 0 {
		limitInt32 := int32(params.Limit)
		pageInt32 := int32(params.Page)
		totalInt32 := int32(0)
		totalPageInt32 := int32(0)

		meta.LimitField = &limitInt32
		meta.PageField = &pageInt32
		meta.TotalRecordsField = &totalInt32
		meta.TotalPagesField = &totalPageInt32

		result.ProductsField = productsResolver
		result.MetaField = &meta

		return &result, nil
	}

	for _, product := range products {
		productsResolver = append(productsResolver, &schema.ProductSchema{Product: product})
	}

	countOutput := r.ProductUsecase.GetTotalProduct(&params)

	if countOutput.Err != nil {
		return nil, countOutput.Err
	}

	total := countOutput.Result.(int)

	totalPage := int(math.Ceil(float64(total) / float64(params.Limit)))

	limitInt32 := int32(params.Limit)
	pageInt32 := int32(params.Page)
	totalInt32 := int32(total)
	totalPageInt32 := int32(totalPage)

	meta.LimitField = &limitInt32
	meta.PageField = &pageInt32
	meta.TotalRecordsField = &totalInt32
	meta.TotalPagesField = &totalPageInt32

	result.ProductsField = productsResolver
	result.MetaField = &meta

	return &result, nil
}
