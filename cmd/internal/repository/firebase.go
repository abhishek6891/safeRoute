package repository

import (
	"context"
	"log"
	"saferoute/internal/domain"

	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var firebaseClient *db.Client

func InitDB(creds string) error {
	opt := option.WithCredentialsJSON([]byte(creds))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	client, err := app.Database(context.Background())
	firebaseClient = client
	return err
}

func SaveHazard(hazard domain.Hazard) error {
	ref := firebaseClient.NewRef("hazards")
	return ref.Push(context.Background(), hazard)
}
