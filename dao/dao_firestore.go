package dao

import (
	"context"
	"fmt"
	"os"
	"price-comparator/logger"

	"cloud.google.com/go/firestore"
)

var firestoreClient *firestore.Client

func initFirestore(ctx context.Context) {
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		projectID = "foobar"
	}
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	firestoreClient = client
}

func shutDownFirestoreClient() {
	err := firestoreClient.Close()
	if err != nil {
		logger.Error(err, "Error closing firestore client")
	}
}
