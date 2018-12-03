package types

import (
	"github.com/graphql-go/graphql"
)

var View = graphql.NewObject(graphql.ObjectConfig{
	Name: "View",
	Fields: graphql.Fields{
		"Id":     &graphql.Field{Type: graphql.String},
		"Post":   &graphql.Field{Type: Post},
		"PostId": &graphql.Field{Type: graphql.String},
		"User":   &graphql.Field{Type: User},
		"UserId": &graphql.Field{Type: graphql.String},
		"Time":   &graphql.Field{Type: graphql.DateTime},
	},
})
