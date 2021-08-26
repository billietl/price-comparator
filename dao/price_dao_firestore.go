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

func (this PriceDAOFirestore) toModel(p *firestorePrice) *model.Price {
	date, _ := time.Parse(time.UnixDate, p.Date)
	return &model.Price{
		Amount:     p.Amount,
		Date:       date,
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
