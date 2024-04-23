package graphql

import (
	"auth-service/internal/models"
	"auth-service/utils"
	"errors"

	"github.com/graphql-go/graphql"
)

var RegistrationMutx = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// username, _ := p.Args["username"].(string)
		// email, _ := p.Args["email"].(string)
		// password, _ := p.Args["password"].(string)

		// hashedPassword, err := hashPassword(password)
		// if err != nil {
		// 	return nil, err
		// }

		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get the resolver from the context")
		}

		// now := time.Now().Format(time.RFC3339)
		user, err := resolver.AddUserResolver(p)
		if err != nil {
			return nil, err
		}

		userObj, ok := user.(*models.User)
		if !ok || userObj.ID == "" {
			return nil, errors.New("user ID is unavailable")
		}

		token, err := utils.CreateToken(userObj.ID)
		if err != nil {
			return nil, errors.New("could not create token")
		}
		return token, nil
	},
}
