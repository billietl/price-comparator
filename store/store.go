package store

import (
	"context"
	"github.com/google/uuid"
	"price-comparator/firestoreclient"
)

type StoreIface interface {
	Save()
}

type Store struct {
	Id      string
    Name    string `firestore:"name"`
	City    string `firestore:"city"`
	Zipcode string `firestore:"zipcode"`
}

func (s Store) Save() {
	return
}

func New() (s Store) {
	s = Store{}
	s.Id = uuid.New().String()
	return
}

func Load(id string) (s Store, err error) {
	ctx := context.Background()
	client, err := firestoreclient.GetFirestoreClient()
	if err != nil {
		return
	}
	d, err := client.Collection("store").Doc(id).Get(ctx)
	if err != nil {
		return
	}
	err = d.DataTo(&s)
	if err != nil {
		return
	}
	s.Id = id
	return
}
