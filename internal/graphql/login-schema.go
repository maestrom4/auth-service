package graphql

import (
	gql "auth-service/internal/types"
	"errors"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

var LoginQuerySchema = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: gql.UserType, // Make sure UserType is correctly defined and imported
			},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get the resolver from the context")
		}
		user, err := resolver.LoginResolver(p)
		if err != nil {
			return nil, err
		}
		log.Debugln("user: ", user)

		return user, nil
	},
}
