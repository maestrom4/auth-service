package graphql

import (
	gql "auth-service/internal/types"
	"errors"

	"github.com/graphql-go/graphql"
)

// var UserMutx = &graphql.Field{
// 	Type: gql.UserType,
// 	Args: graphql.FieldConfigArgument{
// 		"name": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"email": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"hashed_password": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		resolver, ok := p.Context.Value("resolver").(*Resolver)
// 		if !ok {
// 			return nil, errors.New("could not get the resolver from the context")
// 		}
// 		log.Debugln("p: ", p)
// 		return resolver.AddUserResolver(p)
// 	},
// }

var UserUpdateMutx = &graphql.Field{
	Type: gql.UserType,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// Retrieve resolver from GraphQL context
		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get resolver from context")
		}
		id, _ := p.Args["_id"].(string)
		name, _ := p.Args["name"].(string)
		email, _ := p.Args["email"].(string)
		return resolver.UserRepository.UpdateUser(p.Context, id, name, email)
	},
}

var UserDelMutx = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get resolver from context")
		}
		id, _ := p.Args["_id"].(string)
		return true, resolver.UserRepository.DeleteUser(p.Context, id)
	},
}
