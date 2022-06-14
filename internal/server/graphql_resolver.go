package server

import (
	cd "github.com/musobarlab/gorengan/internal/modules/category/delivery"
	pd "github.com/musobarlab/gorengan/internal/modules/product/delivery"
)

type graphqlMutation struct {
	product  *pd.GraphQLProductMutationHandler
	category *cd.GraphQLCategoryMutationHandler
}

type graphqlQuery struct {
	product *pd.GraphQLProductQueryHandler
}

func (g graphqlQuery) ProductQuery() *pd.GraphQLProductQueryHandler {
	return g.product
}

func (g graphqlMutation) ProductMutation() *pd.GraphQLProductMutationHandler {
	return g.product
}

func (g graphqlMutation) CategoryMutation() *cd.GraphQLCategoryMutationHandler {
	return g.category
}

// embedding all graphql resolver/ handler anonymously
type graphqlResolver struct {
	graphqlMutation
	graphqlQuery
}
