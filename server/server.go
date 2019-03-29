package server

import (
	"fmt"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/database"
	"github.com/musobarlab/gorengan/middleware"
	cd "github.com/musobarlab/gorengan/modules/category/delivery"
	cr "github.com/musobarlab/gorengan/modules/category/repository"
	cu "github.com/musobarlab/gorengan/modules/category/usecase"
	pd "github.com/musobarlab/gorengan/modules/product/delivery"
	pr "github.com/musobarlab/gorengan/modules/product/repository"
	pu "github.com/musobarlab/gorengan/modules/product/usecase"
	"github.com/musobarlab/gorengan/schema"
)

// embedding all graphql resolver/ handler anonymously
type graphqlHandlers struct {
	pd.GraphQLProductHandler
	cd.GraphQLCategoryHandler
}

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

	graphqlSchema, err := schema.LoadGraphQLSchema()
	if err != nil {
		return nil, err
	}

	productRepository := pr.NewProductRepositoryGorm(db)
	categoryRepository := cr.NewCategoryRepositoryGorm(db)

	productUsecase := pu.NewProductUsecaseImpl(productRepository, productRepository, categoryRepository)
	categoryUsecase := cu.NewCategoryUsecaseImpl(categoryRepository, categoryRepository)

	productGraphQLHandler := pd.GraphQLProductHandler{ProductUsecase: productUsecase}
	categoryGraphQLHandler := cd.GraphQLCategoryHandler{CategoryUsecase: categoryUsecase}

	// create graphql resolver
	var graphqlResolver graphqlHandlers

	graphqlResolver.GraphQLProductHandler = productGraphQLHandler
	graphqlResolver.GraphQLCategoryHandler = categoryGraphQLHandler

	gqlSchema := graphql.MustParseSchema(graphqlSchema, &graphqlResolver)

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
