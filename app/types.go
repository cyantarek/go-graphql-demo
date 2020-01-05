package app

import gq "github.com/graphql-go/graphql"

var emailType = gq.NewObject(gq.ObjectConfig{
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

var genderType = gq.NewEnum(gq.EnumConfig{
	Name: "Gender",
	Values: gq.EnumValueConfigMap{
		"MALE": &gq.EnumValueConfig{
			Value: "MALE",
		},
		"FEMALE": &gq.EnumValueConfig{
			Value: "FEMALE",
		},
		"OTHER": &gq.EnumValueConfig{
			Value: "OTHER",
		},
	},
})

var friendInputType = gq.NewInputObject(gq.InputObjectConfig{
	Name: "UserInput",
	Fields: gq.InputObjectConfigFieldMap{
		"first_name": &gq.InputObjectFieldConfig{
			Type: gq.NewNonNull(gq.String),
		},
		"last_name": &gq.InputObjectFieldConfig{
			Type: gq.NewNonNull(gq.String),
		},
		"gender": &gq.InputObjectFieldConfig{
			Type: gq.NewNonNull(genderType),
		},
		"language": &gq.InputObjectFieldConfig{
			Type: gq.NewNonNull(gq.String),
		},
	},
})

var friendType = gq.NewObject(gq.ObjectConfig{
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
			Type: genderType,
		},
		"language": &gq.Field{
			Type: gq.String,
		},
		"emails": &gq.Field{
			Type: gq.NewList(emailType),
		},
	},
})
