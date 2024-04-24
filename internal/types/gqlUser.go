package gqltypes

import (
	"auth-service/internal/models"
	"errors"

	"github.com/graphql-go/graphql"
)

var StringType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "String",
	Description: "Custom String type to avoid using graphql.ID",
	Serialize: func(value interface{}) interface{} {
		if str, ok := value.(string); ok {
			return str
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		if str, ok := value.(string); ok {
			return str
		}
		return nil
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"hashed_password": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.HashedPassword, nil
				}
				return nil, errors.New("expected a user object")
			},
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"lastLogin": &graphql.Field{
			Type: graphql.String,
		},
	},
})
