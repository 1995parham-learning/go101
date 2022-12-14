package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
)

const diceFaces = 6

func main() {
	rand.Seed(time.Now().Unix())

	// defining the schema for graphql
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"diceRoll": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
			Args: graphql.FieldConfigArgument{
				"count": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var l []int
				n := 2

				if count, ok := p.Args["count"].(int); ok {
					n = count
				}

				for i := 0; i < n; i++ {
					// nolint: gosec
					l = append(l, rand.Intn(diceFaces)+1)
				}

				return l, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// query in graphql
	query := `
	{
		diceRoll
	}
	`

	r := graphql.Do(graphql.Params{Schema: schema, RequestString: query})
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("json marshal failed %s", err)
	}

	fmt.Printf("client request: %s\n", query)
	fmt.Printf("server answer: %s\n", rj)
}
