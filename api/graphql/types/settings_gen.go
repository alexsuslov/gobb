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
	f, err := os.Create("settings.go")
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

	Setting, err := gql.TypeToGqlType("Setting", models.Setting{})
	_, err = fmt.Fprint(w, Setting)

	w.Flush()

}
