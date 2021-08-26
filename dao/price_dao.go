package dao

import (
	"context"
	"price-comparator/model"
)

type PriceDAO interface {
	Load(ctx context.Context, id string) (*model.Price, error)
	Upsert(ctx context.Context, price *model.Price) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, p *model.Price) (*[]model.Price, error)
	fromModel(p *model.Price) *firestorePrice
	toModel(p *firestorePrice) *model.Price
}
