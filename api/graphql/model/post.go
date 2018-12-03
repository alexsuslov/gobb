package gqlmodel

import (
	"github.com/alexsuslov/gobb/api/graphql/types"
	"github.com/graphql-go/graphql"
	"github.com/alexsuslov/gobb/api/models"
)

var Post = graphql.Field{
	Type: types.Post,
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ID, _ := p.Args["ID"].(string)
		return models.GetPost(ID)
	},
}
