package dao

import (
	"context"
	"price-comparator/model"
)

type ProductDAO interface {
	Load(ctx context.Context, id string) (*model.Product, error)
	Upsert(ctx context.Context, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, product *model.Product) error
	Search(ctx context.Context, p *model.Product) (*[]model.Product, error)
}
