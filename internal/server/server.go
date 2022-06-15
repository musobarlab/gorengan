package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
	"gorm.io/gorm"

	"github.com/musobarlab/gorengan/config"
	"github.com/musobarlab/gorengan/database"
	"github.com/musobarlab/gorengan/internal/server/middleware"
	"github.com/musobarlab/gorengan/pkg/shared"
)

// HTTPServer struct
type HTTPServer struct {
	port             int
	graphQLHandlerV1 *relay.Handler
	db               *gorm.DB
}

// NewHTTPServer echo server constructor
func NewHTTPServer(port int) (*HTTPServer, error) {
	db, err := database.GetGormConn(config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBPort)
	if err != nil {
		return nil, err
	}

	// schema V1
	graphQLHandlerV1, err := initGraphqlSchemaV1(db)
	if err != nil {
		return nil, err
	}

	return &HTTPServer{
		port:             port,
		graphQLHandlerV1: graphQLHandlerV1,
		db:               db,
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
	mux.Handle("/graphql/v1", middleware.BasicAuth(
		middleware.NewBasicAuthConfig(config.BasicAuthUsername, config.BasicAuthPassword),
		s.graphQLHandlerV1,
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
