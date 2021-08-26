package dao

import (
	"context"
	"price-comparator/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
)

func init() {
	randomstring.Seed()
	initFirestore(context.Background())
}

func generatePriceTestData(t *testing.T) (price *model.Price) {
	price = model.GenerateRandomPrice()

	ctx := context.Background()
	dao := NewPriceDAOFirestore()
	doc, _, err := firestoreClient.Collection(firestoreProductCollection).Add(ctx, dao.fromModel(price))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	price.ID = doc.ID

	return
}

func TestPriceDAOFirestoreCreate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	priceDAO := NewPriceDAOFirestore()

	// Upsert new price
	createdPrice := model.GenerateRandomPrice()

	err := priceDAO.Upsert(ctx, createdPrice)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	defer firestoreClient.Collection(firestorePriceCollection).Doc(createdPrice.ID).Delete(ctx)

	// Reload product
	doc, err := firestoreClient.Collection(firestorePriceCollection).Doc(createdPrice.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	docDate, _ := time.Parse(time.UnixDate, docData["date"].(string))
	assert.Equal(t, createdPrice.Amount, docData["amount"])
	assert.Equal(t, createdPrice.Date, docDate)
	assert.Equal(t, createdPrice.Product_ID, docData["product_id"])
	assert.Equal(t, createdPrice.Store_ID, docData["store_id"])
}
