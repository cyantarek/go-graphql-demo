package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gq "github.com/graphql-go/graphql"
)

func main() {
	rootQuery := gq.NewObject(gq.ObjectConfig{
		Name: "Query",
		Fields: gq.Fields{
			"hello": &gq.Field{
				Type: gq.String,
				Resolve: func(p gq.ResolveParams) (interface{}, error) {
					return "Hi, I'm Manny", nil
				},
			},
		},
	})

	schema, _ := gq.NewSchema(gq.SchemaConfig{
		Query:    rootQuery,
		Mutation: nil,
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := gq.Do(gq.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("GraphQL server is up and running")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
