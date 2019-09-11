package main


import {
	"github.com/graphql-go/graphql"
}
/*
Simple GraphQL Server
*/
func main() {
	// schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Field: fields}
	schemaConfig := graphqlSchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err :=- graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create a nre schema, error: %v", err)
	}
}
