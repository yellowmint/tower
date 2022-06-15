package fbase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type Clients struct {
	Auth      *auth.Client
	Firestore *firestore.Client
}

func NewClients(gcpProjectId string) (*Clients, error) {
	ctx := context.Background()

	cfg := firebase.Config{
		ProjectID: gcpProjectId,
	}

	app, err := firebase.NewApp(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &Clients{
		Auth:      authClient,
		Firestore: firestoreClient,
	}, nil
}

func (clients *Clients) Cleanup() error {
	return clients.Firestore.Close()
}
