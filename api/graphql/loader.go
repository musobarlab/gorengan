package schema

import (
	"io/ioutil"
	"strings"
)

const (
	// GraphQLSchemaPath schema paths
	GraphQLSchemaPath        = "./api/graphql/graphql_schema.graphql"
	GraphQLProductSchemaPath = "./api/graphql/product.graphql"
	GraphQLMetaSchemaPath    = "./api/graphql/meta.graphql"
)

// LoadGraphQLSchema will read graphql schema from file
func LoadGraphQLSchema() (string, error) {
	var schemaBuilder strings.Builder

	graphqlMetaSchema, err := ioutil.ReadFile(GraphQLMetaSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlMetaSchema)

	graphqlProductSchema, err := ioutil.ReadFile(GraphQLProductSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlProductSchema)

	graphqlSchema, err := ioutil.ReadFile(GraphQLSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlSchema)

	return schemaBuilder.String(), nil
}
