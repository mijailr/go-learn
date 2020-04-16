package api

import (
	. "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = NewObject(
	ObjectConfig{
		Name: "Query",
		Fields: Fields{
			"alert":  &getAlert,
			"alerts": &listAlerts,
		},
	},
)

var schema, _ = NewSchema(
	SchemaConfig{
		Query: queryType,
	},
)

func Handler() *handler.Handler {
	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
}
