package query

import "github.com/graphql-go/graphql"

import "graphql-study/query/field"

// Object represents a Graphql object for query.
var Object = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &field.Hello,
	},
})
