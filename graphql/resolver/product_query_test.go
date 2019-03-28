package resolver

import (
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	productMock "github.com/musobarlab/gorengan/modules/product/usecase/mock"

	"golang.org/x/net/context"
)

func TestProductQuery(t *testing.T) {

	t.Run("should return success test query get product", func(t *testing.T) {
		productUsecaseMock := productMock.NewProductUsecaseMock()
		categoryUsecaseMock := productMock.NewCategoryUsecaseMock()

		resolver := &Resolver{
			ProductUsecase:  productUsecaseMock,
			CategoryUsecase: categoryUsecaseMock,
		}

		ctx := context.Background()

		product, err := resolver.Product(ctx, &ProductQueryArgs{graphql.ID("1")})

		if err != nil {
			t.Error("get product query should return success")
		}

		if product == nil {
			t.Error("product query result should not be nil")
		}

	})

	t.Run("should return success test query get products", func(t *testing.T) {
		productUsecaseMock := productMock.NewProductUsecaseMock()
		categoryUsecaseMock := productMock.NewCategoryUsecaseMock()

		resolver := &Resolver{
			ProductUsecase:  productUsecaseMock,
			CategoryUsecase: categoryUsecaseMock,
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

		products, err := resolver.Products(ctx, productsQueryArgs)

		if err != nil {
			t.Error("get products query should return success")
		}

		if len(products.Products()) <= 0 {
			t.Error("products query results should greater than 0")
		}

	})

}
