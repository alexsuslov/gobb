//go:generate (cd ./types; go run user_gen.go)

package gobbgraphql

import (
	"github.com/alexsuslov/gobb/api/graphql/model"
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hi": helloWorld,
		"GetPost": &gqlmodel.Post,
		//"code": &code.Code,
		//"codes": &code.Codes,
		//"phone": &code.Phone,
		//"region": &region.One,
		//"regions": &region.Find,
		//"operator": &operator.One,
		//"operators": &operator.Find,
	},
})

var Schema = graphql.SchemaConfig{
	Query: RootQuery,
}

func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(Schema)
}

var helloWorld = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "Hello World!", nil
	},
}
