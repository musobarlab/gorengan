package usecase

import (
	"testing"

	categoryMock "github.com/musobarlab/gorengan/internal/modules/category/repository/mock"
	"github.com/musobarlab/gorengan/internal/modules/product/domain"
	productMock "github.com/musobarlab/gorengan/internal/modules/product/repository/mock"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// TODO
// add more test
func TestProductUsecase(t *testing.T) {

	t.Run("should success test create product", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "3", Name: "Drum", CategoryID: "2", Quantity: 3})

		if output.Err != nil {
			t.Errorf("create product should return nil")
		}
	})

	t.Run("should success test create product already exist", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "1", Name: "Drum", CategoryID: "2", Quantity: 3})

		if output.Err == nil {
			t.Errorf("create product should return error")
		}
	})

	t.Run("should error test create product when category not found", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "3", Name: "Drum", CategoryID: "3", Quantity: 3})

		if output.Err == nil {
			t.Errorf("create product should return error")
		}
	})

	t.Run("should success test remove product", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.RemoveProduct("1")

		if output.Err != nil {
			t.Errorf("remove product should return success")
		}

		productDeleted := output.Result
		if !productDeleted.IsDeleted {
			t.Errorf("product should set is delete to true after deleting")
		}
	})

	t.Run("should error test remove product when product not found", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.RemoveProduct("3")

		if output.Err == nil {
			t.Errorf("remove product should return error")
		}
	})

	t.Run("should success test get product", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.GetProduct("1")

		if output.Err != nil {
			t.Errorf("get product should return success")
		}

		product := output.Result
		if product.ID != "1" {
			t.Errorf("get product should return product with ID 1")
		}
	})

	t.Run("should error test get product when product get by not found ID", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.GetProduct("3")

		if output.Err == nil {
			t.Errorf("get product should return error")
		}

		if output.Result != nil {
			t.Errorf("get product should return nil result")
		}
	})

	t.Run("should success test get products", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := categoryMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.GetAllProduct(&shared.Parameters{})

		if output.Err != nil {
			t.Errorf("get product should return success")
		}

		products := output.Result

		if len(products) <= 0 {
			t.Errorf("get products list with lenght greater than 0")
		}
	})
}
