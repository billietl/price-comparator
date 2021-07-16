package dao

import (
	"context"
	"price-comparator/model"
)

type StoreDAO interface {
	Load(ctx context.Context, id string) (*model.Store, error)
	Upsert(ctx context.Context, store *model.Store) (*model.Store, error)
	Delete(ctx context.Context, store *model.Store) error
	Search(ctx context.Context, s *model.Store) (*[]model.Store, error)
}
