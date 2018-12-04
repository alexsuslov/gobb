package types

import(
	"github.com/graphql-go/graphql"
)
	

var Board =  graphql.NewObject(graphql.ObjectConfig{
	Name: "Board",
	Fields: graphql.Fields{
		"Id": &graphql.Field{ Type: graphql.Int},
		"Title": &graphql.Field{ Type: graphql.String},
		"Description": &graphql.Field{ Type: graphql.String},
		"Order": &graphql.Field{ Type: graphql.Int},
		
	},
})

var BoardLatest =  graphql.NewObject(graphql.ObjectConfig{
	Name: "BoardLatest",
	Fields: graphql.Fields{
		"Op": &graphql.Field{ Type: Post},
		"Latest": &graphql.Field{ Type: Post},
		
	},
})

var JoinBoardView =  graphql.NewObject(graphql.ObjectConfig{
	Name: "JoinBoardView",
	Fields: graphql.Fields{
		"Board": &graphql.Field{ Type: Board},
		"Id": &graphql.Field{ Type: graphql.Int},
		"Title": &graphql.Field{ Type: graphql.String},
		"Description": &graphql.Field{ Type: graphql.String},
		"Order": &graphql.Field{ Type: graphql.Int},
		"ViewedOn": &graphql.Field{ Type: graphql.DateTime},
		
	},
})

var JoinThreadView =  graphql.NewObject(graphql.ObjectConfig{
	Name: "JoinThreadView",
	Fields: graphql.Fields{
		"Thread": &graphql.Field{ Type: Post},
		"Id": &graphql.Field{ Type: graphql.Int},
		"BoardId": &graphql.Field{ Type: graphql.Int},
		"Author": &graphql.Field{ Type: User},
		"AuthorId": &graphql.Field{ Type: graphql.Int},
		"Title": &graphql.Field{ Type: graphql.String},
		"CreatedOn": &graphql.Field{ Type: graphql.DateTime},
		"LatestReply": &graphql.Field{ Type: graphql.DateTime},
		"Sticky": &graphql.Field{ Type: graphql.Boolean},
		"Locked": &graphql.Field{ Type: graphql.Boolean},
		"ViewedOn": &graphql.Field{ Type: graphql.DateTime},
		
	},
})
