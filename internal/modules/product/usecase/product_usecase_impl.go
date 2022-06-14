package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	categoryRepo "github.com/musobarlab/gorengan/internal/modules/category/repository"
	"github.com/musobarlab/gorengan/internal/modules/product/domain"
	"github.com/musobarlab/gorengan/internal/modules/product/repository"
	"github.com/musobarlab/gorengan/pkg/shared"
	"gorm.io/gorm"
)

// ProductUsecaseImpl struct
type ProductUsecaseImpl struct {
	productRepositoryRead  repository.ProductRepository
	productRepositoryWrite repository.ProductRepository
	categoryRepository     categoryRepo.CategoryRepository
}

// NewProductUsecaseImpl function
func NewProductUsecaseImpl(productRepositoryRead, productRepositoryWrite repository.ProductRepository,
	categoryRepository categoryRepo.CategoryRepository) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{
		productRepositoryRead:  productRepositoryRead,
		productRepositoryWrite: productRepositoryWrite,
		categoryRepository:     categoryRepository,
	}
}

// CreateProduct function
func (u *ProductUsecaseImpl) CreateProduct(product *domain.Product) shared.Output[*domain.Product] {
	productOutput := u.productRepositoryRead.FindByID(product.ID)
	if productOutput.Err != nil && productOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output[*domain.Product]{Err: productOutput.Err}
	}

	if productOutput.Result != nil {
		productExist := productOutput.Result

		if productExist != nil {
			return shared.Output[*domain.Product]{Err: fmt.Errorf("product with id %s already exist", product.ID)}
		}
	}

	categoryOutput := u.categoryRepository.FindByID(product.CategoryID)
	if categoryOutput.Err != nil {
		if categoryOutput.Err == gorm.ErrRecordNotFound {
			return shared.Output[*domain.Product]{Err: fmt.Errorf("category with id %s not found", product.CategoryID)}
		}
		return shared.Output[*domain.Product]{Err: productOutput.Err}
	}

	err := product.Validate()

	if err != nil {
		return shared.Output[*domain.Product]{Err: err}
	}

	productSaveOutput := u.productRepositoryWrite.Save(product)
	if productSaveOutput.Err != nil {
		return shared.Output[*domain.Product]{Err: productSaveOutput.Err}
	}

	productSaved := productSaveOutput.Result

	return shared.Output[*domain.Product]{Result: productSaved}
}

// RemoveProduct function
func (u *ProductUsecaseImpl) RemoveProduct(id string) shared.Output[*domain.Product] {
	productResult := u.productRepositoryRead.FindByID(id)
	if productResult.Err != nil {
		return shared.Output[*domain.Product]{Err: productResult.Err}
	}

	product := productResult.Result

	// set flag as deleted
	product.IsDeleted = true
	product.Deleted = time.Now()

	productSaveOutput := u.productRepositoryWrite.Save(product)
	if productSaveOutput.Err != nil {
		return shared.Output[*domain.Product]{Err: productSaveOutput.Err}
	}

	productSaved := productSaveOutput.Result
	return shared.Output[*domain.Product]{Result: productSaved}
}

// GetProduct function
func (u *ProductUsecaseImpl) GetProduct(id string) shared.Output[*domain.Product] {
	productResult := u.productRepositoryRead.FindByID(id)
	if productResult.Err != nil {
		return shared.Output[*domain.Product]{Err: productResult.Err}
	}

	product := productResult.Result

	return shared.Output[*domain.Product]{Result: product}
}

// GetAllProduct function
func (u *ProductUsecaseImpl) GetAllProduct(params *shared.Parameters) shared.Output[domain.Products] {
	params.Page = 1

	if len(params.StrPage) > 0 {
		page, err := strconv.Atoi(params.StrPage)
		if err != nil {
			return shared.Output[domain.Products]{Err: shared.NewErrorAllowNumericOnly("page")}
		}

		params.Page = page
	}

	params.Limit = 10
	if len(params.StrLimit) > 0 {
		limit, err := strconv.Atoi(params.StrLimit)
		if err != nil {
			return shared.Output[domain.Products]{Err: shared.NewErrorAllowNumericOnly("limit")}
		}

		params.Limit = limit
	}

	params.Offset = (params.Page - 1) * params.Limit

	if len(params.OrderBy) > 0 {
		if !shared.StringInSlice(params.OrderBy, shared.AllowedSortFields) {
			return shared.Output[domain.Products]{Err: fmt.Errorf(shared.ErrorParameterInvalid, "order by")}
		}
	} else {
		params.OrderBy = "name"
	}

	if len(params.Sort) > 0 {
		if !shared.StringInSlice(params.Sort, []string{"asc", "desc"}) {
			return shared.Output[domain.Products]{Err: fmt.Errorf("parameter %s allow input asc and desc only", "sort")}

		}
	} else {
		params.Sort = "asc"
	}

	params.OrderBy = fmt.Sprintf(`"%s" %s`, strings.ToUpper(params.OrderBy), params.Sort)

	productResult := u.productRepositoryRead.FindAll(params)
	if productResult.Err != nil {
		return shared.Output[domain.Products]{Err: productResult.Err}
	}

	products := productResult.Result

	return shared.Output[domain.Products]{Result: products}
}

// GetTotalProduct function
func (u *ProductUsecaseImpl) GetTotalProduct(params *shared.Parameters) shared.Output[int64] {
	productResult := u.productRepositoryRead.Count(params)
	if productResult.Err != nil {
		return shared.Output[int64]{Err: productResult.Err}
	}

	totalProduct := productResult.Result

	return shared.Output[int64]{Result: totalProduct}
}
