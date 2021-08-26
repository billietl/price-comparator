package model

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Price struct {
	ID         string     `json:"id" bson:"id"`
	Amount     float64    `json:"amount" bson:"amount"`
	Date       *time.Time `json:"date" bson:"date"`
	Product_ID string     `json:"product_id" bson:"product_id"`
	Store_ID   string     `json:"store_id" bson:"store_id"`
}

func NewPriceNow(product *Product, store *Store, amount float64) *Price {
	now := time.Now()
	return NewPrice(product, store, amount, &now)
}

func NewPrice(product *Product, store *Store, amount float64, date *time.Time) *Price {
	return &Price{
		ID:         uuid.New().String(),
		Amount:     amount,
		Date:       date,
		Product_ID: product.ID,
		Store_ID:   store.ID,
	}
}

func (this *Price) GenerateID() {
	this.ID = uuid.New().String()
}

func (this Price) Equals(price *Price) bool {
	return this.ID == price.ID
}

func (this Price) ValueEquals(price *Price) bool {
	return this.Amount == price.Amount &&
		this.Date.Equal(*price.Date) &&
		this.Product_ID == price.Product_ID &&
		this.Store_ID == price.Store_ID
}

func GenerateRandomPrice() *Price {
	date := Randate()
	return NewPrice(
		GenerateRandomProduct(),
		GenerateRandomStore(),
		rand.Float64(),
		&date,
	)
}

func Randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
