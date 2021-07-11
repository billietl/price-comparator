package dto

import (
	"context"
	"google.golang.org/api/iterator"
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

func (s *Store) Search() (storeList []Store, err error) {
	ctx := context.Background()
	// Build query
	query := firestoreClient.Collection(firestoreStoreCollection).Select()
	if s.Name != "" {
		query = query.Where("name", "==", s.Name)
	}
	if s.City != "" {
		query = query.Where("city", "==", s.City)
	}
	if s.Zipcode != "" {
		query = query.Where("zipcode", "==", s.Zipcode)
	}
	// Retrieve documents
	docs := query.Documents(ctx)
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			docs.Stop()
			return []Store{}, err
		}
		docID := doc.Ref.ID
		newStore := Store{
			ID: docID,
		}
		newStore.Load()
		storeList = append(storeList, newStore)
	}
	// Cleanup
	docs.Stop()
	return
}
