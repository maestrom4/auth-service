package graphql

import (
	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user":  UserQuerySchema,
		"login": LoginQuerySchema,
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"register":   RegistrationMutx,
		"addUser":    UserMutx,
		"updateUser": UserUpdateMutx,
		"deleteUser": UserDelMutx,
	},
})
