package dao

import (
	"context"
	"price-comparator/model"
	"time"
)

const firestorePriceCollection = "price"

type firestorePrice struct {
	Amount     float64 `firestore:"amount"`
	Date       string  `firestore:"date"`
	Product_ID string  `firestore:"product_id"`
	Store_ID   string  `firestore:"store_id"`
}

type PriceDAOFirestore struct{}

func NewPriceDAOFirestore() *PriceDAOFirestore {
	return &PriceDAOFirestore{}
}

func (this PriceDAOFirestore) Load(ctx context.Context, id string) (price *model.Price, err error) {
	doc, err := firestoreClient.Collection(firestorePriceCollection).Doc(id).Get(ctx)
	if err != nil {
		return
	}
	priceEntity := firestorePrice{}
	err = doc.DataTo(&priceEntity)
	if err != nil {
		return
	}
	price = this.toModel(&priceEntity)
	price.ID = id
	return
}

func (this PriceDAOFirestore) Upsert(ctx context.Context, price *model.Price) (err error) {
	pf := this.fromModel(price)
	_, err = firestoreClient.
		Collection(firestorePriceCollection).
		Doc(price.ID).
		Set(ctx, pf)
	return
}

func (this PriceDAOFirestore) Delete(ctx context.Context, id string) (err error) {
	_, err = firestoreClient.Collection(firestorePriceCollection).Doc(id).Delete(ctx)
	return
}

func (this PriceDAOFirestore) toModel(p *firestorePrice) *model.Price {
	date, _ := time.Parse(time.UnixDate, p.Date)
	return &model.Price{
		Amount:     p.Amount,
		Date:       &date,
		Product_ID: p.Product_ID,
		Store_ID:   p.Store_ID,
	}
}

func (this PriceDAOFirestore) fromModel(price *model.Price) *firestorePrice {
	return &firestorePrice{
		Amount:     price.Amount,
		Date:       price.Date.Format(time.UnixDate),
		Product_ID: price.Product_ID,
		Store_ID:   price.Store_ID,
	}
}

func (this PriceDAOFirestore) Search(ctx context.Context, price *model.Price) (*[]model.Price, error) {
	// Build query
	query := firestoreClient.Collection(firestorePriceCollection).Select()
	if price.Amount != 0 {
		query = query.Where("amount", "==", price.Amount)
	}
	if price.Date != nil {
		query = query.Where("date", "==", price.Date)
	}
	if price.Store_ID != "" {
		query = query.Where("store_id", "==", price.Store_ID)
	}
	if price.Product_ID != "" {
		query = query.Where("product_id", "==", price.Product_ID)
	}
	// Retrieve documents
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	result := make([]model.Price, 0, len(docs))
	for _, doc := range docs {
		newPrice, err := this.Load(ctx, (*doc).Ref.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, *newPrice)
	}
	return &result, nil
}
