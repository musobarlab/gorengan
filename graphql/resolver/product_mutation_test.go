package resolver

import (
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	productMock "github.com/musobarlab/gorengan/modules/product/usecase/mock"

	"golang.org/x/net/context"
)

func TestProductMutation(t *testing.T) {

	t.Run("should return success test mutation create product", func(t *testing.T) {
		productUsecaseMock := productMock.NewProductUsecaseMock()
		categoryUsecaseMock := productMock.NewCategoryUsecaseMock()

		resolver := &Resolver{
			ProductUsecase:  productUsecaseMock,
			CategoryUsecase: categoryUsecaseMock,
		}

		ctx := context.Background()

		productInputArgs := &ProductInputArgs{
			Product: ProductInput{
				ID:       "3",
				Name:     "Drum",
				Category: "2",
				Quantity: 3,
			},
		}

		productCreated, err := resolver.CreateProduct(ctx, productInputArgs)

		if err != nil {
			t.Error("create product mutation should return success")
		}

		if productCreated == nil {
			t.Error("product created should not be nil")
		}

	})

	t.Run("should return success test mutation delete product", func(t *testing.T) {
		productUsecaseMock := productMock.NewProductUsecaseMock()
		categoryUsecaseMock := productMock.NewCategoryUsecaseMock()

		resolver := &Resolver{
			ProductUsecase:  productUsecaseMock,
			CategoryUsecase: categoryUsecaseMock,
		}

		ctx := context.Background()

		productDeleted, err := resolver.DeleteProduct(ctx, &DeleteProductArgs{graphql.ID("1")})

		if err != nil {
			t.Error("delete product mutation should return success")
		}

		if productDeleted == nil {
			t.Error("product deleted should not be nil")
		}

	})

}
