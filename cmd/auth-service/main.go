// File: cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string {
	return "Hello, world!"
}

func main() {
	schema := `
        type Query {
            hello: String!
        }
    `

	parsedSchema, err := graphql.ParseSchema(schema, &query{})
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", &relay.Handler{Schema: parsedSchema})
	log.Println("GraphQL server running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
