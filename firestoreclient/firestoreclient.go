package firestoreclient

import (
	"context"
	"cloud.google.com/go/firestore"
)

func GetFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "price-comparator-dev")
	return client, err
}
