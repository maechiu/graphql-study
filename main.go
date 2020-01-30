package main

import (
	"graphql-study/query"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	config := graphql.SchemaConfig{
		Query: query.Object,
	}

	schema, _ := graphql.NewSchema(config)

	http.Handle("/", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))
	http.ListenAndServe(":8080", nil)
}
