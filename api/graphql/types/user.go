package types

import(
	"github.com/graphql-go/graphql"
)
	

var User =  graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"Id": &graphql.Field{ Type: graphql.Int},
		"GroupId": &graphql.Field{ Type: graphql.Int},
		"CreatedOn": &graphql.Field{ Type: graphql.DateTime},
		"Username": &graphql.Field{ Type: graphql.String},
		"Avatar": &graphql.Field{ Type: graphql.String},
		"Signature": &graphql.Field{ Type: graphql.String},
		"StylesheetUrl": &graphql.Field{ Type: graphql.String},
		"UserTitle": &graphql.Field{ Type: graphql.String},
		"LastSeen": &graphql.Field{ Type: graphql.DateTime},
		"HideOnline": &graphql.Field{ Type: graphql.Boolean},
		"LastUnreadAll": &graphql.Field{ Type: graphql.DateTime},
		
	},
})
