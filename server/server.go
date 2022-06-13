package server

import (
	"fmt"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	otelgraphql "github.com/graph-gophers/graphql-go/trace/otel"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/database"
	"github.com/musobarlab/gorengan/middleware"
	cd "github.com/musobarlab/gorengan/modules/category/delivery"
	cr "github.com/musobarlab/gorengan/modules/category/repository"
	cu "github.com/musobarlab/gorengan/modules/category/usecase"
	pd "github.com/musobarlab/gorengan/modules/product/delivery"
	pr "github.com/musobarlab/gorengan/modules/product/repository"
	pu "github.com/musobarlab/gorengan/modules/product/usecase"
	graphqlSchemaApi "github.com/musobarlab/gorengan/api/graphql"
)

// EchoServer struct
type EchoServer struct {
	port           int
	graphQLHandler *relay.Handler
}

// NewEchoServer echo server constructor
func NewEchoServer(port int) (*EchoServer, error) {
	db, err := database.GetGormConn(config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBPort)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	// load graphql schema file, and convert to string
	graphqlSchema, err := graphqlSchemaApi.LoadGraphQLSchema()
	if err != nil {
		return nil, err
	}

	// initial repository
	productRepository := pr.NewProductRepositoryGorm(db)
	categoryRepository := cr.NewCategoryRepositoryGorm(db)

	// initial usecase
	productUsecase := pu.NewProductUsecaseImpl(productRepository, productRepository, categoryRepository)
	categoryUsecase := cu.NewCategoryUsecaseImpl(categoryRepository, categoryRepository)

	// initial graphql handler/ resolver
	productGraphQLQueryHandler := &pd.GraphQLProductQueryHandler{ProductUsecase: productUsecase}
	productGraphQLMutationHandler := &pd.GraphQLProductMutationHandler{ProductUsecase: productUsecase}
	categoryGraphQLMutationHandler := &cd.GraphQLCategoryMutationHandler{CategoryUsecase: categoryUsecase}

	// create graphql resolver
	var graphqlResolver graphqlResolver

	graphqlResolver.graphqlMutation.product = productGraphQLMutationHandler
	graphqlResolver.graphqlMutation.category = categoryGraphQLMutationHandler
	graphqlResolver.graphqlQuery.product = productGraphQLQueryHandler

	// parse grapqhql schema to code
	gqlSchema := graphql.MustParseSchema(
		graphqlSchema, 
		&graphqlResolver,
		graphql.UseStringDescriptions(),
		graphql.UseFieldResolvers(),

		// tracing
		graphql.Tracer(otelgraphql.DefaultTracer()),
	)

	graphQLHandler := &relay.Handler{Schema: gqlSchema}

	return &EchoServer{
		port:           port,
		graphQLHandler: graphQLHandler,
	}, nil
}

// Run function
func (s *EchoServer) Run() {
	e := echo.New()
	e.Use(echoMiddleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// secure graphql route with Basic Auth
	e.POST("/graphql", echo.WrapHandler(s.graphQLHandler), middleware.BasicAuth(config.BasicAuthUsername, config.BasicAuthPassword))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.port)))

}
