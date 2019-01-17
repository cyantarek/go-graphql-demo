package main

import (
	"fmt"

	gq "github.com/graphql-go/graphql"
)

func createFriend(p gq.ResolveParams) (interface{}, error) {
	payload := p.Args["input"]
	fmt.Println(payload)
	frnd := Friend{FirstName: "Cyan"}
	return frnd, nil
}
