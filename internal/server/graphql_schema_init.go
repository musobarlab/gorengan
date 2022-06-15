package server

import (
	"reflect"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	otelgraphql "github.com/graph-gophers/graphql-go/trace/otel"
	"gorm.io/gorm"

	graphqlSchemaApi "github.com/musobarlab/gorengan/api/graphql"
	cdv1 "github.com/musobarlab/gorengan/internal/modules/category/v1/delivery"
	crv1 "github.com/musobarlab/gorengan/internal/modules/category/v1/repository"
	cuv1 "github.com/musobarlab/gorengan/internal/modules/category/v1/usecase"
	pdv1 "github.com/musobarlab/gorengan/internal/modules/product/v1/delivery"
	prv1 "github.com/musobarlab/gorengan/internal/modules/product/v1/repository"
	puv1 "github.com/musobarlab/gorengan/internal/modules/product/v1/usecase"
)

func initGraphqlSchemaV1(db *gorm.DB) (*relay.Handler, error) {
	// load graphql schema file, and convert to string
	graphqlSchema, err := graphqlSchemaApi.LoadGraphQLSchema("v1")
	if err != nil {
		return nil, err
	}

	// initial repository
	productRepository := prv1.NewProductRepositoryGorm(db)
	categoryRepository := crv1.NewCategoryRepositoryGorm(db)

	// initial usecase
	productUsecase := puv1.NewProductUsecaseImpl(productRepository, productRepository, categoryRepository)
	categoryUsecase := cuv1.NewCategoryUsecaseImpl(categoryRepository, categoryRepository)

	// initial graphql handler/ resolver
	productGraphQLQueryHandler := &pdv1.GraphQLProductQueryHandler{ProductUsecase: productUsecase}
	productGraphQLMutationHandler := &pdv1.GraphQLProductMutationHandler{ProductUsecase: productUsecase}
	categoryGraphQLMutationHandler := &cdv1.GraphQLCategoryMutationHandler{CategoryUsecase: categoryUsecase}

	// load schema and resolver
	var resolverFields []reflect.StructField

	resolverModules := make(map[string]graphqlBaseHandler)
	resolverModules[productGraphQLMutationHandler.Name()] = productGraphQLMutationHandler
	resolverModules[categoryGraphQLMutationHandler.Name()] = categoryGraphQLMutationHandler
	resolverModules[productGraphQLQueryHandler.Name()] = productGraphQLQueryHandler

	for name, handler := range resolverModules {
		resolverFields = append(resolverFields, reflect.StructField{
			Name: name,
			Type: reflect.TypeOf(handler),
		})
	}

	resolverVal := reflect.New(reflect.StructOf(resolverFields)).Elem()
	for k, v := range resolverModules {
		val := resolverVal.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}

	resolver := resolverVal.Addr().Interface()

	// end load schema and resolver

	// parse grapqhql schema to code
	gqlSchema := graphql.MustParseSchema(
		graphqlSchema,
		resolver,
		graphql.UseStringDescriptions(),
		graphql.UseFieldResolvers(),

		// tracing
		graphql.Tracer(otelgraphql.DefaultTracer()),
	)

	graphQLHandler := &relay.Handler{Schema: gqlSchema}

	return graphQLHandler, nil
}
