package dto

import (
	"context"
)

const firestoreStoreCollection = "store"

type Store struct {
	ID      string
	Name    string `firestore:"name"`
	City    string `firestore:"city"`
	Zipcode string `firestore:"zipcode"`
}

func (s *Store) Load() (err error) {
	ctx := context.Background()
	d, err := firestoreClient.Collection(firestoreStoreCollection).Doc(s.ID).Get(ctx)
	if err != nil {
		return
	}
	err = d.DataTo(&s)
	if err != nil {
		return
	}
	return
}

func (s *Store) Upsert() (err error) {
	ctx := context.Background()
	if s.ID == "" {
		ref := firestoreClient.Collection(firestoreStoreCollection).NewDoc()
		s.ID = ref.ID
	}
	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(s.ID).Set(ctx, s)
	if err != nil {
		return
	}
	return
}

func (s *Store) Delete() (err error) {
	ctx := context.Background()
	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(s.ID).Delete(ctx)
	if err != nil {
		return
	}
	return
}
