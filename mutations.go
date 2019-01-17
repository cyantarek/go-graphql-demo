package main

import (
	"fmt"

	gq "github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func createFriend(p gq.ResolveParams) (interface{}, error) {
	payload := p.Args["input"]
	result := Friend{}
	mapstructure.Decode(payload, &result)
	fmt.Println(result)
	frnd := Friend{FirstName: "Cyan"}
	return frnd, nil
}
