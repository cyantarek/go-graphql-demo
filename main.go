package main

import (
	"fmt"
	"log"
	"net/http"

	"gq-test/app"

	"github.com/graphql-go/handler"

	gq "github.com/graphql-go/graphql"
)

func main() {

	schema, _ := gq.NewSchema(gq.SchemaConfig{
		Query:    app.RootQuery,
		Mutation: app.RootMutation,
	})
	
	conf := handler.NewConfig()
	conf.Schema = &schema
	h := handler.New(conf)
	http.Handle("/graphql", h)

	fmt.Println("GraphQL server is up and running")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
