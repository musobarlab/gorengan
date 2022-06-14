package delivery

import (
	"time"

	"github.com/musobarlab/gorengan/internal/modules/category/domain"
	"github.com/musobarlab/gorengan/internal/modules/category/graphql/schema"
	"github.com/musobarlab/gorengan/internal/modules/category/usecase"
	"golang.org/x/net/context"
)

// GraphQLCategoryMutationHandler struct
// Handler means Resolver
type GraphQLCategoryMutationHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

// CategoryInputArgs input
type CategoryInputArgs struct {
	Category schema.CategorySchemaInput
}

// CreateCategory mutation
func (r *GraphQLCategoryMutationHandler) CreateCategory(ctx context.Context, args *CategoryInputArgs) (*schema.CategorySchema, error) {
	var category domain.Category
	category.ID = args.Category.ID
	category.Name = args.Category.Name
	category.Created = time.Now()
	category.LastModified = time.Now()

	output := r.CategoryUsecase.CreateCategory(&category)
	if output.Err != nil {
		return nil, output.Err
	}

	categorySaved := output.Result

	return &schema.CategorySchema{Category: categorySaved}, nil
}
