package server

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/database"
	"github.com/musobarlab/gorengan/modules/product/graphql/resolver"
	"github.com/musobarlab/gorengan/modules/product/repository"
	"github.com/musobarlab/gorengan/modules/product/usecase"
	"github.com/musobarlab/gorengan/schema"
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

	graphqlSchema, err := schema.LoadGraphQLSchema()
	if err != nil {
		return nil, err
	}

	productRepository := repository.NewProductRepositoryGorm(db)
	productUsecase := usecase.NewProductUsecaseImpl(productRepository, productRepository)

	gqlSchema := graphql.MustParseSchema(graphqlSchema, &resolver.Resolver{ProductUsecase: productUsecase})

	graphQLHandler := &relay.Handler{Schema: gqlSchema}

	return &EchoServer{
		port:           port,
		graphQLHandler: graphQLHandler,
	}, nil
}

// Run function
func (s *EchoServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/graphql", echo.WrapHandler(s.graphQLHandler))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.port)))

}
