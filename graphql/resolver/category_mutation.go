package resolver

import (
	"time"

	"github.com/musobarlab/gorengan/modules/product/domain"
	"golang.org/x/net/context"
)

// CategoryInputArgs input
type CategoryInputArgs struct {
	Category CategoryInput
}

// CreateCategory mutation
func (r *Resolver) CreateCategory(ctx context.Context, args *CategoryInputArgs) (*CategoryResolver, error) {
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

	return &CategoryResolver{categorySaved}, nil
}
