package query

import "github.com/graphql-go/graphql"

var fields = graphql.Fields{
	"hello": &graphql.Field{
		Description: "Say hello to somebody.",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "world",
				Description:  "Name to say hello",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Args["name"], nil
		},
	},
}

var config = graphql.ObjectConfig{
	Name:   "Query",
	Fields: fields,
}

// Object represents a Graphql object for query.
var Object = graphql.NewObject(config)
