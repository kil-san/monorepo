package connection

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	log "unknwon.dev/clog/v2"
)

func NewFirestoreClient(ctx context.Context) *firestore.Client {
	conf := &firebase.Config{ProjectID: os.Getenv("GOOGLE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatal("%v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("%v", err)
	}

	return client
}
