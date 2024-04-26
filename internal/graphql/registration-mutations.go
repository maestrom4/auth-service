package graphql

import (
	cfg "auth-service/internal/config"
	"auth-service/internal/models"
	gql "auth-service/internal/types"
	"auth-service/utils"
	"errors"

	"github.com/graphql-go/graphql"
)

var RegistrationMutx = &graphql.Field{
	Type: RegistrationResponseType,
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

		user, err := resolver.AddUserResolver(p)
		if err != nil {
			return nil, err
		}

		userObj, ok := user.(*models.User)
		if !ok || userObj.ID == "" {
			return nil, errors.New("user ID is unavailable")
		}

		token, err := utils.CreateToken(userObj.ID, cfg.SaltPassKey)
		if err != nil {
			return nil, errors.New("could not create token")
		}
		// return token, nil
		response := gql.RegistrationResponse{
			Token:  token,
			UserId: userObj.ID,
		}
		return response, nil
	},
}

var RegistrationResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RegistrationResponse",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"userId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
