package server

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	otelgraphql "github.com/graph-gophers/graphql-go/trace/otel"
	"gorm.io/gorm"

	graphqlSchemaApi "github.com/musobarlab/gorengan/api/graphql"
	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/database"
	cd "github.com/musobarlab/gorengan/internal/modules/category/delivery"
	cr "github.com/musobarlab/gorengan/internal/modules/category/repository"
	cu "github.com/musobarlab/gorengan/internal/modules/category/usecase"
	pd "github.com/musobarlab/gorengan/internal/modules/product/delivery"
	pr "github.com/musobarlab/gorengan/internal/modules/product/repository"
	pu "github.com/musobarlab/gorengan/internal/modules/product/usecase"
	"github.com/musobarlab/gorengan/internal/server/middleware"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// HTTPServer struct
type HTTPServer struct {
	port           int
	graphQLHandler *relay.Handler
	db             *gorm.DB
}

// NewHTTPServer echo server constructor
func NewHTTPServer(port int) (*HTTPServer, error) {
	db, err := database.GetGormConn(config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBPort)
	if err != nil {
		return nil, err
	}

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

	// load schema and resolver
	var resolverFields []reflect.StructField

	resolverModules := make(map[string]interface{})
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

	return &HTTPServer{
		port:           port,
		graphQLHandler: graphQLHandler,
		db:             db,
	}, nil
}

// Run function
func (s *HTTPServer) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shared.BuildJSONResponse(w, shared.Response[shared.EmptyJSON]{
			Success: true,
			Code:    200,
			Message: "server up and running",
			Data:    shared.EmptyJSON{},
		}, 200)
	})

	// secure graphql route with Basic Auth
	mux.Handle("/graphql", middleware.BasicAuth(
		middleware.NewBasicAuthConfig(config.BasicAuthUsername, config.BasicAuthPassword),
		s.graphQLHandler,
	))

	log.Printf("Http server running on port %d\n", s.port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux))

}

func (s *HTTPServer) Exit() {
	log.Print("exiting Http server\n")

	db, err := s.db.DB()
	if err != nil {
		log.Printf("error loading DB %v\n", err)
	}

	log.Print("closing database\n")
	db.Close()
}
