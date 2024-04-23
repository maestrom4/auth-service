package graphql

import (
	gql "auth-service/internal/types"
	"errors"

	"github.com/graphql-go/graphql"
)

var UserQuerySchema = &graphql.Field{
	Type: gql.UserType,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["_id"].(string)
		if !ok {
			return nil, errors.New("invalid user ID")
		}

		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get resolver from context")
		}

		return resolver.UserRepository.GetUserByID(p.Context, id)
	},
}
