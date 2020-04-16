package api

import (
	. "github.com/graphql-go/graphql"
	"github.com/mijailr/go-learn/pkg/database"
	. "github.com/mijailr/go-learn/pkg/model"
	"github.com/satori/go.uuid"
)

var alertType = NewObject(ObjectConfig{
	Name:        "Alert",
	Description: "Alert generated",
	Fields: Fields{
		"id": &Field{
			Type:        ID,
			Description: "Alert unique identifier",
		},
		"title": &Field{
			Type:        String,
			Description: "Alert title",
		},
		"content": &Field{
			Type:        String,
			Description: "Alert content",
		},
		"date": &Field{
			Type:        DateTime,
			Description: "Alert date",
		},
		"reported_at": &Field{
			Type:              DateTime,
			Description:       "Reported date",
		},
	},
})

var db = database.Connect()

var getAlert = Field{
	Name:        "Single Alert",
	Description: "Find alert by ID",
	Type:        alertType,
	Resolve:     resolveAlert,
	Args: FieldConfigArgument{
		"id": &ArgumentConfig{
			Type: ID,
		},
	},
}

var listAlerts = Field{
	Name:        "List of Alerts",
	Description: "Return a list of alerts",
	Type:        NewList(alertType),
	Resolve:     resolveAlerts,
}

var alert = Alert{}

func resolveAlert(params ResolveParams) (interface{}, error) {
	id, error := uuid.FromString(params.Args["id"].(string))

	if error != nil {
		return nil, nil

	}
	return alert.GetAlert(db, id)
}

func resolveAlerts(ResolveParams) (interface{}, error) {
	return alert.GetAllAlerts(db)
}
