package dao

import (
	"context"
	"price-comparator/model"
)

const firestoreStoreCollection = "store"

type firestoreStore struct {
	Name    string `firestore:"name"`
	City    string `firestore:"city"`
	Zipcode string `firestore:"zipcode"`
}

type StoreDAOFirestore struct{}

func NewStoreDAOFirestore() *StoreDAOFirestore {
	return &StoreDAOFirestore{}
}

func (dao StoreDAOFirestore) Load(ctx context.Context, id string) (s *model.Store, err error) {
	d, err := firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if err != nil {
		return
	}
	fs := firestoreStore{}
	err = d.DataTo(&fs)
	if err != nil {
		return
	}
	s = &model.Store{
		ID:      id,
		Name:    fs.Name,
		City:    fs.City,
		Zipcode: fs.Zipcode,
	}
	return
}

func (dao StoreDAOFirestore) Upsert(ctx context.Context, store *model.Store) (err error) {
	sf := firestoreStore{
		Name:    store.Name,
		City:    store.City,
		Zipcode: store.Zipcode,
	}
	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(store.ID).Set(ctx, sf)
	return
}

func (dao StoreDAOFirestore) Delete(ctx context.Context, id string) (err error) {
	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(id).Delete(ctx)
	return
}

func (dao StoreDAOFirestore) Search(ctx context.Context, s *model.Store) (storeList *[]model.Store, err error) {
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
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return
	}
	result := make([]model.Store, 0, len(docs))
	for _, doc := range docs {
		newStore, err := dao.Load(ctx, (*doc).Ref.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, *newStore)
	}
	return &result, nil
}

func (dao StoreDAOFirestore) fromModel(s *model.Store) (store *firestoreStore) {
	store = &firestoreStore{
		Name:    s.Name,
		City:    s.City,
		Zipcode: s.Zipcode,
	}
	return
}

func (dao StoreDAOFirestore) toModel(s *firestoreStore) (store *model.Store) {
	store = &model.Store{
		Name:    s.Name,
		City:    s.City,
		Zipcode: s.Zipcode,
	}
	return
}
