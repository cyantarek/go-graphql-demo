package main

import (
	"fmt"
	"log"
	"net/http"

	"app"

	"github.com/graphql-go/handler"

	gq "github.com/graphql-go/graphql"
)

func main() {

	schema, _ := gq.NewSchema(gq.SchemaConfig{
		Query:    app.rootQuery,
		Mutation: app.rootMutation,
	})

	// http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	// 	result := gq.Do(gq.Params{
	// 		Schema:        schema,
	// 		RequestString: r.URL.Query().Get("query"),
	// 	})

	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(result)
	// })

	conf := handler.NewConfig()
	conf.Schema = &schema
	h := handler.New(conf)
	http.Handle("/graphql", h)

	fmt.Println("GraphQL server is up and running")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
