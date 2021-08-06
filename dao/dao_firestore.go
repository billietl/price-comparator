package dao

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
)

var firestoreClient *firestore.Client

func initFirestore() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "foobar")
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
