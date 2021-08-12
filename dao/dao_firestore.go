package dao

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var firestoreClient *firestore.Client

func initFirestore(ctx context.Context) {
	client, err := firestore.NewClient(ctx, "foobar")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	firestoreClient = client
}

func shutDownFirestoreClient() {
	err := firestoreClient.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}
