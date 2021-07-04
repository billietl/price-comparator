package dto

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
)

var firestoreClient *firestore.Client

func init() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "price-comparator-dev")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	firestoreClient = client

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		firestoreClient.Close()
		done <- true
	}()
}
