package dao

import (
	"context"
	"math/rand"
	"price-comparator/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func init() {
	randomstring.Seed()
	initFirestore(context.Background())
}

func generatePriceTestData(t *testing.T) (price *model.Price) {
	price = model.GenerateRandomPrice()

	ctx := context.Background()
	dao := NewPriceDAOFirestore()
	doc, _, err := firestoreClient.Collection(firestorePriceCollection).Add(ctx, dao.fromModel(price))
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

	// Reload price
	doc, err := firestoreClient.Collection(firestorePriceCollection).Doc(createdPrice.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	docDate, _ := time.Parse(time.UnixDate, docData["date"].(string))
	assert.Equal(t, createdPrice.Amount, docData["amount"])
	assert.Equal(t, true, createdPrice.Date.Equal(docDate))
	assert.Equal(t, createdPrice.Product_ID, docData["product_id"])
	assert.Equal(t, createdPrice.Store_ID, docData["store_id"])
}

func TestPriceDAOFirestoreRead(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	// Setup test data
	testPrice := generatePriceTestData(t)
	defer firestoreClient.Collection(firestorePriceCollection).Doc(testPrice.ID).Delete(ctx)

	priceDAO := NewPriceDAOFirestore()

	loadedPrice, err := priceDAO.Load(ctx, testPrice.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testPrice.Amount, loadedPrice.Amount, "Didn't find the right price amount")
	assert.Equal(t, testPrice.Date, loadedPrice.Date, "Didn't find the right price date")
	assert.Equal(t, testPrice.Product_ID, loadedPrice.Product_ID, "Didn't find the right product price")
	assert.Equal(t, testPrice.Store_ID, loadedPrice.Store_ID, "Didn't find the right store price")
	assert.NotEqual(t, "", loadedPrice.ID, "Loaded store didn't have ID")
}

func TestPriceDAOFirestoreUpdate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	priceDAO := NewPriceDAOFirestore()

	// Setup test data
	testPrice := generatePriceTestData(t)
	defer firestoreClient.Collection(firestorePriceCollection).Doc(testPrice.ID).Delete(ctx)

	price, err := priceDAO.Load(ctx, testPrice.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	newDate := model.Randate()
	price.Amount = rand.Float64()
	price.Date = &newDate
	price.Product_ID = uuid.New().String()
	price.Store_ID = uuid.New().String()
	priceDAO.Upsert(ctx, price)

	// Reload data
	doc, err := firestoreClient.Collection(firestorePriceCollection).Doc(testPrice.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, testPrice.Amount, docData["amount"])
	assert.NotEqual(t, testPrice.Date, docData["date"])
	assert.NotEqual(t, testPrice.Product_ID, docData["product_id"])
	assert.NotEqual(t, testPrice.Store_ID, docData["store_id"])
}

func TestPriceDAOFirestoreDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	priceDAO := NewPriceDAOFirestore()

	// Setup test data
	testPrice := generatePriceTestData(t)

	// Delete data
	err := priceDAO.Delete(ctx, testPrice.ID)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(testPrice.ID).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestPriceDAOFirestoreSearch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	PriceDAO := NewPriceDAOFirestore()

	// Setup test data
	testPrice := generatePriceTestData(t)
	defer firestoreClient.Collection(firestorePriceCollection).Doc(testPrice.ID).Delete(ctx)

	searchedByProductPrice := model.Price{
		Product_ID: testPrice.Product_ID,
	}
	PriceList, err := PriceDAO.Search(ctx, &searchedByProductPrice)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if !assert.Equal(t, 1, len(*PriceList), "Didn't find the right amount of Prices") {
		t.Fail()
		return
	}
	assert.Equal(t, testPrice.Amount, (*PriceList)[0].Amount, "Didn't find the right price amount")
	assert.Equal(t, true, testPrice.Date.Equal(*((*PriceList)[0].Date)))
	// assert.Equal(t, time.Parse(time.UniwDate, *testPrice.Date), time.Parse(time.UnixDate, *((*PriceList)[0].Date)), "Didn't find the right price date")
	assert.Equal(t, testPrice.Product_ID, (*PriceList)[0].Product_ID, "Didn't find the right product price")
	assert.Equal(t, testPrice.Store_ID, (*PriceList)[0].Store_ID, "Didn't find the right store price")
}
