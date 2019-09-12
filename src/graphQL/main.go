package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	// communicates with sqlite3 database
	_ "github.com/mattn/go-sqlite3"
)

// Tutorial: in memory tutorial
type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

// Author: author
type Author struct {
	Name      string
	Tutorials []int
}

// Comment: comment s
type Comment struct {
	Body string
}

// Populate: returns an arrat of Tutorials
func Populate() []Tutorial {
	author := &Author{Name: "Elliot Forbes", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "My first comment"},
			Comment{Body: "My second comment"},
		},
	}
	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	return tutorials
}

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				// we'll use NewList to deal with an array
				// of int values
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

// tutorials - Global Variable
var tutorials []Tutorial = Populate()

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        tutorialType,
			Description: "Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				tutorial := Tutorial{
					Title: params.Args["title"].(string),
				}
				tutorials = append(tutorials, tutorial)
				return tutorial, nil
			},
		},
	},
})

/*
Simple GraphQL Server
*/
func main() {
	// schema
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type: tutorialType,
			// good practice to add description
			Description: "Get Tutorial By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			}, Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					// parse tutorial for matching id
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		// list endpoint returns all tutorials
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db, err := sql.Open("sqlite3", "./foo.db")
				if err != nil {
					log.Fatat(err)
				}
				defer db.Close()
				var tutorials []Tutorial
				results, err := db.Query("SELECT * FROM tutorials")
				if err != nil {
					fmt.Println(err)
				}
				// iterate through all the results 
				for results.Next() {
					var tutorial Tutorial 
					err = results.Scane(&tutorial.ID, &tutorial.Title)
					if err != nil {
						fmt.Println(err)
					}
					log.Println(tutorial)
					tutorials = append(tutorials, tutorial)
				}
				return tutorials , nil 
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create a schema, error: %v", err)
	}

	// query 1
	// query := `
	// 	{
	// 		list {
	// 			id
	// 			title
	// 			comments {
	// 				body
	// 			}
	// 			author {
	// 				Name
	// 				Tutorials
	// 			}
	// 		}
	// 	}
	// `

	// query2
	// query := `
	// 	mutation {
	// 		create (title: "Hello World") {
	// 			title
	// 		}
	// 	}
	// `

	// query3
	query := `
	 	{
			list {
				id 
				title
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, error %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
