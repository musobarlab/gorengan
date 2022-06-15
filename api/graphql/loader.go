package schema

import (
	"errors"
	"io/ioutil"
	"strings"
)

// LoadGraphQLSchema will read graphql schema from file
func LoadGraphQLSchema(version string) (string, error) {
	switch version {
	case "v1":
		return loadV1Schema()
	}

	return "", errors.New("error loading graphql schema")
}

func loadV1Schema() (string, error) {

	var (
		schemaBuilder strings.Builder

		graphQLSchemaPath        string = "./api/graphql/v1/graphql_schema.graphql"
		graphQLProductSchemaPath string = "./api/graphql/v1/product.graphql"
		graphQLMetaSchemaPath    string = "./api/graphql/v1/meta.graphql"
	)

	graphqlMetaSchema, err := ioutil.ReadFile(graphQLMetaSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlMetaSchema)

	graphqlProductSchema, err := ioutil.ReadFile(graphQLProductSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlProductSchema)

	graphqlSchema, err := ioutil.ReadFile(graphQLSchemaPath)
	if err != nil {
		return "", err
	}

	schemaBuilder.Write(graphqlSchema)

	return schemaBuilder.String(), nil
}
