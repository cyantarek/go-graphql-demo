package main

import (
	gq "github.com/graphql-go/graphql"
)

var friends = []Friend{
	Friend{1, "aaa", "aaa", "MALE", "aaa", []Email{Email{ID: 1, Address: "Hisd"}}},
	Friend{2, "bbb", "bbb", "MALE", "bbb", []Email{Email{ID: 2, Address: "Hisd"}}},
	Friend{3, "ccc", "ccc", "FEMALE", "ccc", []Email{Email{ID: 5, Address: "Hisd"}}},
	Friend{4, "ddd", "ddd", "OTHER", "ddd", []Email{Email{ID: 10, Address: "Hisss"},
		Email{ID: 12, Address: "Hi2"}}},
}

// HelloResolver defines
func HelloResolver(p gq.ResolveParams) (interface{}, error) {
	return "Hi, I'm Manny", nil
}

// FriendsResolver defines
func FriendsResolver(p gq.ResolveParams) (interface{}, error) {

	return friends, nil
}

// FriendResolver defines
func FriendResolver(p gq.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int)
	// idInt, _ := strconv.Atoi(id)
	for _, v := range friends {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, nil
}
