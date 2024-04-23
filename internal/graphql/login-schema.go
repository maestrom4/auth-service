package graphql

import (
	gql "auth-service/internal/types"
	"errors"

	"github.com/graphql-go/graphql"
)

var LoginQuerySchema = &graphql.Field{
	Type: gql.UserType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username, ok := p.Args["username"].(string)
		if !ok {
			return nil, errors.New("invalid user ID")
		}
		// id, ok := p.Args["password"].(string)
		// if !ok {
		// 	return nil, errors.New("invalid user ID")
		// }

		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get resolver from context")
		}

		return resolver.UserRepository.GetUserByID(p.Context, username)
	},
}
