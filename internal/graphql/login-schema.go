package graphql

import (
	cfg "auth-service/internal/config"
	"errors"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

var LoginQuerySchema = &graphql.Field{
	Type: LoginResponseType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		resolver, ok := p.Context.Value(string(cfg.ResolverKey)).(*Resolver)
		userID, _ := p.Context.Value(string(cfg.UserIDKey)).(string)
		log.Debugln("LoginQuerySchema userID: ", userID)
		if !ok {
			return nil, errors.New("could not get the resolver from the context")
		}
		user, err := resolver.LoginResolver(p)
		if err != nil {
			return nil, err
		}
		log.Debugln("LoginQuerySchema user: ", user)

		return user, nil
	},
}

var LoginResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"isLoggedIn": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
