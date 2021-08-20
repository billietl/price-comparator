package dao

import (
	"context"
	"price-comparator/model"
)

const firestoreProductCollection = "product"

type firestoreProduct struct {
	Name string `firestore:"name"`
	Bio  bool   `firestore:"bio"`
	Vrac bool   `firestore:"vrac"`
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
	p = dao.toModel(&fp)
	p.ID = id
	return
}

func (dao ProductDAOFirestore) Upsert(ctx context.Context, product *model.Product) (err error) {
	pf := dao.fromModel(product)
	_, err = firestoreClient.
		Collection(firestoreProductCollection).
		Doc(product.ID).
		Set(ctx, pf)
	return
}

func (dao ProductDAOFirestore) Delete(ctx context.Context, id string) (err error) {
	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(id).Delete(ctx)
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

func (dao ProductDAOFirestore) toModel(p *firestoreProduct) (product *model.Product) {
	product = &model.Product{
		Name: p.Name,
		Bio:  p.Bio,
		Vrac: p.Vrac,
	}
	return
}

func (dao ProductDAOFirestore) fromModel(p *model.Product) (product *firestoreProduct) {
	product = &firestoreProduct{
		Name: p.Name,
		Bio:  p.Bio,
		Vrac: p.Vrac,
	}
	return
}
