package field

import "github.com/graphql-go/graphql"

func friends(p graphql.ResolveParams) (interface{}, error) {
	return []string{}, nil
}

// Friends returns a empty list.
var Friends = graphql.Field{
	Type: &graphql.List{
		OfType: graphql.String,
	},
	Resolve:     friends,
	Description: "Get all friends.",
}
