package app

import (
	"fmt"

	gq "github.com/graphql-go/graphql"
)

var RootMutation = gq.NewObject(gq.ObjectConfig{
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

func createFriend(p gq.ResolveParams) (interface{}, error) {
	payload := p.Args["input"]
	fmt.Println(payload)
	frnd := Friend{FirstName: "Cyan"}
	return frnd, nil
}
