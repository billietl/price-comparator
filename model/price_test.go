package model

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPriceFields(t *testing.T) {
	t.Parallel()

	result_amount := rand.Float64()
	result_date := Randate()
	result_product := GenerateRandomProduct()
	result_store := GenerateRandomStore()

	price := NewPrice(result_product, result_store, result_amount, result_date)

	assert.NotEqual(t, "", price.ID)
	assert.Equal(t, result_amount, price.Amount)
	assert.Equal(t, result_date, price.Date)

	assert.Equal(t, result_product.ID, price.Product_ID)
	assert.Equal(t, result_store.ID, price.Store_ID)
}

func TestNewPriceNowDate(t *testing.T) {
	t.Parallel()

	price1 := GenerateRandomPrice()
	time.Sleep(1)
	price2 := GenerateRandomPrice()

	assert.NotEqual(t, price1.Date, price2.Date)
}

func TestNewPriceUUID(t *testing.T) {
	t.Parallel()

	product := GenerateRandomProduct()
	store := GenerateRandomStore()
	today := time.Now()

	price1 := NewPrice(product, store, 1.0, today)
	price2 := NewPrice(product, store, 1.0, today)

	assert.NotEqual(t, price1.ID, price2.ID)
}

func TestPriceGenerateID(t *testing.T) {
	t.Parallel()

	price := GenerateRandomPrice()

	id1 := price.ID
	price.GenerateID()
	id2 := price.ID

	assert.NotEqual(t, id1, id2)
}

func TestPriceEquals(t *testing.T) {
	t.Parallel()

	price1 := &Price{
		ID:         uuid.New().String(),
		Amount:     rand.Float64(),
		Date:       Randate(),
		Product_ID: GenerateRandomProduct().ID,
		Store_ID:   GenerateRandomStore().ID,
	}
	price2 := &Price{
		ID:         uuid.New().String(),
		Amount:     price1.Amount,
		Date:       price1.Date,
		Product_ID: price1.Product_ID,
		Store_ID:   price1.Store_ID,
	}
	price3 := &Price{
		ID:         price1.ID,
		Amount:     rand.Float64(),
		Date:       Randate(),
		Product_ID: GenerateRandomProduct().ID,
		Store_ID:   GenerateRandomStore().ID,
	}

	assert.Equal(t, false, price1.Equals(price2))
	assert.Equal(t, true, price1.Equals(price3))

	assert.Equal(t, true, price1.ValueEquals(price2))
	assert.Equal(t, false, price1.ValueEquals(price3))
}
