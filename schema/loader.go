package schema

import "io/ioutil"

const (
	// GraphQLSchemaPath schema paths
	GraphQLSchemaPath = "./schema/graphql_schema.graphql"
)

// LoadGraphQLSchema will read graphql schema from file
func LoadGraphQLSchema() (string, error) {
	b, err := ioutil.ReadFile(GraphQLSchemaPath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
