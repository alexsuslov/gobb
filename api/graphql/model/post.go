package gqlmodel

import (
	"github.com/alexsuslov/gobb/api/graphql/types"
	"github.com/alexsuslov/gobb/api/models"
	"github.com/graphql-go/graphql"
	"log"
)

var Post = graphql.Field{
	Type: types.Post,
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		ID := p.Args["ID"].(int)
		log.Println(ID)
		m, err :=models.GetPost(ID)
		log.Println(m, err)
		return m, err
	},
}
