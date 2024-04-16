package app

import (
	"log"

	"github.com/graph-gophers/graphql-go"
)

var Schema *graphql.Schema

func Initialize() {
	schema := `
        type Query {
            hello: String!
        }
    `
	var err error
	Schema, err = graphql.ParseSchema(schema, &query{})
	if err != nil {
		log.Fatal(err)
	}
}

type query struct{}

func (_ *query) Hello() string {
	return "Hello, world!"
}
