package types

import (
	"github.com/graphql-go/graphql"
)

var Setting = graphql.NewObject(graphql.ObjectConfig{
	Name: "Setting",
	Fields: graphql.Fields{
		"Key":   &graphql.Field{Type: graphql.String},
		"Value": &graphql.Field{Type: graphql.String},
	},
})
