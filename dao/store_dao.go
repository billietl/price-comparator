package dao

import (
	"context"
	"price-comparator/model"
)

type StoreDAO interface {
	Load(ctx context.Context, id string) (*model.Store, error)
	Upsert(ctx context.Context, store *model.Store) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, s *model.Store) (*[]model.Store, error)
	List(ctx context.Context, p *Paginator) (*[]model.Store, error)
	Count(ctx context.Context) (int, error)
	fromModel(s *model.Store) *firestoreStore
	toModel(s *firestoreStore) *model.Store
}
