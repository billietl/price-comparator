package dao

import (
	"context"
	"price-comparator/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyproto/randomstring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func init() {
	randomstring.Seed()
}

func generateStoreTestData(t *testing.T) (id string, result map[string]string) {
	result = map[string]string{}
	result["name"] = randomstring.HumanFriendlyString(10)
	result["city"] = randomstring.HumanFriendlyString(10)
	result["zipcode"] = randomstring.HumanFriendlyString(5)

	ctx := context.Background()
	doc, _, err := firestoreClient.Collection(firestoreStoreCollection).Add(ctx, result)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	id = doc.ID

	return
}

func cleanupStoreTestData(t *testing.T, id string) {
	ctx := context.Background()
	firestoreClient.Collection(firestoreStoreCollection).Doc(id).Delete(ctx)
}

func TestStoreDAOFirestoreCreate(t *testing.T) {
	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Upsert new store
	createdStore := model.NewStore(
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(10),
		randomstring.HumanFriendlyString(5),
	)

	_, err := storeDAO.Upsert(ctx, createdStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	// Reload data
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(createdStore.ID).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.Equal(t, createdStore.Name, docData["name"])
	assert.Equal(t, createdStore.City, docData["city"])
	assert.Equal(t, createdStore.Zipcode, docData["zipcode"])

	// Cleanup test data
	cleanupStoreTestData(t, createdStore.ID)
}

func TestStoreDAOFirestoreRead(t *testing.T) {
	ctx := context.Background()
	// Setup test data
	id, testData := generateStoreTestData(t)

	storeDAO := NewStoreDAOFirestore()

	loadedStore, err := storeDAO.Load(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, testData["name"], loadedStore.Name, "Didn't find the right store name")
	assert.Equal(t, testData["city"], loadedStore.City, "Didn't find the right store city")
	assert.Equal(t, testData["zipcode"], loadedStore.Zipcode, "Didn't find the right store zipcode")
	assert.NotEqual(t, "", loadedStore.ID, "Loaded store didn't have ID")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestStoreDAOFirestoreUpdate(t *testing.T) {
	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	id, testData := generateStoreTestData(t)

	store, err := storeDAO.Load(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	store.Name = randomstring.HumanFriendlyString(10)
	store.City = randomstring.HumanFriendlyString(10)
	store.Zipcode = randomstring.HumanFriendlyString(5)
	storeDAO.Upsert(ctx, store)

	// Reload data
	doc, err := firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	docData := doc.Data()
	assert.NotEqual(t, docData["name"], testData["name"])
	assert.NotEqual(t, docData["city"], testData["city"])
	assert.NotEqual(t, docData["zipcode"], testData["zipcode"])

	// Cleanup test data
	cleanupStoreTestData(t, id)
}

func TestStoreDAOFirestoreDelete(t *testing.T) {
	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	id, _ := generateStoreTestData(t)

	err := storeDAO.Delete(ctx, id)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	_, err = firestoreClient.Collection(firestoreStoreCollection).Doc(id).Get(ctx)
	if grpc.Code(err) != codes.NotFound {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestStoreDAOFirestoreSearch(t *testing.T) {
	ctx := context.Background()
	storeDAO := NewStoreDAOFirestore()

	// Setup test data
	id, testData := generateStoreTestData(t)

	searchedByNameStore := model.Store{
		Name: testData["name"],
	}
	storeList, err := storeDAO.Search(ctx, &searchedByNameStore)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assert.Equal(t, 1, len(*storeList), "Didn't find the right amount of stores")
	assert.Equal(t, testData["name"], (*storeList)[0].Name, "Didn't find the right store name")
	assert.Equal(t, testData["city"], (*storeList)[0].City, "Didn't find the right store city")
	assert.Equal(t, testData["zipcode"], (*storeList)[0].Zipcode, "Didn't find the right store zipcode")

	// Cleanup test data
	cleanupStoreTestData(t, id)
}
