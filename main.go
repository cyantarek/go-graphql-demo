package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gq "github.com/graphql-go/graphql"
)

type Friend struct {
	ID        int     `json: "id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    string  `json:"gender"`
	Language  string  `json:"language"`
	Emails    []Email `json:"emails"`
}

type Email struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {
	emailType := gq.NewObject(gq.ObjectConfig{
		Name: "email",
		Fields: gq.Fields{
			"id": &gq.Field{
				Type: gq.ID,
			},
			"address": &gq.Field{
				Type: gq.String,
			},
		},
	})

	friendInputType := gq.NewInputObject(gq.InputObjectConfig{
		Name: "UserInput",
		Fields: gq.InputObjectConfigFieldMap{
			"first_name": &gq.InputObjectFieldConfig{
				Type: gq.NewNonNull(gq.String),
			},
			"last_name": &gq.InputObjectFieldConfig{
				Type: gq.NewNonNull(gq.String),
			},
			"gender": &gq.InputObjectFieldConfig{
				Type: gq.NewNonNull(gq.String),
			},
			"language": &gq.InputObjectFieldConfig{
				Type: gq.NewNonNull(gq.String),
			},
		},
	})

	friendType := gq.NewObject(gq.ObjectConfig{
		Name: "friend",
		Fields: gq.Fields{
			"id": &gq.Field{
				Type: gq.ID,
			},
			"first_name": &gq.Field{
				Type: gq.String,
			},
			"last_name": &gq.Field{
				Type: gq.String,
			},
			"gender": &gq.Field{
				Type: gq.String,
			},
			"language": &gq.Field{
				Type: gq.String,
			},
			"emails": &gq.Field{
				Type: gq.NewList(emailType),
			},
		},
	})

	rootQuery := gq.NewObject(gq.ObjectConfig{
		Name: "Query",
		Fields: gq.Fields{
			"hello": &gq.Field{
				Type:    gq.String,
				Resolve: HelloResolver,
			},
			"friends": &gq.Field{
				Type:    gq.NewList(friendType),
				Resolve: FriendsResolver,
			},
			"friend": &gq.Field{
				Type:    friendType,
				Resolve: FriendResolver,
				Args: gq.FieldConfigArgument{
					"id": &gq.ArgumentConfig{
						Type: gq.NewNonNull(gq.Int),
					},
				},
			},
		},
	})

	rootMutation := gq.NewObject(gq.ObjectConfig{
		Name: "Mutation",
		Fields: gq.Fields{
			"createFriend": &gq.Field{
				Type: friendType,
				Args: gq.FieldConfigArgument{
					"input": &gq.ArgumentConfig{
						Type: gq.NewNonNull(friendInputType),
					},
				},
				Resolve: createFriend,
			},
		},
	})

	schema, _ := gq.NewSchema(gq.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
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
