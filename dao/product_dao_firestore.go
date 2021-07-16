package dao

import (
	"context"
	"price-comparator/model"
)

const firestoreProductCollection = "product"

type firestoreProduct struct {
	Name string `firestore:"name"`
	Bio  bool   `firestore:"bio"`
}

type ProductDAOFirestore struct{}

func NewProductDAOFirestore() *ProductDAOFirestore {
	return &ProductDAOFirestore{}
}

func (dao ProductDAOFirestore) Load(ctx context.Context, id string) (p *model.Product, err error) {
	d, err := firestoreClient.Collection(firestoreProductCollection).Doc(id).Get(ctx)
	if err != nil {
		return
	}
	fp := firestoreProduct{}
	err = d.DataTo(&fp)
	if err != nil {
		return
	}
	p = &model.Product{
		ID:   id,
		Name: fp.Name,
		Bio:  fp.Bio,
	}
	return
}

func (dao ProductDAOFirestore) Upsert(ctx context.Context, product *model.Product) (p *model.Product, err error) {
	pf := firestoreProduct{
		Name: product.Name,
		Bio:  product.Bio,
	}
	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(product.ID).Set(ctx, pf)
	if err != nil {
		return
	}
	return
}

func (dao ProductDAOFirestore) Delete(ctx context.Context, id string) (err error) {
	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(id).Delete(ctx)
	if err != nil {
		return
	}
	return
}

func (dao ProductDAOFirestore) Search(ctx context.Context, p *model.Product) (*[]model.Product, error) {
	// Build query
	query := firestoreClient.Collection(firestoreProductCollection).Select()
	if p.Name != "" {
		query = query.Where("name", "==", p.Name)
	}
	query = query.Where("bio", "==", p.Bio)
	// Retrieve documents
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	result := make([]model.Product, 0, len(docs))
	for _, doc := range docs {
		newProduct, err := dao.Load(ctx, (*doc).Ref.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, *newProduct)
	}
	return &result, nil
}
