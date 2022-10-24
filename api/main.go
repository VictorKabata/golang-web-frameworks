package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

type Note struct {
	ID          int    `json: "ID"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

/**Mock database*/
var notes = []Note{
	{1, "Title 1", "Description 1"},
	{2, "Title 2", "Description 2"},
	{3, "Title 3", "Description 3"},
}

func main() {
	fmt.Println("Starting server...")

	noteType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Note",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	fields := graphql.Fields{

		"note": &graphql.Field{
			Type:        noteType,
			Description: "Get note by ID",
			Args:        graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.Int}},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, isOkay := p.Args["id"].(int)
				if isOkay {
					for _, note := range notes {
						if note.ID == id {
							return note, nil
						}
					}
				}
				return nil, nil
			},
		},

		"notes": &graphql.Field{
			Type:        graphql.NewList(noteType),
			Description: "Get list of all notes",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return notes, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", handler)
	http.Handle("/sandbox", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { w.Write(sandboxHTML) }))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
new window.EmbeddedSandbox({
  target: "#sandbox",
  // Pass through your server href if you are embedding on an endpoint.
  // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
  initialEndpoint: "http://localhost:8080/graphql",
});
// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body>
 
</html>`)
