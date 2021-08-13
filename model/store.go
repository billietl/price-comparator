package model

import (
	"github.com/google/uuid"
)

type Store struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	City    string `json:"city" bson:"city"`
	Zipcode string `json:"zipcode" bson:"zipcode"`
}

func NewStore(name, city, zipcode string) *Store {
	return &Store{
		ID:      uuid.New().String(),
		Name:    name,
		City:    city,
		Zipcode: zipcode,
	}
}

func (s *Store) GenerateID() {
	s.ID = uuid.New().String()
}

func (s Store) Equals(store *Store) bool {
	return s.ID == store.ID
}

func (s Store) ValueEquals(store *Store) bool {
	return s.Name == store.Name &&
		s.City == store.City &&
		s.Zipcode == store.Zipcode
}
