package dto

import (
	"context"
	"google.golang.org/api/iterator"
)

const firestoreProductCollection = "product"

type Product struct {
	ID   string
	Name string `firestore:"name"`
	Bio  bool   `firestore:"bio"`
}

func (p *Product) Load() (err error) {
	ctx := context.Background()
	d, err := firestoreClient.Collection(firestoreProductCollection).Doc(p.ID).Get(ctx)
	if err != nil {
		return
	}
	err = d.DataTo(&p)
	if err != nil {
		return
	}
	return
}

func (p *Product) Upsert() (err error) {
	ctx := context.Background()
	if p.ID == "" {
		ref := firestoreClient.Collection(firestoreProductCollection).NewDoc()
		p.ID = ref.ID
	}
	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(p.ID).Set(ctx, p)
	if err != nil {
		return
	}
	return
}

func (p *Product) Delete() (err error) {
	ctx := context.Background()
	_, err = firestoreClient.Collection(firestoreProductCollection).Doc(p.ID).Delete(ctx)
	if err != nil {
		return
	}
	return
}

func (p *Product) Search() (productList []Product, err error) {
	ctx := context.Background()
	// Build query
	query := firestoreClient.Collection(firestoreProductCollection).Select()
	if p.Name != "" {
		query = query.Where("name", "==", p.Name)
	}
	if p.Bio != false {
		query = query.Where("bio", "==", p.Bio)
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
			return []Product{}, err
		}
		docID := doc.Ref.ID
		newProduct := Product{
			ID: docID,
		}
		newProduct.Load()
		productList = append(productList, newProduct)
	}
	// Cleanup
	docs.Stop()
	return
}
