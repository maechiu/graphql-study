package field

import "github.com/graphql-go/graphql"

var args = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "world",
		Description:  "Name to say hello",
	},
}

func sayHello(p graphql.ResolveParams) (interface{}, error) {
	return p.Args["name"], nil
}

// Hello says hello to somebody.
var Hello = graphql.Field{
	Type:        graphql.String,
	Args:        args,
	Resolve:     sayHello,
	Description: "Say hello to somebody.",
}
