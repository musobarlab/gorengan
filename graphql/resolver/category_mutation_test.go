package resolver

import (
	"testing"

	productMock "github.com/musobarlab/gorengan/modules/product/usecase/mock"

	"golang.org/x/net/context"
)

func TestCategoryMutation(t *testing.T) {

	t.Run("should return success test mutation create category", func(t *testing.T) {
		productUsecaseMock := productMock.NewProductUsecaseMock()
		categoryUsecaseMock := productMock.NewCategoryUsecaseMock()

		resolver := &Resolver{
			ProductUsecase:  productUsecaseMock,
			CategoryUsecase: categoryUsecaseMock,
		}

		ctx := context.Background()

		categoryInputArgs := &CategoryInputArgs{
			Category: CategoryInput{
				ID:   "1",
				Name: "Music",
			},
		}

		categoryCreated, err := resolver.CreateCategory(ctx, categoryInputArgs)

		if err != nil {
			t.Error("create category mutation should return success")
		}

		if categoryCreated == nil {
			t.Error("category created should not be nil")
		}

	})

}
