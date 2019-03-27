package usecase

import (
	"testing"

	"github.com/musobarlab/gorengan/modules/product/domain"
	productMock "github.com/musobarlab/gorengan/modules/product/repository/mock"
)

// TODO
// add more test
func TestProductUsecase(t *testing.T) {

	t.Run("should success test create product", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := productMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "3", Name: "Drum", CategoryID: "2", Quantity: 3})

		if output.Err != nil {
			t.Errorf("create product should return nil")
		}
	})

	t.Run("should success test create product already exist", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := productMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "1", Name: "Drum", CategoryID: "2", Quantity: 3})

		if output.Err == nil {
			t.Errorf("create product should return error")
		}
	})

	t.Run("should error test create product when category not found", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := productMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.CreateProduct(&domain.Product{ID: "3", Name: "Drum", CategoryID: "3", Quantity: 3})

		if output.Err == nil {
			t.Errorf("create product should return error")
		}
	})

	t.Run("should success test remove product", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := productMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.RemoveProduct("1")

		if output.Err != nil {
			t.Errorf("remove product should return success")
		}

		productDeleted := output.Result.(*domain.Product)
		if !productDeleted.IsDeleted {
			t.Errorf("product should set is delete to true after deleting")
		}
	})

	t.Run("should error test remove product when product not found", func(t *testing.T) {
		productRepositoryMock := productMock.NewProductRepositoryMock()
		categoryRepositoryMock := productMock.NewCategoryRepositoryMock()

		productUsecase := NewProductUsecaseImpl(productRepositoryMock, productRepositoryMock, categoryRepositoryMock)

		output := productUsecase.RemoveProduct("3")

		if output.Err == nil {
			t.Errorf("remove product should return error")
		}
	})
}
