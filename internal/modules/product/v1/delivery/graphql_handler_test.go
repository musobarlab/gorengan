package delivery

import (
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/internal/modules/product/v1/graphql/schema"
	usecaseMock "github.com/musobarlab/gorengan/internal/modules/product/v1/usecase/mock"
	"golang.org/x/net/context"
)

func TestGraphQLHandler(t *testing.T) {

	t.Run("should return success test mutation create product", func(t *testing.T) {
		productUsecaseMock := usecaseMock.NewProductUsecaseMock()

		handler := &GraphQLProductMutationHandler{
			ProductUsecase: productUsecaseMock,
		}

		ctx := context.Background()

		productInputArgs := &ProductInputArgs{
			Product: schema.ProductSchemaInput{
				ID:       "3",
				Name:     "Drum",
				Category: "2",
				Quantity: 3,
			},
		}

		productCreated, err := handler.CreateProduct(ctx, productInputArgs)

		if err != nil {
			t.Error("create product mutation should return success")
		}

		if productCreated == nil {
			t.Error("product created should not be nil")
		}

	})

	t.Run("should return success test mutation delete product", func(t *testing.T) {
		productUsecaseMock := usecaseMock.NewProductUsecaseMock()

		handler := &GraphQLProductMutationHandler{
			ProductUsecase: productUsecaseMock,
		}

		ctx := context.Background()

		productDeleted, err := handler.DeleteProduct(ctx, &DeleteProductArgs{graphql.ID("1")})

		if err != nil {
			t.Error("delete product mutation should return success")
		}

		if productDeleted == nil {
			t.Error("product deleted should not be nil")
		}

	})

	t.Run("should return success test query get product", func(t *testing.T) {
		productUsecaseMock := usecaseMock.NewProductUsecaseMock()

		handler := &GraphQLProductQueryHandler{
			ProductUsecase: productUsecaseMock,
		}

		ctx := context.Background()

		product, err := handler.Product(ctx, &ProductQueryArgs{graphql.ID("1")})

		if err != nil {
			t.Error("get product query should return success")
		}

		if product == nil {
			t.Error("product query result should not be nil")
		}

	})

	t.Run("should return success test query get products", func(t *testing.T) {
		productUsecaseMock := usecaseMock.NewProductUsecaseMock()

		handler := &GraphQLProductQueryHandler{
			ProductUsecase: productUsecaseMock,
		}

		ctx := context.Background()

		query := "samsung"
		limit := float64(1)
		page := float64(1)
		orderBy := "created"
		sort := "asc"
		productsQueryArgs := &ProductsQueryArgs{
			ProductsArgs: &ProductsArgs{
				Query:   &query,
				Limit:   &limit,
				Page:    &page,
				OrderBy: &orderBy,
				Sort:    &sort,
			},
		}

		products, err := handler.Products(ctx, productsQueryArgs)

		if err != nil {
			t.Error("get products query should return success")
		}

		if len(products.Products()) <= 0 {
			t.Error("products query results should greater than 0")
		}

	})

}
