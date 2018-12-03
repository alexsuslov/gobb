package types

import (
	"github.com/graphql-go/graphql"
)

var Post = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"Id":          &graphql.Field{Type: graphql.String},
		"BoardId":     &graphql.Field{Type: graphql.String},
		"ParentId":    &graphql.Field{Type: graphql.String},
		"Author":      &graphql.Field{Type: User},
		"AuthorId":    &graphql.Field{Type: graphql.String},
		"Title":       &graphql.Field{Type: graphql.String},
		"Content":     &graphql.Field{Type: graphql.String},
		"CreatedOn":   &graphql.Field{Type: graphql.DateTime},
		"LatestReply": &graphql.Field{Type: graphql.DateTime},
		"LastEdit":    &graphql.Field{Type: graphql.DateTime},
		"Sticky":      &graphql.Field{Type: graphql.Boolean},
		"Locked":      &graphql.Field{Type: graphql.Boolean},
	},
})
