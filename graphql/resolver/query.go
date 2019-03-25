package resolver

import (
	"math"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/shared"
	"golang.org/x/net/context"
)

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

// Product Query function
func (r *Resolver) Product(ctx context.Context, args *ProductQueryArgs) (*ProductResolver, error) {
	output := r.ProductUsecase.GetProduct(string(args.ID))

	if output.Err != nil {
		return nil, output.Err
	}

	product := output.Result.(*domain.Product)

	return &ProductResolver{product}, nil

}

// Products Query function
func (r *Resolver) Products(ctx context.Context, args *ProductsQueryArgs) (*ProductListResolver, error) {
	var (
		params           shared.Parameters
		productsResolver []*ProductResolver
		meta             MetaResolver
		result           ProductListResolver
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

		meta.limit = &limitInt32
		meta.page = &pageInt32
		meta.totalRecords = &totalInt32
		meta.totalPages = &totalPageInt32

		result.products = productsResolver
		result.meta = &meta

		return &result, nil
	}

	for _, product := range products {
		productsResolver = append(productsResolver, &ProductResolver{product})
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

	meta.limit = &limitInt32
	meta.page = &pageInt32
	meta.totalRecords = &totalInt32
	meta.totalPages = &totalPageInt32

	result.products = productsResolver
	result.meta = &meta

	return &result, nil
}
