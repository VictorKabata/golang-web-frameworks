package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"noteIt/api/models"
)

/**Mock database*/
var notes = []models.Note{
	{1, "Title 1", "Description 1"},
	{2, "Title 2", "Description 2"},
	{3, "Title 3", "Description 3"},
}

func main() {
	fmt.Println("Starting server...")

	fields := graphql.Fields{

		"home": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "home", nil
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

	query := `{
home
}`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Error")
	}

	rJson, _ := json.Marshal(r)
	fmt.Printf("Response: %s\n", rJson)
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
