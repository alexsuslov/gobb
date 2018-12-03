// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates. It can be invoked by running
// go generate

package main

import (
	"bufio"
	"fmt"
	"github.com/alexsuslov/gobb/api/models"
	"github.com/alexsuslov/gobb/pkg/gql"
	"os"
)

func main() {
	f, err := os.Create("board.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = fmt.Fprintln(w, `package types

import(
	"github.com/graphql-go/graphql"
)
	`)

	Board, err := gql.TypeToGqlType("Board", models.Board{})
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(w, Board)

	//BoardLatest
	BoardLatest, err := gql.TypeToGqlType("BoardLatest", models.BoardLatest{})
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(w, BoardLatest)

	//JoinBoardView
	JoinBoardView, err := gql.TypeToGqlType("JoinBoardView", models.JoinBoardView{})
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(w, JoinBoardView)

	//JoinThreadView
	JoinThreadView, err := gql.TypeToGqlType("JoinThreadView", models.JoinThreadView{})
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(w, JoinThreadView)

	w.Flush()

}
