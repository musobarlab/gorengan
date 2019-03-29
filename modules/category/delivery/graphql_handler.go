package delivery

import (
	"time"

	"github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/category/graphql/schema"
	"github.com/musobarlab/gorengan/modules/category/usecase"
	"golang.org/x/net/context"
)

// GraphQLCategoryHandler struct
// Handler means Resolver
type GraphQLCategoryHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

// CategoryInputArgs input
type CategoryInputArgs struct {
	Category schema.CategorySchemaInput
}

// CreateCategory mutation
func (r *GraphQLCategoryHandler) CreateCategory(ctx context.Context, args *CategoryInputArgs) (*schema.CategorySchema, error) {
	var category domain.Category
	category.ID = args.Category.ID
	category.Name = args.Category.Name
	category.Created = time.Now()
	category.LastModified = time.Now()

	output := r.CategoryUsecase.CreateCategory(&category)
	if output.Err != nil {
		return nil, output.Err
	}

	categorySaved := output.Result.(*domain.Category)

	return &schema.CategorySchema{Category: categorySaved}, nil
}
